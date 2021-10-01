# // как я понял, когда мы пишем make в консоли, булут выполняться
# // все "команды" которые прописаны ниже

.PHONY: build
build:
		go build -v ./cmd/serverwithapi


.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
