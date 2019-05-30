#!/usr/bin/env bash

if [ -z "$TRAVIS_PULL_REQUEST" ] || [ "$TRAVIS_PULL_REQUEST" == "false" ]; then
    if [ "$TRAVIS_BRANCH" == "master" ]; then
        docker --version
        pip install --user awscli
        export PATH=$PATH:$HOME/.local/bin
        eval $(aws ecr get-login --no-include-email --region $AWS_DB_REGION)
        docker tag "$APP:latest" "$AWS_ECR/$APP:$TRAVIS_COMMIT"
        docker tag "$APP:latest" "$AWS_ECR:$APP:latest"
        docker push "$AWS_ECR/$APP:$TRAVIS_COMMIT"
        docker push "$AWS_ECR/$APP:latest"
    fi
fi
