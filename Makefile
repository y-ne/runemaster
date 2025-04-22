APP_NAME = runemaster

.PHONY: build run dev clean

build:
	go build -o bin/$(APP_NAME) cmd/server/main.go

run:
	./bin/$(APP_NAME)

dev:
	go run cmd/server/main.go

clean:
	rm -rf bin/
