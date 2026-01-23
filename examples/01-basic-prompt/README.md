# Example: Basic Prompt Composition

## Problem

Documentation drifts across 50+ files. Prompts are copied, pasted, and diverge silently. No single source of truth exists.

Engineers paste prompts into different tools, tweak them for specific contexts, and lose track of what changed. Two engineers solving the same problem use different prompt strategies, producing inconsistent documentation quality.

## Why this exists

Ad-hoc prompt management fails because:

- **No version control**: Prompt changes live in copy-paste buffers, not git
- **No composition**: Each prompt is a monolith — can't share behavioral modules
- **No determinism**: Same inputs to LLM produce different outputs based on prompt drift
- **No audit trail**: Can't trace which prompt version produced which artifact

PPC replaces giant prompt files with small, composable modules that compile deterministically.

## What this example demonstrates

- PPC as a prompt compiler (not a prompt manager)
- Modular behavior composition (base, mode, traits, contracts)
- Deterministic ordering (same inputs → same output)
- Exclusive tag groups (risk, output) for conflict-free compilation
- Profile-based execution (preconfigured behavior sets)

## How to run

```bash
ppc doctor --strict
ppc explore --profile explore-creative
ppc explore --profile explore-conservative
```

## Output

The compiled prompt defines a technical writer with exploration mode and either creative or conservative behavior.

Sample output (truncated):

```markdown
## Agent Identity

You are a senior technical writer compiling project documentation.

You value:
- clarity over cleverness
- consistency over novelty
- accuracy over completeness
- actionable guidance over theory

## Primary Objective

Compile clear, accurate documentation that helps engineers build effectively.

## Documentation Standards

- Use present tense for descriptions
- Prefer imperative mood for procedures
- Link to relevant source code
- Include minimal, working examples

## Mode: Explore

Generate multiple viable approaches, call out tradeoffs, then recommend.
Prefer breadth first, then narrow to the best option.

## Exploration Framework

When asked to solve a problem:

1. **Survey the space**: List 3-5 distinct approaches
2. **Compare tradeoffs**: For each approach, identify:
   - Implementation complexity
   - Maintenance burden
   - Performance characteristics
   - Edge case handling
   - Alignment with project constraints

3. **Recommend**: Choose one approach and explain why

[... output truncated ...]
```

To see the full output:

```bash
ppc explore --profile explore-creative --out compiled.md
```

The creative variant will emphasize novel approaches and idea breadth. The conservative variant will emphasize stability, standard patterns, and minimal risk.

## What to copy into your project

Copy these into your repository:

- `prompts/` — module structure (base, modes, traits, contracts)
- `profiles/` — preconfigured behavior sets
- `rules.yml` — tag conflict rules

Adapt the `base.md` identity to match your domain (technical writer → your role).

## Common failure modes

**Tag conflict:**

If you try to load both creative and conservative:

```yaml
mode: explore
contract: markdown
traits:
  - traits/creative
  - traits/conservative  # CONFLICT: both tag risk:high and risk:low
vars: {}
```

PPC will fail with:

```
error: conflicting tags in group risk: [risk:high risk:low]
```

This is intentional — PPC fails loudly rather than silently resolving conflicts.

**Cycle detection:**

If `base.md` requires `explore.md` and `explore.md` requires `base.md`:

```bash
ppc doctor --strict
# ERROR: circular dependency detected: base → explore → base
```

PPC validates dependency graphs before compilation.

## CI

Copy `examples/workflows/prompt-lint.yml` into your repo at:

`.github/workflows/prompt-lint.yml`

This ensures all prompt modules pass `ppc doctor --strict` before merge.

## Key takeaway

> "This is what PPC replaces: one giant prompt file."
