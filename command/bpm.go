package command

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

type Bpm struct {
	markup.Component `class:"cli-handler"`
}

func (inst *Bpm) _Impl() cli.Handler {
	return inst
}

func (inst *Bpm) Init(service cli.Service) error {
	service.RegisterHandler("bpm", inst)
	return nil
}

func (inst *Bpm) Handle(ctx *cli.TaskContext) error {
	cc := ctx.Context
	client := ctx.Service.GetClient(cc)
	args1 := ctx.CurrentTask.Arguments
	cmd2, args2, err := inst.parseArgs(args1)
	if err != nil {
		return err
	}
	if cmd2 == "" {
		cmd2 = "help"
	}
	return client.Execute("bpm-"+cmd2, args2)
}

func (inst *Bpm) parseArgs(args []string) (string, []string, error) {
	if args == nil {
		return "", nil, errors.New("no args")
	}
	const index = 1 // index of c2
	c2 := ""
	a2 := []string{}
	if index < len(args) {
		c2 = args[index]
		a2 = args[index+1:]
	}
	return c2, a2, nil
}

func getPackageNameListFromArgs(args []string) []string {
	list := make([]string, 0)
	if args == nil {
		return list
	}
	const wantIndex = 1
	if len(args) <= wantIndex {
		return list
	}
	src := args[wantIndex:]
	for _, item := range src {
		if strings.HasPrefix(item, "-") {
			continue
		}
		list = append(list, item)
	}
	return list
}
