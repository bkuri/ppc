# GitHub Actions Integration

This document provides copy‑paste‑ready GitHub Actions workflows for using
**PPC (Prompt Policy Compiler)** as a CI linting step for prompt repositories.

PPC runs entirely offline and requires no LLM access.

---

## What this checks

Running:

```bash
ppc doctor --strict --json
```

will fail CI if:

- modules have missing or invalid frontmatter
- `requires` references are missing
- circular dependencies exist
- exclusive tag groups conflict
- unreachable modules are detected (strict mode)

---

## Recommended: Download from Releases

This approach does **not** require Go to be installed.

```yaml
name: Prompt Policy Check

on: [push, pull_request]

jobs:
  doctor:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Install PPC
        run: |
          set -euo pipefail
          VERSION=v0.2.0
          curl -fsSL -o ppc.tar.gz \
            https://github.com/bkuri/ppc/releases/download/${VERSION}/ppc_${VERSION}_linux_amd64.tar.gz
          tar -xzf ppc.tar.gz
          chmod +x ppc
          ./ppc --version

      - name: Prompt policy lint
        run: |
          ./ppc doctor --strict --json --out report.json

      - uses: actions/upload-artifact@v4
        with:
          name: doctor-report
          path: report.json
```

---

## Alternative: Build from source

```yaml
- uses: actions/setup-go@v5
  with:
    go-version: '1.22'

- run: go build -o ppc ./cmd/build-prompt
- run: ./ppc doctor --strict --json --out report.json
```

---

## Optional: Dependency graph artifact

```yaml
- name: Dependency graph
  run: ./ppc doctor --graph --out deps.dot

- uses: actions/upload-artifact@v4
  with:
    name: ppc-graph
    path: deps.dot
```

---

## Summary

- deterministic
- reproducible
- zero-runtime dependencies

Prompts become build artifacts.
