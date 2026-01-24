---
id: response-format
title: Response Format Requirements
description: Structure that every response must follow
requires: [base]
tags: []
---

Every response from team members must follow a consistent structure. This ensures that feedback is substantive, clear, and actionable.

## Every Team Response Must Include Three Parts

### Part 1: Perspective

State what you see, know, or understand about the requirement.

- Be specific: "I see a need for real-time updates, which means infrastructure for websockets or polling."
- Be brief: One sentence to one paragraph.
- Ground it in what you know: "Based on similar projects, I expect..."

### Part 2: Concern or Question

Identify what might go wrong, what's missing, or what's unclear.

- **Technical concerns:** "Real-time updates require infrastructure investment. Have we budgeted for this?"
- **Design concerns:** "Mobile-first might conflict with desktop analytics views."
- **Process concerns:** "We haven't discussed data retention requirements. Does this data need to be retained indefinitely?"

### Part 3: Recommendation

Suggest what we should do about it.

- "Suggest we start with update frequency of 5 seconds, not 1 second, to reduce infrastructure cost."
- "Recommend we defer the third tenant to phase 2, so we can focus on core functionality."
- "Ask the client if WCAG 2.1 AA compliance is required, or if basic accessibility is sufficient."

## Examples

**Good response:**

> **Perspective:** I understand you want a dashboard that updates in real-time. Based on similar projects, this means websocket infrastructure.
>
> **Concern:** Real-time updates can get expensive if we have many concurrent users. Have we estimated the expected concurrent user count?
>
> **Recommendation:** Suggest we clarify the expected scale. If it's under 1,000 concurrent users, websockets are feasible. If it's more, we may need to consider polling or server-sent events.

**Bad response:**

> Real-time updates are hard. We should think about scaling.

This response is vague and lacks actionable insight. It doesn't identify a specific concern or offer a concrete recommendation.

## Every Decision Must Record

When a decision is made, record:

- **What we decided:** The actual scope or path.
- **Rationale:** Why this path.
- **What we deferred:** Scope that's out of scope, and why.
- **Risks we're accepting:** Tradeoffs and constraints we're living with.

**Example:**

> **Decided:** Build MVP with two core views (project list, task detail). Mobile views deferred to phase 2.
>
> **Rationale:** Two views cover the most critical user journeys. Mobile adds significant design effort that would delay the MVP.
>
> **Deferred:** Mobile views, team analytics dashboard.
>
> **Risks we're accepting:** Users may expect mobile support. We're accepting the risk that some users will access from mobile devices in phase 1.
