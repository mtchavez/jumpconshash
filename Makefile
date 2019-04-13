ROOT := $(CURDIR)
GOPKGS = \
		golang.org/x/tools/cmd/cover \
		github.com/golang/lint/golint \
		github.com/golang/dep/cmd/dep

default: test

deps:
	@go get -u -v $(GOPKGS)
	@if [ -z $(GO111MODULE) ]; then dep ensure; else go build; fi

lint:
	@echo "[Lint] running golint"
	@golint

vet:
	@echo "[Vet] running go vet"
	@go vet

ci: deps vet lint test

test:
	@echo "[Test] running tests"
	@go test -v ./... -cover -coverprofile=c.out

.PHONY: default golint test vet deps
