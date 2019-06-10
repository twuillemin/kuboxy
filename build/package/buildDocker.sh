#!/bin/bash

# Keep the original directory
ORIGIN_DIR=$(pwd)

# Warning this won't work if the script is under symlinked path, but that is not the case here
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Go the root of the project
cd ../..

# Define the name of the app
declare -r IMAGE_NAME="uxxu/kuboxy"
declare -r IMAGE_TAG="latest"

echo "Building image '$IMAGE_NAME:$IMAGE_TAG'..."
docker build -t ${IMAGE_NAME}:${IMAGE_TAG} -f build/package/Dockerfile .

# Go back to the origin dir
cd $ORIGIN_DIR
