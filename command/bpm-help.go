package command

import (
	"io"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter-cli/handlers"
	"github.com/bitwormhole/starter/markup"
)

type BpmHelp struct {
	markup.Component `class:"cli-handler"`
}

func (inst *BpmHelp) _Impl() cli.Handler {
	return inst
}

func (inst *BpmHelp) Init(service cli.Service) error {

	//	service.RegisterHandler("bpm-help", inst)

	h := &handlers.Help{}
	service.RegisterHandler("bpm-help", h)

	return nil
}

func (inst *BpmHelp) Handle(ctx *cli.TaskContext) error {

	out := ctx.Console.Output()

	inst.println(out, "usage: bpm [command] [...args]")
	inst.println(out, "  [command] = (update|install|run|upgrade|help)")

	return nil
}

func (inst *BpmHelp) print(o io.Writer, msg string) {
	o.Write([]byte(msg))
}

func (inst *BpmHelp) println(o io.Writer, msg string) {
	inst.print(o, msg+"\n")
}
