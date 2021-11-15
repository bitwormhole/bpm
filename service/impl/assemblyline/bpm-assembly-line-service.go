package assemblyline

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
)

// AssemblyLineServiceImpl ...
type AssemblyLineServiceImpl struct {
	markup.Component ` id:"bpm-assembly-line-service" class:"bpm-service" `
}

func (inst *AssemblyLineServiceImpl) _Impl() service.AssemblyLineService {
	return inst
}

// Assembly ...
func (inst *AssemblyLineServiceImpl) Assembly(ctx context.Context, targetNames []string) error {
	for _, target := range targetNames {
		task := assemblyLineTask{}
		err := task.init(ctx, target)
		if err != nil {
			return err
		}
		err = task.run()
		if err != nil {
			return err
		}
	}
	return nil
}

///////////////////////////////////////////////////////////////////////

type assemblyLineCopyFile struct {
	parent *assemblyLineTask
	from   fs.Path
	to     fs.Path
}

func (inst *assemblyLineCopyFile) copy() error {
	from := inst.from
	to := inst.to
	console := inst.parent.console
	console.WriteString(fmt.Sprintln("copy file from ", from.Path()))
	console.WriteString(fmt.Sprintln("            to ", to.Path()))
	dir := to.Parent()
	if !dir.Exists() {
		dir.Mkdirs()
	}
	return from.CopyTo(to)
}

///////////////////////////////////////////////////////////////////////

type assemblyLineTask struct {
	context context.Context
	console cli.Console

	now             int64
	timestamp       string
	assemblyBpmConf collection.Properties // the 'assembly.bpm' props

	targetName      string
	packageName     string
	packageVersion  string
	packageRevision string // (string format for int)
	packagePlatform string
	packageState    string // (alpha|beta|stable)
	packageFileName string

	wkdir           fs.Path
	assemblyBpmFile fs.Path // the 'assembly.bpm' file

	distributionDir          fs.Path
	distributionSourceFile   fs.Path // the 'package-list.source' file
	distributionBpmFile      fs.Path
	distributionBpmPropsFile fs.Path

	packagingDir                fs.Path
	packagingEtcAppInitFile     fs.Path // the '.bpm/files/etc/???/init' file
	packagingBpmConfigFile      fs.Path // the '.bpm.config' file
	packagingOutputBpmFile      fs.Path // the 'dist/????.bpm' file
	packagingOutputBpmPropsFile fs.Path // the 'dist/????.bpm.properties' file

	sourceDir   fs.Path
	sourceFiles []*assemblyLineCopyFile
}

func (inst *assemblyLineTask) init(ctx context.Context, targetName string) error {
	console, err := cli.GetConsole(ctx)
	if err != nil {
		return err
	}
	now := util.CurrentTimestamp()
	timestamp := strconv.FormatInt(now, 10)
	wkdir := console.GetWorkingPath()
	inst.wkdir = wkdir
	inst.context = ctx
	inst.console = console
	inst.targetName = targetName
	inst.now = now
	inst.timestamp = timestamp
	return nil
}

func (inst *assemblyLineTask) run() error {
	const nl = "\n"
	inst.console.WriteString("bpm Assembly" + nl)
	inst.console.WriteString("  wkdir=" + inst.wkdir.Path() + nl)

	err := inst.findAssemblyBpmFile()
	if err != nil {
		return err
	}

	err = inst.loadAssemblyBpmConfig()
	if err != nil {
		return err
	}

	err = inst.locateFiles()
	if err != nil {
		return err
	}

	err = inst.injectProperties()
	if err != nil {
		return err
	}

	err = inst.loadSourceFileList()
	if err != nil {
		return err
	}

	err = inst.doCopyFilesFromSource()
	if err != nil {
		return err
	}

	err = inst.doCopyFilesFromSource()
	if err != nil {
		return err
	}

	err = inst.doBpmMake()
	if err != nil {
		return err
	}

	err = inst.doCopyProductsToDist()
	if err != nil {
		return err
	}

	err = inst.updatePackageList()
	if err != nil {
		return err
	}

	return nil
}

// findAssemblyBpmFile find 'assembly.bpm', the config for this command
func (inst *assemblyLineTask) findAssemblyBpmFile() error {
	targetName := inst.targetName + ".bpm.assembly"
	wkdir := inst.wkdir
	dir := wkdir
	for ; dir != nil; dir = dir.Parent() {
		file := dir.GetChild(targetName)
		if file.IsFile() {
			inst.assemblyBpmFile = file
			return nil
		}
	}
	return errors.New("no file [" + targetName + "] in path [" + wkdir.Path() + "]")
}

func (inst *assemblyLineTask) loadProperties(file fs.Path) (collection.Properties, error) {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return collection.ParseProperties(text, nil)
}

func (inst *assemblyLineTask) importExtendConfig(props1 collection.Properties) (collection.Properties, error) {

	const prefix = "import."
	const suffix = ".from"

	props2 := collection.CreateProperties()
	props2.Import(props1.Export(nil))
	err := collection.ResolvePropertiesVar(props1)
	if err != nil {
		return nil, err
	}
	ids := inst.listPropertyID(props1, prefix, suffix)

	for _, id := range ids {
		value, err := props1.GetPropertyRequired(prefix + id + suffix)
		if err != nil {
			return nil, err
		}
		file := fs.Default().GetPath(value)
		text, err := file.GetIO().ReadText(nil)
		if err != nil {
			return nil, err
		}
		props3, err := collection.ParseProperties(text, nil)
		if err != nil {
			return nil, err
		}
		props2.Import(props3.Export(nil))
	}

	return props2, nil
}

func (inst *assemblyLineTask) loadAssemblyBpmConfig() error {

	file := inst.assemblyBpmFile
	props, err := inst.loadProperties(file)
	if err != nil {
		vlog.Error("in file ", file.Path())
		return err
	}

	props, err = inst.importExtendConfig(props)
	if err != nil {
		vlog.Error("in file ", file.Path())
		return err
	}

	err = collection.ResolvePropertiesVar(props)
	if err != nil {
		vlog.Error("in file ", file.Path())
		return err
	}

	const (
		keySourceDir       = "dir.source.path"
		keyPackagingDir    = "dir.packaging.path"
		keyDistributionDir = "dir.distribution.path"
		keyPackageName     = "package.name"
		keyPackageVer      = "package.version"
		keyPackageRev      = "package.revision"
		keyPackagePlatform = "package.platform"
		keyPackageState    = "package.state"
		keyPackageFilename = "package.filename"
	)

	defaultFS := fs.Default()
	getter := props.Getter()

	strSourceDir := getter.GetString(keySourceDir, "")
	strPackageDir := getter.GetString(keyPackagingDir, "")
	strDistDir := getter.GetString(keyDistributionDir, "")

	inst.packageName = getter.GetString(keyPackageName, "")
	inst.packageVersion = getter.GetString(keyPackageVer, "")
	inst.packageRevision = getter.GetString(keyPackageRev, "")
	inst.packagePlatform = getter.GetString(keyPackagePlatform, "")
	inst.packageState = getter.GetString(keyPackageState, "")
	inst.packageFileName = getter.GetString(keyPackageFilename, "")

	err = getter.Error()
	if err != nil {
		vlog.Error("in file " + file.Path())
		return err
	}

	inst.sourceDir = defaultFS.GetPath(strSourceDir)
	inst.packagingDir = defaultFS.GetPath(strPackageDir)
	inst.distributionDir = defaultFS.GetPath(strDistDir)

	inst.assemblyBpmConf = props
	return nil
}

func (inst *assemblyLineTask) locateFiles() error {

	filename := inst.packageFileName

	// source
	dir := inst.sourceDir

	// packaging
	dir = inst.packagingDir
	inst.packagingBpmConfigFile = inst.GetChild(dir, ".bpm.config")
	inst.packagingEtcAppInitFile = inst.GetChild(dir, ".bpm/files/etc/${package.name}/init")
	inst.packagingOutputBpmFile = inst.GetChild(dir, "dist/"+filename)
	inst.packagingOutputBpmPropsFile = inst.GetChild(dir, "dist/"+filename+".properties")

	// distribution
	dir = inst.distributionDir
	inst.distributionBpmFile = inst.GetChild(dir, "packages/${package.name}/${package.version}/"+filename)
	inst.distributionBpmPropsFile = inst.GetChild(dir, "packages/${package.name}/${package.version}/"+filename+".properties")
	inst.distributionSourceFile = inst.GetChild(dir, "etc/bpm/sources/packages-${package.platform}-${package.state}.source")

	return nil
}

func (inst *assemblyLineTask) injectProperties() error {
	list := make([]fs.Path, 0)
	list = append(list, inst.packagingEtcAppInitFile)
	list = append(list, inst.packagingBpmConfigFile)
	for _, file := range list {
		err := inst.injectPropertiesToFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *assemblyLineTask) injectPropertiesToFile(file fs.Path) error {

	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	platform, err := props.GetPropertyRequired("package.platform")
	if err != nil {
		vlog.Error("in file " + file.Path())
		return err
	}
	if platform != inst.packagePlatform {
		return errors.New(fmt.Sprintln("bad platform, want:", inst.packagePlatform, " have:", platform))
	}

	props.SetProperty("package.state", inst.packageState)
	props.SetProperty("package.version", inst.packageVersion)
	props.SetProperty("package.revision", inst.packageRevision)

	text = collection.FormatPropertiesWithSegment(props)
	return file.GetIO().WriteText(text, nil, false)
}

func (inst *assemblyLineTask) GetChild(dir fs.Path, path string) fs.Path {

	path = strings.ReplaceAll(path, "${package.name}", inst.packageName)
	path = strings.ReplaceAll(path, "${package.version}", inst.packageVersion)
	path = strings.ReplaceAll(path, "${package.revision}", inst.packageRevision)
	path = strings.ReplaceAll(path, "${package.platform}", inst.packagePlatform)
	path = strings.ReplaceAll(path, "${package.state}", inst.packageState)

	return dir.GetChild(path)
}

func (inst *assemblyLineTask) doBpmMake() error {
	c := exec.Command("bpm", "make")
	c.Dir = inst.packagingDir.Path()
	return inst.exec(c)
}

func (inst *assemblyLineTask) listPropertyID(props collection.Properties, prefix, suffix string) []string {
	ids := make([]string, 0)
	table := props.Export(nil)
	for name := range table {
		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, suffix) {
			id := name[len(prefix) : len(name)-len(suffix)]
			if len(id) > 0 {
				ids = append(ids, id)
			}
		}
	}
	return ids
}

func (inst *assemblyLineTask) loadSourceFileList() error {

	const prefix = "copy."
	const suffix = ".to"

	props := inst.assemblyBpmConf
	list := make([]*assemblyLineCopyFile, 0)
	getter := props.Getter()
	deffs := fs.Default()
	ids := inst.listPropertyID(props, prefix, suffix)

	for _, id := range ids {
		from := getter.GetString(prefix+id+".from", "")
		to := getter.GetString(prefix+id+".to", "")
		item := &assemblyLineCopyFile{}
		item.from = deffs.GetPath(from)
		item.to = deffs.GetPath(to)
		list = append(list, item)
	}

	err := getter.Error()
	if err != nil {
		return err
	}
	inst.sourceFiles = list
	return nil
}

func (inst *assemblyLineTask) doCopyFilesFromSource() error {
	list := inst.sourceFiles
	for _, item := range list {
		item.parent = inst
		err := item.copy()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *assemblyLineTask) exec(c *exec.Cmd) error {

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Start()
	if err != nil {
		return err
	}

	err = c.Wait()
	if err != nil {
		return err
	}

	code := c.ProcessState.ExitCode()
	if code != 0 {
		return errors.New(fmt.Sprint("exit with code ", code))
	}

	return nil
}

// doCopy2 from packaging to dist
func (inst *assemblyLineTask) doCopyProductsToDist() error {

	list := make([]*assemblyLineCopyFile, 0)

	item := &assemblyLineCopyFile{}
	item.parent = inst
	item.from = inst.packagingOutputBpmFile
	item.to = inst.distributionBpmFile
	list = append(list, item)

	item = &assemblyLineCopyFile{}
	item.parent = inst
	item.from = inst.packagingOutputBpmPropsFile
	item.to = inst.distributionBpmPropsFile
	list = append(list, item)

	for _, item := range list {
		err := item.copy()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *assemblyLineTask) updatePackageList() error {

	src := inst.distributionBpmPropsFile
	dst := inst.distributionSourceFile

	srcProps, err := inst.loadProperties(src)
	if err != nil {
		return err
	}

	dstProps, err := inst.loadProperties(dst)
	if err != nil {
		return err
	}

	dstProps.Import(srcProps.Export(nil))
	trimmer := distPackageListTrimmer{}
	dstProps = trimmer.trim(dstProps)

	text := collection.FormatPropertiesWithSegment(dstProps)
	return dst.GetIO().WriteText(text, nil, false)
}
