.PHONY:
.SILENT:

build: -o ./.bin/bot cmd/main.go

run: build
     ./bin/bot
build-image:

start-container:
      docker run --name telegram-bot -p 80:80 --env-file .env