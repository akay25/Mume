package main

import (
	"embed"

	"Mume/backend"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS
var logger = backend.Logger

func main() {
	// Load/create config file
	store, err := backend.NewConfigStore()
	if err != nil {
		logger.Panicf("Could not initialize the config store: %v", err)
		return
	}
	logger.Debugf("Config path found to be: %v", store)

	// Read existing values
	cfg, err := store.Config()
	if err != nil {
		logger.Panicf("Could not retrieve the configuration: %v from config path %v", err, store)
		return
	}

	// Set loglevel
	switch cfg.LogLevel {
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
	case "INFO":
		logger.SetLevel(logrus.InfoLevel)
	case "WARN":
		logger.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
	}

	// Bases on the config file, launch the app

	// Create an instance of the app structure
	app := backend.NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "Mume",
		Width:     1000,
		Height:    563,
		MinWidth:  1000,
		MinHeight: 563,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.StartUp,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
