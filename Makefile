PKGSWITHOUTEXAMPLES := $(shell go list ./... | grep -v 'examples/')
GO_FILES := $(shell find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" -print0 | xargs -0)

check: overalls vet lint misspell staticcheck cyclo const veralls

checkLocal: start overalls vet lint misspell staticcheck cyclo const end

start: wright consul

end:
	@echo "end"
	rm -rf a.txt
	consul leave

consul:
	@echo "consul"
	nohup consul agent -dev &

wright:
	@echo "wright"
	echo "this is my test\n" > a.txt

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
	misspell -source=text -error $(GO_FILES)

staticcheck:
	@echo "staticcheck"
	staticcheck $(PKGSWITHOUTEXAMPLES)

cyclo:
	@echo "gocyclo"
	gocyclo -over 10 $(GO_FILES)
#	gocyclo -top 10 $(GO_FILES)

const:
	@echo "goconst"
	goconst $(PKGSWITHOUTEXAMPLES)

veralls:
	@echo "goveralls"
	goveralls -coverprofile=overalls.coverprofile -service=travis-ci -repotoken $(COVERALLS_TOKEN)

test:
	@echo "test"
	go test -v -cover $(PKGSWITHOUTEXAMPLES)