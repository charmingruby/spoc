###################
# Build           #
###################
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/job/main.go

###################
# Testing         #
###################
.PHONY: test
test: 
	go test ./... 

###################
# Linting         #
###################
.PHONY: lint
lint:
	go fmt ./...
	golangci-lint run --fix ./...