# Prompt Policy Compiler (PPC)

## Product Requirements Document (PRD)
**Version:** v0.1  
**Status:** Draft  
**Primary goal:** Define a minimal, deterministic, high‑performance CLI tool for composing LLM behavior via modular Markdown templates.

---

## 1. Problem Statement

LLM prompts are currently:

- written manually
- duplicated across tools
- difficult to keep consistent
- hard to reason about when layered

Existing tools (e.g. Fabric) focus on **task prompts** — *what the model should do*.

This project targets a different layer:

> **How the model is allowed to behave while doing it.**

The Prompt Policy Compiler (PPC) compiles small, composable Markdown behavior modules into a single deterministic system prompt.

---

## 2. Non‑Goals

This project explicitly does **not**:

- execute LLMs
- manage API keys
- replace prompt libraries
- perform templating logic beyond simple variable substitution
- introduce a DSL
- implement memory, RAG, or orchestration

PPC produces **text only**.

---

## 3. Design Principles

1. **Simplicity over features**
2. **Determinism over convenience**
3. **Composable behaviors, not giant prompts**
4. **Unix‑style stdout first**
5. **Markdown as the only content format**
6. **Minimal metadata**
7. **Fast startup, single binary**

---

## 4. Primary Use Cases

### 4.1 Generate a prompt via CLI

```bash
ppc --conservative --revisions 1 ship
```

Outputs compiled prompt to stdout.

---

### 4.2 Pipe into other tools

```bash
ppc --conservative --revisions 1 ship \
  | fabric --pattern analyze_code
```

---

### 4.3 Write prompt to file

```bash
ppc --creative --out AGENTS.md explore
```

---

### 4.4 Inspect behavior resolution

```bash
ppc --conservative --explain ship
```

---

## 5. High‑Level Architecture

```
CLI flags
   ↓
Module selection
   ↓
Frontmatter parsing
   ↓
Requires expansion
   ↓
Circular dependency detection
   ↓
Tag conflict validation
   ↓
Module ordering
   ↓
Variable substitution
   ↓
Final Markdown output
```

---

## 6. Repository Structure

```
.
├── prompts/
│   ├── base.md
│   ├── modes/
│   │   ├── explore.md
│   │   ├── build.md
│   │   └── ship.md
│   ├── traits/
│   │   ├── conservative.md
│   │   ├── creative.md
│   │   ├── terse.md
│   │   └── verbose.md
│   ├── policies/
│   │   ├── revisions.md
│   │   └── self_score.md
│   ├── contracts/
│   │   ├── markdown.md
│   │   └── code.md
│   └── rules.yml
│
├── cmd/build-prompt/
│   └── main.go
│
└── README.md
```

---

## 7. Module Format

### 7.1 Markdown body

All prompt content is plain Markdown.

---

### 7.2 Optional YAML frontmatter

Frontmatter appears at the top of a module file:

```md
---
id: traits/conservative
desc: Prefer stable, boring solutions.
priority: 50
tags: [risk:low]
requires: [policies/self_score]
---
```

---

### 7.3 Supported frontmatter fields (v0.1)

| Field | Type | Required | Purpose |
|------|------|----------|---------|
| `id` | string | yes | Stable module identifier |
| `desc` | string | no | Shown in `--list` |
| `priority` | int | no | Ordering within layer |
| `tags` | array[string] | no | Behavior classification |
| `requires` | array[string] | no | Dependency inclusion |

---

## 8. Tags System

Tags define behavioral properties.

### 8.1 Keyed tag format

```
<group>:<value>
```

Examples:

- `risk:low`
- `risk:high`
- `tone:terse`
- `tone:verbose`
- `output:code`

---

### 8.2 Purpose

Tags are used to:

- detect incompatible behavior combinations
- avoid explicit override lists
- centralize policy enforcement

Modules do **not** define conflicts themselves.

---

## 9. Rules File

Location:

```
prompts/rules.yml
```

---

### 9.1 Minimal v0.1 schema

```yaml
exclusive_groups:
  - risk
  - tone
  - output
```

Meaning:

> Only one tag per group may be present in the final module set.

---

### 9.2 Example conflict

If final modules include:

- `risk:low`
- `risk:high`

Compiler error:

```
conflicting tags in group "risk": low, high
```

---

## 10. Requires Semantics

### 10.1 Definition

```
requires: [module_id]
```

Means:

> If this module is selected, its required modules must also be included.

---

### 10.2 Rules

- dependencies are transitive
- duplicate requires are ignored
- missing module → error
- circular dependencies → error

---

### 10.3 Circular dependency detection

Implemented via DFS with:

- unvisited
- visiting
- done

Cycle errors must report full path:

```
a -> b -> c -> a
```

---

## 11. Resolution Order

1. CLI‑selected modules loaded
2. Frontmatter parsed
3. Requires closure expanded
4. Circular dependency detection
5. Tag conflict validation
6. Final module ordering
7. Markdown concatenation
8. Variable substitution

---

## 12. Module Ordering Rules

Modules are sorted by:

1. layer (folder order)
2. priority (ascending)
3. id (lexicographic)

Layer precedence:

```
base
→ modes
→ traits
→ policies
→ contracts
```

---

## 13. Variable Substitution

Simple string replacement only:

```
{{mode}}
{{revisions}}
```

No logic.
No conditionals.
No templating engine.

---

## 14. CLI Interface

### 14.1 Basic usage

```bash
ppc [flags] <mode>
```

---

### 14.2 Supported flags (v0.1)

| Flag | Description |
|-----|-------------|
| `--conservative` | include conservative trait |
| `--creative` | include creative trait |
| `--terse` | terse output |
| `--verbose` | verbose output |
| `--revisions N` | revision budget |
| `--contract code|markdown` | output contract |
| `--out FILE` | write to file |
| `--list` | list available modules |
| `--explain` | show resolution steps |

---

## 15. Output Behavior

### Default

- writes compiled Markdown to stdout

### With `--out`

- writes output to file
- still prints to stdout unless `--quiet`

---

## 16. Determinism Guarantees

Given:

- same prompt directory
- same CLI flags

The tool must produce:

- identical output text
- identical module order

across machines and executions.

---

## 17. Language & Performance

- Implementation language: **Go**
- Target startup time: < 20 ms
- Memory usage: negligible
- Distribution: single static binary

---

## 18. Future (Out of Scope for v0.1)

- conditional rules
- module phases
- remote module registries
- TOML/JSON frontmatter
- prompt diffing
- policy versioning
- execution backends

---

## 19. Summary

The Prompt Policy Compiler is:

- a deterministic prompt behavior builder
- composable via small Markdown modules
- governed by centralized policy rules
- compatible with Fabric and other tooling
- fast, inspectable, and boring by design

Prompts are not written.

**They are compiled.**

