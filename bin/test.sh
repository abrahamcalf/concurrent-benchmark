#!/bin/bash

# https://superuser.com/a/1310112/859165
time=$(
  (
    time ./bin/concurrent-benchmark \
      'node fixtures/timeout.js' \
      'node fixtures/timeout.js' \
      'node fixtures/timeout.js' \
      'node fixtures/timeout.js'
  ) 2>&1 >/dev/null
)

# time taken `0m2.330s`
time_taken=$(echo $time | grep real | awk '{print $2}')

# extract seconds
time_taken_seconds=$(echo $time_taken | cut -d '.' -f 1 | cut -d 'm' -f 2)

if [ $time_taken_seconds -eq 2 ]; then
  echo "PASS"
  exit 0
else
  echo "FAILED"
  exit 1
fi
