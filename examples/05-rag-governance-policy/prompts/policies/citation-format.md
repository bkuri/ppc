---
id: policies/citation-format
desc: Standardize citation format.
priority: 11
tags: []
---
## Policy: Citation Format

All information retrieved from RAG systems must use standardized citation format.

## Standard Format

**Format:** \`[source-id:confidence:timestamp]\`

**Components:**
- \`source-id\`: Unique identifier (e.g., \`doc-1234\`, \`policy-567\`)
- \`confidence\`: Confidence level (\`high\`, \`medium\`, \`low\`)
- \`timestamp\`: ISO 8601 date (YYYY-MM-DD)

**Examples:**
- \`[doc-1234:high:2024-01-15]\` — High confidence, verified Jan 15 2024
- \`[policy-567:medium:2023-11-20]\` — Medium confidence, verified Nov 20 2023
- \`[faq-890:low:2024-01-10]\` — Low confidence, verified Jan 10 2024

## Citation Placement

### Inline Citations

Use inline citations immediately after factual claims:

> "The product supports OAuth 2.0 authentication [doc-1234:high:2024-01-15] and SSO integration [policy-567:medium:2023-11-20]."

### Multiple Citations

If multiple sources support a claim, list all:

> "OAuth 2.0 is the recommended authentication method [doc-1234:high:2024-01-15][rfc-6749:high:2023-12-01]."

### Source-Specific Citations

If citing specific sections or features:

> "OAuth 2.0 supports refresh tokens [doc-1234:high:2024-01-15#refresh-tokens] for extending access."

## Citation Best Practices

### Do

- **Cite immediately**: Place citation right after the claim it supports
- **Cite explicitly**: Don't group citations at end of response
- **Cite accurately**: Ensure citation matches the specific claim
- **Cite completely**: Include all three components (id, confidence, timestamp)
- **Cite recent sources**: Prefer sources verified within recency thresholds

### Don't

- Don't cite general background (well-known facts don't need citations)
- Don't cite the same source repeatedly (cite once per unique claim)
- Don't use vague references (e.g., "according to the docs" without ID)
- Don't omit confidence level or timestamp
- Don't mix citation formats

## Confidence Level in Citations

### When to Use Each Level

**High confidence:**
- Direct quotes from official sources
- Official documentation, policies, procedures
- Authoritative external sources (RFCs, standards)
- Recently verified (within thresholds)

**Medium confidence:**
- Paraphrased or summarized information
- Moderately recent sources
- Reputable but potentially dated information

**Low confidence:**
- Inferred or synthesized information
- Dated sources
- Community or less authoritative sources
- **Client-facing audience:** Prohibited

## Timestamp Format

**Always use ISO 8601 (YYYY-MM-DD):**

- Correct: \`[doc-1234:high:2024-01-15]\`
- Incorrect: \`[doc-1234:high:Jan 15, 2024]\`
- Incorrect: \`[doc-1234:high:01/15/2024]\`

**Use last verification date, not creation date:**

- If source was created 2020-01-01 but verified 2024-01-15, use 2024-01-15
- If source has never been verified, use creation date with warning

## Source ID Format

**Recommended patterns:**

- **Documents:** \`doc-1234\`, \`doc-5678\`
- **Policies:** \`policy-123\`, \`policy-456\`
- **FAQs:** \`faq-789\`, \`faq-012\`
- **External:** \`rfc-6749\`, \`w3c-html5\`
- **Issues:** \`issue-3456\`

**Requirements:**
- Unique across all sources
- Descriptive (indicates source type)
- Human-readable (for audit and debugging)

## Missing or Incomplete Citations

If information lacks proper citations:

**Do not:**
- Make claims without sources
- Use generic references ("the documentation says...")
- Omit confidence level or timestamp
- Fabricate or hallucinate source IDs

**Do:**
- Mark uncertainty explicitly: "Based on available sources..."
- Provide caveat: "This information could not be verified"
- Escalate to human review (client-facing audience)
- Request clarification or additional context

## Citation Verification

Before including a citation, verify:

- [ ] Source ID exists and is unique
- [ ] Source is accessible (not deleted or archived)
- [ ] Confidence level is appropriate for claim and audience
- [ ] Timestamp is accurate (last verification date)
- [ ] Citation format matches \`[source-id:confidence:timestamp]\`
- [ ] Source supports the specific claim being made

## Common Errors

**Missing components:**
- Wrong: \`[doc-1234]\` — missing confidence and timestamp
- Wrong: \`[doc-1234:high]\` — missing timestamp
- Correct: \`[doc-1234:high:2024-01-15]\`

**Incorrect format:**
- Wrong: \`(doc-1234, high, 2024-01-15)\` — wrong delimiter
- Wrong: \`[doc-1234: high: 2024-01-15]\` — extra spaces
- Correct: \`[doc-1234:high:2024-01-15]\`

**Wrong timestamp:**
- Wrong: \`[doc-1234:high:Jan 15 2024]\` — not ISO 8601
- Wrong: \`[doc-1234:high:2024-01-15T10:30:00Z]\` — too precise
- Correct: \`[doc-1234:high:2024-01-15]\`

## Audience-Specific Requirements

### Internal Audience
- Can use all confidence levels with appropriate warnings
- Can cite community and unverified sources with explicit warnings
- Timestamp warnings required if > 6 months old

### Client-Facing Audience
- Must use \`high\` or \`medium\` confidence for factual claims
- Cannot cite community or unverified sources
- Timestamp warnings required if > 3 months old
- Legal/compliance topics require disclaimer and professional consultation recommendation
