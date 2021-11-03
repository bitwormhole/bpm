package service

import (
	"context"
	"errors"

	"github.com/bitwormhole/bpm/data/convert"
	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/po"
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/tools"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// DeployPackageFilter 在部署时过滤安装包
type DeployPackageFilter interface {
	AcceptDeploy(prev *entity.InstalledPackageInfo, next *entity.AvailablePackageInfo) bool
}

// DeployService 部署已缓存的bpm包
type DeployService interface {
	Deploy(ctx context.Context, in *vo.Deploy, out *vo.Deploy) error

	DeployPackage(ctx context.Context, pack *entity.AvailablePackageInfo, filter DeployPackageFilter) error
	DeployPackages(ctx context.Context, packs []*entity.AvailablePackageInfo, filter DeployPackageFilter) error
	DeployByNames(ctx context.Context, names []string, filter DeployPackageFilter) error
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
	return inst.DeployByNames(ctx, in.PackageNames, nil)
}

// DeployPackage ...
func (inst *DeployServiceImpl) DeployPackage(ctx context.Context, pack *entity.AvailablePackageInfo, filter DeployPackageFilter) error {

	console, err := cli.GetConsole(ctx)
	if err != nil {
		return err
	}

	task := deployServiceTask{}
	task.parent = inst
	task.context = ctx
	task.console = console
	task.pack81 = pack
	task.filter = filter

	return task.run()
}

// DeployPackages ...
func (inst *DeployServiceImpl) DeployPackages(ctx context.Context, packs []*entity.AvailablePackageInfo, filter DeployPackageFilter) error {
	for _, pack := range packs {
		err := inst.DeployPackage(ctx, pack, filter)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

// DeployByNames ...
func (inst *DeployServiceImpl) DeployByNames(ctx context.Context, names []string, filter DeployPackageFilter) error {
	packs, err := inst.PM.SelectAvailablePackages(names)
	if err != nil {
		return err
	}
	return inst.DeployPackages(ctx, packs, filter)
}

////////////////////////////////////////////////////////////////////////////////

type deployServiceTask struct {
	parent  *DeployServiceImpl
	context context.Context
	console cli.Console
	filter  DeployPackageFilter

	pack80 *entity.InstalledPackageInfo // 已安装的包
	pack81 *entity.AvailablePackageInfo // 将要安装的包
	pack82 *entity.InstalledPackageInfo // 完成安装的包

	bpmFiles *LocalBpmFiles
	manifest po.Manifest

	tmpdir           fs.Path
	tmpDotBpmDir     fs.Path
	tmpSignatureFile fs.Path
	tmpManifestFile  fs.Path

	err error
}

func (inst *deployServiceTask) run() error {

	err := inst.init()
	if err != nil {
		return err
	}

	inst.loadPack80()
	inst.logTodo()

	// 检查是否需要skip当前的部署
	accepted := inst.filter.AcceptDeploy(inst.pack80, inst.pack81)
	if !accepted {
		inst.console.WriteString(" ... skip deploying.\n")
		return nil
	}

	err = inst.prepareDirs()
	if err != nil {
		return err
	}
	defer inst.clearTempDir()

	err = inst.preparePack82()
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

	err = inst.saveManifestAndSignature()
	if err != nil {
		return err
	}

	err = inst.saveMetaToInstalled()
	if err != nil {
		return err
	}

	// @defer	inst.clearTempDir()
	return inst.done()
}

func (inst *deployServiceTask) init() error {
	if inst.filter == nil {
		inst.filter = inst
	}
	return nil
}

func (inst *deployServiceTask) done() error {
	if inst.err == nil {
		msg := " ... deploying is success."
		inst.console.WriteString(msg + "\n")
	}
	return inst.err
}

func (inst *deployServiceTask) logTodo() {
	pkg := inst.pack81
	name := pkg.Name
	version := pkg.Version
	msg := "deploy package: " + name + "@" + version
	inst.console.WriteString(msg + " ...\n")
}

func (inst *deployServiceTask) AcceptDeploy(installed *entity.InstalledPackageInfo, available *entity.AvailablePackageInfo) bool {
	return true
}

func (inst *deployServiceTask) loadPack80() error {
	pack1 := inst.pack81
	name := pack1.Name
	namelist := []string{name}
	packs, err := inst.parent.PM.SelectInstalledPackages(namelist)
	if err == nil {
		for _, item := range packs {
			inst.pack80 = item
		}
	}
	return nil
}

func (inst *deployServiceTask) preparePack82() error {

	pack1 := inst.pack81
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

	inst.pack82 = pack2
	return nil
}

func (inst *deployServiceTask) prepareDirs() error {

	files := inst.parent.Env.GetLocalBpmFiles(&inst.pack81.BasePackageInfo)
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
	return tools.Unzip(from, to)
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

// todo...
func (inst *deployServiceTask) saveManifestAndSignature() error {

	// to read
	manifest1 := inst.tmpManifestFile
	manifestData, err := manifest1.GetIO().ReadBinary(nil)
	if err != nil {
		return err
	}

	sign1 := inst.tmpSignatureFile
	signData, err := sign1.GetIO().ReadBinary(nil)
	if err != nil {
		return err
	}

	// to compute
	sum := tools.ComputeSHA256sumForBytes(manifestData)

	// to write
	manifest2 := inst.bpmFiles.Manifest
	manifest3 := inst.parent.Env.GetManifestDir().GetChild(sum)
	sign2 := inst.bpmFiles.Signature
	sign3 := inst.parent.Env.GetSignatureDir().GetChild(sum)

	err = manifest2.GetIO().WriteBinary(manifestData, nil, true)
	if err != nil {
		return err
	}

	err = manifest3.GetIO().WriteBinary(manifestData, nil, true)
	if err != nil {
		return err
	}

	err = sign2.GetIO().WriteBinary(signData, nil, true)
	if err != nil {
		return err
	}

	err = sign3.GetIO().WriteBinary(signData, nil, true)
	if err != nil {
		return err
	}

	return nil
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
			// if node.IsDir() {
			// 	continue
			// } else {
			// 	return errors.New("dir not found, path=" + item.Path)
			// }
			if !node.Exists() {
				node.Mkdirs()
			}
		} else {
			// for file
			wantSum := item.SHA256
			wantSize := item.Size

			haveSum, err := tools.ComputeSHA256sum(node)
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
		haveSum, err := tools.ComputeSHA256sum(dstNode)
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
	pack0 := inst.pack80
	pack2 := inst.pack82
	manifest := &inst.manifest

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

	// value of newer
	if pack0 != nil {
		pack2.AutoUpgrade = pack0.AutoUpgrade
	}
	pack2.Main = manifest.Meta.Main

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
