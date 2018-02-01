.PHONY: clean
clean:
	rm -rf */data/ */*.exe

.PHONY: test
test:
	go test ./...