#!/bin/bash

# Get the directory of the app
APP_DIR="/path/to/app"
# Set the command to build the app
BUILD_CMD="make run"
# Set the delay to wait before rebuilding the app
BUILD_DELAY=5

# Get the file pattern to ignore
IGNORE_FILE_PATTERN="(\.tmp$)|(\.log$)"

# Watch the directory for changes
inotifywait -mrqe modify --exclude "${IGNORE_FILE_PATTERN}" "${APP_DIR}" |
# Loop through the events
while read -r event; do
  # Log the event
  echo "File changed - ${event}"

  # Wait for the delay before rebuilding the app
  sleep ${BUILD_DELAY} && {
    # Build the app
    echo "Rebuilding..."
    ${BUILD_CMD}
    # Log that the app was rebuilt
    echo "Rebuilt successfully"
  }
done