package backend

import (
	"fmt"

	"github.com/gen2brain/beeep"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// Go/system based functions
func (a *App) GetLibraryPath() string {
	return a.cfg.LibraryPath
}

func (a *App) SetLibraryPath(path string) {
	a.cfg.LibraryPath = path
	a.updateConfig()
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CloseMe() {
	a.Shutdown(a.ctx)
	rt.Quit(a.ctx)
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
