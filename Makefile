GOCMD = go
PKG   = ./...

default: setup fmt deps linter test install

setup:
	chmod +x ./scripts/*.sh
fmt:
	$(GOCMD) fmt $(PKG)
deps:
	./scripts/deps.sh
linter:
	./scripts/linter.sh
test: deps
	./scripts/tests.sh
install: deps
	$(GOCMD) install