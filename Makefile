.PHONY: all clean

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

LDFLAGS := -X "main.Version=$(VERSION)" -X "main.Build=$(shell git rev-parse --short=7 HEAD)"
TAGS ?=
SERVICENAME ?= "identity"
DD := "docker"

.PHONY: all
all:
	go build -o $(SERVICENAME) $(LDFLAGS)

.PHONY: all
clean:
	go clean -i ./...
	rm -rf $(SERVICENAME)
	$(DD) rmi "carprks/$(SERVICENAME):$(VERSION)"
	$(DD) rmi "carprks/$(SERVICENAME):latest"

.PHONY: osx
osx:
	GOOS=darwin go build -v tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)' -o $(SERVICENAME)

.PHONY: docker
docker:
	docker build -t "carprks/$(SERVICENAME):$(VERSION)" \
		--build-arg build=$(BUILD) \
		--build-arg version=$(VERSION) \
		--build-arg serviceName=$(SERVICENAME) \
		--build-arg AWS_DB_REGION=$(AWS_DB_REGION) \
		--build-arg AWS_DB_ENDPOINT=$(AWS_DB_ENDPOINT) \
		--build-arg AWS_DB_TABLE=$(AWS_DB_TABLE) \
		--build-arg AWS_ACCESS_KEY_ID=$(AWS_ACCESS_KEY_ID) \
		--build-arg AWS_SECRET_ACCESS_KEY=$(AWS_SECRET_ACCESS_KEY) \
		--build-arg DATABASE_DYNAMO=$(DATABASE_DYNAMO) \
		--build-arg SERVICE_NAME=$(SERVICENAME) \
		--build-arg SERVICE_DEPENDENCIES=$(SERVICE_DEPENDENCIES) \
		--build-arg SITE_PREFIX=$(SITE_PREFIX) \
		-f Dockerfile .

