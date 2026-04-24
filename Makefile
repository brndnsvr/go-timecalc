BINARY := timecalc
PREFIX ?= /usr/local

.PHONY: build install-local test clean

build:
	go build -o $(BINARY) .

install-local: build
	install -d $(PREFIX)/bin
	install -m 0755 $(BINARY) $(PREFIX)/bin/$(BINARY)

test:
	go test ./...

clean:
	rm -f $(BINARY)
