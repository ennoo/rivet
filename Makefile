PKGSWITHEXAMPLES := $(shell go list ./...)
PKGSWITHOUTEXAMPLES := $(shell go list ./... | grep -v 'examples/')
TXT_FILES := $(shell find * -type f -not -path 'vendor/**')
TESTFLAG=-race -cover
Type := $(shell find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" -print0 | xargs -0)
COVERALLS_TOKEN := ff2BrkJczedGPzmKWFaOBClvTZrJ2b67e

check: overalls vet lint misspell staticcheck cyclo const veralls

overalls:
	@echo "overalls"
	overalls -project=github.com/ennoo/rivet -covermode=count -ignore='.git,_vendor'

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

cyclo:
	@echo "gocyclo"
	gocyclo -over 12 $(Type)

const:
	@echo "goconst"
	goconst $(PKGSWITHOUTEXAMPLES)

veralls:
	@echo "goveralls"
	goveralls -coverprofile=overalls.coverprofile -service=travis-ci -repotoken $COVERALLS_TOKEN

test:
	go test -v -cover $(PKGSWITHOUTEXAMPLES)