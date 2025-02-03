package main

import (
	"github.com/newonne/gopportunities/config"
	"github.com/newonne/gopportunities/router"
)

func main() {
	logger := config.GetLogger("main")

	// inicializa configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
	}
	router.Initialize()
}
