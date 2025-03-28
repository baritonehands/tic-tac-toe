package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"syscall/js"
	"tic-tac-toe/game"
	"tic-tac-toe/game/models"
)

var global = js.Global()

func getSolverCache(this js.Value, args []js.Value) any {
	cache := make(map[string]any, len(models.SolverCache))
	for k, v := range models.SolverCache {
		cache[fmt.Sprint(k)] = js.ValueOf(fmt.Sprint(v))
	}
	return js.ValueOf(cache)
}

func interopTest(this js.Value, args []js.Value) any {
	return js.ValueOf([]any{"foo", "bar"})
}

func main() {
	game.InitApp()

	global.Set("getSolverCache", js.FuncOf(getSolverCache))
	global.Set("interopTest", js.FuncOf(interopTest))

	app.RunWhenOnBrowser()
}
