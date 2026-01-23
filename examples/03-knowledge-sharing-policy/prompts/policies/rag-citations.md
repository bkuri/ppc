---
id: policies/rag-citations
desc: Require source attribution for RAG-retrieved information.
priority: 12
tags: []
---
## Policy: RAG Citations

All information retrieved from RAG systems must be cited with source and confidence.

## Citation Format

**Format:** \`[source:confidence]\`

**Source types:**
- \`prd\`: Product requirements document
- \`design\`: Design documentation
- \`code\`: Codebase or commit history
- \`ticket\`: Bug tracker or feature requests
- \`external\`: Public documentation or RFCs
- \`meeting\`: Meeting notes or recordings

**Confidence levels:**
- \`high\`: Direct quote, unambiguous
- \`medium\`: Paraphrased, minor interpretation
- \`low\`: Inferred, significant interpretation

## Examples

**Direct quote:**
> According to the architecture spec [design:high], "all services must implement circuit breakers."

**Paraphrased:**
> The performance requirements suggest we need horizontal scaling [prd:medium], though exact numbers vary by service tier.

**Inferred:**
> Based on past incidents, this pattern likely causes race conditions [code:low].

## When to Cite

Cite RAG sources when:
- Referencing specific requirements or constraints
- Explaining technical tradeoffs
- Citing previous design decisions
- Referencing code patterns or examples
- Bringing in external best practices
- Quoting documentation or RFCs

## Confidence Guidelines

**High confidence:**
- Direct quotes from source
- Code references with line numbers
- Exact requirement text
- Official documentation

**Medium confidence:**
- Summarized requirements
- Paraphrased design docs
- Described code patterns
- Multiple sources agreeing

**Low confidence:**
- Interpreted requirements
- Inferred patterns from multiple sources
- Contradictory sources requiring synthesis
- Historical context with gaps

## Missing Citations

If information lacks RAG citations:

**Do not:**
- Present information as fact without source
- Reference "the codebase" generically
- Use "we've discussed this before" without artifact link
- Assume knowledge is common across team

**Do:**
- Flag: "Source required"
- Ask: "What document or code supports this?"
- Pause until source is provided
- If source doesn't exist, create artifact

## Verification

Before finalizing content, check:

- [ ] Every factual claim has citation
- [ ] Confidence level is appropriate to claim
- [ ] Source is specific (not generic)
- [ ] Citations link to actual artifacts
- [ ] Contradictory information is flagged

## Example: Well-Cited Decision

```
Decision: Implement OAuth 2.0 for authentication

Requirements:
- [prd:high] Requires "industry-standard authentication"
- [ticket:medium] Customer request for SSO support

Approach:
- Use standard OAuth 2.0 flow per RFC 6749 [external:high]
- Follow security recommendations from OAuth 2.1 draft [external:high]
- Implement as middleware similar to existing service-auth [code:high]

Tradeoffs:
- Complexity over basic auth (justified by [prd:high] requirement)
- Token management overhead (acceptable per [design:medium])

Related:
- Previous auth discussion [meeting:medium] from March sprint
```

## Anti-Patterns

You must not:
- Use "the team knows" as source
- Cite "hallway conversations" or informal chats
- Present speculation without labeling as low confidence
- Assume all team members have same context
- Reference deleted or obsolete artifacts
