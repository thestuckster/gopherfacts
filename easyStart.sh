#!/usr/bin/env bash

if [ -z "$TOKEN" ]; then
  echo "Error: TOKEN environment variable is not set."
  exit 1
fi

git pull
go build .
./gopherfacts
