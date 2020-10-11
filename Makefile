GOCMD := go
GOBUILD := $(GOCMD) build
GOINSTALL := $(GOCMD) install
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
NAME := wareki
CURRENT := $(shell pwd)
BUILDDIR := ./build
BINDIR := $(BUILDDIR)/bin
PKGDIR := $(BUILDDIR)/pkg
DISTDIR := $(BUILDDIR)/dist

VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X 'main.version=$(VERSION)'
GOXOS := "darwin windows linux"
GOXARCH := "386 amd64"
GOXOSARCH := "darwin/amd64 windows/386 windows/amd64 linux/386 linux/amd64"
GOXOUTPUT := "$(PKGDIR)/$(NAME)_{{.OS}}_{{.Arch}}/{{.Dir}}"

export GO111MODULE=on

.PHONY: deps
## Install dependencies
deps:
	GO111MODULE=off $(GOGET) golang.org/x/tools/cmd/goimports \
	golang.org/x/lint/golint \
	github.com/urfave/cli \
	github.com/Songmu/make2help/cmd/make2help \
	github.com/mitchellh/gox \
	github.com/tcnksm/ghr

.PHONY: build
## Build binaries
build: deps
	$(GOBUILD) -ldflags "$(LDFLAGS)" -o $(BINDIR)/$(NAME)

.PHONY: cross-build
## Cross build binaries
cross-build:
	rm -rf $(PKGDIR)
	gox -osarch=$(GOXOSARCH) -ldflags "$(LDFLAGS)" -output=$(GOXOUTPUT) ./cmd/wareki

.PHONY: package
## Make package
package: cross-build
	rm -rf $(DISTDIR)
	mkdir $(DISTDIR)
	pushd $(PKGDIR) > /dev/null && \
		for P in `ls | xargs basename`; do zip -r $(CURRENT)/$(DISTDIR)/$$P.zip $$P; done && \
		popd > /dev/null

.PHONY: release
## Release package to Github
release: package
	ghr $(VERSION) $(DISTDIR)

.PHONY: install
## compile and install
install:
	$(GOINSTALL) -ldflags "$(LDFLAGS)"

.PHONY: test
## Run tests
test: deps
	$(GOTEST) -v ./...

.PHONY: lint
## Lint
lint: deps
	go vet ./...
	golint -set_exit_status ./...

.PHONY: fmt
## Format source codes
fmt: deps
	find . -name "*.go" -not -path "./vendor/*" | xargs goimports -w

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf $(BUILDDIR)

.PHONY: dockerfile
## Update Dockerfile
dockerfile:
	sed -e "s/<VERSION>/$(VERSION)/g" Dockerfile.base > Dockerfile

.PHONY: help
## Show help
help:
	@make2help $(MAKEFILE_LIST)
