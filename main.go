package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/kkamperschroer/GoHome/config"
	_ "github.com/kkamperschroer/GoHome/core"
)

var (
	configFilePath = flag.String("c", "config.json", "Path to the config file.")
	debug          = flag.Bool("d", false, "Debug level logging")
)

func main() {
	flag.Parse()

	fmt.Println("Welcome! Thanks for using GoHome!")

	config, err := config.LoadConfigData(*configFilePath)

	if err != nil {
		fmt.Println("Failed to load your config:\n\t" + err.Error())
	}

	if *debug {
		bytes, err := json.MarshalIndent(config, "", "  ")

		if err == nil {
			fmt.Println(string(bytes))
		}
	}

}
