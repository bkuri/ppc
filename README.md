# PPC — Prompt Policy Compiler (v0.1)

PPC compiles small Markdown behavior modules into a single deterministic prompt (stdout-first).

## Build

```bash
go build -o ppc ./cmd/build-prompt
```

## Subcommands

PPC uses a subcommand-based CLI:

```bash
ppc <subcommand> [flags]
```

### Mode Subcommands

Generate prompts for specific modes:

```bash
./ppc explore --conservative --revisions 1 --contract markdown
./ppc build --conservative --revisions 1 --contract code --explain
./ppc ship --creative --out AGENTS.md --hash
```

### Doctor Subcommand

Validate module structure and dependencies:

```bash
./ppc doctor --strict
./ppc doctor --json
./ppc doctor --prompts custom/
```

### Global Flags

```bash
./ppc --list        # List all available modules
./ppc --help        # Show help
```

### Flags Per Subcommand

Each mode subcommand supports:

```
--conservative          Include conservative trait (boring, stable)
--creative              Include creative trait (novelty encouraged)
--terse                 Include terse trait (brief, concise)
--verbose               Include verbose trait (detailed, expansive)
--revisions N           Enable policies/revisions with budget N
--contract TYPE         Output contract: code|markdown (default: markdown)
--out PATH              Write output to file (default: stdout)
--explain               Print resolution details to stderr
--hash                  Prepend SHA256 prompt-id header
--prompts DIR           Prompts directory (default: prompts)
```

## Layout

- `prompts/` contains Markdown modules with optional YAML frontmatter.
- `prompts/rules.yml` defines `exclusive_groups` for keyed tags (`group:value`).

## Notes

- Deterministic output: same inputs -> same output.
- Fails loudly on: missing modules, tag conflicts, circular requires.

