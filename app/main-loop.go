package app

import (
	"os"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

type MainLoop struct {
	markup.Component `class:"looper"`

	ClientFactory cli.ClientFactory   `inject:"#cli-client-factory"`
	Context       application.Context `inject:"context"`
}

func (inst *MainLoop) _Impl() application.Looper {
	return inst
}

func (inst *MainLoop) Loop() error {

	ctx, err := prepareContext(inst.Context)
	if err != nil {
		return err
	}

	client := inst.ClientFactory.CreateClient(ctx)
	return client.Execute("bpm", os.Args[1:])
}
