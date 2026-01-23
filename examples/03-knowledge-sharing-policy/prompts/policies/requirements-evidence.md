---
id: policies/requirements-evidence
desc: Require PRD citations for all decisions.
priority: 11
tags: []
---
## Policy: Requirements Evidence

All decisions must cite PRD requirements that justify them.

## Citation Mandate

Every decision must reference:
- At least one PRD requirement
- The specific section if applicable
- The requirement ID or title

## Citation Format

**Format:** \`[PRD-ID]\` or \`[PRD-ID:section]\`

**Examples:**
> "Selected approach A because [PRD-45] requires support for 10k concurrent users."

> "This caching strategy satisfies [PRD-23:performance] which mandates <100ms p99 latency."

> "Simplified architecture due to [PRD-89:timeline] which requires shipping by Q3."

## What to Cite

Cite PRD requirements when:
- Explaining why a solution is chosen
- Justifying tradeoff decisions
- Rejecting alternative approaches
- Defining acceptance criteria
- Setting scope boundaries

## Missing Citations

If a decision lacks PRD citations:

**Do not:**
- Accept implicit "business requirements"
- Allow "it seems like a good idea"
- Use previous decisions as justification
- Rely on hallway conversation context

**Do:**
- Flag: "Missing PRD citation"
- Ask: "What requirement drives this?"
- Pause decision until cited
- Escalate to product owner if unclear

## Ambiguous Requirements

If PRD requirements are unclear:

1. **Flag explicitly**: "PRD-[ID] is ambiguous"
2. **Quote the requirement**: Copy exact text
3. **Explain ambiguity**: What's unclear?
4. **Create clarification artifact**: Document interpretation
5. **Get product owner sign-off**: Explicit approval

## Verification

Before finalizing decision, check:

- [ ] Every major design choice has PRD citation
- [ ] Tradeoffs reference specific requirements
- [ ] Scope boundaries cite PRD "non-goals" or out-of-scope sections
- [ ] Acceptance criteria map to PRD success metrics
- [ ] Ambiguities are documented and clarified

## Traceability

Good traceability example:

```
Decision: Use Redis for caching

Why:
- [PRD-45] requires 10k concurrent users, in-memory caching necessary
- [PRD-23:performance] mandates <100ms p99 latency, DB cannot meet alone
- [PRD-67:reliability] requires 99.9% uptime, Redis cluster provides HA

Tradeoffs accepted:
- Added operational complexity (justified by [PRD-45] scale requirements)
- Additional infrastructure cost (accepted per [PRD-89] resource budget)
```

## Anti-Patterns

You must not:
- Use generic "business needs" as citation
- Cite previous decisions without original requirement
- Assume team knows the PRD (cite explicitly)
- Accept "we need this" without requirement backing
- Skip citations for "obvious" decisions (nothing is obvious)
