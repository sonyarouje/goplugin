package main

import (
	"fmt"
	"goplugin/writers"
)

func main() {
	fmt.Println("Print Hello world")
	ws, err := writers.Load()
	if err != nil {
		panic(err)
	}
	util := NewPluginUtils()
	for _, w := range ws {
		w.Write(util)
	}
}

type PluginUtil struct {
}

func NewPluginUtils() PluginUtil {
	return PluginUtil{}
}

func (p PluginUtil) Print(a ...interface{}) {
	fmt.Println(a...)
}
