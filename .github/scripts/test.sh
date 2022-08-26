#!/bin/bash

echo $GITHUB_REPOSITORY
type=$GITHUB_EVENT_NAME
file=$GITHUB_EVENT_PATH
if [ $type == "pull_request" ]; then
  echo $GITHUB_EVENT.NUMBER
  num=$(cat $file | jr -r .pull_request.changed_files)
  echo $num
  if [ $num > 1 ]; then
    echo fail
  fi
fi

