APP_NAME = grpc-app
PROG = ${APP_NAME}
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

.PHONY: build
build: ## Builds application artifacts
	# go build -ldflags="-X main.version=${VERSION} -X main.gitSHA=${GIT_SHA}" -o ${PROG} cmd/server.go
	go build -o ${PROG} cmd/server/main.go

db.rebuild: ## Rebuilds the DBs
	@echo 'Rebuilding the DB...'
	@sh resources/rebuild-db.sh run
	@echo 'Done.'
.PHONY: rebuild-dbs

db.console: ## Open the db-console
	sqlite3 db/todo-db.sqlt
.PHONY: db-console


help: ## Prints this help command
	@grep -E '^[a-zA-Z0-9\._-]+:.*?## .*$$' $(MAKEFILE_LIST) |\
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help
