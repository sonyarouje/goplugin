package main

import "fmt"

type Plugin struct {
}

func (d Plugin) Write() {
	fmt.Println("Hallo Welt")
}

var P Plugin
