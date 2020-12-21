package component

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pelly-ryu/minim/app/internal"
)

type NoteList struct {
	app.Compo

	opened bool
}

func NewNoteList() *NoteList {
	return &NoteList{
		opened: false,
	}
}

func (l *NoteList) Render() app.UI {
	if !l.opened {
		return app.Aside().ID("note-list").Class("closed").
			OnClick(func(ctx app.Context, e app.Event) {
				l.Toggle()
			})
	}

	ids, err := internal.StorageListNoteId()
	if err != nil {
		panic(err)
	}

	var noteListEl []app.UI
	for _, id := range ids {
		n, err := internal.StorageGetNote(id)
		if err != nil {
			panic(err)
		}

		short := n.Body
		if len(n.Body) > 50 {
			short = n.Body[:50]
		}

		el := app.Div().Class("note-list-item note-list-item-selected").Body(
			app.H5().Class("note-list-name").Text("Tilo Mitra"),
			app.H4().Class("note-list-subject").Text(n.Title),
			app.P().Class("note-list-desc").Text(short),
		)

		noteListEl = append(noteListEl, el)
	}

	return app.Aside().ID("note-list").Class("pure-u-1").Body(
		noteListEl...
	)
}

func (l *NoteList) Toggle() {
	l.opened = !l.opened
	l.Update()
}

func (l *NoteList) Opened() bool {
	return l.opened
}
