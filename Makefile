
PWD=$(pwd)

help:
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

backend-generate:
	cd ./backend
	go generate ./...
	cd $PWD

backend-build: ## go build
	cd ./backend
	go build -o bin/social-network ./
	cd $PWD

frontend-lint: ## frontend lint
	cd ./frontend
	npm run lint
	cd $PWD

frontend-lint-fix: ## frontend lint fix
	cd ./frontend
	npm run lint:fix
	cd $PWD

test-e2e: ## e2e tests
	./scripts/run-test-e2e.sh

test-units: ## units tests
	./scripts/run-test-e2e.sh
