GOCMD = go
PKG   = ./...

.PHONY: %

default: fmt deps test build

all: build
build: deps
	$(GOCMD) install
fmt:
	$(GOCMD) fmt $(PKG)
test: deps
	$(GOCMD) test $(go list ./... | grep -v /vendor/)
deps:
	curl https://glide.sh/get | sh
	glide install