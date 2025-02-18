.PHONY: test test-verbose coverage build install

test:
	go test ./...

test-verbose:
	o test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

build:
	go build -o stm

install:
	go install

# Run all tests and show coverage
test-all: test coverage 