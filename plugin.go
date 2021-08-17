package main

import "fmt"

type plugin string

func (g plugin) Clean() {
	fmt.Println("Hello Universe")
}

var Clean plugin
