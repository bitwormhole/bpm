package command

import (
	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type BpmAssemblyLine struct {
	markup.Component `class:"cli-handler"`

	Service service.AssemblyLineService `inject:"#bpm-assembly-line-service"`
}

func (inst *BpmAssemblyLine) _Impl() cli.Handler {
	return inst
}

func (inst *BpmAssemblyLine) Init(service cli.Service) error {
	service.RegisterHandler("bpm-assembly", inst)
	return nil
}

func (inst *BpmAssemblyLine) Handle(ctx *cli.TaskContext) error {

	// in := vo.Upgrade{}
	// out := vo.Upgrade{}

	args := ctx.CurrentTask.Arguments[1:]

	if len(args) == 0 {
		ctx.Console.WriteString("usage: bpm assembly {targets...}\n")
		return nil
	}

	return inst.Service.Assembly(ctx.Context, args)
}
