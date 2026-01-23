---
id: policies/acceptance-criteria
desc: Define measurable success criteria.
priority: 10
tags: []
---
## Policy: Acceptance Criteria

All PRDs must include measurable acceptance criteria.

## Criteria Requirements

Acceptance criteria must be:
- **Specific**: Clear what "done" means
- **Measurable**: Can be tested or verified
- **Achievable**: Realistic given constraints
- **Relevant**: Directly addresses user problem
- **Time-bound**: Clear when to evaluate

## Format

Use this format for each acceptance criterion:

```
### AC-<number>: <Brief description>

**What:**
<One-sentence summary of requirement>

**How measured:**
<Specific metric or test>

**Success threshold:**
<Numeric value or pass/fail criteria>

**Owner:**
<Who validates this>

**Related to PRD:**
<Section or requirement being satisfied>
```

## Examples

**Good:**
```
### AC-1: Users can complete checkout in under 60 seconds

**What:**
Checkout flow from cart to payment confirmation

**How measured:**
Automated e2e test measuring time from "Checkout" click to "Order confirmed" page load

**Success threshold:**
p50 < 45s, p95 < 60s in load test with 100 concurrent users

**Owner:**
QA Lead

**Related to PRD:**
"Performance Requirements" section
```

**Bad:**
```
### AC-1: Checkout is fast

**What:**
Users shouldn't wait too long

**How measured:**
We'll check it manually

**Success threshold:**
Feels responsive

**Owner:**
Someone
```

## Coverage Criteria

Acceptance criteria must cover:

**Functional requirements:**
- Core user workflows
- Edge cases
- Error conditions
- Integration points

**Non-functional requirements:**
- Performance (response times, throughput)
- Reliability (uptime, error rates)
- Security (auth, data protection)
- Scalability (concurrent users, data volume)

**Quality requirements:**
- Accessibility (WCAG 2.1 AA)
- Usability (task completion rate)
- Documentation (API docs, user guides)

## Anti-Patterns

You must not accept acceptance criteria that:

- Use vague language ("fast", "responsive", "good user experience")
- Cannot be tested or measured
- Lack clear thresholds or pass/fail criteria
- Don't specify owners or validation method
- Reference "business value" without connecting to user outcome
- Include "nice to have" or "if time permits" language

## Verification

Before approving PRD, check:

- [ ] Each criterion is specific and measurable
- [ ] Success thresholds are numeric or binary (pass/fail)
- [ ] Measurement method is explicit
- [ ] Owners are assigned
- [ ] All functional requirements have criteria
- [ ] All non-functional requirements have criteria
- [ ] No "nice to have" language
- [ ] Criteria are testable (automated or manual)

## Revision Process

If acceptance criteria are incomplete:

1. **Flag specific deficiencies**: "AC-3 lacks success threshold"
2. **Request revision**: "Add measurable threshold for AC-3"
3. **Resubmit**: Include updated acceptance criteria

Revision counts against PRD revision budget.
