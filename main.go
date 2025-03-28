package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log"
	"tic-tac-toe/game"
)

func main() {
	game.InitApp()

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
