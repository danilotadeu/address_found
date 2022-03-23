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