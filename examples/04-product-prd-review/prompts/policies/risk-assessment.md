---
id: policies/risk-assessment
desc: Evaluate and mitigate project risks.
priority: 12
tags: []
---
## Policy: Risk Assessment

All PRDs must include comprehensive risk assessment with mitigation plans.

## Risk Categories

Every PRD must address:

### 1. Technical Risks
- Feasibility (can we build this?)
- Performance (will it meet requirements?)
- Scalability (will it handle load?)
- Security (are there vulnerabilities?)
- Integration (will third-party systems work?)
- Technical debt (are we creating problems for future work?)

### 2. Business Risks
- Market fit (do users actually need this?)
- Competition (will someone beat us to market?)
- Pricing (will users pay?)
- Regulatory (are there compliance issues?)
- Dependencies (do we rely on external factors?)

### 3. Operational Risks
- Resource availability (do we have enough people?)
- Timeline feasibility (can we ship on schedule?)
- On-call burden (will this create operational pain?)
- Support load (will we be overwhelmed with tickets?)
- Monitoring (can we observe system health?)

## Risk Rating Matrix

Use this matrix to classify risks:

| Likelihood | Impact | Risk Level | Response Required |
|-------------|---------|-------------|------------------|
| High | High | Critical | Immediate mitigation or project cancellation |
| High | Medium | High | Mitigation before approval |
| High | Low | Medium | Document, monitor |
| Medium | High | High | Mitigation before approval |
| Medium | Medium | Medium | Document, monitor |
| Medium | Low | Low | Document |
| Low | High | Medium | Document, monitor |
| Low | Medium | Low | Document |
| Low | Low | Low | Document |

## Risk Format

Use this format for each risk:

```markdown
### Risk-<number>: <Brief title>

**Category:**
<Technical, Business, Operational>

**Likelihood:**
<High, Medium, Low>

**Impact:**
<High, Medium, Low>

**Risk Level:**
<Critical, High, Medium, Low>

**Description:**
<What could go wrong?>

**Consequence if it happens:**
<What's the worst-case outcome?>

**Mitigation plan:**
<How to prevent or reduce likelihood?>

**Contingency plan:**
<What to do if it happens anyway?>

**Owner:**
<Who owns tracking this?>

**Review frequency:**
<Weekly, Biweekly, Per milestone>
```

## Examples

**Good risk:**

```markdown
### Risk-1: Database performance under load

**Category:**
Technical

**Likelihood:**
Medium

**Impact:**
High

**Risk Level:**
High

**Description:**
Current query patterns may not scale to projected 10k concurrent users.

**Consequence if it happens:**
System timeouts, poor user experience, potential revenue loss.

**Mitigation plan:**
- Load test with 2x projected load before launch
- Add query performance monitoring
- Prepare read replica architecture

**Contingency plan:**
- Implement query caching layer
- Add circuit breakers to prevent cascade failures

**Owner:**
Tech Lead

**Review frequency:**
Per milestone
```

**Bad risk:**

```markdown
### Risk-1: Performance issues

**Description:**
It might be slow.

**Mitigation:**
We'll make it fast.
```

## Risk Level Requirements

**High-risk products require:**
- Minimum 8 risks documented
- At least 2 critical risks identified (if applicable)
- All critical and high risks have mitigation AND contingency plans
- Weekly review frequency for critical risks

**Medium-risk products require:**
- Minimum 5 risks documented
- All high risks have mitigation plans
- Biweekly review frequency

**Low-risk products require:**
- Minimum 3 risks documented
- All documented risks have owner assigned

## Anti-Patterns

You must not accept risk assessments that:

- Use vague descriptions ("it might fail")
- Lack mitigation plans (especially for critical risks)
- Don't assign owners
- Fail to identify consequences
- Have no review frequency
- Ignore technical feasibility
- Omit common failure modes for similar products

## Verification

Before approving PRD, check risk assessment includes:

- [ ] At least 3 risks (low risk), 5 (medium), 8 (high)
- [ ] Each risk has category, likelihood, impact, level
- [ ] All critical/high risks have mitigation AND contingency plans
- [ ] Each risk has assigned owner
- [ ] Review frequency specified for critical risks
- [ ] Technical feasibility addressed
- [ ] Integration risks identified (if applicable)
- [ ] Security risks addressed (if handling user data)
- [ ] Operational risks (on-call, support) considered

## Risk Reassessment

Risks must be reassessed at:
- Each stage gate (exploration, definition, ship)
- When new requirements added
- After major architectural changes
- When external dependencies change
- If any risk materializes

If risk level increases to critical:
1. Immediately flag to stakeholders
2. Require mitigation plan before proceeding
3. Consider project cancellation if unmitigatable
