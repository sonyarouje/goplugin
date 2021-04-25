package main

import (
	"goplugin/writers/common"
)

type Plugin struct {
}

func (d Plugin) Write(plugger common.Plugger) {
	// fmt.Println("Hallo Welt")
	plugger.Print("Hallo Welt")
}

var P Plugin
