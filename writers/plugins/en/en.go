package main

import "fmt"

type Plugin struct {
}

func (d Plugin) Write() {
	fmt.Println("Hello world")
}

var P Plugin
