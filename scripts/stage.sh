#!/usr/bin/env bash

set -euo pipefail

major=$1
minor=$2
patch=$3
label=${4-}

if [[ -z "$major" || -z "$minor" || -z "$patch" ]]; then
  echo "usage: $0 <major> <minor> <patch> [label]"
  exit 1
fi

ver="${major}.${minor}.${patch}"
if [[ -n "$label" ]]; then
  ver="${ver}-${label}"
fi

RELEASE_PACKAGE=github.com/digitalocean/doctl/cmd/doctl
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

OUTPUT_DIR="${DIR}/../builds/${ver}"
if [[ -d $OUTPUT_DIR ]]; then
  rm -rf "$OUTPUT_DIR"
fi

STAGE_DIR=$OUTPUT_DIR/stage
RELEASE_DIR=$OUTPUT_DIR/release
mkdir -p "$STAGE_DIR" "$RELEASE_DIR"

if [[ -z ${SKIPBUILD-} ]]; then
  echo "building doctl"

  baseFlag="-X github.com/digitalocean/doctl"
  ldflags="${baseFlag}.Build=$(git rev-parse --short HEAD)"
  ldflags="${ldflags} $baseFlag.Major=${major} $baseFlag.Minor=${minor} $baseFlag.Patch=${patch} $baseFlag.Label=release"
  if [[ -n "$label" ]]; then
    ldflags="${ldflags} $baseFlag.Label=${label}"
  fi

  cd "$DIR"/../cmd/doctl

  # ugly, but soon to be replaced by goreleaser
  GO111MODULE=on GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -ldflags "${ldflags}" -o "${STAGE_DIR}/doctl-${ver}-linux-amd64"
  GO111MODULE=on GOOS=linux GOARCH=386 GOFLAGS=-mod=vendor go build -ldflags "${ldflags}" -o "${STAGE_DIR}/doctl-${ver}-linux-386"
  GO111MODULE=on GOOS=darwin GOARCH=amd64 GOFLAGS=-mod=vendor go build -ldflags "${ldflags}" -o "${STAGE_DIR}/doctl-${ver}-darwin-amd64"
  GO111MODULE=on GOOS=windows GOARCH=amd64 GOFLAGS=-mod=vendor go build -ldflags "${ldflags}" -o "${STAGE_DIR}/doctl-${ver}-windows-amd64.exe"
  GO111MODULE=on GOOS=windows GOARCH=386 GOFLAGS=-mod=vendor go build -ldflags "${ldflags}" -o "${STAGE_DIR}/doctl-${ver}-windows-386.exe"
fi

cd "$RELEASE_DIR"

for f in "$STAGE_DIR"/*; do
  distfile_basename=$(basename "${f%".exe"}")

  if [[ $f == *"windows"* ]]; then
    distfile=${distfile_basename}.zip
  else
    distfile=${distfile_basename}.tar.gz
  fi
  
  distbin=$(basename $RELEASE_PACKAGE)
  if [[ $f == *.exe ]]; then
    distbin=$distbin.exe
  fi
  
  bin=$STAGE_DIR/$distbin
  cp "$f" "$bin"
  
  if [[ $distfile == *.zip ]]; then
    zip -j "$distfile" "$bin"
  else
    tar cvzhf "$distfile" -C "$STAGE_DIR" "$distbin"
  fi
  
  pushd "$STAGE_DIR"
  shasum -a 256 "$(basename "$distbin")" > "${RELEASE_DIR}/${distfile_basename}.sha256"
  popd
  
  rm "$bin"
done
