.PHONY:
.SILENT:

build: -o ./.bin/bot cmd/main.go

run: build
     ./bin/bot