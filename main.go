package main

import (
	"embed"

	"Mume/backend"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
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
	cfg, err := store.GetConfig()
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
	app := backend.NewApp(&cfg, store)

	// Create application with options
	err = wails.Run(&options.App{
		Title:                    backend.APPLICATION_NAME,
		Width:                    cfg.WindowWidth,
		Height:                   cfg.WindowHeight,
		MinWidth:                 backend.MIN_WINDOW_WIDTH,
		MinHeight:                backend.MIN_WINDOW_HEIGHT,
		Frameless:                true,
		LogLevel:                 wailsLogger.INFO,
		LogLevelProduction:       wailsLogger.ERROR,
		BackgroundColour:         &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		EnableDefaultContextMenu: false,
		OnStartup:                app.StartUp,
		OnDomReady:               app.DomReady,
		OnShutdown:               app.Shutdown,
		OnBeforeClose:            app.BeforeClose,
		Bind: []interface{}{
			app,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
