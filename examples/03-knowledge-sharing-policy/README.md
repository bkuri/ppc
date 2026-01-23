# Example: Knowledge Sharing Policy

## Problem

Knowledge conversations drift. No citation standards. PRD requirements are referenced implicitly or not at all. Decisions are made in hallway chats and lost.

Teams waste time re-discussing the same topics because:
- Previous decisions aren't documented
- No requirement traceability exists
- RAG-retrieved information isn't cited
- Conversation periods are unbounded

## Why this exists

Ad-hoc knowledge management fails because:

- **No process governance**: Conversations never conclude, decisions stall
- **No traceability**: Can't link decisions to requirements
- **No audit trail**: Can't reference why decisions were made
- **No structure**: Discussions vary wildly in format and depth

PPC defines conversation governance and traceability by encoding process rules into prompt itself.

## What this example demonstrates

- Conversational governance (7-day conversation window)
- Requirements traceability (mandatory PRD citations)
- RAG source attribution (mandatory source:confidence format)
- Contracts as process definition (not just output contract)
- Artifact-driven development (discussions produce referenceable artifacts)

## How to run

```bash
ppc doctor
ppc explore --profile knowledge-sharing
```

## Output

The compiled prompt defines a knowledge management facilitator enforcing structured conversations with 7-day windows, PRD citations, and RAG source attribution.

Sample output (truncated):

```markdown
## Agent Identity

You are a knowledge management facilitator helping teams capture and organize technical discussions.

You value:
- documentation over memory
- searchable decisions over hallway conversations
- traceable requirements over implicit assumptions
- explicit reasoning over tacit knowledge

## Contract: Knowledge-Sharing Conversation

This conversation follows a structured governance process with a fixed 7-day window.

## Conversation Lifecycle

Knowledge-sharing conversations have these stages:

1. **Problem Framing (Day 1-2)**
   - Define the problem clearly
   - Cite relevant PRD requirements
   - Identify constraints and assumptions

2. **Exploration (Day 2-5)**
   - Discuss 2-3 viable approaches
   - Surface tradeoffs explicitly
   - Capture all perspectives

3. **Decision (Day 5-6)**
   - Select one approach with explicit rationale
   - Document tradeoffs being accepted
   - Identify conditions that would invalidate this choice

4. **Consolidation (Day 6-7)**
   - Create referenceable artifact
   - Assign action items
   - Link to related artifacts

## Policy: Conversation Period

Knowledge-sharing conversations have a maximum duration of 7 days.

## Policy: Requirements Evidence

All decisions must cite PRD requirements that justify them.

## Citation Mandate

Every decision must reference:
- At least one PRD requirement
- The specific section if applicable
- The requirement ID or title

[... output truncated ...]
```

To see the full output:

```bash
ppc explore --profile knowledge-sharing --out compiled.md
```

## What to copy into your project

Copy these into your repository:

- \`prompts/\` — module structure (contracts, policies)
- \`profiles/\` — knowledge-sharing profile
- \`rules.yml\` — tag conflict rules

Adjust the 7-day window in \`contracts/knowledge-sharing.md\` if your organization requires different timelines.

## Common failure modes

**Missing PRD citation:**

If a decision is made without citing requirements:

```
Decision: Use Redis for caching

Why:
- It's fast
- It scales well

ERROR: Missing PRD citations
- Every major design choice must have [PRD-ID] citation
- Tradeoffs must reference specific requirements
```

**Missing RAG source:**

If information is presented without source:

```
According to the codebase, we should use this pattern...

ERROR: Missing RAG source citation
- Information must include [source:confidence]
- Specify source type: prd, design, code, ticket, external, meeting
- Specify confidence: high, medium, low
```

## CI

Copy \`examples/workflows/prompt-lint.yml\` into your repo at:

\`.github/workflows/prompt-lint.yml\`

This ensures all prompt modules pass \`ppc doctor\` before merge.

## Key takeaway

> "PPC can encode organizational process — not just prompt text."
