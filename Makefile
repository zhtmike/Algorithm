all: install

.PHONY: install
install:
	go install ./...

.PHONY: clean
clean:
	go clean ./...

.PHONY: test
test:
	go test ./...
