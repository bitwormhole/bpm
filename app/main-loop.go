package app

import (
	"errors"
	"os"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
)

type Runner struct {
	i application.Initializer
}

func (inst *Runner) Init(i application.Initializer) {
	inst.i = i
}

func (inst *Runner) Run() error {

	rt, err := inst.i.RunEx()
	if err != nil {
		return err
	}

	appctx := rt.Context()
	ctx, err := PrepareContext(appctx)
	if err != nil {
		return err
	}

	o1, err := appctx.GetComponent("#cli-client-factory")
	if err != nil {
		return err
	}

	o2, ok := o1.(cli.ClientFactory)
	if !ok {
		return errors.New("o1.(cli.ClientFactory) cast fail")
	}

	args := os.Args[1:]
	client := o2.CreateClient(ctx)
	err = client.Execute("bpm", args)
	if err != nil {
		return err
	}

	return rt.Exit()
}
