package main

import (
	"context"
	"fmt"
	"strconv"
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

func (a *App) Add(onScreenNumber string, memory string) float64 {
	onScreenNumberFloat, _ := strconv.ParseFloat(onScreenNumber, 64)
	memoryFloat, _ := strconv.ParseFloat(memory, 64)
	return onScreenNumberFloat + memoryFloat
}

func (a *App) Subtract(onScreenNumber string, memory string) float64 {
	onScreenNumberFloat, _ := strconv.ParseFloat(onScreenNumber, 64)
	memoryFloat, _ := strconv.ParseFloat(memory, 64)
	return memoryFloat - onScreenNumberFloat
}

func (a *App) Multiply(onScreenNumber string, memory string) float64 {
	onScreenNumberFloat, _ := strconv.ParseFloat(onScreenNumber, 64)
	memoryFloat, _ := strconv.ParseFloat(memory, 64)
	return onScreenNumberFloat * memoryFloat
}

func (a *App) Divide(onScreenNumber string, memory string) float64 {
	onScreenNumberFloat, _ := strconv.ParseFloat(onScreenNumber, 64)
	memoryFloat, _ := strconv.ParseFloat(memory, 64)
	return memoryFloat / onScreenNumberFloat
}

func (a *App) DigitOperation(operation string, onScreenNumber string, memory string) []string {
	if memory == "" {
		memory = onScreenNumber + operation
		onScreenNumber = "0"
		return []string{onScreenNumber, memory}
	} else {
		// swith case with operation
		sinal := memory[len(memory)-1:]
		// removing last character from memory
		memory = memory[:len(memory)-1]
		switch sinal {
		case "+":
			memory = fmt.Sprintf("%f", a.Add(onScreenNumber, memory))
			onScreenNumber = "0"
		case "-":
			memory = fmt.Sprintf("%f", a.Subtract(onScreenNumber, memory))
			onScreenNumber = "0"
		case "*":
			memory = fmt.Sprintf("%f", a.Multiply(onScreenNumber, memory))
			onScreenNumber = "0"
		case "/":
			memory = fmt.Sprintf("%f", a.Divide(onScreenNumber, memory))
			onScreenNumber = "0"
		}
		memory += operation
	}
	return []string{onScreenNumber, memory}
}
