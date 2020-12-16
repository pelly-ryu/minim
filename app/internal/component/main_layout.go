package component

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type MainLayout struct {
	app.Compo
}

func (l *MainLayout) Render() app.UI {
	return app.Div().ID("layout").Class("content").Body(
		NewNoteList(),
		NewArticle(),
	)
}
