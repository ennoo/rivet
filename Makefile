PKGSWITHOUTEXAMPLES := $(shell go list ./... | grep -v 'examples/\|utils/')
GO_FILES := $(shell find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" -print0 | xargs -0)

export SQLName=circle_test
export DBUser=root
export DBName=mysql
checkTravis:export DB_PORT=33061
checkTravis:export DB_URL=127.0.0.1:${DB_PORT}
checkTravis:export DB_PASS=secret
checkCircle:export DB_PORT=3306
checkCircle:export DB_URL=127.0.0.1:${DB_PORT}
checkCircle:export DB_PASS=
checkLocal:export DB_PORT=3306
checkLocal:export DB_URL=127.0.0.1:${DB_PORT}
checkLocal:export DB_PASS=secret

checkTravis: start overalls vet lint misspell staticcheck cyclo const veralls test end

checkCircle: wright consul overalls vet lint misspell staticcheck cyclo const test end

checkLocal: start overalls vet lint misspell staticcheck cyclo const end

start: wright mysql consul

end:
	@echo "end"
	rm -rf a.txt
	docker rm -f ${SQLName}
	consul leave

mysql:
	@echo "docker run mysql"
	docker run --name ${SQLName} -e MYSQL_ROOT_PASSWORD=${DB_PASS} -d -p ${DBPort}:3306 mysql:5.7.26

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