---
id: contracts/prd-review
desc: Multi-stage PRD review process with gates.
priority: 0
tags: []
---
## Contract: PRD Review Flow

This review enforces structured evaluation of {{product_name}} for {{target_user}} at {{risk_level}} risk level.

## Product Context

**Product:** {{product_name}}
**Target User:** {{target_user}}
**Risk Level:** {{risk_level}}

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

**Decision outcomes:**
- **Approve for definition:** Clear problem, reasonable approach
- **Request changes:** Missing information, unclear solution
- **Reject:** Wrong problem, infeasible solution, out of scope

**Acceptable revisions:** 2
**Reviewers:** Product manager + tech lead

### Stage 2: Definition Review
**Purpose:** Validate requirements, scope, and feasibility
**Duration:** 3-5 days

**Review criteria:**
- Acceptance criteria are measurable
- Non-goals are explicit (no scope creep)
- Dependencies are identified
- Timeline is realistic
- Risk assessment is complete

**Decision outcomes:**
- **Approve for implementation:** Ready for engineering
- **Request changes:** Missing acceptance criteria, unrealistic timeline
- **Reject:** Infeasible, insufficient risk mitigation

**Acceptable revisions:** 2
**Reviewers:** Product manager + tech lead + engineering manager

### Stage 3: Ship Readiness Review ({{risk_level}} risk only)
**Purpose:** Final sign-off before production release
**Duration:** 1-2 days

**Review criteria:**
- All acceptance criteria met
- No critical blockers
- Rollback plan documented
- Monitoring in place
- Post-launch plan defined

**Decision outcomes:**
- **Ship to production:** Safe to release
- **Ship to staging:** Minor concerns, canary deployment
- **Hold:** Critical blockers, cannot ship

**Acceptable revisions:** 1
**Reviewers:** Product manager + engineering manager + stakeholder

## Stage-Specific Requirements

### Exploration Stage
PRD must include:
- Problem statement with user quotes or data
- Solution hypothesis
- 2-3 alternative approaches considered
- Initial scope boundaries
- Success hypothesis

### Definition Stage
PRD must include:
- Detailed functional requirements
- Non-functional requirements (performance, reliability, security)
- Explicit non-goals section
- Acceptance criteria (measurable)
- Risk assessment with mitigation plans
- Dependencies (internal, external)
- Timeline with milestones
- Resource estimates

### Ship Stage ({{risk_level}} risk only)
Release must include:
- Testing summary (unit, integration, e2e)
- Performance metrics (load test results)
- Security review sign-off
- Rollback plan tested
- Monitoring and alerting configured
- Post-launch success metrics

## Revision Budgets

Each stage has a revision budget:
- Exploration: 2 revisions
- Definition: 2 revisions
- Ship: 1 revision

If revision budget is exceeded:
1. Escalate to product director
2. Create new PRD or significantly simplify
3. Re-enter process at appropriate stage

## Risk Level Requirements

**{{risk_level}} risk products require:**

### All risk levels:
- Problem validation evidence (user research, data, stakeholder input)
- Measurable success criteria
- Explicit non-goals

### Medium+ risk:
- Technical feasibility assessment
- Risk mitigation plans
- Dependencies explicitly mapped
- Post-mortem triggers defined

### High risk:
- Ship readiness review required
- Security and compliance review
- Legal review if applicable
- Executive approval
- External customer beta if applicable

## Gate Enforcement

A PRD cannot advance stages unless:
- All review criteria met
- Required reviewers have approved
- Revision budget not exceeded
- Risk-level-specific requirements satisfied

## Failure Modes

If review process fails:
- **Rejection:** Clear rationale given, can resubmit with changes
- **Stall:** Escalation required, process blocked
- **Exception:** Product director override (documented, rare)

## Artifacts

Each stage produces:
- **Review comments:** Structured feedback with action items
- **Decision record:** Approved, changes requested, or rejected
- **PRD version:** Versioned artifact at stage gate
- **Metrics:** Time to decision, revision count, reviewer agreement

## Reviewer Responsibilities

**Product manager:**
- Validate problem and user value
- Ensure business case is sound
- Evaluate market fit

**Tech lead:**
- Assess technical feasibility
- Evaluate architecture approach
- Identify technical risks

**Engineering manager:**
- Validate timeline estimates
- Assess resource availability
- Identify capacity risks

**Stakeholder:**
- Validate alignment with strategy
- Identify dependencies on other work
- Provide domain expertise
