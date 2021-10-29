package command

import (
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type BpmMake struct {
	markup.Component `class:"cli-handler"`

	Service service.MakeService `inject:"#bpm-make-service"`
}

func (inst *BpmMake) _Impl() (cli.Handler, cli.CommandHelper) {
	return inst, inst
}

func (inst *BpmMake) GetHelpInfo() *cli.CommandHelpInfo {
	info := &cli.CommandHelpInfo{}
	info.Name = "make"
	info.Title = "生成BPM包"
	info.Description = "usage: bpm make"
	info.Content = "在当前目录下查找【.bpm】文件夹，并生成其BPM包"
	return info
}

func (inst *BpmMake) Init(service cli.Service) error {
	service.RegisterHandler("bpm-make", inst)
	return nil
}

func (inst *BpmMake) Handle(ctx *cli.TaskContext) error {
	in := vo.Make{}
	out := vo.Make{}
	return inst.Service.Make(ctx.Context, &in, &out)
}
