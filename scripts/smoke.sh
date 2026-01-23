#!/usr/bin/env bash
set -euo pipefail

go build -ldflags="-s -w" -o ppc ./cmd/build-prompt

echo "== version =="
./ppc --version

echo "== list =="
./ppc --list | sed -n '1,20p'

echo "== compile =="
./ppc explore --conservative --revisions 1 --contract markdown >/tmp/ppc_prompt.md
wc -c /tmp/ppc_prompt.md

echo "== explain (should go to stderr) =="
./ppc ship --conservative --revisions 1 --contract code --explain >/dev/null
