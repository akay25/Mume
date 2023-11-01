package backend

import (
	"context"
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
)

// App struct
type App struct {
	ctx      context.Context
	cfg      *Config
	cfgStore *ConfigStore
}

// NewApp creates a new App application struct
func NewApp(cfg *Config, store *ConfigStore) *App {
	app := &App{}
	app.cfg = cfg
	app.cfgStore = store
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) StartUp(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetConfig() *Config {
	return a.cfg
}

func (a *App) updateConfig() {
	a.cfgStore.UpdateConfig(*a.cfg)
}

func (a *App) Shutdown(ctx context.Context) {
	a.cfg.WindowHeight = 500
	a.updateConfig()
}

func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	a.Shutdown(ctx)
	return false
}

// Go/system based functions
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CloseMe() {
	a.Shutdown(a.ctx)
	os.Exit(0)
}

func (a *App) Notify(title string, body string, imagePath string) {
	err := beeep.Notify(title, body, imagePath)
	if err != nil {
		panic(err)
	}
}

func (a *App) Alert(title string, body string, imagePath string) {
	err := beeep.Alert(title, body, imagePath)
	if err != nil {
		panic(err)
	}
}
