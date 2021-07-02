all: deps test

deps:
	go build -o /dev/null ./...
	go mod tidy

test:
	go test ./...

test-cov:
	go test -cover -coverprofile=c.out ./...

test-cov-html: test-cov
	go tool cover -html=c.out

lint:
	golint ./...

vet:
	go vet ./...
