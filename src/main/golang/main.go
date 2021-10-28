package main

import (
	"github.com/bitwormhole/bpm"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.Use(bpm.Module())
	i.Run()
}
