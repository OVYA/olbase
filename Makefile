.PHONY: all
all:glide-install test

.PHONY: glide-clean
glide-clean:
	@glide install

.PHONY: glide-install
glide-install:
	@glide install

.PHONY: test
test:
	go test -v --race `go list ./... | grep -v 'vendor'`
