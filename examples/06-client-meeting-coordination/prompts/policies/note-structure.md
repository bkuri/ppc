---
id: note-structure
title: Meeting Notes Markdown Structure
description: How meeting notes must be formatted
requires: [base]
tags: []
---

Meeting notes are the artifact of truth. They must be structured, complete, and actionable.

## Meeting Notes Format

### Header Section

```
# Kickoff Meeting: {{project_name}}
**Date:** {{meeting_date}}
**Attendees:** [List participants with roles]
**Meeting Duration:** {{total_time_mins}} minutes
**Topics:** {{num_topics}}
```

### Per-Topic Section

Each topic covers one item under discussion. Repeat this structure for each topic.

```
## Topic {{N}}: {{topic_title}}

### Phase 1: Requirements Refinement
**Client Initial Requirement:**
> [Raw requirement from client]

**Refined Requirement:**
[Coordinator-refined version with specifics, after Q&A]
```

```
### Phase 2: Team Review
**Requirements Clarifier:**
- Perspective: [What they see]
- Concern: [What might go wrong]
- Recommendation: [What to do about it]

**Technical Advisor:**
- Perspective: [What they see]
- Concern: [What might go wrong]
- Recommendation: [What to do about it]

**Design Advisor:**
- Perspective: [What they see]
- Concern: [What might go wrong]
- Recommendation: [What to do about it]
```

```
### Phase 3: Followup Questions & Answers
**Q (Requirements Clarifier):** [Question]
**A (Client):** [Answer]

**Q (Technical Advisor):** [Question]
**A (Client):** [Answer]
```

```
### Phase 4: Decision
**Decided:** [What we're building]

**Rationale:** [Why this path]

**Deferred:** [What's out of scope and why]

**Risks Accepted:** [Constraints we're living with]
```

### Footer Section (One per Meeting, Not Per Topic)

```
## Overall Summary

**Final Scope:** [Consolidated scope across all topics]

**Blockers:** [If any exist, list them with owners]

**Next Steps:**
- Task 1 → Owner: [Person], Due: [Date]
- Task 2 → Owner: [Person], Due: [Date]
```

## Formatting Rules

- Use headings (`##`, `###`) to structure the document hierarchically.
- Use blockquotes (`>`) for direct quotes from participants.
- Use lists (`-`) for items like concerns, questions, recommendations.
- Use bold (`**`) for key labels like "Decided:", "Rationale:", etc.
- Leave blank lines between sections for readability.
- No trailing whitespace.

## Completeness Rules

The meeting notes are incomplete if:

- Any topic is missing the Decision section.
- The Decision section is missing Rationale, Deferred, or Risks Accepted.
- The Overall Summary is missing Next Steps with owners.
- Blockers exist but are not listed.

If any of these are missing, the meeting notes must be revised before being considered final.
