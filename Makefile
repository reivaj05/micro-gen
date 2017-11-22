GOCMD = go
PKG   = ./...

default: fmt deps linter test build

fmt:
	$(GOCMD) fmt $(PKG)
deps:
	./scripts/deps.sh
linter:
	./scripts/linter.sh
test: deps
	./scripts/tests.sh
build: deps
	$(GOCMD) install