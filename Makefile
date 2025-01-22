ifeq ($(OS),Windows_NT)
  ifeq ($(ComSpec),$(shell echo %ComSpec%))
    GO_BASE := $(shell cd)
  else
    GO_BASE := $(shell cygpath -w $(shell pwd))
  endif
else
  GO_BASE := $(shell pwd)
endif

GO_BIN := $(GO_BASE)/bin

MIGRATION_PATH := ./db/migrations
WIRE_PATH := ./internal/wire

create-migration:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

wire:
	@wire gen $(WIRE_PATH)

build: wire
	@go build -o $(GO_BIN)/api/main cmd/api/main.go

run: build
	@$(GO_BIN)/api/main

clean:
	@rm -rf $(GO_BIN)

.PHONY: build run clean wire create-migration