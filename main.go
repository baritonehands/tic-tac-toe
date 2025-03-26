package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log"
	"tic-tac-toe/gh-pages/components"
)

// App is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type App struct {
	app.Compo
	boardComponent *components.Board
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *App) Render() app.UI {
	player := h.boardComponent.Model.CurPlayer
	text := fmt.Sprintf("Current Player: %c", player)
	if winner := h.boardComponent.Model.Winner; winner != 0 {
		text = fmt.Sprintf("Winner: %c", player)
	} else if h.boardComponent.Model.GameOver {
		text = "Cat's Game!"
	}
	return app.Div().Class("container").Body(
		app.H1().Text(text),
		h.boardComponent,
	)
}

func main() {
	appInstance := App{
		boardComponent: components.NewBoard(),
	}
	app.Route("/", func() app.Composer { return &appInstance })
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite("gh-pages/tic-tac-toe", &app.Handler{
		Name:        "Tic Tac Toe",
		Description: "A tic-tac-toe example",
		Resources:   app.GitHubPages("tic-tac-toe"),
		Styles:      []string{"/web/tic-tac-toe.css"},
	})

	if err != nil {
		log.Fatal(err)
	}
}
