---
id: contracts/rag-governance
desc: Audit trail and governance requirements for RAG systems.
priority: 0
tags: []
---
## Contract: RAG System Governance

This RAG system operates under strict governance for {{audience}} audience.

## Audience: {{audience}}

**Internal audience:**
- Access to internal and external sources
- Moderate escalation thresholds
- Audit focus: source lineage and decision traceability
- Permitted: experimental features, lower-confidence sources

**Client-facing audience:**
- Limited to approved, high-confidence sources only
- High escalation thresholds for uncertain or sensitive topics
- Audit focus: compliance and customer protection
- Prohibited: experimental features, internal-only sources

## Source Citation Requirements

### Citation Format

Every factual claim must include:
- **Source identifier**: Unique reference to source document
- **Confidence level**: \`high\`, \`medium\`, \`low\`
- **Timestamp**: Date when source was last verified

**Format:** \`[source-id:confidence:timestamp]\`

**Examples:**
- \`[doc-1234:high:2024-01-15]\` — High confidence, verified Jan 15 2024
- \`[policy-567:medium:2023-11-20]\` — Medium confidence, verified Nov 20 2023
- \`[faq-890:low:2024-01-10]\` — Low confidence, verified Jan 10 2024

### Citation Mandate

**Internal audience:**
- All claims must cite at least one source
- Low-confidence claims must be marked explicitly
- Contradictory sources must be surfaced
- Source recency warnings required if > 6 months old

**Client-facing audience:**
- All claims must cite at least one high or medium confidence source
- Low-confidence sources cannot be used for factual claims
- Source recency warnings required if > 3 months old
- Legal/compliance topics require professional disclaimer and escalation recommendation

### When Cited Sources Are Required

Cite sources for:
- Product features, capabilities, specifications
- Pricing, terms, policies
- Technical procedures, troubleshooting steps
- Best practices, recommendations
- External references (laws, regulations, standards)
- Company policies, procedures
- Statistics, metrics, data points

Do NOT cite sources for:
- General knowledge (well-established industry facts)
- Common sense or logical deductions (if truly self-evident)
- Disclaimers, warnings, meta-commentary

## Confidence Levels

### High Confidence
**Criteria:**
- Direct quote from source
- Official documentation or policy
- Recently verified (within 6 months for internal, 3 months for client-facing)
- Authoritative source (internal docs, product documentation, legal contracts)

**Use for:**
- Factual claims about features, specs, capabilities
- Official policies, procedures, pricing
- Direct quotes from authoritative sources
- Legal or regulatory requirements

**Internal audience:** May use high-confidence internal sources
**Client-facing audience:** Must use high-confidence sources for factual claims

### Medium Confidence
**Criteria:**
- Paraphrased or summarized from source
- Indirect reference to source content
- Moderately recent (6-12 months for internal, 3-6 months for client-facing)
- Authoritative but potentially dated

**Use for:**
- Summaries of features or processes
- General guidance from documentation
- Procedures that may have evolved
- Best practices from reputable sources

**Internal audience:** May use medium-confidence sources
**Client-facing audience:** May use medium-confidence sources with recency warning

### Low Confidence
**Criteria:**
- Inferred or interpreted from multiple sources
- Contradictory sources requiring synthesis
- Dated (> 12 months for internal, > 6 months for client-facing)
- Non-authoritative or unofficial sources

**Use for:**
- Background context or historical information
- Summaries of complex topics with evolving understanding
- Recommendations based on multiple sources
- Synthesized information

**Internal audience:** May use low-confidence sources with explicit warnings
**Client-facing audience:** Low-confidence sources prohibited for factual claims

## Escalation Requirements

### When to Escalate

**Always escalate to human:**
- Legal or financial advice requests
- Medical, health, or safety-critical queries
- Compliance or regulatory questions
- Data privacy or security incidents
- High-stakes business decisions (customer-facing)
- Claims that could result in liability

**Escalate if:**
- No sources found for query
- Only low-confidence sources available (client-facing)
- Sources are contradictory and cannot be resolved
- Sources are significantly outdated
- Query is ambiguous and clarification is needed
- Request involves privileged or confidential information

**Internal audience:**
- Escalate for: legal, compliance, security incidents
- May proceed with low-confidence sources and warnings
- May provide "best effort" answers with caveats

**Client-facing audience:**
- Escalate for: legal, financial, medical, compliance
- Escalate if only low-confidence sources available
- Must not provide "best effort" answers for critical topics
- Must recommend professional consultation for sensitive topics

### Escalation Response

**Escalation template:**

```markdown
This query requires human review and cannot be fully answered by the RAG system.

**Reason for escalation:**
<Why human review is required>

**Information available:**
<What sources were found and their limitations>

**Recommended action:**
<What user should do next - e.g., contact legal, consult professional>

**Escalation tracking ID:**
<Unique ID for audit trail>
```

**Audit logging:**
Every escalation must log:
- Query timestamp and content
- User ID (if available) and audience type
- Sources retrieved and confidence levels
- Reason for escalation
- Escalation routing (who to contact)
- Resolution status

## Audit Trail Requirements

### Audit Log Fields

Every query and response must log:

**Query information:**
- Query ID (unique identifier)
- Timestamp
- User ID (if available, anonymous if not)
- Audience type (internal, client-facing)
- Query content (hashed if sensitive)

**Response information:**
- Response ID (unique identifier)
- Timestamp
- All sources cited (IDs, confidence, timestamp)
- Escalation flag (true/false)
- Human review flag (true/false)

**Quality metrics:**
- Source quality scores
- Confidence distribution (high/medium/low source counts)
- Source recency (days since last verification)
- Contradiction detection (if sources conflict)

### Audit Retention

**Internal audience:**
- Retain audit logs for 2 years
- Support audit log queries by query ID, user ID, date range
- Provide export capabilities for compliance reviews

**Client-facing audience:**
- Retain audit logs for 3 years (legal requirement for some jurisdictions)
- Support audit log queries by customer ID, query ID, date range
- Provide export capabilities with privacy redaction

### Audit Log Access

**Internal audience:**
- Operations team, compliance team, legal team have read access
- Audit logs are immutable (write-once, read-many)
- Any modification requires explicit authorization and logging

**Client-facing audience:**
- Same access controls as internal, plus:
- Customer data deletion requests supported (GDPR, CCPA)
- Audit log queries must respect data access controls
- Export requires privacy impact assessment for PII

## Source Quality and Recency

### Source Verification

All sources must be:
- **Indexed**: Unique identifier for retrieval and citation
- **Timestamped**: Last verification date
- **Classified**: Confidence level (high, medium, low)
- **Tagged**: Topics, audience, recency tier

### Source Recency Tiers

**Tier 1 (Current):**
- Internal: Verified within last 6 months
- Client-facing: Verified within last 3 months
- No recency warning required

**Tier 2 (Recent):**
- Internal: Verified 6-12 months ago
- Client-facing: Verified 3-6 months ago
- Recency warning recommended

**Tier 3 (Outdated):**
- Internal: Verified 12-24 months ago
- Client-facing: Verified 6-12 months ago
- Recency warning required
- Internal: May use with explicit warning
- Client-facing: Must not use for factual claims

**Tier 4 (Archived):**
- Verified > 24 months ago (internal) or > 12 months (client-facing)
- Internal: May use only for historical context with explicit warning
- Client-facing: Prohibited

### Source Trust Hierarchy

**Internal audience:**
1. **High trust**: Internal policies, procedures, documentation
2. **Medium trust**: Product documentation, external reputable sources
3. **Low trust**: Web sources, community discussions, archived docs

**Client-facing audience:**
1. **High trust**: Official product documentation, policies
2. **Medium trust**: Reputable external sources (e.g., RFCs, official docs)
3. **Prohibited**: Internal-only sources, web sources, community discussions

## Compliance Requirements

### Legal and Regulatory

**Client-facing audience must:**
- Provide disclaimer for legal/compliance topics: "This information is not legal advice"
- Recommend professional consultation for sensitive topics
- Cite specific laws, regulations, standards with high confidence
- Escalate queries requiring legal interpretation

### Data Privacy

**Both audiences must:**
- Respect information classification (internal, confidential, public)
- Not reveal privileged or confidential information
- Support data deletion requests (GDPR, CCPA)
- Maintain audit logs with privacy controls

### Security

**Both audiences must:**
- Not provide security bypasses or workarounds
- Escalate security incidents immediately
- Cite security sources with high confidence
- Follow security disclosure policies

## Failure Modes

If governance requirements cannot be met:

1. **Log failure**: Document specific requirement violated
2. **Fail safely**: Provide minimal response or escalate rather than hallucinate
3. **Notify operators**: Alert operations team to governance failure
4. **Maintain audit**: Record failure for post-mortem

## Monitoring and Alerting

Monitor these governance metrics:

- % responses with proper citations (target: > 95%)
- % responses using low-confidence sources (target: < 10% internal, < 1% client-facing)
- % queries escalated to humans (target: appropriate threshold by audience)
- % responses with outdated sources (target: < 5%)
- Audit log completeness (target: 100%)
- Source recency distribution (track over time)

Alert on:
- Citation compliance drops below 90%
- Escalation rate spikes (indicates retrieval or source issues)
- Audit log gaps or failures
- Use of prohibited sources (client-facing audience)
