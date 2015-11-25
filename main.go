package main

import (
	"flag"
	"os"

	"github.com/kkamperschroer/GoHome/config"
	"github.com/kkamperschroer/GoHome/core"

	"github.com/pivotal-golang/lager"
)

var (
	configFilePath = flag.String("c", "config.json", "Path to the config file.")
	debug          = flag.Bool("d", false, "Debug level logging")
)

func main() {
	flag.Parse()

	logger := lager.NewLogger("GoHome")
	logLevel := lager.INFO

	if *debug {
		logLevel = lager.DEBUG
	}

	logger.RegisterSink(lager.NewWriterSink(os.Stdout, logLevel))

	logger.Info("Welcome! Thanks for using GoHome!")

	config, err := config.LoadConfigData(*configFilePath)

	if err != nil {
		logger.Fatal("Failed to load your config", err)
	}

	goHome, err := core.NewGoHome(config, logger)

	if err != nil {
		logger.Fatal("Failed to initialize GoHome", err)
	}

	err = goHome.Go()

	if err != nil {
		logger.Fatal("Error running GoHome", err)
	}
}
