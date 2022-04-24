## start migration
export

## end 

VERSION = $(shell git branch --show-current)

help:  ## show this help
	@echo "usage: make [target]"
	@echo ""
	@egrep "^(.+)\:\ .*##\ (.+)" ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

run: ## run it will instance server 
	VERSION=$(VERSION) go run main.go

run-watch: ## run-watch it will instance server with reload
	VERSION=$(VERSION) nodemon --exec go run main.go --signal SIGTERM

mock:
	rm -rf ./mocks
	~/go/bin/mockgen -source=./app/transaction/transaction.go -destination=./mocks/transaction_app_mock.go -package=mocks -mock_names=App=TransactionApp &

migrateup:
	migrate -path db_pismo/db/migration -database "mysql://go_test:pismo123@tcp(localhost:3306)/pismo?multiStatements=true" -verbose up

migratedown:
	migrate -path db_pismo/db/migration -database "mysql://go_test:pismo123@tcp(localhost:3306)/pismo?multiStatements=true" -verbose down