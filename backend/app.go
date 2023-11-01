package backend

import (
	"context"
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
