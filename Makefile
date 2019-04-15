GO           ?= go
STATICCHECK  := $(GOPATH)/bin/staticcheck

.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -o bin/monster

run:
	dep ensure -v
	go run main.go

test:
	dep ensure -v
	go test

clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: staticcheck
staticcheck: $(STATICCHECK)
	@echo ">> running staticcheck"
	$(STATICCHECK) -ignore "$(STATICCHECK_IGNORE)" $(pkgs)

.PHONY: $(STATICCHECK)
$(STATICCHECK):
	GOOS= GOARCH= $(GO) get -u honnef.co/go/tools/cmd/staticcheck