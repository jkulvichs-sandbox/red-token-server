package main

import (
	"flag"
	"fmt"
	"github.com/jkulvichs-sandbox/red-token/config"
	"github.com/jkulvichs-sandbox/red-token/logger"
	"github.com/jkulvichs-sandbox/red-token/storage"
	"os"
	"os/signal"
)

func main() {

	// Flags parsing
	confPath := flag.String("config", "config.yml", "Path to config file")
	confCreateDefault := flag.Bool("config-create-default", true, "Create default file if doesn't exists")
	flag.Parse()

	// Config loading
	conf, err := config.LoadConfig(*confPath, *confCreateDefault)
	if err != nil {
		panic(err)
	}

	// Configuring a logger
	log, err := logger.New(conf.Logger)
	if err != nil {
		panic(fmt.Sprintf("can't configure logger: %s", err))
	}

	// Connecting to Storage
	log.Debug("connecting to store ...")
	store, err := storage.Connect(conf.Storage)
	if err != nil {
		log.Fatalf("can't configure storage: %s", err)
	}
	log.Infof("store configured as %s", store.Type)

	// Waiting for exit
	log.Info("the app is running and waiting for exit signal")
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Kill, os.Interrupt)
	<-exitChan
	log.Info("exit signal was gotten, waiting for resources dispose")

	// Closing handlers and resources disposing
	if err := store.Disconnect(); err != nil {
		log.Errorf("can't dispose the store resources correctly: %s", err)
	}
}
