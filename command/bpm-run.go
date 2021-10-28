package command

import (
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type BpmRun struct {
	markup.Component `class:"cli-handler"`

	Service service.RunService `inject:"#bpm-run-service"`
}

func (inst *BpmRun) _Impl() cli.Handler {
	return inst
}

func (inst *BpmRun) Init(service cli.Service) error {
	service.RegisterHandler("bpm-run", inst)
	return nil
}

func (inst *BpmRun) Handle(ctx *cli.TaskContext) error {

	in := vo.Run{}
	out := vo.Run{}
	in.PackageName = ""
	in.Arguments = nil
	return inst.Service.Run(ctx.Context, &in, &out)
}
