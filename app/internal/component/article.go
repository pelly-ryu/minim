package component

import (
	"errors"
	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pelly-ryu/minim/app/internal"
)

type Article struct {
	app.Compo

	noteId string
}

func NewArticle() *Article {
	return &Article{
		noteId: uuid.New().String(),
	}
}

func LoadArticle(noteId string) *Article {
	panic("not implemented")
}

func (a *Article) Render() app.UI {
	return app.Div().ID("article").Body(
		app.Div().Class("article-header pure-g").Body(
			app.Div().Class("pure-u-1-2").Body(
				app.H1().Class("article-title").
					ContentEditable(true).
					OnInput(a.onTitleChange).
					Text("New note..."),
				app.P().Class("article-subtitle").Body(
					app.Span().Text("3:56pm, April 3, 2012"),
				),
			),
		),
		app.Div().Class("article-body").
			ContentEditable(true).
			OnInput(a.onBodyChange).
			Body(app.P().Text("")),
	)
}

func (a *Article) onTitleChange(ctx app.Context, _ app.Event) {
	title := ctx.JSSrc.Get("innerText").String()
	body := ""
	note, err := internal.StorageGetNote(a.noteId)
	if err != nil {
		if !errors.Is(err, internal.ErrKeyNotExist) {
			panic("unexpected failure of getting note:" + err.Error())
		}
	} else {
		body = note.Body
	}

	err = internal.StorageSet(a.noteId, internal.Note{
		Title: title,
		Body:  body,
	})
	if err != nil {
		panic("unexpected failure of setting note:" + err.Error())
	}
}

func (a *Article) onBodyChange(ctx app.Context, _ app.Event) {
	title := ""
	body := ctx.JSSrc.Get("innerHTML").String()
	note, err := internal.StorageGetNote(a.noteId)
	if err != nil {
		if !errors.Is(err, internal.ErrKeyNotExist) {
			panic("unexpected failure of getting note:" + err.Error())
		}
	} else {
		title = note.Title
	}

	err = internal.StorageSet(a.noteId, internal.Note{
		Title: title,
		Body:  body,
	})
	if err != nil {
		panic("unexpected failure of setting note:" + err.Error())
	}
}
