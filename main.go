package main

import(
	"fmt"
	)

func main() {
	fmt.Println("Welcome! Thanks for using GoHome!");

	var goHome Automator

	goHome.Init()
	goHome.Run()
}