include .env_example
export

WORK_DIR = $(PWD)
EXEC_DIR ?= ${WORK_DIR}/gobin

all: install-tools

.PHONY:
install-tools: export GOBIN=$(EXEC_DIR)
install-tools: 
	@mkdir -p $(EXEC_DIR)
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
	@go install mvdan.cc/gofumpt@v0.6.0
	@go install github.com/vektra/mockery/v2@v2.43.1


.PHONY:
lint:
	@${EXEC_DIR}/golangci-lint run ./...

.PHONY:
install-modules:
	@go mod download

.PHONY:
test:
	@go test -v ./...

.PHONY:
fmt-imports:
	@gofumpt -l -w .

.PHONY:
local-run: install-modules
	@go run ./cmd/main.go

.PHONY:
build: install-modules lint test
	@go build ./cmd/main.go

.PHONY:
docker-run:
	@docker-compose run --rm rover_app