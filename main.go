package main

import (
	"github.com/gg-amorim/opportunities/config"
	"github.com/gg-amorim/opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Init Congfigs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initializion error: %v", err)
		return
	}
	// Init Router
	router.Initialize()
}
