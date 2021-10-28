package command

import (
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type BpmUpdate struct {
	markup.Component `class:"cli-handler"`

	Service service.UpdateService `inject:"#bpm-update-service"`
}

func (inst *BpmUpdate) _Impl() cli.Handler {
	return inst
}

func (inst *BpmUpdate) Init(service cli.Service) error {
	service.RegisterHandler("bpm-update", inst)
	return nil
}

func (inst *BpmUpdate) Handle(ctx *cli.TaskContext) error {

	in := vo.Update{}
	out := vo.Update{}
	return inst.Service.Update(ctx.Context, &in, &out)
}
