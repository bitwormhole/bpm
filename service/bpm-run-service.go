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
)

type RunService interface {
	Run(ctx context.Context, in *vo.Run, out *vo.Run) error
}

type RunServiceImpl struct {
	markup.Component `id:"bpm-run-service" class:"bpm-service"`

	PM  PackageManager `inject:"#bpm-package-manager"`
	Env EnvService     `inject:"#bpm-env-service"`
}

func (inst *RunServiceImpl) _Impl() RunService {
	return inst
}

func (inst *RunServiceImpl) Run(ctx context.Context, in *vo.Run, out *vo.Run) error {

	console, err := cli.GetConsole(ctx)
	if err != nil {
		return err
	}

	task := myRunServiceTask{}
	task.console = console
	task.context = ctx
	task.parent = inst
	task.packageName = in.PackageName
	task.scriptName = in.ScriptName
	return task.run()
}

////////////////////////////////////////////////////////////////////////////////

type myRunServiceTask struct {
	parent  *RunServiceImpl
	context context.Context
	console cli.Console

	theBitwormholeHome fs.Path

	mainConfigFile    fs.Path
	mainConfigProps   collection.Properties
	mainConfigData    []byte
	mainConfigSumWant string
	mainConfigSumHave string

	manifestFile    fs.Path
	manifestProps   collection.Properties
	manifestData    []byte
	manifestSumWant string
	manifestSumHave string
	manifestModel   po.Manifest

	exeFile    fs.Path
	exeSumWant string
	exeSumHave string

	scriptName  string
	packageName string
	packInfo    entity.InstalledPackageInfo
}

func (inst *myRunServiceTask) run() error {

	err := inst.init()
	if err != nil {
		return err
	}

	err = inst.loadPackInfo()
	if err != nil {
		return err
	}

	err = inst.loadManifest()
	if err != nil {
		return err
	}

	err = inst.loadMainConfig()
	if err != nil {
		return err
	}

	err = inst.loadExecutableInfo()
	if err != nil {
		return err
	}

	return nil
}

func (inst *myRunServiceTask) init() error {
	home := inst.parent.Env.GetBitwormholeHome()
	inst.theBitwormholeHome = home
	return nil
}

func (inst *myRunServiceTask) lookUpForExeFile() (fs.Path, error) {

	//todo
	return nil, nil
}

func (inst *myRunServiceTask) lookUpForManifestFile() (fs.Path, error) {
	files := inst.parent.Env.GetLocalBpmFiles(&inst.packInfo.BasePackageInfo)
	file := files.Manifest
	return file, nil
}

func (inst *myRunServiceTask) lookUpForMainConfigFile() (fs.Path, error) {
	main := inst.manifestModel.Meta.Main
	home := inst.theBitwormholeHome
	file := home.GetChild(main)
	return file, nil
}

func (inst *myRunServiceTask) loadExecutableInfo() error {

	file, err := inst.lookUpForExeFile()
	if err != nil {
		return err
	}

	sum, err := tools.ComputeSHA256sum(file)
	if err != nil {
		return err
	}

	inst.exeSumHave = sum
	inst.exeFile = file
	return nil
}

func (inst *myRunServiceTask) loadManifest() error {

	file, err := inst.lookUpForManifestFile()
	if err != nil {
		return err
	}

	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return err
	}

	text := string(data)
	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	err = convert.LoadPackageManifest(&inst.manifestModel, props)
	if err != nil {
		return err
	}

	sum := tools.ComputeSHA256sumForBytes(data)

	inst.manifestFile = file
	inst.manifestData = data
	inst.manifestProps = props
	inst.manifestSumHave = sum
	return nil
}

func (inst *myRunServiceTask) loadMainConfig() error {

	file, err := inst.lookUpForMainConfigFile()
	if err != nil {
		return err
	}

	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return err
	}

	text := string(data)
	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	sum := tools.ComputeSHA256sumForBytes(data)

	// todo : convert.load

	inst.mainConfigData = data
	inst.mainConfigProps = props
	inst.mainConfigFile = file
	inst.mainConfigSumHave = sum
	return nil
}

func (inst *myRunServiceTask) loadPackInfo() error {

	pkgName := inst.packageName
	list := []string{pkgName}

	packs, err := inst.parent.PM.SelectInstalledPackages(list)
	if err != nil {
		return err
	}

	for _, pack := range packs {
		inst.packInfo = *pack
		return nil
	}

	return errors.New("no installed package named: " + pkgName)
}

////////////////////////////////////////////////////////////////////////////////
