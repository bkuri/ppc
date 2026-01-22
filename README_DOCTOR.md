# `build-prompt doctor`

Repo-level validation / lint for PPC prompt modules.

## Usage

```bash
build-prompt doctor
build-prompt doctor --prompts prompts
build-prompt doctor --strict
```

## What it checks (v0.1)

- all module YAML frontmatter parses and has `id`
- all tags are keyed (`group:value`)
- `rules.yml` exists and parses
- missing `requires` targets
- circular `requires` (reports the chain)
- warns about exclusive groups that never appear in any module tags
- warns about modules unreachable from entrypoints (`base`, `modes/*`, `contracts/*`)
