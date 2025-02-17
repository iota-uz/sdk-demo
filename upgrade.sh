#!/bin/bash

# Repository URL
REPO="https://github.com/iota-uz/iota-sdk"

# Fetch the latest commit hash
LATEST_COMMIT=$(git ls-remote $REPO | grep refs/heads/main | awk '{print $1}')

if [ -z "$LATEST_COMMIT" ]; then
    echo "Failed to fetch the latest commit hash"
    exit 1
fi

echo "Latest commit hash: $LATEST_COMMIT"

# Use go get with the latest commit hash
go get github.com/iota-uz/iota-sdk@$LATEST_COMMIT

# Verify installation
go mod tidy
