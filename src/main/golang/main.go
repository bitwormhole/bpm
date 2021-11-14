package main

import (
	"strconv"

	"github.com/bitwormhole/bpm"
	"github.com/bitwormhole/bpm/app"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/collection"
)

// bpm 是一个命令行工具，用来管理本地的 BPM (Bitwormhole Package Manager) 软件包
func main() {

	mod := bpm.Module()
	props := collection.CreateProperties()
	props.SetProperty("application.module", mod.GetName())
	props.SetProperty("application.version", mod.GetVersion())
	props.SetProperty("application.revision", strconv.Itoa(mod.GetRevision()))

	i := starter.InitApp()
	i.UseMain(mod)
	i.UseProperties(props)

	r := app.Runner{}
	r.Init(i)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
