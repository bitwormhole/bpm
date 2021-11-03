package command

import (
	"strings"

	"github.com/bitwormhole/bpm/service"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type BpmAutoUpgrade struct {
	markup.Component `class:"cli-handler"`

	Service service.UpgradeService `inject:"#bpm-upgrade-service"`
}

func (inst *BpmAutoUpgrade) _Impl() cli.Handler {
	return inst
}

func (inst *BpmAutoUpgrade) Init(service cli.Service) error {
	service.RegisterHandler("bpm-auto-upgrade", inst)
	return nil
}

func (inst *BpmAutoUpgrade) Handle(ctx *cli.TaskContext) error {
	args := ctx.CurrentTask.Arguments
	doUpdate := inst.getArgDoUpdate(args)
	return inst.Service.UpgradeAuto(ctx.Context, doUpdate)
}

func (inst *BpmAutoUpgrade) getArgDoUpdate(args []string) bool {
	for _, str := range args {
		if strings.Contains(str, "update") {
			return true
		}
	}
	return false
}
