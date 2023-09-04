.PHONY: build run install

BINARY=crypto-ss
OUTDIR=bin

all: build

build:
	@echo "Building..."
	@mkdir -p $(OUTDIR)
	@go build -o $(OUTDIR)/$(BINARY) cmd/main.go

run:
	@echo "Running..."
	@go run cmd/main.go

install:
	@echo "Installing..."
	@cp $(OUTDIR)/$(BINARY) /usr/local/bin/
