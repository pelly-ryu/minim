package component

import "github.com/maxence-charriere/go-app/v7/pkg/app"

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
		return app.Aside().ID("note-list").Class("closed")
	}

	return app.Aside().ID("note-list").Class("pure-u-1").Body(
		app.Div().Class("email-item email-item-selected").Body(
			app.H5().Class("email-name").Text("Tilo Mitra"),
			app.H4().Class("email-subject").Text("Hello from Toronto"),
			app.P().Class("email-desc").Text("Hey, I just wanted to check in with you from Toronto. I got here earlier today."),
		),
		app.Div().Class("email-item email-item-unread").Body(
			app.H5().Class("email-name").Text("Tilo Mitra"),
			app.H4().Class("email-subject").Text("Hello from Toronto"),
			app.P().Class("email-desc").Text("Hey, I just wanted to check in with you from Toronto. I got here earlier today."),
		),
		app.Div().Class("email-item").Body(
			app.H5().Class("email-name").Text("Tilo Mitra"),
			app.H4().Class("email-subject").Text("Hello from Toronto"),
			app.P().Class("email-desc").Text("Hey, I just wanted to check in with you from Toronto. I got here earlier today."),
		),
	)
}

func (l *NoteList) Toggle() {
	l.opened = !l.opened
}

func (l *NoteList) Opened() bool {
	return l.opened
}
