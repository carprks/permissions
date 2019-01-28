#!/usr/bin/env bash

dockerClean=$(docker ps -a --format '{{.Names}} {{.Status}}' | grep 'Exited' | awk '{print $1}' | xargs docker rm)

# Set the name of the service
dockerPath=$(basename ${PWD##*/} | tr 'A-Z' 'a-z')
if [[ ! -z "$1" ]]; then
    dockerPath=$1
fi

# Set the version of the service
build=$(git rev-parse --short=7 HEAD)
version=1.0.$build
if [[ ! -z "$2" ]]; then
    version=$2
fi

# Export the vars
export SERVICENAME=$dockerPath
export VERSION=$version

# Stop the old image
if [[ -e $(pwd)/docker-compose.yml ]]; then
    docker-compose stop
    yes | docker-compose rm

else
    docker stop $dockerPath
    docker rmi $dockerPath
    dockerClean
fi

# Start the new image
make docker SERVICENAME=$dockerPath VERSION=$version
docker tag carprks/$dockerPath:$version carprks/$dockerPath:latest
if [[ -e $(pwd)/docker-compose.yml ]]; then
    docker-compose build
    docker-compose up -d
else
    docker run -P --rm -d -it --name $dockerPath carprks/$dockerPath:$version
fi
