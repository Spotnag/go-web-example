all: build

build: templ
	@go build -o bin/app cmd/main.go

templ:
	templ generate

run:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false