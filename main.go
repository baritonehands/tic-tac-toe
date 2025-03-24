package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log"
	"tic-tac-toe/gh-pages/components"
)

// App is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type App struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *App) Render() app.UI {
	return app.Div().Class("app-container").Body(
		app.H1().Text("Hello World!"),
		&components.Board{},
	)
}

func main() {
	app.Route("/", func() app.Composer { return &App{} })
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite("gh-pages/tic-tac-toe", &app.Handler{
		Name:        "Tic Tac Toe",
		Description: "A tic-tac-toe example",
		Resources:   app.GitHubPages("tic-tac-toe"),
		Styles:      []string{"/tic-tac-toe.css"},
	})

	if err != nil {
		log.Fatal(err)
	}
}
