#!make

# default variables
#
NAME=wizworld


# main entries
# 
run:
	@go run cmd/main.go

build:
	@go build -o build/$(NAME) cmd/main.go

.PHONY: run build

# test
#
# unit test
# 
test:
	@go test ./internal/... ./cmd/... -coverprofile cover.out
	go tool cover -func cover.out
	@rm cover.out

.PHONY: test

# mockery
#
# generate mocks for a given interface  | usage: make name=ElixirService mock
mock:
	@mockery --name=$(name) --recursive
