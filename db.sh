#!/usr/bin/env bash
if ! type "localstack" > /dev/null; then
    docker-compose stop db
    docker-clean
    docker ps -a --format '{{.Names}} {{.Status}}' | grep 'Exited' | awk '{print $1}' | xargs docker rm
    docker-compose up -d db
fi
#STACK=$(SERVICES=dynamodb TMPDIR=private$TMPDIR localstack start --docker)

aws dynamodb create-table --table-name permissions --attribute-definitions AttributeName=identity,AttributeType=S --key-schema AttributeName=identity,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 --endpoint-url http://docker.devel:4569
