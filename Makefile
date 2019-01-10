.PHONY: errgen

PKGS := $(shell go list ./... | grep -v go-errgen)

check: test lint vet fmt-check

test: errgen
	go test -v $(PKGS)

lint:
	golint $(PKGS)

vet:
	go vet $(PKGS)

fmt-check:
	gofmt -l -s *.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi; \
	goimports -l *.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi \

fmt:
	gofmt -w -s *.go
	goimports -w *.go

installdeps:
	GO111MODULE=on go mod vendor
	GO111MODULE=on go mod tidy

bootstrap: installdeps
	git submodule init
	git submodule update
	(cd go-errgen && git checkout 1.3.0)

build-errgen:
	if [ ! -f author/bin/errgen ]; then \
		go build -o author/bin/errgen go-errgen/cmd/errgen/errgen.go; \
		fi

errgen: build-errgen
	./author/errgen.sh
