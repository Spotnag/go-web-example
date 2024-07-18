all: build

build: templ
	@go build -o bin/app cmd/main.go

templ:
	templ generate

run:
	templ generate --watch --proxy="http://172.0.0.1:3000" --open-browser=false