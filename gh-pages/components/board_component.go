package components

import (
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"maps"
	"tic-tac-toe/gh-pages/models"
)

type Board struct {
	app.Compo
	Model *models.Board
}

func NewBoard() *Board {
	return &Board{
		Model: models.NewBoard(),
	}
}

func (b *Board) Cell(x, y int) app.UI {
	text := "\u00A0"
	classes := []string{"cell"}
	if b.Model.At(x, y) == 'X' {
		text = "X"
		classes = append(classes, "X")
	} else if b.Model.At(x, y) == 'O' {
		text = "O"
		classes = append(classes, "O")
	}
	return app.Div().Class(classes...).Text(text).OnClick(b.cellClickHandler(x, y))
}

func (b *Board) Row(y int) app.UI {
	return app.Div().Class(fmt.Sprintf("row%d", y+1)).Body(
		b.Cell(0, y),
		app.Div().Class("vertical").Text("\u00A0"),
		b.Cell(1, y),
		app.Div().Class("vertical").Text("\u00A0"),
		b.Cell(2, y),
	)
}

func (b *Board) Render() app.UI {
	if !b.Model.GameOver {
		solver := models.Solver{Board: b.Model, AsPlayer: b.Model.CurPlayer, Level: 9 - b.Model.Taken}
		score := solver.Score()
		strScore := maps.Collect(it.Map2(maps.All(score), func(k byte, v float64) (byte, string) {
			return k, fmt.Sprintf("%0.2f%%", v*100.0)
		}))
		fmt.Printf("Player %c: %v\n", solver.AsPlayer, strScore)
	}
	return app.Div().Body(
		b.Row(0),
		app.Div().Class("horizontal").Text("\u00A0"),
		b.Row(1),
		app.Div().Class("horizontal").Text("\u00A0"),
		b.Row(2),
	)
}

func (b *Board) cellClickHandler(x, y int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		err := b.Model.Move(x, y)
		if err != nil {
			ctx.Reload()
		}
	}
}
