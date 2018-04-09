
.PHONY: test
test: ## test all services
	@./scripts/test

.PHONY: provision
provision: ## provision all services
	@./scripts/provision

.PHONY: deploy
deploy: ## deploy all services
	@./scripts/deploy


.PHONY: build
build: ## buildthe frdm app
	@./scripts/build
