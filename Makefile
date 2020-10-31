GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=portfolio
BINARY_UNIX=$(BINARY_NAME)_unix

.DEFAULT_GOAL := build-and-generate

.PHONY: build-and-generate
build-and-generate: clean build generate

.PHONY: generate
generate:
	./$(BINARY_NAME)

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME)

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX)

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	find public | grep "[^index].html" | xargs rm

.PHOHY: deploy
deploy:
	git subtree push --prefix public origin gh-pages