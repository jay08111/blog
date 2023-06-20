package main

import (
	"server/domain"
	"server/env"
)

func main() {
	env.Setup()

	domain.Init()

	var echoRouter Router
	echoRouter.Init()

	_ = echoRouter.Run(env.ServerSetting.Port)
}
