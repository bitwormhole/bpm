package main

import (
	"strconv"

	"github.com/bitwormhole/bpm"
	"github.com/bitwormhole/bpm/app"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/collection"
)

func main() {

	mod := bpm.Module()
	props := collection.CreateProperties()
	props.SetProperty("application.module", mod.GetName())
	props.SetProperty("application.version", mod.GetVersion())
	props.SetProperty("application.revision", strconv.Itoa(mod.GetRevision()))

	i := starter.InitApp()
	i.Use(mod)
	i.UseProperties(props)

	r := app.Runner{}
	r.Init(i)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
