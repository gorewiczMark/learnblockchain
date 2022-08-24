#!/bin/bash

echo $GITHUB_REPOSITORY
type=$GITHUB_EVENT_NAME
if [ $type == "pull_request" ]; then
  echo $GITHUB_EVENT.NUMBER
fi
cat $GITHUB_EVENT_PATH
