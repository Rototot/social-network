
PWD=$(pwd)

help:
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

init:	## init app
	cp backend/env.example backend/.env

start: ## start app

build: backend-generate backend-build frontend-build	## build

migrate-create: ## create migrate
	migrate create -ext sql -dir backend/db/migrate "$@"

migrate-up:	## run migrate
	cd ./backend && ./bin/social-network migrate

backend-generate:
	cd ./backend && go generate ./...

backend-build:
	cd ./backend && go build -o bin/social-network ./

frontend-build:
	#cd ./frontend && npm install

frontend-lint: ## frontend lint
	cd ./frontend && npm run lint

frontend-lint-fix: ## frontend lint fix
	cd ./frontend && npm run lint:fix

test-e2e: ## e2e tests
	./scripts/run-test-e2e.sh

test-units: ## units tests
	./scripts/run-test-e2e.sh
