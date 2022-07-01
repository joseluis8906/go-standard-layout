.PHONY: build
build:
	@if [ -d "./vendor" ]; then echo 'vendor exists' >> /dev/null; else go mod vendor; fi
	@go build -mod=vendor -o bin/go-standard-layout-http cmd/http/main.go

.PHONY: clean
clean:
	@rm -rf bin/*

.PHONY: run
run: build
	@./bin/go-standard-layout-http

.PHONY: test
test:
	@go test ./... -cover
