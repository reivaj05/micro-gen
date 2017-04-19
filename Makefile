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
	./scripts/tests.sh
deps:
	./scripts/deps.sh