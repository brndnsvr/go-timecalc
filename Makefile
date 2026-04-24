BINARY := timecalc
UNAME_M := $(shell uname -m)
DEFAULT_BINDIR := $(HOME)/.local/bin

ifeq ($(UNAME_M),arm64)
ifneq ($(wildcard /opt/homebrew/bin),)
DEFAULT_BINDIR := /opt/homebrew/bin
endif
else
ifneq ($(wildcard /usr/local/bin),)
DEFAULT_BINDIR := /usr/local/bin
endif
endif

BINDIR ?= $(DEFAULT_BINDIR)

.PHONY: build install-local test clean

build:
	go build -o $(BINARY) .

install-local: build
	install -d $(BINDIR)
	install -m 0755 $(BINARY) $(BINDIR)/$(BINARY)
	@echo "Installed $(BINDIR)/$(BINARY)"

test:
	go test ./...

clean:
	rm -f $(BINARY)
