.PHONY: all clean

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

BUILD=$(shell git rev-parse --short=7 HEAD)
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

all:
	go build -o $(SERVICENAME) $(LDFLAGS)

clean:
	-rm $(SERVICENAME) \
        && docker rmi "carprks/$(SERVICENAME):$(VERSION)" \
        && docker rmi "carprks/$(SERVICENAME):latest"

docker:
	docker build -t "carprks/$(SERVICENAME):$(VERSION)" \
		--build-arg build=$(BUILD) --build-arg version=$(VERSION) --build-arg serviceName=$(SERVICENAME) \
		-f Dockerfile .

