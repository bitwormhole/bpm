package command

import (
	"errors"
	"strings"

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

	// bpm-run [package:script]

	pack, script, args, err := inst.parseArgs(ctx.CurrentTask.Arguments)
	if err != nil {
		return err
	}
	in := vo.Run{}
	out := vo.Run{}
	in.PackageName = pack
	in.ScriptName = script
	in.Arguments = args
	return inst.Service.Run(ctx.Context, &in, &out)
}

// @return ( package, script, args2, error )
func (inst *BpmRun) parseArgs(args []string) (string, string, []string, error) {

	// args like 'bpm run package:script [args2...]'
	const psIndex = 1 // the index of 'package:script'

	if psIndex < len(args) {
		ps := args[psIndex]
		psArray := strings.Split(ps, ":")
		if len(psArray) == 2 {
			pack := psArray[0]
			script := psArray[1]
			return pack, script, args[psIndex+1:], nil
		}
	}

	msg := "bad format of arguments"
	want := ", want: 'bpm run {package}:{script} [args...]'"
	have := strings.Builder{}
	have.WriteString(", have: '")
	for _, item := range args {
		have.WriteString(item)
		have.WriteRune(' ')
	}
	have.WriteString("'")
	return "", "", nil, errors.New(msg + want + have.String())
}
