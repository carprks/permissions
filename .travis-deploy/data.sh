#!/usr/bin/env bash
DEPLOY_ENV=dev

injectIt()
{
    S3_FOLDER=$S3_BUCKET-$DEPLOY_ENV
    AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY aws s3 cp s3://"$S3_FOLDER"/$SERVICE_NAME/data.json data.json
    for i in $(jq -r '.keys[] | "\(.key),\(.expires),\(.service)"' keys.json); do
        IFS=','
        read -ra SPL <<< "$i"
        AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY aws dynamodb put-item --table-name $TABLE_NAME --item "{\"authKey\":{\"S\":\"${SPL[0]}\"},\"expires\":{\"N\":\"${SPL[1]}\"},\"service\":{\"S\":\"${SPL[2]}\"}}" --region $AWS_REGION
        IFS=' '
    done
}


if [[ -z "$TRAVIS_PULL_REQUEST" ]] || [[ "$TRAVIS_PULL_REQUEST" == "false" ]]; then
    AWS_ACCESS_KEY_ID=$DEV_AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY=$DEV_AWS_SECRET_ACCESS_KEY
    TABLE_NAME="$SERVICE_NAME-dynamo-dev"

    echo "Injecting Dev"
    injectIt
    echo "Injected Dev"

    # Master has an extra step to launch into live
    if [[ -z "$SKIP_LIVE" ]] || [[ "$SKIP_LIVE" == "false" ]]; then
        if [[ "$TRAVIS_BRANCH" == "master" ]]; then
            DEPLOY_ENV=live
            AWS_ACCESS_KEY_ID=$LIVE_AWS_ACCESS_KEY_ID
            AWS_SECRET_ACCESS_KEY=$LIVE_AWS_SECRET_ACCESS_KEY
            TABLE_NAME="$SERVICE_NAME-dynamo-live"

            echo "Injecting Live"
            injectIt
            echo "Injected Live"
        fi
    fi
fi