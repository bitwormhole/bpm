package command

import (
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type BpmPackInfo struct {
	markup.Component `class:"cli-handler"`

	Service service.PackInfoService `inject:"#bpm-pack-info-service"`
}

func (inst *BpmPackInfo) _Impl() cli.Handler {
	return inst
}

func (inst *BpmPackInfo) Init(service cli.Service) error {
	service.RegisterHandler("bpm-pack-info", inst)
	return nil
}

func (inst *BpmPackInfo) Handle(ctx *cli.TaskContext) error {

	a2, err := parseArguments(ctx.CurrentTask.Arguments)
	if err != nil {
		return err
	}

	in := vo.PackInfo{}
	out := vo.PackInfo{}
	in.PackageNames = a2.Packages
	return inst.Service.DisplayPackInfo(ctx.Context, &in, &out)
}
