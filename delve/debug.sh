#!/bin/bash

APP_DIR="/path/to/app"
BUILD_CMD="make run"
BUILD_DELAY=5

IGNORE_FILE_PATTERN="(\.tmp$)|(\.log$)"

inotifywait -mrqe modify --exclude "${IGNORE_FILE_PATTERN}" "${APP_DIR}" |
while read -r event; do
  echo "File changed - ${event}"

  sleep ${BUILD_DELAY} && {
    echo "Rebuilding..."
    ${BUILD_CMD}
    echo "Rebuilt successfully"
  }
done