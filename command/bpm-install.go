package command

import (
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type BpmInstall struct {
	markup.Component `class:"cli-handler"`

	Service service.InstallService `inject:"#bpm-install-service"`
}

func (inst *BpmInstall) _Impl() cli.Handler {
	return inst
}

func (inst *BpmInstall) Init(service cli.Service) error {
	service.RegisterHandler("bpm-install", inst)
	return nil
}

func (inst *BpmInstall) Handle(ctx *cli.TaskContext) error {

	a2, err := parseArguments(ctx.CurrentTask.Arguments)
	if err != nil {
		return err
	}

	in := vo.Install{}
	out := vo.Install{}
	in.PackageNames = a2.Packages
	return inst.Service.Install(ctx.Context, &in, &out)
}
