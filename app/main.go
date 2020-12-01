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
type mainLayout struct {
	app.Compo
}

func (h *mainLayout) Render() app.UI {
	textarea := app.Textarea()
	textarea.OnChange(func(ctx app.Context, e app.Event) {
		fmt.Println(ctx.JSSrc.Get("value").String())
	})

	return app.Div().ID("layout").Class("content pure-g").Body(
		app.Div().ID("list").Class("pure-u-1").Body(
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
		),
		app.Div().ID("main").Class("pure-u-1").Body(
			app.Div().Class("email-content").Body(
				app.Div().Class("email-content-header pure-g").Body(
					app.Div().Class("pure-u-1-2").Body(
						app.H1().Class("email-content-title").Text("Hello from Toronto"),
						app.P().Class("email-content-subtitle").Body(
							app.Text("From "),
							app.A().Text("Tilo Mitra"),
							app.Text(" at "),
							app.Span().Text("3:56pm, April 3, 2012"),
						),
					),
					app.Div().Class("email-content-controls pure-u-1-2").Body(
						app.Button().Class("secondary-button pure-button").Text("Reply"),
						app.Button().Class("secondary-button pure-button").Text("Forward"),
						app.Button().Class("secondary-button pure-button").Text("Move to"),
					),
				),
				app.Div().Class("email-content-body").Body(
					app.P().Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."),
					app.P().Text("Duis aute irure dolor in reprehenderit in voluptate velit essecillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
					app.P().Text("Aliquam ac feugiat dolor. Proin mattis massa sit amet enim iaculis tincidunt. Mauris tempor mi vitae sem aliquet pharetra. Fusce in dui purus, nec malesuada mauris. Curabitur ornare arcu quis mi blandit laoreet. Vivamus imperdiet fermentum mauris, ac posuere urna tempor at. Duis pellentesque justo ac sapien aliquet egestas. Morbi enim mi, porta eget ullamcorper at, pharetra id lorem."),
					app.P().Text("Donec sagittis dolor ut quam pharetra pretium varius in nibh. Suspendisse potenti. Donec imperdiet, velit vel adipiscing bibendum, leo eros tristique augue, eu rutrum lacus sapien vel quam. Nam orci arcu, luctus quis vestibulum ut, ullamcorper ut enim. Morbi semper erat quis orci aliquet condimentum. Nam interdum mauris sed massa dignissim rhoncus."),
					app.P().Body(
						app.Text("Regards,"),
						app.Br(),
						app.Text("Tilo"),
					),
				),
			),
		),
	)
}

func main() {
	app.Route("/", &mainLayout{})
	app.Run()
}
