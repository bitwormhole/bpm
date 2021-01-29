package main

import (
	"fmt"

	"github.com/bitwormhole/go-wormhole-bpm/bpm"
)

func tryRunMain() error {

	factory := &bpm.CommandContextFactory{}
	cc, err := factory.Create()
	if err != nil {
		return err
	}

	return cc.Handler.Execute(cc)
}

func main() {
	fmt.Println("Bitwormhole Package Manager v1")
	err := tryRunMain()
	if err != nil {
		fmt.Errorf("Error: %+v", err)
	}
}
