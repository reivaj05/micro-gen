#!/bin/bash

IMAGE_NAME=micro-gen
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
if [ $BRANCH_NAME == "master" ]; then
	BRANCH_NAME="latest"
fi
DOCKER_IMAGE=$IMAGE_NAME:$BRANCH_NAME

REPO_NAME=javiersv05
REPO_IMAGE=$REPO_NAME/$DOCKER_IMAGE

if [ -z $DOCKER_USERNAME ]; then
	echo "Missing DOCKER_USERNAME env var"
	exit 1
fi
if [ -z $DOCKER_PASSWORD ]; then
	echo "Missing DOCKER_PASSWORD env var"
	exit 1
fi

docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
docker push $REPO_IMAGE
