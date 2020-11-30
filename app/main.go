// +build wasm

// The UI is running only on a web browser. Therefore, the build instruction
// above is to compile the code below only when the program is built for the
// WebAssembly (wasm) architecture.

package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type rootWrapper struct {
	app.Compo
}

func (h *rootWrapper) Render() app.UI {
	textarea := app.Textarea()
	textarea.OnChange(func(ctx app.Context, e app.Event) {
		fmt.Println(ctx.JSSrc.Get("value").String())
	})

	return app.Div().Body(
		app.H1().Text("Hello World!"),
		textarea,
	)
}

func main() {
	app.Route("/", &rootWrapper{})
	app.Run()
}
