---
id: policies/conversation-period
desc: Enforce 7-day conversation window.
priority: 10
tags: []
---
## Policy: Conversation Period

Knowledge-sharing conversations have a maximum duration of 7 days.

## Timebox Rules

**Initiation timestamp**: When the first message is posted

**Hard deadline**: 7 days (168 hours) from initiation

**No exceptions**: 7 days is absolute, not negotiable

## Why Timebox?

Unbounded conversations fail because:
- Decisions never happen — analysis paralysis
- Stakeholders lose interest
- Requirements change during discussion
- Context is lost as timeline extends
- Action items have no urgency

## Time Enforcement

**Day 0-2 (Problem Framing):**
- Define problem clearly
- Cite PRD requirements
- Identify constraints

**Day 2-5 (Exploration):**
- Discuss 2-3 approaches
- Surface tradeoffs
- Capture perspectives

**Day 5-6 (Decision):**
- Select approach
- Document rationale
- Accept tradeoffs

**Day 6-7 (Consolidation):**
- Create artifact
- Assign actions
- Link to related work

## If Deadline Approaches

**At 6 days:**
- Flag conversation: "Decision required by [time]"
- Summarize remaining blockers
- Force decision or escalation

**At 7 days without decision:**
- Automatically mark "stalled"
- Require technical lead intervention
- Cannot be used as precedent

## Escalation Path

If consensus cannot be reached:

1. **Technical lead makes provisional decision**
   - Documents dissenting perspectives
   - Marks decision as "provisional"
   - Sets revisit window (14 days)

2. **Product owner clarifies requirements**
   - If issue is ambiguous requirements
   - Creates requirements artifact
   - Conversation resumes

3. **Architecture owner intervenes**
   - If no viable technical solution exists
   - Redefines constraints
   - Restarts conversation

## Anti-Patterns

You must not:
- Extend deadline for "more discussion"
- Allow "we'll figure it out later"
- Accept decisions without explicit tradeoffs
- Mark incomplete conversations as final

## Measurement

Track these metrics:
- % conversations concluded within 7 days
- % conversations requiring escalation
- Average time to decision
- % decisions referenced in future work
