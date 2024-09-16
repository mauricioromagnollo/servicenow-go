# =====================================
# GENERAL COMMANDS
# =====================================
help: ## Print available commands
	$(info ========================================)
	$(info Available Commands:)
	@grep '^[[:alnum:]_-]*:.* ##' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS=":.* ## "}; {printf "\t%-25s %s\n", $$1, $$2};'
	$(info ========================================)
	$(info make <command>)
	$(info )
.PHONY: help

clear: ## Stop containers, remove images, networks, and volumes
	@docker compose down --rmi all --volumes --remove-orphans
.PHONY: clear

lint: ## Run linter
	docker run -t --rm -v ${PWD}/:/app -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v
.PHONY: lint

test: ## Run all tests
	go test -v ./...
.PHONY: tests

test-coverage: ## Run test coverage
	go test -v -race -coverprofile=cover.out ./...
	go tool cover -func=cover.out
.PHONY: test-coverage

test-coverage-web: ## Run test coverage and show in browser
	go test -v -race -coverprofile=cover.out ./... && go tool cover -html=cover.out
.PHONY: test-coverage-web

test-race: # Run data race tests
	go test -race ./...
.PHONY: test-race
