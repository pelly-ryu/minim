build-wasm:
	GOARCH=wasm GOOS=js go build -o web/app.wasm app/main.go
