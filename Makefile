PKGSWITHEXAMPLES := $(shell go list ./...)
PKGSWITHOUTEXAMPLES := $(shell go list ./... | grep -v 'examples/')
TXT_FILES := $(shell find * -type f -not -path 'vendor/**')
TESTFLAG=-race -cover

check: vet lint misspell staticcheck

vet:
	@echo "vet"
	go vet $(PKGSWITHOUTEXAMPLES)

lint:
	@echo "golint"
	golint -set_exit_status $(PKGSWITHOUTEXAMPLES)

misspell:
	@echo "misspell"
	misspell -source=text -error $(TXT_FILES)

staticcheck:
	@echo "staticcheck"
	staticcheck $(PKGSWITHOUTEXAMPLES)

test:
	go test -v -cover $(PKGSWITHOUTEXAMPLES)