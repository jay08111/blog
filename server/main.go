package main

import (
	"flag"
	"server/domain"
	"server/env"
)

func main() {
	configName := flag.String("config", "", "config file name")

	flag.Parse()

	env.Setup(*configName)

	// var echoRouter Router
	// echoRouter.Init()

	domain.Init()

	// defer domain.Close()

	// _ = echoRouter.Run(env.ServerSetting.Port)
}
