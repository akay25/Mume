package backend

import (
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
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
