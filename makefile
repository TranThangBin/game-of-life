clean:
	rm ./build/game_of_life 2>/dev/null || true

.PHONY: build

build:
	go build -o ./build/game_of_life ./cmd/game_of_life/main.go

test:
	go test ./internal/game
