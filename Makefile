GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=portfolio
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME)

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX)

.PHONY: run
run: build
	./$(BINARY_NAME)

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f ./public/about.html
	rm -f ./public/projects.html

.PHOHY: deploy
deploy:
	git subtree push --prefix public origin gh-pages