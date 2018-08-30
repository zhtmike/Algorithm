.PHONY: install clean test

install:
	go install ./...

clean:
	go clean ./...

test:
	go test ./...
