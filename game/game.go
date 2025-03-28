package game

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"tic-tac-toe/game/components"
)

type TicTacToe struct {
	app.Compo
	BoardComponent *components.Board
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *TicTacToe) Render() app.UI {
	player := h.BoardComponent.Model.CurPlayer
	text := fmt.Sprintf("Current Player: %c", player)
	if winner := h.BoardComponent.Model.Winner; winner != 0 {
		text = fmt.Sprintf("Winner: %c", player)
	} else if h.BoardComponent.Model.GameOver {
		text = "Cat's Game!"
	}
	return app.Div().Class("container").Body(
		app.H1().Text(text),
		h.BoardComponent,
	)
}

func InitApp() {
	appInstance := TicTacToe{
		BoardComponent: components.NewBoard(),
	}
	app.Route("/", func() app.Composer { return &appInstance })
}
