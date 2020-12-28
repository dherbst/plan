.PHONY:all test build clean install pull lint lint-in-container sec sec-in-container

GOLANG := golang:1.15
GOOS := darwin

VERSION := $(shell git rev-parse --short HEAD)
TODAY := $(shell date "+%Y-%m-%d-%H:%S-%Z")

all: clean pull test build install

pull:
	docker pull $(GOLANG)

clean:
	mkdir -p bin
	rm bin/plan || true

lint:
	docker run -i --rm -v  ${PWD}:/go/src/github.com/dherbst/plan -w /go/src/github.com/dherbst/plan ${GOLANG} make lint-in-container

lint-in-container:
	go get -u golang.org/x/lint/golint
	golint github.com/dherbst/plan/...
	golint github.com/dherbst/plan/cmd/plan/...

sec:
	docker run -it --rm -v ${PWD}:/go/src/github.com/dherbst/plan -w /go/src/github.com/dherbst/plan ${GOLANG} make sec-in-container

sec-in-container:
	go get -u github.com/securego/gosec/cmd/gosec
	gosec .

test:
	docker run -it --rm -v ${PWD}:/go/src/github.com/dherbst/plan -w /go/src/github.com/dherbst/plan ${GOLANG} make test-in-container

test-in-container:
	go test -ldflags "-X github.com/dherbst/plan.Version=TESTVERSION" -coverprofile=coverage.out github.com/dherbst/plan
	go tool cover -html=coverage.out -o coverage.html

build:
	docker run -i --rm -v "$(PWD)":/go/src/github.com/dherbst/plan -w /go/src/github.com/dherbst/plan ${GOLANG} make build-in-container

build-in-container:
	GOOS=darwin go build -o bin/plan -ldflags "-X github.com/dherbst/plan.Version=$(VERSION)-$(TODAY)" cmd/plan/*.go

install:
	cp bin/plan ~/bin
