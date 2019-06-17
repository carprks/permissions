#!/usr/bin/env bash
pushIt()
{
    export PATH=$PATH:$HOME/.local/bin
    eval $(AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY aws ecr get-login --no-include-email --region $AWS_REGION)
    docker tag "$SERVICE_NAME:latest" "$AWS_ECR/$SERVICE_NAME:$TRAVIS_COMMIT"
    docker tag "$SERVICE_NAME:latest" "$AWS_ECR/$SERVICE_NAME:latest"
    docker push "$AWS_ECR/$SERVICE_NAME:$TRAVIS_COMMIT"
    docker push "$AWS_ECR/$SERVICE_NAME:latest"
}

buildIt()
{
    docker build -t $SERVICE_NAME:latest \
        --build-arg build=$TRAVIS_COMMIT \
        --build-arg version=$TRAVIS_COMMIT \
        --build-arg serviceName=$SERVICE_NAME \
        -f Dockerfile .
}

if [ -z "$TRAVIS_PULL_REQUEST" ] || [ "$TRAVIS_PULL_REQUEST" == "false" ]; then
    AWS_ECR=$DEV_AWS_ECR
    AWS_ACCESS_KEY_ID=$DEV_AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY=$DEV_AWS_SECRET_ACCESS_KEY

    echo "Push Dev"
    buildIt
    pushIt
    echo "Pushed Dev"

    if [ "$TRAVIS_BRANCH" == "master" ]; then
        if [[ -z "$SKIP_LIVE" ]] || [[ "$SKIP_LIVE" == "false" ]]; then
            AWS_ECR=$DEV_AWS_ECR
            AWS_ACCESS_KEY_ID=$DEV_AWS_ACCESS_KEY_ID
            AWS_SECRET_ACCESS_KEY=$DEV_AWS_SECRET_ACCESS_KEY

            echo "Push Live"
            pushIt
            echo "Pushed Live"
        fi
    fi
fi
