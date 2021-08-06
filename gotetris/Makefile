BINARY_NAME=gotetris
GOFILES=$(wildcard *.go)

## help: show help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## run: run this program
run:
	go run ${GOFILES}

## build: build this program
build:
	go build -o ${BINARY_NAME} -ldflags "-s -w" ${GOFILES}

## fmt: format go files
fmt:
	go fmt ./...

## clean: clean
clean:
	go clean
	rm -f ${BINARY_NAME}

.PHONY: help run build fmt clean
