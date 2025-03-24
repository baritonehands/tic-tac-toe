package components

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Board struct {
	app.Compo
}

func (b *Board) Render() app.UI {
	return app.Div().Text("My Div")
}
