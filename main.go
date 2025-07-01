package main

import (
	"github.com/job-nunes/students/config"
	"github.com/job-nunes/students/router"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
