test:
	# -gcflags -m : dosplay compiler decisions about optimisations
	GOPATH=$(GOPATH) go test #-gcflags -m=2

bench:
	GOPATH=$(GOPATH) go test -bench=.

build:
	GOPATH=$(GOPATH) go build

get:
	GOPATH=$(GOPATH) go get

init:
	GOPATH=$(GOPATH) go mod init radix
