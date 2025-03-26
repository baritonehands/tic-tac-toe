package components

import (
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"maps"
	"math"
	"tic-tac-toe/gh-pages/models"
)

type Board struct {
	app.Compo
	Model  *models.Board
	Solver *models.Solver
	Scores map[byte]float64
}

func NewBoard() *Board {
	boardModel := models.NewBoard()
	return &Board{
		Model:  boardModel,
		Solver: models.NewSolver(boardModel),
	}
}

func (b *Board) Cell(x, y int) app.UI {
	text := "\u00A0"
	classes := []string{"cell"}
	styles := map[string]string{}
	idx := models.CoordToIndex(x, y)
	if b.Model.At(x, y) == 'X' {
		text = "X"
		classes = append(classes, "X")
	} else if b.Model.At(x, y) == 'O' {
		text = "O"
		classes = append(classes, "O")
	} else {
		score := b.Scores[byte(idx)] * 100.0
		styles["opacity"] = fmt.Sprintf("%0.2f%%", math.Abs(score))
		if score < 0 {
			classes = append(classes, "bad")
		} else {
			classes = append(classes, "good")
		}
	}
	return app.Div().Class(classes...).Styles(styles).Text(text).OnClick(b.cellClickHandler(x, y))
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
		b.Scores = b.Solver.Score()
		strScore := maps.Collect(it.Map2(maps.All(b.Scores), func(k byte, v float64) (byte, string) {
			return k, fmt.Sprintf("%0.2f%%", v*100.0)
		}))
		fmt.Printf("Player %c: %v\n", b.Solver.AsPlayer, strScore)
	} else {
		b.Scores = map[byte]float64{}
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
		b.Solver.BoardUpdated()
	}
}
