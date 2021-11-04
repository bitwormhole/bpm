package command

import (
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type BpmFetch struct {
	markup.Component `class:"cli-handler"`

	Service service.FetchService `inject:"#bpm-fetch-service"`
}

func (inst *BpmFetch) _Impl() cli.Handler {
	return inst
}

func (inst *BpmFetch) Init(service cli.Service) error {
	service.RegisterHandler("bpm-fetch", inst)
	return nil
}

func (inst *BpmFetch) Handle(ctx *cli.TaskContext) error {

	a2, err := parseArguments(ctx.CurrentTask.Arguments)
	if err != nil {
		return err
	}

	in := vo.Fetch{}
	out := vo.Fetch{}
	in.PackageNames = a2.Packages
	return inst.Service.Fetch(ctx.Context, &in, &out)
}
