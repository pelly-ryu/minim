# minim

> "Every deep thinker is more afraid of being understood than of being misunderstood" - Friedrich Nietzsche

A cross-platform syncing note app for who has secrets, makes mistakes, and has no time.

## How To Build

*This can be changed in future*

### Requirements
- go 1.14 or further

### Local Build
```shell script
GOARCH=wasm GOOS=js go build -o web/app.wasm app/main.go
go build server/main.go
PORT=[port number] ./main # if you don't pass $PORT, default is 8000

# visit localhost:[port number] via a web browser
```

### Documentation

https://www.notion.so/modars/Minim-e36bfd7a184c4bafb851b4a079118393
