// +build wasm

// The UI is running only on a web browser. Therefore, the build instruction
// above is to compile the code below only when the program is built for the
// WebAssembly (wasm) architecture.

package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pelly-ryu/minim/app/internal/component"
)

func main() {
	layout := &component.MainLayout{}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Go program panic:", r)
			app.Window().Call("alert", "Fatal error occurred. Minim stops here.\nmsg:"+fmt.Sprint(r))
		}
	}()

	app.Route("/", layout)
	app.Run()
}
