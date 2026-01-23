# Example: Product PRD Review Flow

## Problem

PRD reviews are inconsistent. Scope creeps. Acceptance criteria are vague. Non-goals are missing or weak. Risk assessments are superficial.

Engineers receive PRDs with:
- No clear success metrics ("fast enough", "good user experience")
- Missing scope boundaries ("we'll add more features if time permits")
- No acknowledgment of tradeoffs
- Unrealistic timelines or technical claims
- No risk mitigation plans

## Why this exists

Ad-hoc PRD reviews fail because:

- **No process gates**: Reviews vary wildly in rigor
- **No scope containment**: Features creep in during implementation
- **No measurable success**: "Done" is subjective, not objective
- **No tradeoff awareness**: PRDs present everything as possible
- **No risk mitigation**: Surprises emerge during implementation

PPC formalizes multi-stage review with explicit gates, acceptance criteria, and risk assessment.

## What this example demonstrates

- Multi-stage review process (exploration, definition, ship)
- Artifact-driven development (PRDs as first-class artifacts)
- Variable substitution (\`{{product_name}}\`, \`{{target_user}}\`, \`{{risk_level}}\`)
- Non-goals enforcement (scope creep prevention)
- Risk assessment with mitigation plans
- Risk-level-specific requirements

## How to run

```bash
ppc doctor
ppc explore --profile explore
ppc explore --profile ship
```

## Output

The compiled prompt defines a product manager enforcing structured PRD reviews with variable-substituted context.

Sample output (truncated):

```markdown
## Agent Identity

You are a senior product manager conducting structured PRD reviews.

You value:
- clarity over completeness
- constraints over possibilities
- evidence over opinions
- decisions over discussions
- shipped products over perfect plans

## Contract: PRD Review Flow

This review enforces structured evaluation of NewFeature for Developers at low risk level.

## Product Context

**Product:** NewFeature
**Target User:** Developers
**Risk Level:** low

**Risk Definitions:**
- **low**: Internal tool, user base < 100, no revenue impact
- **medium**: External product, user base 100-1000, minor revenue impact
- **high**: Critical product, user base > 1000, major revenue impact

## Review Stages

PRDs progress through these stages:

### Stage 1: Exploration Review
**Purpose:** Evaluate problem statement and solution approach
**Duration:** 2-3 days

**Review criteria:**
- Problem is real and worth solving
- Solution approach is viable
- Alternatives were considered
- Initial scope is realistic

[... output truncated ...]
```

To see the full output:

```bash
ppc explore --profile explore --out compiled.md
```

The \`explore\` profile substitutes \`NewFeature\`, \`Developers\`, \`low risk\` — producing a brainstorming-focused review with lighter requirements.

The \`ship\` profile substitutes \`ProductionRelease\`, \`Customers\`, \`high risk\` — producing a strict review with ship readiness gates and executive approval requirements.

## What to copy into your project

Copy these into your repository:

- \`prompts/\` — module structure (contracts, policies)
- \`profiles/\` — explore and ship profiles
- \`rules.yml\` — tag conflict rules

Customize variable values in profiles to match your products:
- \`product_name\`: Your product or feature name
- \`target_user\`: Your primary user segment
- \`risk_level\`: Product risk level (low, medium, high)

## Common failure modes

**Missing variable:**

If profile doesn't define required variable:

```yaml
mode: explore
contract: markdown
traits: []
vars: {}  # Missing product_name, target_user, risk_level
```

PPC will fail with:

```
error: variable not defined: product_name
```

Add all variables to profile:

```yaml
vars:
  product_name: "YourProduct"
  target_user: "YourUsers"
  risk_level: "low"
```

**Missing acceptance criteria:**

If PRD lacks measurable acceptance criteria:

```
### AC-1: Good performance

**What:**
System should be fast

**How measured:**
We'll test it

ERROR: Acceptance criteria must be measurable
- Use specific metrics and thresholds
- Example: p95 < 100ms, not "fast enough"
```

**Weak non-goals:**

If non-goals use vague language:

```
## Non-Goals

We'll keep it simple and not add too many features.

ERROR: Non-goals must be explicit
- List specific out-of-scope items
- Explain why each is excluded
- State tradeoffs explicitly
- Avoid "nice to have" or "if time permits"
```

## CI

Copy \`examples/workflows/prompt-lint.yml\` into your repo at:

\`.github/workflows/prompt-lint.yml\`

This ensures all prompt modules pass \`ppc doctor\` before merge.

## Key takeaway

> "PPC can formalize how products are designed, not just how prompts are written."
