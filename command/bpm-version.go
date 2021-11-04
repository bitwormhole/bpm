package command

import (
	"strings"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// BpmVersion ...
type BpmVersion struct {
	markup.Component `class:"cli-handler"`

	Context application.Context `inject:"context"`
}

func (inst *BpmVersion) _Impl() cli.Handler {
	return inst
}

// Init ...
func (inst *BpmVersion) Init(service cli.Service) error {
	service.RegisterHandler("bpm-version", inst)
	service.RegisterHandler("bpm-about", inst)
	return nil
}

// Handle ...
func (inst *BpmVersion) Handle(ctx *cli.TaskContext) error {

	const bitHome = "BITWORMHOLE_HOME"

	console := ctx.Console
	context := inst.Context
	props := context.GetProperties()

	version := props.GetProperty("application.version", "0.0.0")
	home, _ := context.GetEnvironment().GetEnv(bitHome)

	console.WriteString("\nBitWormhole Package Manager (version:" + version + ")")
	if inst.isAbout(ctx) {
		console.WriteString("\n" + bitHome + "=" + home)
	}
	console.WriteString("\n\n")

	return nil
}

func (inst *BpmVersion) isAbout(ctx *cli.TaskContext) bool {
	args := inst.Context.GetArguments()
	all := args.Export()
	for _, item := range all {
		if strings.TrimSpace(item) == "about" {
			return true
		}
	}
	return false
}
