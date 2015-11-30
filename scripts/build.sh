#!/bin/sh

set -o pipefail

ver=$1

if [[ -z "$ver" ]]; then
	echo "usage: $0 <version>"
	exit 1
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
OUTPUT_DIR="${DIR}/../builds/${ver}"

mkdir -p $OUTPUT_DIR

xgo \
	--dest $OUTPUT_DIR \
	--targets='windows/*,darwin/*,linux/*' \
	-out doit-0.6.0 github.com/bryanl/doit/cmd/doit

for f in $(find ${OUTPUT_DIR} -maxdepth 1 -perm -111 -type f); do
	echo "generating sha256 checksum for $f"
	openssl dgst -sha256 < $f > ${f}.sha256
done

