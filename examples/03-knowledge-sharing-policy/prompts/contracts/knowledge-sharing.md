---
id: contracts/knowledge-sharing
desc: Define conversation period and artifact requirements.
priority: 0
tags: []
---
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

## Conversation Period

All discussions must conclude within 7 days of initiation.

**Why 7 days?**
- Forces progress and prevents analysis paralysis
- Balances thoroughness with velocity
- Ensures decisions don't stall indefinitely
- Aligns with typical sprint cadence

**Exceptions:**
- Security incidents (24 hours)
- Production outages (immediate)
- Regulatory compliance (as required by policy)

## Artifact Requirements

Every conversation must produce an artifact with:

**Required sections:**
- Problem statement with PRD citations
- Requirements list with source links
- Options analysis (2-3 approaches)
- Decision with rationale
- Accepted tradeoffs
- Next steps with owners
- Related artifacts (links)

**Citation format:**
- PRD requirements: `[PRD-123]` or `[PRD-123:section]`
- RAG sources: `[source:confidence]` where confidence is \`high\`, \`medium\`, \`low\`
- Code references: \`src/path/to/file.ts:45\`

## Failure Modes

If the conversation cannot reach consensus within 7 days:
1. Escalate to technical lead
2. Make provisional decision with documented dissent
3. Schedule revisit window (14 days)
4. Document blockers for future reference

If requirements are ambiguous:
1. Pause and flag to product owner
2. Create requirements clarification artifact
3. Resume only after clarification

## Enforcement

Conversations that exceed 7 days without decision:
- Are marked as "stalled" in knowledge base
- Require technical lead approval to continue
- Cannot be referenced as precedent

Artifacts missing required sections:
- Cannot be marked as "final"
- Must be completed before close
- Block dependent decisions
