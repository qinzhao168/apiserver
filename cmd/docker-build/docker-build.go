package main

import (
	"apiserver/cmd/docker-build/app"
	"apiserver/pkg/util/log"
)

func main() {
	dockerBuild := app.NewDockerBuild()
	if err := app.Run(dockerBuild); err != nil {
		log.Fatalf("start dockerBuild err: %v", err)
	}
}