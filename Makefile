
PWD=$(pwd)

help:
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

migrate-create: ## create migrate
	migrate create -ext sql -dir backend/db/migrate "$@"

migrate-up:	## run migrate
	migrate create -ext sql -dir backend/db/migrate

build: backend-generate backend-build frontend-build	## build


backend-generate:
	cd ./backend
	go generate ./...
	cd $PWD

backend-build:
	cd ./backend
	go build -o bin/social-network ./
	cd $PWD

frontend-build:
	cd ./frontend
	npm install
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
