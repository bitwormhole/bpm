package bpm

import (
	"embed"

	"github.com/bitwormhole/bpm/gen"
	"github.com/bitwormhole/starter"
	startercli "github.com/bitwormhole/starter-cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/bpm"
	myModuleVer  = "v0"
	myModuleRev  = 0
)

//go:embed src/main/resources
var theMainRes embed.FS

func Module() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName).Version(myModuleVer).Revision(myModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	mb.OnMount(gen.ExportConfigBPM)

	mb.Dependency(starter.Module())
	mb.Dependency(startercli.Module())

	return mb.Create()
}
