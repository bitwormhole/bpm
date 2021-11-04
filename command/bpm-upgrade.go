package command

import (
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type BpmUpgrade struct {
	markup.Component `class:"cli-handler"`

	Service service.UpgradeService `inject:"#bpm-upgrade-service"`
}

func (inst *BpmUpgrade) _Impl() cli.Handler {
	return inst
}

func (inst *BpmUpgrade) Init(service cli.Service) error {
	service.RegisterHandler("bpm-upgrade", inst)
	return nil
}

func (inst *BpmUpgrade) Handle(ctx *cli.TaskContext) error {

	args, err := parseArguments(ctx.CurrentTask.Arguments)
	if err != nil {
		return err
	}

	in := vo.Upgrade{}
	out := vo.Upgrade{}
	in.PackageNames = args.Packages
	return inst.Service.Upgrade(ctx.Context, &in, &out)
}
