package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bitwormhole/bpm"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter-cli/terminal"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

// bpm-installer 是一个简单的命令，用来安装（或升级）BPM
func main() {

	mod := bpm.Module()

	props := collection.CreateProperties()
	props.SetProperty("application.module", mod.GetName())
	props.SetProperty("application.version", mod.GetVersion())
	props.SetProperty("application.revision", strconv.Itoa(mod.GetRevision()))

	fmt.Println("BPM Installer " + mod.GetVersion())

	i := starter.InitApp()
	i.UseMain(mod)
	i.UseProperties(props)

	r := bpmInstaller{}
	r.Init(i)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type bpmInstaller struct {
	args []string
	i    application.Initializer
}

func (inst *bpmInstaller) Init(i application.Initializer) error {
	inst.args = os.Args
	inst.i = i
	return nil
}

func (inst *bpmInstaller) Run() error {

	rt, err := inst.i.RunEx()
	if err != nil {
		return err
	}

	appctx := rt.Context()
	ctx, err := terminal.Prepare(appctx)
	if err != nil {
		return err
	}

	o1, err := appctx.GetComponent("#cli-client-factory")
	if err != nil {
		return err
	}
	o2 := o1.(cli.ClientFactory)
	client := o2.CreateClient(ctx)

	if inst.argsContains("install") {
		// bpm install
		err := inst.doBpmInstall(client)
		if err != nil {
			return err
		}
	} else if inst.argsContains("update") || inst.argsContains("upgrade") {
		// bpm upgrade
		err := inst.doBpmUpgrade(client)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("usage: bpm-installer  [install|update|upgrade]")
	}

	return rt.Exit()
}

func (inst *bpmInstaller) doBpmUpdate(client cli.Client) error {
	return client.Execute("bpm update", nil)
}

func (inst *bpmInstaller) doBpmInstall(client cli.Client) error {

	err := inst.doBpmUpdate(client)
	if err != nil {
		return err
	}

	return client.Execute("bpm install bpm", nil)
}

func (inst *bpmInstaller) doBpmUpgrade(client cli.Client) error {

	err := inst.doBpmUpdate(client)
	if err != nil {
		return err
	}

	return client.Execute("bpm upgrade bpm", nil)
}

func (inst *bpmInstaller) argsContains(text string) bool {
	args := inst.args
	for _, item := range args {
		item = strings.TrimSpace(item)
		if text == item {
			return true
		}
	}
	return false
}
