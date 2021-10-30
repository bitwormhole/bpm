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

	// bpm-run [package] [script]

	args := ctx.CurrentTask.Arguments

	in := vo.Run{}
	out := vo.Run{}
	in.PackageName = inst.getArgAt(1, args)
	in.ScriptName = inst.getArgAt(2, args)
	in.Arguments = args
	return inst.Service.Run(ctx.Context, &in, &out)
}

func (inst *BpmRun) getArgAt(index int, args []string) string {
	if index < len(args) {
		return args[index]
	}
	return ""
}
