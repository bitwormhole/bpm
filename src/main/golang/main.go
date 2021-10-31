package main

import (
	"strconv"

	"github.com/bitwormhole/bpm"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/collection"
)

func main() {

	app := bpm.Module()
	props := collection.CreateProperties()

	props.SetProperty("application.module", app.GetName())
	props.SetProperty("application.version", app.GetVersion())
	props.SetProperty("application.revision", strconv.Itoa(app.GetRevision()))

	i := starter.InitApp()
	i.Use(app)
	i.UseProperties(props)
	i.Run()
}
