.PHONY: clean
clean:
	\rm -rf bin/*

build: clean
	go build -o ./bin/gitsast main.go

run:
	go run main.go

start:
	./bin/gitsast

build-all: clean
	GOOS=linux GOARCH=amd64 go build -o ./bin/gitsast-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o ./bin/gitsast-linux-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build -o ./bin/gitsast-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o ./bin/gitsast-darwin-arm64 main.go

deps-cleancache:
	go clean -modcache

mock:
	mockgen -source=internal/repository/service.go \
		-package testutil \
		-destination=testutil/mocks/repository/service.go
	mockgen -source=internal/model/repository.go \
		-package testutil \
		-destination=testutil/mocks/model/repository.go
	mockgen -source=internal/model/report.go \
		-package testutil \
		-destination=testutil/mocks/model/report.go
	mockgen -source=internal/model/rule.go \
		-package testutil \
		-destination=testutil/mocks/model/rule.go
	mockgen -source=internal/queue/handler.go \
		-package testutil \
		-destination=testutil/mocks/queue/handler.go
	mockgen -source=internal/queue/task/analyzer/git/client.go \
		-package testutil \
		-destination=testutil/mocks/queue/analyzer/client.go
	mockgen -source=internal/queue/task/analyzer/detector.go \
		-package testutil \
		-destination=testutil/mocks/queue/analyzer/detector.go
	mockgen -source=internal/queue/task/analyzer/scanner.go \
		-package testutil \
		-destination=testutil/mocks/queue/analyzer/scanner.go
	