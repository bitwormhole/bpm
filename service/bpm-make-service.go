package service

import (
	"context"
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/bitwormhole/bpm/data/convert"
	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/po"
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/tools"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
)

// MakeService ...
type MakeService interface {
	Make(ctx context.Context, in *vo.Make, out *vo.Make) error
	MakePackage(ctx context.Context, pwd fs.Path) error
}

////////////////////////////////////////////////////////////////////////////////

// MakeServiceImpl ...
type MakeServiceImpl struct {
	markup.Component `id:"bpm-make-service" class:"bpm-service"`

	// Fetch  FetchService  `inject:"#bpm-fetch-service"`
	// Deploy DeployService `inject:"#bpm-deploy-service"`
}

func (inst *MakeServiceImpl) _Impl() MakeService {
	return inst
}

// Make 生成安装包
func (inst *MakeServiceImpl) Make(ctx context.Context, in *vo.Make, out *vo.Make) error {
	return inst.MakePackage(ctx, nil)
}

// MakePackage 生成安装包
func (inst *MakeServiceImpl) MakePackage(ctx context.Context, pwd fs.Path) error {

	console, err := cli.GetConsole(ctx)
	if err != nil {
		return err
	}

	if pwd == nil {
		pwd = console.GetWorkingPath()
	}

	task := myMakeServiceTask{}
	task.context = ctx
	task.console = console
	task.parent = inst
	task.pwd = pwd
	return task.run()
}

////////////////////////////////////////////////////////////////////////////////

type myMakeServiceTask struct {
	// contexts
	context context.Context
	console cli.Console
	parent  *MakeServiceImpl

	// files
	pwd              fs.Path
	theDotBpmDir     fs.Path
	theBpmConfigFile fs.Path
	theFilesDir      fs.Path
	theManifestFile  fs.Path
	theFilelistFile  fs.Path
	theSignatureFile fs.Path
	theZipFile       fs.Path
	thePackInfoFile  fs.Path

	// data
	timestamp     int64
	itemIndexer   int
	packInfo      po.AvailablePackages
	signature     po.Signature
	manifest      po.Manifest // 生成的清单
	config        po.Manifest // 配置的清单
	overrideTable map[string]bool
}

func (inst *myMakeServiceTask) run() error {

	err := inst.init()
	if err != nil {
		return err
	}

	err = inst.scanFiles()
	if err != nil {
		return err
	}

	err = inst.makeManifest()
	if err != nil {
		return err
	}

	err = inst.makeSignature()
	if err != nil {
		return err
	}

	err = inst.makeFileList()
	if err != nil {
		return err
	}

	err = inst.zip()
	if err != nil {
		return err
	}

	err = inst.makeZipMeta()
	if err != nil {
		return err
	}

	inst.logDone()
	return nil
}

func (inst *myMakeServiceTask) init() error {

	inst.timestamp = util.CurrentTimestamp()
	inst.itemIndexer = 100000

	dotbpm, err := inst.findDotBpm()
	if err != nil {
		return err
	}

	err = inst.initFiles(dotbpm)
	if err != nil {
		return err
	}

	return nil
}

func (inst *myMakeServiceTask) initFiles(dotBpm fs.Path) error {

	pdir := dotBpm.Parent() // parent-dir

	// input
	inst.theDotBpmDir = dotBpm
	inst.theFilesDir = dotBpm.GetChild("files")
	inst.theBpmConfigFile = pdir.GetChild(".bpm.config")

	// load config
	err := inst.loadConfig()
	if err != nil {
		return err
	}
	zipFileName := inst.makeOutputZipFileName()

	// output
	inst.theFilelistFile = dotBpm.GetChild("filelist")
	inst.theManifestFile = dotBpm.GetChild("manifest")
	inst.theSignatureFile = dotBpm.GetChild("signature")
	inst.theZipFile = pdir.GetChild("dist/" + zipFileName)
	inst.thePackInfoFile = pdir.GetChild("dist/" + zipFileName + ".properties")

	// log
	inst.logPath("pwd", inst.pwd)
	inst.logPath(".bpm", inst.theDotBpmDir)
	inst.logPath("files", inst.theFilesDir)
	inst.logPath("manifest", inst.theManifestFile)
	inst.logPath("signature", inst.theSignatureFile)
	inst.logPath("config", inst.theBpmConfigFile)
	inst.logPath("zip", inst.theZipFile)
	inst.logPath("packinfo", inst.thePackInfoFile)

	return nil
}

func (inst *myMakeServiceTask) makeOutputZipFileName() string {

	// meta := &inst.config.Meta
	// name := meta.Name
	// version := meta.Version
	// platform := meta.Platform

	// builder := strings.Builder{}
	// builder.WriteString(name)
	// builder.WriteString("-")
	// builder.WriteString(version)
	// builder.WriteString("-")
	// builder.WriteString(platform)
	// builder.WriteString(".bpm.zip")

	return inst.config.Meta.Filename

	// return builder.String()
}

func (inst *myMakeServiceTask) logPath(tag string, path fs.Path) {
	const wantLen = 20
	for {
		curLen := len(tag)
		if curLen < wantLen {
			tag = " " + tag
		} else {
			break
		}
	}
	vlog.Debug(tag, " = ", path.Path())
}

func (inst *myMakeServiceTask) findDotBpm() (fs.Path, error) {
	const target = ".bpm"
	pwd := inst.pwd
	p := pwd
	for ; p != nil; p = p.Parent() {
		want := p.GetChild(target)
		if want.IsDir() {
			p = want
			break
		}
	}
	if p == nil {
		return nil, errors.New("cannot find .bpm dir in the pwd path " + pwd.Path())
	}
	inst.console.WriteString("find [.bpm] dir at " + p.Path() + "\n")
	return p, nil
}

func (inst *myMakeServiceTask) scanFiles() error {
	return inst.scanDir(inst.theFilesDir, "", 0)
}

func (inst *myMakeServiceTask) scanDir(dir fs.Path, shortPath string, depth int) error {

	if !dir.IsDir() {
		return errors.New("the path is not a dir, path=" + dir.Path())
	}

	if depth > 99 {
		return errors.New("the path is too deep, path=" + dir.Path())
	}

	if shortPath != "" {
		if !strings.HasSuffix(shortPath, "/") {
			shortPath = shortPath + "/"
		}
	}

	// for dir self
	err := inst.scanOnDir(dir, shortPath)
	if err != nil {
		return err
	}

	// for items
	names := dir.ListNames()
	sort.Strings(names)
	for _, name := range names {
		child := dir.GetChild(name)
		shortPath2 := shortPath + name
		if child.IsFile() {
			err = inst.scanOnFile(child, shortPath2)
		} else if child.IsDir() {
			err = inst.scanDir(child, shortPath2, depth+1)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myMakeServiceTask) scanOnFile(file fs.Path, shortPath string) error {

	sum, err := tools.ComputeSHA256sum(file)
	if err != nil {
		return err
	}
	item := &entity.ManifestItem{}

	item.ID = inst.nextItemID()
	item.IsDir = false
	item.Name = file.Name()
	item.Path = shortPath
	item.SHA256 = sum
	item.Size = file.Size()
	item.IsOverride = inst.overrideTable[shortPath]

	inst.addManifestItem(item)
	return nil
}

func (inst *myMakeServiceTask) scanOnDir(dir fs.Path, shortPath string) error {
	item := &entity.ManifestItem{}
	item.ID = inst.nextItemID()
	item.IsDir = true
	item.Name = dir.Name()
	item.Path = shortPath
	item.SHA256 = ""
	item.Size = 0
	inst.addManifestItem(item)
	return nil
}

func (inst *myMakeServiceTask) addManifestItem(item *entity.ManifestItem) {
	list := inst.manifest.Items
	list = append(list, item)
	inst.manifest.Items = list
}

func (inst *myMakeServiceTask) nextItemID() string {
	inst.itemIndexer++
	idx := inst.itemIndexer
	return strconv.Itoa(idx)
}

func (inst *myMakeServiceTask) loadConfig() error {
	file := inst.theBpmConfigFile
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	err = tools.ResolveConfig(props)
	if err != nil {
		return err
	}

	err = inst.loadOverrideTable(props)
	if err != nil {
		return err
	}

	return convert.LoadPackageManifest(&inst.config, props)
}

func (inst *myMakeServiceTask) loadOverrideTable(props collection.Properties) error {
	const (
		overYes = "override.yes."
		overNo  = "override.no."
	)
	table := make(map[string]bool)
	all := props.Export(nil)
	for key, value := range all {
		if strings.HasPrefix(key, overYes) {
			table[value] = true
		} else if strings.HasPrefix(key, overNo) {
			table[value] = false
		}
	}
	inst.overrideTable = table
	return nil
}

func (inst *myMakeServiceTask) makeFileList() error {

	files := make([]string, 0)
	dirs := make([]string, 0)
	builder := strings.Builder{}
	manifest := &inst.manifest
	items := manifest.Items

	for _, item := range items {
		if item.IsDir {
			dirs = append(dirs, item.Path)
		} else {
			files = append(files, item.Path)
		}
	}

	sort.Strings(dirs)
	sort.Strings(files)
	var index int = 0

	builder.WriteString("[dirs]\n")
	for _, path := range dirs {
		index++
		builder.WriteString(strconv.Itoa(index))
		builder.WriteString("=")
		builder.WriteString(path)
		builder.WriteString("\n")
	}

	builder.WriteString("[files]\n")
	for _, path := range files {
		index++
		builder.WriteString(strconv.Itoa(index))
		builder.WriteString("=")
		builder.WriteString(path)
		builder.WriteString("\n")
	}

	text := builder.String()
	file := inst.theFilelistFile
	return file.GetIO().WriteText(text, nil, true)
}

func (inst *myMakeServiceTask) makeManifest() error {

	cfg := &inst.config
	manifest := &inst.manifest
	now := inst.timestamp
	nowTime := util.Int64ToTime(now)

	manifest.Meta = cfg.Meta
	manifest.Meta.Date = now
	manifest.Meta.DateString = nowTime.String()

	props := collection.CreateProperties()
	err := convert.SavePackageManifest(manifest, props)
	if err != nil {
		return err
	}
	text := collection.FormatPropertiesWithSegment(props)
	file := inst.theManifestFile
	return file.GetIO().WriteText(text, nil, true)
}

func (inst *myMakeServiceTask) makeSignature() error {

	sign := &inst.signature
	cfg := &inst.config
	manifest := &inst.manifest

	// compute sum of manifest
	sum, err := tools.ComputeSHA256sum(inst.theManifestFile)
	if err != nil {
		return err
	}

	// make
	sign.Info.BasePackageInfo = manifest.Meta.BasePackageInfo
	sign.Info.Algorithm = cfg.Meta.SignatureAlgorithm
	sign.Info.PublicFinger = cfg.Meta.SignaturePublicFinger
	sign.Info.Secret = "todo:99999999999"
	sign.Info.Plain = "sha256sum(manifest):" + sum

	// save
	props := collection.CreateProperties()
	err = convert.SavePackageSignature(sign, props)
	if err != nil {
		return err
	}
	text := collection.FormatPropertiesWithSegment(props)
	file := inst.theSignatureFile
	return file.GetIO().WriteText(text, nil, true)
}

func (inst *myMakeServiceTask) zip() error {
	from := inst.theDotBpmDir
	to := inst.theZipFile
	return tools.Zip(from, to, true)
}

func (inst *myMakeServiceTask) makeZipMeta() error {

	zipfile := inst.theZipFile
	sum, err := tools.ComputeSHA256sum(zipfile)
	if err != nil {
		return err
	}

	// prepare structs
	packInfoItem := &entity.AvailablePackageInfo{}
	packInfoList := &po.AvailablePackages{}
	packInfoList.Packages = []*entity.AvailablePackageInfo{packInfoItem}

	packInfoItem.BasePackageInfo = inst.manifest.Meta.BasePackageInfo
	packInfoItem.SHA256 = sum
	packInfoItem.Size = zipfile.Size()
	packInfoItem.URL = inst.config.Meta.URL
	packInfoItem.Filename = inst.config.Meta.Filename

	// save
	props := collection.CreateProperties()
	err = convert.SaveAvailablePackages(packInfoList, props)
	if err != nil {
		return err
	}
	text := collection.FormatPropertiesWithSegment(props)
	file := inst.thePackInfoFile
	return file.GetIO().WriteText(text, nil, true)
}

func (inst *myMakeServiceTask) logDone() error {
	file := inst.theZipFile
	msg := "The BPM file is created, path=" + file.Path()
	inst.console.WriteString(msg + "\n")
	return nil
}

////////////////////////////////////////////////////////////////////////////////
