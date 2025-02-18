.PHONY: test test-verbose coverage build install snapshot

test:
	go test ./src/commands/...

test-verbose:
	go test -v ./src/commands/...

coverage:
	go test -coverprofile=coverage.out ./src/commands/...
	go tool cover -html=coverage.out

build:
	cd src && go build -o ../bin/stm

install:
	go install

snapshot:
	goreleaser release --snapshot --clean

# Run all tests and show coverage
test-all: test coverage 