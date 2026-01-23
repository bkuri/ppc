#!/usr/bin/env bash
set -e

BUILDDIR=${BUILDDIR:-dist}
CHECKSUM_FILE=$BUILDDIR/checksums.txt

echo "Generating checksums for release archives..."
cd "$BUILDDIR"
sha256sum *.tar.gz > checksums.txt
cat checksums.txt
echo ""
echo "Checksums written to $CHECKSUM_FILE"
