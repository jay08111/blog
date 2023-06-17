package main

import (
	"server/domain"
	"server/env"
)

func main() {
	env.Setup()

	var echoRouter Router
	echoRouter.Init()

	domain.Init()

	// defer domain.Close()

	_ = echoRouter.Run(env.ServerSetting.Port)
}
