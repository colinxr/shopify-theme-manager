.PHONY: test test-verbose coverage build install snapshot

test:
	cd src && go test ./commands/...

test-verbose:
	cd src && go test -v ./commands/...

coverage:
	cd src && go test -coverprofile=coverage.out ./commands/...
	cd src && go tool cover -html=coverage.out

build:
	cd src && go build -o ../bin/stm

install:
	go install

snapshot:
	goreleaser release --snapshot --clean --skip-publish

# Run all tests and show coverage
test-all: test coverage
