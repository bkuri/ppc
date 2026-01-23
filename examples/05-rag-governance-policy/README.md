# Example: RAG Governance Policy

## Problem

RAG systems hallucinate. No audit trails. Compliance risks. Sources are misattributed or uncited.

Enterprise AI systems fail because:
- Hallucinations produce false information
- No source attribution makes claims unverifiable
- Legal/compliance queries lack proper oversight
- Audit trails don't exist for regulatory review
- Low-confidence sources are presented as facts

## Why this exists

Ad-hoc RAG governance fails because:

- **No source hierarchy**: All sources treated equally, regardless of trust
- **No citation standards**: Inconsistent or missing attribution
- **No escalation rules**: Critical queries get automated responses
- **No audit trails**: Can't trace decisions or comply with regulatory requests
- **No compliance enforcement**: Legal/financial queries lack human oversight

PPC enforces RAG governance by encoding citation rules, source hierarchies, and escalation policies into prompt.

## What this example demonstrates

- Multiple exclusive tag groups (\`audience\`, \`output\`)
- Enterprise AI governance (audit trails, compliance)
- Source ranking and trust hierarchy (5 levels)
- Mandatory citation format (\`[source-id:confidence:timestamp]\`)
- Escalation rules (mandatory vs conditional)
- Audience-specific requirements (internal vs client-facing)

## How to run

```bash
ppc doctor
ppc explore --profile internal
ppc explore --profile client-facing
```

## Output

The compiled prompt defines an enterprise RAG systems operator enforcing source attribution and audit trails for either internal or client-facing audiences.

Sample output (truncated):

```markdown
## Agent Identity

You are an enterprise AI systems operator ensuring RAG system compliance and auditability.

You value:
- transparency over speed
- source attribution over convenience
- human oversight over automation
- audit trails over operational efficiency
- compliance over cleverness

## Contract: RAG System Governance

This RAG system operates under strict governance for internal audience.

## Audience: internal

**Internal audience:**
- Access to internal and external sources
- Moderate escalation thresholds
- Audit focus: source lineage and decision traceability
- Permitted: experimental features, lower-confidence sources

**Client-facing audience:**
- Limited to approved, high-confidence sources only
- High escalation thresholds for uncertain or sensitive topics
- Audit focus: compliance and customer protection
- Prohibited: experimental features, internal-only sources

## Source Citation Requirements

### Citation Format

Every factual claim must include:
- **Source identifier**: Unique reference to source document
- **Confidence level**: \`high\`, \`medium\`, \`low\`
- **Timestamp**: Date when source was last verified

**Format:** \`[source-id:confidence:timestamp]\`

**Examples:**
- \`[doc-1234:high:2024-01-15]\` — High confidence, verified Jan 15 2024
- \`[policy-567:medium:2023-11-20]\` — Medium confidence, verified Nov 20 2023
- \`[faq-890:low:2024-01-10]\` — Low confidence, verified Jan 10 2024

[... output truncated ...]
```

To see the full output:

```bash
ppc explore --profile internal --out compiled.md
```

The \`internal\` profile substitutes \`internal\` audience — allowing access to internal sources, lower escalation thresholds, and moderate risk tolerance.

The \`client-facing\` profile substitutes \`client-facing\` audience — restricting to approved sources only, requiring higher escalation thresholds, and enforcing stricter compliance requirements.

## What to copy into your project

Copy these into your repository:

- \`prompts/\` — module structure (contracts, policies)
- \`profiles/\` — internal and client-facing profiles
- \`rules.yml\` — tag conflict rules (multiple exclusive groups)

Adjust source hierarchy and trust levels in \`policies/source-ranking.md\` to match your organization's data sources and risk tolerance.

## Common failure modes

**Tag conflict:**

If you try to load both internal and client-facing:

```yaml
mode: explore
contract: markdown
traits: []
vars:
  audience: "internal"     # CONFLICT: both audience tags would be active
```

PPC will fail with:

```
error: conflicting tags in group audience
```

**Missing variable:**

If profile doesn't define \`audience\`:

```yaml
mode: explore
contract: markdown
traits: []
vars: {}  # Missing audience variable
```

PPC will fail with:

```
error: variable not defined: audience
```

Add to profile:

```yaml
vars:
  audience: "internal"  # or "client-facing"
```

**Improper citation format:**

If sources are cited incorrectly:

```
According to the documentation (doc-1234), the feature is supported.

ERROR: Citation format must be [source-id:confidence:timestamp]

Correct:
According to the documentation [doc-1234:high:2024-01-15], the feature is supported.
```

**Escalation bypass:**

If legal query gets automated response:

```
Q: Is this legal interpretation correct?

A: Based on the contract language, this interpretation appears valid.

ERROR: Legal queries must escalate to human review

Correct escalation:
This query requires human review and cannot be fully answered by the RAG system.

**Reason for escalation:**
This query involves legal interpretation and advice.

**Recommended action:**
Contact your legal department or consult with qualified legal counsel.
```

## CI

Copy \`examples/workflows/prompt-lint.yml\` into your repo at:

\`.github/workflows/prompt-lint.yml\`

This ensures all prompt modules pass \`ppc doctor\` before merge.

## Key takeaway

> "PPC can define AI governance policies suitable for regulated environments."
