# PPC — Prompt Policy Compiler (v0.1)

PPC compiles small Markdown behavior modules into a single deterministic prompt (stdout-first).

## Build

```bash
go build -o ppc ./cmd/build-prompt
```

## Run

From repo root:

```bash
./ppc --conservative --revisions 1 --contract markdown explore
./ppc --conservative --revisions 1 --contract code --explain ship
./ppc --creative --out AGENTS.md explore
```

**Important**: Flags must come before the mode argument.

## Layout

- `prompts/` contains Markdown modules with optional YAML frontmatter.
- `prompts/rules.yml` defines `exclusive_groups` for keyed tags (`group:value`).

## Notes

- Deterministic output: same inputs -> same output.
- Fails loudly on: missing modules, tag conflicts, circular requires.

