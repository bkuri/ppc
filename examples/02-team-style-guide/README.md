# Example: Team Style Guide Policy

## Problem

Inconsistent tone across 15 engineers. Design reviews vary wildly — some are one-sentence approvals, others are multi-page treatises. New team members don't know the communication standards.

PRs contain conflicting documentation styles. One engineer writes terse summaries, another writes verbose explanations. External stakeholders receive inconsistent communication about design decisions.

## Why this exists

Ad-hoc style guides fail because:

- **Not enforced**: Style guides as wikis are suggestions, not constraints
- **Not composed**: Each prompt adds its own style, diverging from team standards
- **Not versioned**: Style drifts as new engineers join and leave
- **Not contextual**: Different contexts (internal vs external) need different voice

PPC enforces team-wide behavioral policy by encoding style rules into the prompt itself. The compiled prompt always uses the correct tone and formatting.

## What this example demonstrates

- Policy enforcement (style guide, formatting standards)
- Exclusive tag groups (\`tone\` for voice variants)
- Organization-wide behavioral consistency
- Profile-based execution (default, design-reviews)
- Separation of voice rules from content logic

## How to run

```bash
ppc doctor
ppc explore --profile default
ppc explore --profile design-reviews
```

## Output

The compiled prompt defines a design systems engineer with company voice standards and either terse (default) or verbose (design reviews) communication style.

Sample output (truncated):

```markdown
## Agent Identity

You are a senior design systems engineer building reusable UI components.

You value:
- consistency over customization
- accessibility-first implementation
- semantic HTML over div soup
- component composition over inheritance
- documented patterns over implicit conventions

## Primary Objective

Create UI components that are:
- Accessible by default (WCAG 2.1 AA)
- Themeable via CSS variables
- Composable and nestable
- Performant and lightweight
- Well-documented with live examples

## Style Guide Policy

All communication must follow company voice standards.

## Voice Characteristics

Our voice is:
- **Direct but not blunt**: Say what you mean, but with respect
- **Technical but not jargon-heavy**: Use precise terminology, explain unfamiliar concepts
- **Pragmatic but not dismissive**: Focus on what works, acknowledge tradeoffs
- **Authoritative but not arrogant**: Speak from expertise, leave room for other perspectives
- **Helpful but not condescending**: Assume intelligence, not prior knowledge

[... output truncated ...]
```

To see the full output:

```bash
ppc explore --profile default --out compiled.md
```

The \`default\` profile produces terse, direct communication. The \`design-reviews\` profile produces detailed explanations with context and rationale.

## What to copy into your project

Copy these into your repository:

- \`prompts/\` — module structure (base, traits, policies, contracts)
- \`profiles/\` — preconfigured tone variants
- \`rules.yml\` — tag conflict rules

Adapt the \`base.md\` identity to match your team. Customize \`policies/style-guide.md\` to match your company voice.

## Common failure modes

**Tag conflict:**

If you try to load both terse and verbose:

```yaml
mode: explore
contract: markdown
traits:
  - traits/terse
  - traits/verbose  # CONFLICT: both tag tone:terse and tone:verbose
vars: {}
```

PPC will fail with:

```
error: conflicting tags in group tone: [tone:terse tone:verbose]
```

**Missing policy:**

If you remove \`policies/style-guide.md\` but references remain:

```bash
ppc doctor
# error: missing required entrypoint module: base (if base requires style-guide)
```

## CI

Copy \`examples/workflows/prompt-lint.yml\` into your repo at:

\`.github/workflows/prompt-lint.yml\`

This ensures all prompt modules pass \`ppc doctor\` before merge.

## Key takeaway

> "PPC can enforce team-wide communication standards."
