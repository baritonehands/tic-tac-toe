package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World!")
}

func main() {
	app.Route("/", func() app.Composer { return &hello{} })
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite("gh-pages/tic-tac-toe", &app.Handler{
		Name:        "Tic Tac Toe",
		Description: "A tic-tac-toe example",
		Resources:   app.GitHubPages("tic-tac-toe"),
	})

	if err != nil {
		log.Fatal(err)
	}
}
