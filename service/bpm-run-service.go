package service

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/bitwormhole/bpm/data/convert"
	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/po"
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/tools"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

type RunService interface {
	Run(ctx context.Context, in *vo.Run, out *vo.Run) error
}

type RunServiceImpl struct {
	markup.Component `id:"bpm-run-service" class:"bpm-service"`

	PM      PackageManager      `inject:"#bpm-package-manager"`
	Env     EnvService          `inject:"#bpm-env-service"`
	Context application.Context `inject:"context"`
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

	scriptName  string
	packageName string

	theBitwormholeHome fs.Path

	mainConfigFile    fs.Path
	mainConfigProps   collection.Properties
	mainConfigData    []byte
	mainConfigSumWant string
	mainConfigSumHave string
	mainConfigModel   po.AppMain

	manifestFile    fs.Path
	manifestProps   collection.Properties
	manifestData    []byte
	manifestSumWant string
	manifestSumHave string
	manifestModel   po.Manifest

	exeFile       fs.Path
	exeWorkingDir fs.Path
	exeSumWant    string
	exeSumHave    string

	targetScript *entity.MainScript
	packInfo     entity.InstalledPackageInfo
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

	err = inst.selectTargetScript()
	if err != nil {
		return err
	}

	err = inst.loadExecutableInfo()
	if err != nil {
		return err
	}

	err = inst.checkFiles()
	if err != nil {
		return err
	}

	inst.printScriptParams()

	return inst.execute()
}

func (inst *myRunServiceTask) init() error {
	home := inst.parent.Env.GetBitwormholeHome()
	inst.theBitwormholeHome = home
	return nil
}

func (inst *myRunServiceTask) lookUpForWorkingDir() (fs.Path, error) {
	selector := inst.targetScript
	dir := fs.Default().GetPath(selector.WorkingDirectory)
	return dir, nil
}

func (inst *myRunServiceTask) lookUpForExeFile() (fs.Path, error) {
	selector := inst.targetScript
	file := fs.Default().GetPath(selector.Executable)
	return file, nil
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

	wd, err := inst.lookUpForWorkingDir()
	if err != nil {
		return err
	}

	sum, err := tools.ComputeSHA256sum(file)
	if err != nil {
		return err
	}

	inst.exeSumHave = sum
	inst.exeFile = file
	inst.exeWorkingDir = wd
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

func (inst *myRunServiceTask) selectTargetScript() error {

	selector := strings.TrimSpace(inst.scriptName)
	model := &inst.mainConfigModel
	scripts := model.Scripts
	scriptPtr := inst.targetScript

	if selector == "" {
		selector = model.Main.Script // use the default script
	}

	for _, script := range scripts {
		if script.Name == selector {
			scriptPtr = script
			break
		}
	}

	if scriptPtr == nil {
		return errors.New("no script named: " + selector)
	}

	inst.targetScript = scriptPtr
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

	inst.injectEnvToProperties(props)
	err = tools.ResolveConfig(props)
	if err != nil {
		return err
	}

	err = convert.LoadMainConfig(&inst.mainConfigModel, props)
	if err != nil {
		return err
	}

	inst.mainConfigData = data
	inst.mainConfigProps = props
	inst.mainConfigFile = file
	inst.mainConfigSumHave = sum
	return nil
}

func (inst *myRunServiceTask) injectEnvToProperties(props collection.Properties) {
	env := inst.parent.Context.GetEnvironment()
	tools.InjectEnvToProperties(props, env)
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

func (inst *myRunServiceTask) checkFiles() error {

	manifest := string(inst.manifestData)
	executableSum := inst.exeSumHave
	mainConfigSum := inst.mainConfigSumHave
	manifestSum := inst.manifestSumHave

	if !strings.Contains(manifest, executableSum) {
		return errors.New("executable file is not in manifest")
	}
	if !strings.Contains(manifest, mainConfigSum) {
		return errors.New("main-config file is not in manifest")
	}

	vlog.Info("check ", manifestSum, ": manifest ... ok")
	vlog.Info("check ", mainConfigSum, ": init ... ok")
	vlog.Info("check ", executableSum, ": "+inst.exeFile.Name()+" ... ok")

	// todo ... check signature

	return nil
}

func (inst *myRunServiceTask) printScriptParams() {

	wd := inst.exeWorkingDir
	exe := inst.exeFile
	args := inst.targetScript.Arguments
	console := inst.console

	console.WriteString("\nscript.arguments=" + args)
	console.WriteString("\nscript.working-directory=" + wd.Path())
	console.WriteString("\nscript.executable=" + exe.Path())
	console.WriteString("\n\n")
}

func (inst *myRunServiceTask) execute() error {

	console := inst.console
	wd := inst.exeWorkingDir
	exefile := inst.exeFile

	parser := cli.CommandLineParser{}
	args, err := parser.Parse(inst.targetScript.Arguments)
	if err != nil {
		return err
	}

	cmd := exec.Command(exefile.Path())
	cmd.Dir = wd.Path()
	cmd.Args = args
	cmd.Stdout = console.Output()
	cmd.Stderr = console.Error()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	state := cmd.ProcessState
	code := state.ExitCode()
	os.Exit(code)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
