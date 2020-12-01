// +build !wasm

// The server is a classic Go program that can run on various architecture but
// not on WebAssembly. Therefore, the build instruction above is to exclude the
// code below from being built on the wasm architecture.

package main

import (
	"github.com/pelly-ryu/minim/server/internal"
	"log"
	"net/http"
	"os"
	"strconv"
)

// The main function is the entry of the server. It is where the HTTP handler
// that serves the UI is defined and where the server is started.
//
// Note that because main.go and app.go are built for different architectures,
// this main() function is not in conflict with the one in
// app.go.
func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	port := 8000
	if a := os.Getenv("PORT"); a != "" {
		var err error
		port, err = strconv.Atoi(a)
		if err != nil {
			log.Fatalln("wrong port number", a)
		}
	}

	// app.Handler is a standard HTTP handler that serves the UI and its
	// resources to make it work in a web browser.
	//
	// It implements the http.Handler interface so it can seamlessly be used
	// with the Go HTTP standard library.
	http.Handle("/", internal.NewWebHandler())

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}

	//todo
	//fmt.Println("hello world")
	//
	//
	//// init git repo
	//
	//// get a new note
	//
	//// add and commit into git
}
