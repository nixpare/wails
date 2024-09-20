package main

import (
	_ "embed"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
)

func main() {
	app := application.New(application.Options{
		Name:        "Key Bindings Demo",
		Description: "A demo of the Key Bindings Options",
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		KeyBindings: map[string]func(window application.Window){
			"shift+ctrl+c": func(window application.Window) {
				window.Center()
			},
		},
	})

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Name:  "Window 1",
		Title: "Window 1",
		URL:   "https://wails.io",
		KeyBindings: map[string]func(window *application.WebviewWindow){
			"F12": func(window *application.WebviewWindow) {
				window.OpenDevTools()
			},
		},
	})

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Name:  "Window 2",
		Title: "Window 2",
		URL:   "https://google.com",
		KeyBindings: map[string]func(window *application.WebviewWindow){
			"F12": func(window *application.WebviewWindow) {
				println("Window 2: Toggle Dev Tools")
			},
		},
	})

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}

}
