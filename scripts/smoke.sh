#!/usr/bin/env bash
set -euo pipefail

go build -ldflags="-s -w" -o ppc ./cmd/build-prompt

echo "== list =="
./ppc --list | sed -n '1,20p'

echo "== compile =="
./ppc --conservative --revisions 1 --contract markdown explore >/tmp/ppc_prompt.md
wc -c /tmp/ppc_prompt.md

echo "== explain (should go to stderr) =="
./ppc --conservative --revisions 1 --contract code --explain ship >/dev/null
