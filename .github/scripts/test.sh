#!/bin/bash

echo $GITHUB_REPOSITORY
type=$GITHUB_EVENT_NAME
file=$GITHUB_EVENT_PATH
before=$GITHUB_EVENT_BEFORE
cat $file
if [ $type == "pull_request" ]; then
  num=$(cat $file | jq -r .pull_request.changed_files)
  echo $num
  if [ $num = 11 ]; then
    echo pass
  else
    echo fail
  fi
fi

