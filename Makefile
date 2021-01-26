
# Go parameters set in order to be used accross the Makefile.
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BUILDNAME=doqu

# version linker
SELF := $(shell pwd | sed 's/.*github.com/github.com/')
LDFLAGS := -X github.com/stnrd/doqu-cli/internal/build.Version=$(shell git describe --tags)

build:
	@rm -rf dist/bin
	@echo "Building Doqu CLI -> ./dist/bin"
	$(GOBUILD) -ldflags '$(LDFLAGS)' -gcflags "-N -l" -o ./dist/bin/${BUILDNAME} ./cmd/doqu
	@ls -l ./dist/bin

all: vendor clean build test
	echo "binary file is located under ${BUILDNAME}"

vendor:
	go env -w GOPRIVATE="github.com/stnrd"
	@echo "settings OK"

test:
	$(GOTEST) -v ./... -ldflags '$(LDFLAGS)'

clean:
	if [ -f ${BUILDNAME} ] ; then rm ${BUILDNAME} ; fi

## to be manually invoked to trigger a patch release e.g. 2.4.4 -> 2.4.5
tag-patch: no-pending-tags
	git fetch --tags
	@export VERSION=$$(git tag --sort=v:refname | tail -1) && \
	if git diff --quiet $$VERSION; then echo "No differences since last tag $$VERSION"; exit 1; fi && \
	export PATCH=$$(echo $$VERSION | awk -F. '{ print $$3}') && \
	export MINOR=$$(echo $$VERSION | awk -F. '{ print $$2}') && \
	export MAJOR=$$(echo $$VERSION|tr -d "v"| awk -F. '{ print $$1}') && \
	export NEW_VERSION=v$${MAJOR}.$${MINOR}.$$((PATCH+1)) && \
	echo "$$VERSION -> $$NEW_VERSION" && \
	git tag $$NEW_VERSION && \
	echo "Push tags with: git push --tags"

## to be manually invoked to trigger a minor release e.g. 2.4.4 -> 2.5.0
tag-minor: no-pending-tags
	git fetch --tags
	@export VERSION=$$(git tag --sort=v:refname | tail -1) && \
	if git diff --quiet $$VERSION; then echo "No differences since last tag $$VERSION"; exit 1; fi && \
	export PATCH=$$(echo $$VERSION | awk -F. '{ print $$3}') && \
	export MINOR=$$(echo $$VERSION | awk -F. '{ print $$2}') && \
	export MAJOR=$$(echo $$VERSION|tr -d "v"| awk -F. '{ print $$1}') && \
	export NEW_VERSION=v$${MAJOR}.$$((MINOR+1)).0 && \
	echo "$$VERSION -> $$NEW_VERSION" && \
	git tag $$NEW_VERSION && \
	echo "Push tags with: git push --tags"

## to be manually invoked to trigger a major release e.g. 2.4.4 -> 3.0.0
tag-major: no-pending-tags
	git fetch --tags
	@export VERSION=$$(git tag --sort=v:refname | tail -1) && \
	if git diff --quiet $$VERSION; then echo "No differences since last tag $$VERSION"; exit 1; fi && \
	export PATCH=$$(echo $$VERSION | awk -F. '{ print $$3}') && \
	export MINOR=$$(echo $$VERSION | awk -F. '{ print $$2}') && \
	export MAJOR=$$(echo $$VERSION|tr -d "v"| awk -F. '{ print $$1}') && \
	export NEW_VERSION=v$$((MAJOR+1)).0.0 && \
	echo "$$VERSION -> $$NEW_VERSION" && \
	git tag $$NEW_VERSION && \
	echo "Push tags with: git push --tags"

.PHONY: all vendor build test clean tag-patch tag-minor tag-major no-pending-tags
