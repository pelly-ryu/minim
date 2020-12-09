// +build wasm

// The UI is running only on a web browser. Therefore, the build instruction
// above is to compile the code below only when the program is built for the
// WebAssembly (wasm) architecture.

package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pelly-ryu/minim/app/internal/component"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type mainLayout struct {
	app.Compo
}

func (h *mainLayout) Render() app.UI {
	return app.Div().ID("layout").Class("content").Body(
		component.NewNoteList(),
		&component.Article{},
	)
}

func main() {
	app.Route("/", &mainLayout{})
	app.Run()
}
