.PHONY: test test-verbose coverage build install

test:
	go test ./commands/...

test-verbose:
	go test -v ./commands/...

coverage:
	go test -coverprofile=coverage.out ./commands/...
	go tool cover -html=coverage.out

build:
	go build -o stm

install:
	go install

# Run all tests and show coverage
test-all: test coverage 