package main

import "fmt"

type plugin string

func (g plugin) Cleanup() {
	fmt.Println("Hello Universe")
}

var Cleanup plugin
