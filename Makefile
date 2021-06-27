#build information
LDFLAGS       := -w -s
BINDIR        := $(CURDIR)/bin
OS            ?= darwin
SERVICE_NAME  ?= checkout

LOG_LEVEL ?= debug

.PHONY: all
all: deps lint test build

.PHONY: deps
deps:
	$(info Installing dependencies)
	go mod download && go mod tidy

.PHONY: lint
lint:
	$(info Running Go code checkers and linters)
	@golangci-lint -v run

.PHONY: test
test:
	$(info Running unit tests)
	go test --cover -covermode=atomic --race -short ./...

.PHONY: build
build:
	$(info Building binary to bin/$(SERVICE_NAME))
	@CGO_ENABLED=0 GOOS=$(OS) go build -o $(BINDIR)/$(SERVICE_NAME) -installsuffix cgo -ldflags '$(LDFLAGS)' ./cmd/$(SERVICE_NAME)

.PHONY: run
run: build
	$(info Running LOG_LEVEL=$(LOG_LEVEL) $(BINDIR)/$(SERVICE_NAME))
	@LOG_LEVEL=$(LOG_LEVEL) $(BINDIR)/$(SERVICE_NAME)