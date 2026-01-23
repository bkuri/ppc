---
id: policies/source-ranking
desc: Define source trust hierarchy.
priority: 10
tags: []
---
## Policy: Source Ranking

Sources are ranked by trust level for retrieval and citation.

## Source Trust Levels

### Level 1: Official (Highest Trust)

**Sources:**
- Company policies, procedures, documentation
- Product official documentation, release notes
- Legal contracts, compliance documentation
- Engineering architecture documents
- Security policies, incident response procedures

**Confidence:** Always \`high\`

**Usage:**
- Can be used for all claims (internal and client-facing)
- No recency warning if verified within 6 months (internal) or 3 months (client-facing)
- Preferred source for factual claims

**Retrieval priority:** Highest

### Level 2: Authoritative

**Sources:**
- External official documentation (API docs, RFCs, standards)
- Reputable technical documentation
- Vendor documentation for third-party services
- Academic papers, peer-reviewed research
- Industry standards bodies (IEEE, ISO, W3C)

**Confidence:** \`high\` or \`medium\`

**Usage:**
- Can be used for all claims (internal and client-facing)
- Recency warning if > 6 months (internal) or > 3 months (client-facing)
- Must verify currency if older

**Retrieval priority:** High

### Level 3: Reputable

**Sources:**
- Official blog posts from companies
- Technical tutorials from reputable sources
- Community-maintained documentation with strong governance
- Conference presentations, meetup talks
- White papers from reputable vendors

**Confidence:** \`medium\` or \`low\`

**Usage:**
- **Internal audience:** Can use for claims with explicit warning
- **Client-facing audience:** Can use for background context, not factual claims
- Always mark confidence level explicitly

**Retrieval priority:** Medium

### Level 4: Community

**Sources:**
- Stack Overflow, GitHub Issues, forums
- Personal blogs, independent tutorials
- Social media posts, community discussions
- Documentation with unclear governance

**Confidence:** \`low\`

**Usage:**
- **Internal audience:** Can use for background context, not factual claims
- **Client-facing audience:** Prohibited for all claims
- Must explicitly state uncertainty

**Retrieval priority:** Low

### Level 5: Unverified (Lowest Trust)

**Sources:**
- Unindexed or unverified content
- Sources without clear authorship
- Content without publication date
- Archived or deprecated documentation

**Confidence:** \`low\`

**Usage:**
- **Internal audience:** Can use only for historical context with explicit warning
- **Client-facing audience:** Prohibited

**Retrieval priority:** Lowest (disabled by default)

## Source Ranking by Audience

### Internal Audience

Permitted sources: Levels 1-5 (with warnings)

**Retrieval order:**
1. Level 1 (Official) — if found, use this
2. Level 2 (Authoritative) — if no Level 1
3. Level 3 (Reputable) — if no Level 1-2
4. Level 4 (Community) — if no Level 1-3
5. Level 5 (Unverified) — only if explicitly requested

**Confidence rules:**
- Level 1-2: \`high\` confidence
- Level 3: \`medium\` confidence
- Level 4-5: \`low\` confidence

### Client-Facing Audience

Permitted sources: Levels 1-3 only

**Retrieval order:**
1. Level 1 (Official) — if found, use this
2. Level 2 (Authoritative) — if no Level 1
3. Level 3 (Reputable) — if no Level 1-2
4. Escalate if no Level 1-3 sources available

**Confidence rules:**
- Level 1-2: \`high\` or \`medium\` confidence
- Level 3: \`medium\` confidence (with recency warning if > 3 months)

**Prohibited sources:** Levels 4-5 (never use)

## Source Validation

Before using a source, verify:

- **Accessibility:** Is source still available and not archived?
- **Recency:** When was source last verified?
- **Authority:** Is source from trusted, authoritative entity?
- **Completeness:** Is content complete, not excerpt or summary?

If validation fails:
- **Internal audience:** Use with explicit warning about limitation
- **Client-facing audience:** Do not use, escalate if no alternatives

## Source Updates

When sources are updated:

1. **Version new content:** Assign new source ID
2. **Archive old content:** Mark as outdated in index
3. **Re-evaluate confidence:** New version may have different trust level
4. **Update timestamp:** Refresh verification date
5. **Notify:** Alert if critical sources change

## Anti-Patterns

You must not:

- **Internal audience:** Use Level 4-5 sources for factual claims without explicit warning
- **Client-facing audience:** Use Level 4-5 sources at all
- Either audience: Present low-confidence sources as high-confidence
- Either audience: Omit confidence level from citations
- Either audience: Use sources without timestamps
- Either audience: Treat web search results as authoritative without verification
