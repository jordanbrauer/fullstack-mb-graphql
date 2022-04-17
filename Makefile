.DEFAULT_GOAL := help
.PHONY: $(MAKECMDGOALS)

help: ## Show this help message
	@printf "\033[33mUsage:\033[0m\n  make [target] [arg=\"val\"...]\n\n\033[33mTargets:\033[0m\n"
	@grep -E '^[-a-zA-Z0-9_\.\/]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[32m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build a binary native to your host, useful for development
	@go build -o ./build/fsmb .

windows: ## Build a 64-bit Windows executable
	@GOOS=windows GOARCH=amd64 go build -o ./build/fsmb.exe .

mac: ## Build a 64-bit ARM (Apple Silicon) binary
	@GOOS=darwin GOARCH=arm64 go build -o ./build/fsmb .

linux: ## Build a 64-bit Linux binary
	@GOOS=linux GOARCH=amd64 go build -o ./build/fsmb .

mod: ## Install modules
	@go mod tidy

serve: ## Serve the application (for development only!)
	@go run main.go

types: ## Generate types from the GraphQL schema
	@go run github.com/99designs/gqlgen generate ./..

slides:
	@lookatme --live --style stata-dark SLIDES.md

present:
	@lookatme --style stata-dark SLIDES.md
