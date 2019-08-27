# Build box
FROM golang:1.12 AS build

RUN mkdir -p /home/main
WORKDIR /home/main

# Get Lint
ENV GO111MODULE=auto
RUN go get -u golang.org/x/lint/golint

# Dependencies
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

# Envs
ARG DB_REGION
ARG DB_ENDPOINT
ARG DB_TABLE
ARG AWS_ACCESS_KEY_ID
ARG AWS_SECRET_ACCESS_KEY
ARG DATABASE_DYNAMO
ARG SERVICE_NAME
ARG SERVICE_DEPENDENCIES
ARG SITE_PREFIX

# Lint and Test
COPY . .
#RUN golint -set_exit_status ./...
#RUN go test -short ./...
#RUN go test -race -short ./...

# Build
ARG build
ARG version
ARG serviceName
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${version} -X main.Build=${build}" -o ${serviceName}
RUN cp ${serviceName} /

# Final
FROM alpine
ARG serviceName
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN apk add curl
RUN rm -rf /var/cache/apk/*

# Move 
COPY --from=build /${serviceName} /home/

# Set TimeZone
ENV TZ=Europe/London

# Entrypoint
WORKDIR /home
ENV _SERVICENAME=${serviceName}
RUN echo "#!/bin/bash" > ./entrypoint.sh
RUN echo "./${serviceName}" >> ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

# EntryPoint
ENTRYPOINT ["sh", "./entrypoint.sh"]

# healthcheck
HEALTHCHECK --interval=5s --timeout=2s --retries=12 CMD curl --silent --fail localhost/probe || exit 1

# Expose Port
EXPOSE 80

