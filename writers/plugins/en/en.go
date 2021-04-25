package main

import "goplugin/writers/common"

type Plugin struct {
}

func (d Plugin) Write(plugger common.Plugger) {
	// fmt.Println("Hello world")
	plugger.Print("Hello world")
}

var P Plugin
