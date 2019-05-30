#!/usr/bin/env bash
if [ -z "$TRAVIS_PULL_REQUEST" ] || [ "$TRAVIS_PULL_REQUEST" == "false" ]; then
    if [ "$TRAVIS_BRANCH" == "master" ]; then
        bash .travis-deploy/ecs.sh -c $CLUSTER -n $APP -i "$AWS_ECR/$APP:latest" -r $AWS_DB_REGION -t 240
    fi
fi
