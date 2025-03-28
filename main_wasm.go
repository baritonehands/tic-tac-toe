package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"tic-tac-toe/game"
)

var global = js.Global()

func main() {
	game.InitApp()
	app.RunWhenOnBrowser()
}
