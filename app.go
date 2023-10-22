package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Go/system based functions
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CloseMe() {
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
