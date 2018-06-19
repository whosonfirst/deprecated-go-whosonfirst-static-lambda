CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test ! -d src; then mkdir src; fi
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/algnhsa"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-static"
	mv src/github.com/whosonfirst/go-whosonfirst-static/vendor/github.com/whosonfirst/go-whosonfirst-readwrite-s3 src/github.com/whosonfirst/

# if you're wondering about the 'rm -rf' stuff below it's because Go is
# weird... https://vanduuren.xyz/2017/golang-vendoring-interface-confusion/
# (20170912/thisisaaronland)

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/wof-staticd-lambda cmd/wof-staticd-lambda.go

lambda: self
	if test -f main; then rm -f main; fi
	if test -f deployment.zip; then rm -f deployment.zip; fi
	@GOPATH=$(GOPATH) GOOS=linux go build -o main cmd/wof-staticd-lambda.go
	zip deployment.zip main
	rm -f main