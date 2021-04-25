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
	for _, w := range ws {
		w.Write()
	}
}
