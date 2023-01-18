package main

import (
	"github.com/cucumberjaye/qtsoft/configs"
	"github.com/cucumberjaye/qtsoft/internal/pkg/app"
	"github.com/cucumberjaye/qtsoft/pkg/logger"
)

func init() {
	logger.InitLogger()
	configs.InitConfigs()
}

func main() {
	a := app.New()

	err := a.Run()
	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}
}
