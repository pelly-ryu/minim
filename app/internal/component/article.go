package component

import (
	"errors"
	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pelly-ryu/minim/app/internal"
	"net/url"
	"sync"
)

const (
	articleTitleId = "article-title"
	articleBodyId  = "article-body"
)

type Article struct {
	app.Compo
	sync.RWMutex

	noteId string // fixed when initiated

	isNew bool
	title string
	body  string
}

func NewArticle() *Article {
	return &Article{
		noteId: uuid.New().String(),
		isNew:  true,
	}
}

func (a *Article) Load(noteId string) {
	note, err := internal.StorageGetNote(noteId)
	if err != nil {
		if errors.Is(err, internal.ErrKeyNotExist) {
			panic("unexpected failure of getting note:" + err.Error())
		}
		panic(err)
	}

	a.Lock()
	a.noteId = noteId
	a.isNew = false
	a.title = note.Title
	a.body = note.Body
	a.Unlock()

	a.Update()
}

func (a *Article) Render() app.UI {
	a.RLock()
	isNew := a.isNew
	titleText := a.title
	bodyText := a.body
	a.RUnlock()

	titleEl := app.H1().ID(articleTitleId).
		ContentEditable(true).
		OnInput(a.onTitleChange).
		Text(titleText)

	if isNew {
		titleEl.Class("new")
	} else {
		titleEl.Class("")
	}

	return app.Div().ID("article").Body(
		app.Div().Class("article-header pure-g").Body(
			app.Div().Class("pure-u-1-2").Style("position", "relative").Body(
				titleEl,
				app.P().Class("article-subtitle").Body(
					app.Span().Text("3:56pm, April 3, 2012"),
				),
			),
		),
		app.Div().ID(articleBodyId).
			ContentEditable(true).
			OnInput(a.onBodyChange).
			Body(app.P().Text(bodyText)),
	)
}

func (a *Article) OnNav(_ app.Context, _ *url.URL) {
	a.Lock()
	isNew := a.isNew
	a.isNew = false
	a.Unlock()

	if isNew {
		app.Window().GetElementByID(articleTitleId).Call("focus")
	}
}

func (a *Article) onTitleChange(ctx app.Context, _ app.Event) {
	title := ctx.JSSrc.Get("innerText").String()
	a.Lock()
	a.title = title
	body := a.body
	a.Unlock()

	ctx.JSSrc.Set("className", "")

	err := internal.StorageSetNote(a.noteId, internal.Note{
		Title: title,
		Body:  body,
	})
	if err != nil {
		panic("unexpected failure of setting note:" + err.Error())
	}
}

func (a *Article) onBodyChange(ctx app.Context, _ app.Event) {
	body := ctx.JSSrc.Get("innerHTML").String()
	a.Lock()
	a.body = body
	title := a.title
	a.Unlock()

	err := internal.StorageSetNote(a.noteId, internal.Note{
		Title: title,
		Body:  body,
	})
	if err != nil {
		panic("unexpected failure of setting note:" + err.Error())
	}
}
