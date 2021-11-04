package app

import (
	"context"
	"fmt"
	"os"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/contexts"
)

func PrepareContext(ctx application.Context) (context.Context, error) {

	contexts.SetupContextSetter(&myContextSetter{ac: ctx})
	contexts.SetupApplicationContext(ctx)

	// console
	cli.SetupConsole(ctx, nil)
	console, err := cli.GetConsole(ctx)
	if err != nil {
		return nil, err
	}
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	console.SetOutput(os.Stdout)
	console.SetError(os.Stderr)
	console.SetWD(wd)

	return ctx, nil
}

type myContextSetter struct {
	ac application.Context
}

func (inst *myContextSetter) _Impl() contexts.ContextSetter {
	return inst
}

func (inst *myContextSetter) GetContext() context.Context {
	return inst.ac
}

func (inst *myContextSetter) SetValue(key interface{}, value interface{}) {
	name := inst.stringifyKey(key)
	atts := inst.ac.GetAttributes()
	atts.SetAttribute(name, value)
}

func (inst *myContextSetter) stringifyKey(key interface{}) string {
	o1, ok := key.(string)
	if ok {
		return o1
	}
	return fmt.Sprint(key)
}
