# PPC Roadmap

This roadmap defines a **clean, low-risk path from v0.1 → v0.2+**, prioritizing determinism, maintainability, and long‑term leverage.

It is intentionally ordered to preserve a working compiler at every step.

---

## Guiding Principles

All roadmap items must satisfy:

- identical inputs → identical outputs
- no hidden behavior
- no implicit overrides
- explicit failure over silent resolution
- zero runtime LLM logic

If an item violates these, it does not belong in PPC.

---

# Phase 0 — Baseline (Current State)

Status: **complete** ✅

PPC currently provides:

- deterministic prompt compilation
- module frontmatter parsing
- keyed tag enforcement
- exclusive group rules
- transitive `requires`
- circular dependency detection
- deterministic ordering
- explain mode (`--explain`)
- repo linting (`doctor`)
- machine-readable lint (`doctor --json`)

This is the stability anchor.

No refactors begin until this phase passes golden tests.

---

# Phase 1 — v0.2 Structural Refactor (No Behavior Changes)

**Goal:** prepare PPC for growth without changing output.

> Output must remain byte-identical.

---

## 1.1 Extract Internal Packages

### Motivation

`main.go` currently handles:

- CLI parsing
- module loading
- dependency resolution
- validation
- doctor logic
- rendering

This limits testability and maintainability.

---

### Target Layout

```
cmd/build-prompt/
  main.go

internal/
  model/
    module.go
    rules.go

  loader/
    frontmatter.go
    load.go

  resolver/
    requires.go
    ordering.go
    tags.go

  render/
    render.go

  doctor/
    doctor.go
    report.go
```

---

### Rules

- no logic changes
- no new features
- no output differences
- functions move verbatim

Golden snapshot tests must pass unchanged.

---

## 1.2 Introduce Core Compile API

### Public boundary

```go
type CompileOptions struct {
    Mode     string
    Traits   []string
    Contract string
    Vars     map[string]string
    PromptsDir string
}

func Compile(opts CompileOptions) (string, error)
```

CLI becomes a thin wrapper.

---

### Benefits

- enables editor tooling
- simplifies testing
- allows future binaries
- enforces clean architecture

---

## 1.3 Deterministic Ordering Tests

Add explicit tests asserting ordering precedence:

1. layer
2. priority
3. module id

Any deviation must break a test.

---

# Phase 2 — Patch Series (Incremental Improvements)

Each item should be its own commit or patch.

---

## 2.1 Golden Snapshot Tests

Add byte-for-byte output fixtures:

```
tests/fixtures/
  explore_conservative.md
  ship_code.md
  explore_creative.md
```

Used by:

```bash
go test ./...
```

---

## 2.2 Error Context Enrichment

Enhance errors with:

- file path
- module id
- offending tag or require

Example:

```
prompts/traits/creative.md
  risk:high
conflicts with
prompts/traits/conservative.md
  risk:low
```

---

## 2.3 Doctor Enhancements

Add optional statistics block (JSON only):

```json
"stats": {
  "modules": 18,
  "dead": 2,
  "groups": 3,
  "tags": 7
}
```

Human output remains unchanged.

---

## 2.4 Provenance Header (Optional)

Optional compiled header:

```md
<!--
compiled-from:
  - base
  - modes/ship
  - traits/conservative
  - contracts/code
-->
```

Disabled by default.

---

# Phase 3 — Ecosystem Leverage

Only after Phase 2 stabilizes.

---

## 3.1 Profiles

```
profiles/ship.yml
profiles/explore.yml
```

Example:

```yaml
mode: ship
traits:
  - conservative
  - terse
contract: code
revisions: 1
```

CLI:

```bash
build-prompt --profile ship
```

---

## 3.2 Graph Output

```
build-prompt doctor --graph > deps.dot
```

Visualizes:

- requires graph
- unreachable modules
- cycles

---

## 3.3 GitHub Actions Template

Provide reusable workflow snippet:

- run doctor
- fail on strict
- artifact JSON output

---

# Phase 4 — Distribution

- static releases
- Arch PKGBUILD
- checksum verification

---

# Explicit Non‑Goals

PPC must not include:

- DSLs
- templating engines
- conditionals or loops
- LLM providers
- execution engines
- agent orchestration

Those belong elsewhere.

---

# Success Criteria

PPC is successful when:

- prompts are versioned like code
- behavior is reproducible
- diffs are meaningful
- prompt repos can be linted
- CI can enforce policy

---

## North Star

> Prompts should not be written — they should be compiled.

---

