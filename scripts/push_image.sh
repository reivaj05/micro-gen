#!/bin/bash

IMAGE_NAME=micro-gen
BRANCH_NAME=$TRAVIS_BRANCH

if [ $BRANCH_NAME == "master" ]; then
	BRANCH_NAME="latest"
fi

if [ -z $DOCKER_USERNAME ]; then
	echo "Missing DOCKER_USERNAME env var"
	exit 1
fi
if [ -z $DOCKER_PASSWORD ]; then
	echo "Missing DOCKER_PASSWORD env var"
	exit 1
fi

DOCKER_IMAGE=$IMAGE_NAME:$BRANCH_NAME
REPO_IMAGE=$DOCKER_USERNAME/$DOCKER_IMAGE

docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
docker push $REPO_IMAGE
