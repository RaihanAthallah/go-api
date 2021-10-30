MODULE := go_listofclinics_api
PACKAGES := $(shell go list ./... | grep -v /vendor/)

.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out

.PHONY: build
build: ## build and run for development
	docker-compose -f docker-compose.dev.yml up --build

.PHONY: build-prod
build-prod: ## build and run production version
	docker-compose -f docker-compose.yml up --build -d

.PHONY: run
run: ## run development version (hot reload)
	docker-compose -f docker-compose.dev.yml up 

.PHONY: run-prod
run-prod: ## run production version
	docker-compose -f docker-compose.yml up -d

.PHONY: stop
stop: ## stop docker development version
	docker-compose -f docker-compose.dev.yml stop 

.PHONY: stop-prod
stop-prod: ## stop docker production version
	docker-compose -f docker-compose.yml stop

# If the first argument is "migrate"
ifeq (migrate,$(firstword $(MAKECMDGOALS)))
  MIGRATE_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(MIGRATE_ARGS):;@:)
endif
.PHONY: migrate
migrate: ## Run migration tools
	docker-compose -f docker-compose.dev.yml exec app migrate $(MIGRATE_ARGS)