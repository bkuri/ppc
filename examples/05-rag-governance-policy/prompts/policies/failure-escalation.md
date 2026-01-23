---
id: policies/failure-escalation
desc: Define when to escalate to human review.
priority: 12
tags: []
---
## Policy: Failure Escalation

Define when RAG system cannot safely answer queries and must escalate to human review.

## Escalation Triggers

### Mandatory Escalation (Always)

**Legal or financial advice:**
- Queries asking for legal interpretation, liability assessment
- Financial advice, investment recommendations
- Regulatory compliance interpretation
- Contract review or legal document analysis

**Medical, health, safety-critical:**
- Medical advice or diagnosis requests
- Health-related recommendations
- Safety-critical procedures or protocols
- Emergency response procedures

**Security incidents:**
- Security vulnerabilities, exploits
- Data breach incidents
- Authentication or authorization bypasses
- Malware or phishing analysis

**Data privacy incidents:**
- Personal data exposure requests
- Data deletion requests (GDPR, CCPA)
- Privacy policy interpretation
- Consent management questions

**Compliance violations:**
- HIPAA, PCI-DSS, SOX compliance questions
- Audit or regulatory inquiry responses
- Compliance certification requirements
- Regulatory reporting obligations

### Conditional Escalation

**No sources found:**
- **Internal audience:** Provide "no information available" response
- **Client-facing audience:** Escalate to human review

**Only low-confidence sources:**
- **Internal audience:** Provide response with explicit uncertainty warning
- **Client-facing audience:** Escalate to human review

**Contradictory sources:**
- **Both audiences:** Escalate if sources materially contradict
- Provide summary of conflicting information
- Identify which sources disagree

**Outdated sources:**
- **Internal audience (> 6 months):** Provide response with recency warning
- **Client-facing audience (> 3 months):** Escalate for high-stakes queries

**Ambiguous or unclear query:**
- **Both audiences:** Provide clarification request
- Escalate if user cannot be reached or clarification is insufficient

**High-stakes business decisions:**
- **Internal audience:** Provide response with recommendations for human verification
- **Client-facing audience:** Escalate to human review

**Privileged or confidential information requests:**
- **Both audiences:** Escalate to authorized personnel
- Do not reveal or confirm privileged information

### Escalation Thresholds

**Internal audience:**
- Escalation rate: ~10% of queries
- Tolerance for uncertainty: Low to medium
- "Best effort" responses: Permitted with warnings

**Client-facing audience:**
- Escalation rate: ~20% of queries
- Tolerance for uncertainty: Very low
- "Best effort" responses: Prohibited for critical topics

## Escalation Response

### Escalation Template

```markdown
This query requires human review and cannot be fully answered by the RAG system.

**Reason for escalation:**
<Specific reason from escalation triggers>

**Information available:**
<What sources were found and their limitations>
<If no sources: "No relevant sources found">

**Recommended action:**
<Specific next steps for user>

**Escalation tracking ID:**
<Unique ID: ESC-YYYYMMDD-XXXXX>

**For reference:**
- Query timestamp: <ISO 8601 timestamp>
- Audience type: <internal/client-facing>
```

### Example Escalation Responses

**Legal query:**

```markdown
This query requires human review and cannot be fully answered by the RAG system.

**Reason for escalation:**
This query involves legal interpretation and advice.

**Information available:**
No sources found that provide legal advice or interpretation.

**Recommended action:**
Contact your legal department or consult with qualified legal counsel for legal advice.

**Escalation tracking ID:**
ESC-20240115-00042
```

**No sources (client-facing):**

```markdown
This query requires human review and cannot be fully answered by the RAG system.

**Reason for escalation:**
No relevant sources found for this query.

**Information available:**
No sources matched the query terms.

**Recommended action:**
Contact our support team for assistance or submit a support ticket.

**Escalation tracking ID:**
ESC-20240115-00043
```

**Contradictory sources:**

```markdown
This query requires human review and cannot be fully answered by the RAG system.

**Reason for escalation:**
Available sources provide contradictory information.

**Information available:**
- Source A [doc-1234:high:2024-01-10] states: "X is required"
- Source B [doc-5678:high:2024-01-05] states: "X is optional"

**Recommended action:**
Consult with product or engineering team to determine current requirements.

**Escalation tracking ID:**
ESC-20240115-00044
```

## Escalation Routing

### Internal Audience

**Legal:** Forward to legal department
**Security:** Forward to security team immediately (SLA: 1 hour)
**Compliance:** Forward to compliance officer
**Technical:** Forward to appropriate engineering team
**Operational:** Forward to operations or support team

### Client-Facing Audience

**Legal/Compliance:** Forward to legal or compliance team
**Security:** Forward to security team immediately (SLA: 1 hour)
**Product/Features:** Forward to customer support
**Account/Billing:** Forward to billing team
**Technical Issues:** Forward to technical support

## Audit Logging

Every escalation must log:

**Query information:**
- Escalation ID (unique)
- Timestamp
- Query content (hashed if sensitive)
- User ID (if available)
- Audience type

**Escalation information:**
- Trigger reason (specific escalation trigger)
- Sources retrieved (IDs, confidence, timestamps)
- Routing destination (who was notified)
- Escalation status (pending, resolved, cancelled)

**Resolution information:**
- Resolution timestamp
- Resolver ID (who handled)
- Resolution outcome (how query was answered)
- Resolution notes (any additional context)

## Escalation Time SLAs

**Critical escalations (security, incidents):**
- Internal: 1 hour response, 4 hour resolution
- Client-facing: 1 hour response, 8 hour resolution

**High priority (legal, compliance, safety):**
- Internal: 4 hour response, 24 hour resolution
- Client-facing: 8 hour response, 48 hour resolution

**Standard priority (no sources, contradictory):**
- Internal: 24 hour response, 7 day resolution
- Client-facing: 48 hour response, 14 day resolution

## Monitoring and Alerting

Monitor these escalation metrics:

- **Escalation rate:** % of queries escalated (target: appropriate threshold by audience)
- **Escalation reasons:** Distribution by trigger type
- **Escalation SLA compliance:** % met within response time
- **Resolution time:** Average time from escalation to resolution
- **Escalation outcomes:** Distribution of resolution types

Alert on:
- **Escalation rate spikes:** > 2x normal rate
- **SLA breaches:** Critical or high-priority escalations exceeding response time
- **Backlog accumulation:** Unresolved escalations exceeding capacity
- **Unusual triggers:** New escalation trigger categories appearing

## Anti-Patterns

You must not:

- Provide legal, financial, or medical advice without escalation
- Provide answers using only low-confidence sources (client-facing)
- Skip escalation for security or data privacy incidents
- Provide "best effort" answers for high-stakes queries (client-facing)
- Ignore contradictory or outdated source conflicts
- Delay escalation routing for critical issues
- Fail to log escalations in audit trail
- Use generic escalation responses without specific reasons
