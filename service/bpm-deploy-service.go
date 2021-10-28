package service

import (
	"context"
	"errors"

	"github.com/bitwormhole/bpm/data/convert"
	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/po"
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// DeployService 部署已缓存的bpm包
type DeployService interface {
	Deploy(ctx context.Context, in *vo.Deploy, out *vo.Deploy) error

	DeployPackage(ctx context.Context, pack *entity.AvailablePackageInfo) error
	DeployPackages(ctx context.Context, packs []*entity.AvailablePackageInfo) error
	DeployByNames(ctx context.Context, names []string) error
}

////////////////////////////////////////////////////////////////////////////////

// DeployServiceImpl ...
type DeployServiceImpl struct {
	markup.Component `id:"bpm-deploy-service"`

	PM  PackageManager `inject:"#bpm-package-manager"`
	Env EnvService     `inject:"#bpm-env-service"`
}

// Deploy ...
func (inst *DeployServiceImpl) Deploy(ctx context.Context, in *vo.Deploy, out *vo.Deploy) error {
	return inst.DeployByNames(ctx, in.PackageNames)
}

// DeployPackage ...
func (inst *DeployServiceImpl) DeployPackage(ctx context.Context, pack *entity.AvailablePackageInfo) error {

	console, err := cli.GetConsole(ctx)
	if err != nil {
		return err
	}

	task := deployServiceTask{}
	task.parent = inst
	task.context = ctx
	task.console = console
	task.pack1 = pack

	return task.run()
}

// DeployPackages ...
func (inst *DeployServiceImpl) DeployPackages(ctx context.Context, packs []*entity.AvailablePackageInfo) error {
	for _, pack := range packs {
		err := inst.DeployPackage(ctx, pack)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

// DeployByNames ...
func (inst *DeployServiceImpl) DeployByNames(ctx context.Context, names []string) error {
	packs, err := inst.PM.SelectAvailablePackages(names)
	if err != nil {
		return err
	}
	return inst.DeployPackages(ctx, packs)
}

////////////////////////////////////////////////////////////////////////////////

type deployServiceTask struct {
	parent  *DeployServiceImpl
	context context.Context
	console cli.Console

	pack1 *entity.AvailablePackageInfo
	pack2 *entity.InstalledPackageInfo

	bpmFiles *LocalBpmFiles
	manifest po.Manifest

	tmpdir           fs.Path
	tmpDotBpmDir     fs.Path
	tmpSignatureFile fs.Path
	tmpManifestFile  fs.Path
}

func (inst *deployServiceTask) run() error {

	name := inst.pack1.Name

	if inst.hasInstalled() {
		inst.console.WriteString("package [" + name + "] is installed.\n")
		return nil
	}

	err := inst.prepareDirs()
	if err != nil {
		return err
	}
	defer inst.clearTempDir()

	err = inst.preparePack2()
	if err != nil {
		return err
	}

	err = inst.unzipToTempDir()
	if err != nil {
		return err
	}

	err = inst.loadManifestAndSignature()
	if err != nil {
		return err
	}

	err = inst.checkFilesInPackage()
	if err != nil {
		return err
	}

	err = inst.moveFilesToDestination()
	if err != nil {
		return err
	}

	err = inst.moveManifestAndSignature()
	if err != nil {
		return err
	}

	err = inst.saveMetaToInstalled()
	if err != nil {
		return err
	}

	// 	inst.clearTempDir() @defer
	return nil
}

func (inst *deployServiceTask) hasInstalled() bool {
	name := inst.pack1.Name
	namelist := []string{name}
	packs, err := inst.parent.PM.SelectInstalledPackages(namelist)
	if err == nil {
		if len(packs) > 0 {
			return true
		}
	}
	return false
}

func (inst *deployServiceTask) preparePack2() error {

	pack1 := inst.pack1
	pack2 := &entity.InstalledPackageInfo{}

	pack2.Name = pack1.Name
	pack2.Version = pack1.Version
	pack2.Revision = pack1.Revision
	pack2.SHA256 = pack1.SHA256
	pack2.URL = pack1.URL
	pack2.Type = pack1.Type
	pack2.Size = pack1.Size
	pack2.Dependencies = pack1.Dependencies
	pack2.File = inst.bpmFiles.BPM.Path()
	pack2.AutoUpgrade = false

	inst.pack2 = pack2
	return nil
}

func (inst *deployServiceTask) prepareDirs() error {

	files := inst.parent.Env.GetLocalBpmFiles(&inst.pack1.BasePackageInfo)
	tmpDir := files.Dir.GetChild(".unzip.tmp.d")
	dotbpm := tmpDir.GetChild(".bpm")
	manifest := dotbpm.GetChild("manifest")
	signature := dotbpm.GetChild("signature")

	inst.tmpdir = tmpDir
	inst.bpmFiles = files
	inst.tmpDotBpmDir = dotbpm
	inst.tmpManifestFile = manifest
	inst.tmpSignatureFile = signature

	return nil
}

func (inst *deployServiceTask) unzipToTempDir() error {
	from := inst.bpmFiles.BPM
	to := inst.tmpdir
	if !to.Exists() {
		to.Mkdirs()
	}
	return convert.Unzip(from, to)
}

func (inst *deployServiceTask) loadPropertiesFile(file fs.Path) (string, collection.Properties, error) {

	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return "", nil, err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return "", nil, err
	}

	return text, props, nil

}

func (inst *deployServiceTask) moveManifestAndSignature() error {

	src1 := inst.tmpManifestFile
	dst1 := inst.bpmFiles.Manifest
	err := src1.MoveTo(dst1)
	if err != nil {
		return err
	}

	src2 := inst.tmpSignatureFile
	dst2 := inst.bpmFiles.Signature
	return src2.MoveTo(dst2)
}

func (inst *deployServiceTask) loadManifestAndSignature() error {

	signText, signProps, err := inst.loadPropertiesFile(inst.tmpSignatureFile)
	if err != nil {
		return err
	}

	maniText, maniProps, err := inst.loadPropertiesFile(inst.tmpManifestFile)
	if err != nil {
		return err
	}

	if vlog.Default().IsDebugEnabled() {
		vlog.Debug(maniText)
		vlog.Debug(signText)
		vlog.Debug(signProps)
	}

	maniPO := &inst.manifest
	err = convert.LoadPackageManifest(maniPO, maniProps)
	if err != nil {
		return err
	}

	return nil
}

func (inst *deployServiceTask) checkFilesInPackage() error {
	basedir := inst.tmpDotBpmDir.GetChild("files")
	list := inst.manifest.Items
	for _, item := range list {
		node := basedir.GetChild(item.Path)

		// for dir
		if item.IsDir {
			if node.IsDir() {
				continue
			} else {
				return errors.New("dir not found, path=" + item.Path)
			}
		}

		// for file
		wantSum := item.SHA256
		wantSize := item.Size

		haveSum, err := inst.parent.PM.ComputeSHA256sum(node)
		if err != nil {
			return err
		}

		if wantSize != node.Size() {
			return errors.New("bad file size, path=" + item.Path)
		}

		if wantSum != haveSum {
			return errors.New("bad file sha256sum, path=" + item.Path)
		}
	}
	return nil
}

func (inst *deployServiceTask) moveFilesToDestination() error {

	srcdir := inst.tmpDotBpmDir.GetChild("files")
	dstdir := inst.parent.Env.GetBitwormholeHome()
	list := inst.manifest.Items

	for _, item := range list {
		srcNode := srcdir.GetChild(item.Path)
		dstNode := dstdir.GetChild(item.Path)

		// for dir
		if item.IsDir {
			dstNode.Mkdirs()
			continue
		}

		// for file
		parentdir := dstNode.Parent()
		if !parentdir.Exists() {
			parentdir.Mkdirs()
		}

		err := srcNode.MoveTo(dstNode)
		if err != nil {
			return err
		}

		// check
		haveSum, err := inst.parent.PM.ComputeSHA256sum(dstNode)
		if err != nil {
			return err
		}
		wantSum := item.SHA256
		if wantSum != haveSum {
			return errors.New("bad sha256sum, path=" + dstNode.Path())
		}
	}
	return nil
}

func (inst *deployServiceTask) saveMetaToInstalled() error {

	pm := inst.parent.PM
	nextlist := make([]*entity.InstalledPackageInfo, 0)
	pack2 := inst.pack2

	// load prev
	prev, err := pm.LoadInstalledPackages()
	if err != nil {
		// 可能是文件  “installed” 还未创建
		prev = &po.InstalledPackages{}
	}

	// remove older
	for _, item := range prev.Packages {
		if item.Name != pack2.Name {
			nextlist = append(nextlist, item)
		}
	}

	// append newer
	nextlist = append(nextlist, pack2)

	// save back
	next := &po.InstalledPackages{}
	next.Packages = nextlist
	return pm.SaveInstalledPackages(next)
}

func (inst *deployServiceTask) clearTempDir() error {
	tmpdir := inst.tmpdir
	if tmpdir.IsDir() {
		return clearDir(tmpdir, 99)
	}
	return nil
}

func clearDir(path fs.Path, depthLimit int) error {
	if depthLimit < 0 {
		return errors.New("the path is too deep, path=" + path.Path())
	}
	if path.IsFile() {
		return path.Delete()
	} else if path.IsDir() {
		// continue
	} else {
		return nil
	}
	items := path.ListItems()
	for _, item := range items {
		err := clearDir(item, depthLimit-1)
		if err != nil {
			return err
		}
	}
	return path.Delete()
}
