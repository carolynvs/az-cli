PKG ?= github.com/carolynvs/az-cli
GOPATH ?= $(shell go env GOPATH)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
VERSION ?= $(shell git describe --match=v* --tags --abbrev=7 --dirty 2> /dev/null || echo v0.0.0)

ALL_GOOS = darwin linux windows

ifeq ($(GOOS),windows)
FILE_EXT=.exe
else
FILE_EXT=
endif

build:
	$(MAKE) build-for-$(GOOS)
	cp bin/az-$(GOOS)-$(GOARCH)$(FILE_EXT) bin/az$(FILE_EXT)

build-all: $(addprefix build-for-,$(ALL_GOOS))
build-for-%:
	$(MAKE) GOOS=$* GOARCH=$(GOARCH) VERSION=$(VERSION) xbuild

xbuild:
	mkdir -p bin/
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go build -a -tags netgo -installsuffix netgo \
		-ldflags '-X $(PKG)/pkg.Version=$(VERSION)' \
		-o bin/az-$(GOOS)-$(GOARCH)$(FILE_EXT) .

clean:
	-rm -r bin/

install: build
	cp bin/az$(FILE_EXT) $(GOPATH)/bin/

.PHONY: test
test: build
	./test/test.sh
