#!/usr/bin/env bash

function generate() {
  docker run --volume "${PWD}":/workspace \
    --workdir /workspace \
    bufbuild/buf:1.46.0 generate
}

eval "$@"