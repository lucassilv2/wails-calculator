package main

import (
	"context"
	"fmt"
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

func (a *App) DigitNumber(number int, text string) string {
	if text == "0" {
		text = ""
	}
	return fmt.Sprintf("%s%d", text, number)
}

func (a *App) Clear() string {
	return "0"
}

func (a *App) ClearMemory() string {
	return ""
}

func (a *App) Backspace(text string) string {
	if len(text) == 1 {
		return "0"
	}
	return text[:len(text)-1]
}
