package main

import (
	"fmt"
	"strings"
)

func main() {
	contains()
}

func contains() {
	fmt.Println("きみは薄暗い洞窟の中にいる。")
	const command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Print("洞窟を出る？：", exit)
}
