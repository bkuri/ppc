BIN=ppc

all: build

build:
	go build -ldflags="-s -w" -o $(BIN) ./cmd/build-prompt

test:
	go test ./...

smoke: build
	./scripts/smoke.sh

clean:
	rm -f $(BIN)

.PHONY: all build test smoke clean
