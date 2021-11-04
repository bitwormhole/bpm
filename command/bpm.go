package command

import (
	"errors"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/collection"
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
	action, args2, err := inst.parseArgs(args1)
	if err != nil {
		return err
	}
	return client.Execute("bpm-"+action, args2)
}

// @return( action , args2 , error)
func (inst *Bpm) parseArgs(args []string) (string, []string, error) {
	if args == nil {
		return "", nil, errors.New("no args")
	}
	const actionIndex = 1 // index of c2
	action := "help"
	a2 := []string{}
	if actionIndex < len(args) {
		action = args[actionIndex]
		a2 = args[actionIndex+1:]
	}
	return action, a2, nil
}

////////////////////////////////////////////////////////////////////////////////

type BpmArguments struct {
	Action   string
	Packages []string
}

func parseArguments(args []string) (*BpmArguments, error) {

	dst := &BpmArguments{}
	a2 := collection.InitArguments(args)
	rd := a2.NewReader()

	// for flags
	flagExample := rd.GetFlag("--example")
	if flagExample.Exists() {
		// todo ...
	}

	// for packs
	packs := make([]string, 0)
	for i := 0; ; i++ {
		item, ok := rd.PickNext()
		if ok {
			if i == 0 {
				dst.Action = item
			} else {
				packs = append(packs, item)
			}
		} else {
			break
		}
	}

	dst.Packages = packs
	return dst, nil
}
