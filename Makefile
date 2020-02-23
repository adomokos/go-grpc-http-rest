APP_NAME = grpc-app
CLIENT_NAME = grpc-client
PROG = ${APP_NAME}
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

build: ## Builds application artifacts
	# go build -ldflags="-X main.version=${VERSION} -X main.gitSHA=${GIT_SHA}" -o ${PROG} cmd/server.go
	go build -o $(PROG) cmd/server/main.go
	go build -o $(CLIENT_NAME) cmd/client-grpc/main.go

db.rebuild: ## Rebuilds the DBs
	@echo 'Rebuilding the DB...'
	@sh resources/rebuild-db.sh run
	@echo 'Done.'
.PHONY: rebuild-dbs

db.console: ## Opens the db-console
	sqlite3 db/todo-db.sqlt
.PHONY: db-console

run-server: ## Runs the server
	./$(PROG) -db-file=db/todo-db.sqlt -grpc-port=9090 -http-port=8080
.PHONY: run-server

run-client: ## Runs the client
	./$(CLIENT_NAME) -server=localhost:9090
.PHONY: run-server

gen-files-from-proto: ## Generates swagger JSON and Go files from proto file
	sh third_party/protoc-gen.sh
.PHONY: gen-files-from-proto

get-one-http: ## Executes a CURL request to get a ToDo
	curl -X GET -H "Content-Type: application/json" http://localhost:8080/v1/todo/1
.PHONY: get-one-http

get-all-http: ## Executes a CURL request to get all ToDos
	curl -X GET -H "Content-Type: application/json" http://localhost:8080/v1/todo/all
.PHONY: get-all-http


help: ## Prints this help command
	@grep -E '^[a-zA-Z0-9\._-]+:.*?## .*$$' $(MAKEFILE_LIST) |\
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help
