
PWD=$(pwd)

help:
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

init:	## init app
	cp backend/env.example backend/.env

start: ## start app

build: go-generate go-build	## build

migrate-create: ## create migrate
	migrate create -ext sql -dir backend/db/migrate "$@"

migrate-up:	## run migrate
	./bin/social-network migrate

go-generate:
	go generate ./...

go-build:
	go build -o bin/social-network ./

lint: ## run linter
	docker-compose -f docker-compose.lint.yaml pull \
		&& docker-compose -f docker-compose.lint.yaml run linter

test-e2e: ## e2e tests
	./scripts/run-test-e2e.sh

test-units: ## units tests
	./scripts/run-test-e2e.sh
