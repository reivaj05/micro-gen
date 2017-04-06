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
	$(GOCMD) test -a -v ./...
deps:
	wget -q https://raw.githubusercontent.com/pote/gpm/v1.4.0/bin/gpm
	chmod +x gpm
	./gpm
	rm gpm