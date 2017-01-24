.PHONY: all
all:glide-install test

.PHONY: glide-install
glide-install:
	@glide install

.PHONY: test
test:
	go test -v ./...
