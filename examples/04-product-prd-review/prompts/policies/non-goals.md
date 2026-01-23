---
id: policies/non-goals
desc: Enforce explicit scope boundaries.
priority: 11
tags: []
---
## Policy: Non-Goals

All PRDs must include explicit non-goals section to prevent scope creep.

## Purpose of Non-Goals

Non-goals serve as:
- **Scope boundaries**: What we're NOT building
- **Tradeoff acknowledgment**: What we're giving up
- **Anti-pattern prevention**: Prevents future "nice to have" additions
- **Reference point**: Basis for rejecting feature requests

## Format

Use this format:

```markdown
## Non-Goals

The following are explicitly out of scope:

### Out of Scope: <Feature name>

**Reason:**
<Why not included in this release>

**Tradeoff:**
<What we lose by excluding this>

**Future consideration:**
<When this might be addressed - if ever>

**Explicitly rejected:**
<If decision to never include, state this>
```

## Required Non-Goals Categories

Every PRD must address:

### 1. Features Not in MVP
**Examples:**
- "Single sign-on (OAuth)" for initial release
- "Mobile app" for web-only product
- "Advanced search" for basic search feature

### 2. Performance Targets Beyond Scope
**Examples:**
- "Sub-second response times" for MVP (accept 2-3s)
- "99.999% uptime" for internal tool (accept 99.9%)

### 3. Edge Cases Not Covered
**Examples:**
- "Offline mode" for online-only product
- "Multi-language support" for English-only MVP
- "Enterprise SSO" for consumer product

### 4. Platform Exclusions
**Examples:**
- "iOS and Android" for web-first approach
- "Legacy browser support" (IE11, Safari < 14)

### 5. Administrative Features
**Examples:**
- "Admin dashboard" for MVP (handle via DB directly)
- "Audit logs" for initial release

## Examples

**Good non-goals section:**

```markdown
## Non-Goals

The following are explicitly out of scope:

### Out of Scope: Mobile App

**Reason:**
Web-first approach allows faster iteration, smaller initial investment.

**Tradeoff:**
Lower mobile user experience, no push notifications.

**Future consideration:**
Phase 2 if web usage metrics justify mobile investment.

**Explicitly rejected:**
Native app (will evaluate PWA first)
```

**Bad (or missing):**

```markdown
## Non-Goals

We'll keep it simple and not add too many features.
```

## Anti-Patterns

You must not:
- Omit non-goals section entirely
- Use vague language ("nice to have features")
- List features without explaining why excluded
- Fail to acknowledge tradeoffs
- Leave ambiguous future consideration
- Use "if time permits" (sets wrong expectation)

## Scope Creep Prevention

When features are requested during or after PRD approval:

1. **Check non-goals**: Is this explicitly listed as out of scope?
2. **If yes**: Reference non-goals section, reject with clear rationale
3. **If no**: Evaluate if core to user value
4. **If core**: Update PRD, require new approval cycle
5. **If not core**: Add to non-goals or backlog

## Revision Requirements

If non-goals are missing or insufficient:

1. **Identify categories**: Which required categories are missing?
2. **Request additions**: "Add non-goals for platform exclusions"
3. **Validate tradeoffs**: Are exclusions justified?
4. **Ensure explicit rejection**: If never included, state this clearly

## Verification

Before approving PRD, check non-goals section includes:

- [ ] At least 3 distinct out-of-scope items
- [ ] Each item has reason and tradeoff
- [ ] Explicitly rejected items are marked
- [ ] Future consideration is clear (or "never")
- [ ] No vague language ("nice to have", "if time permits")
- [ ] Platform exclusions addressed (if applicable)
- [ ] Performance boundaries stated (if applicable)
- [ ] Edge cases not covered listed

## Risk Level Impact

**High-risk products** require:
- More comprehensive non-goals (5+ items)
- Explicit "never include" decisions for some items
- Legal or compliance exclusions stated
- Security features not in initial release documented

**Low-risk products** may have:
- Fewer non-goals (3+ items)
- Less formal rejection language
- Simpler future consideration ("later")
