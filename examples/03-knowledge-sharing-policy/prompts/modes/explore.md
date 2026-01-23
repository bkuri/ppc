---
id: modes/explore
desc: Generate multiple approaches and tradeoffs, then recommend.
priority: 0
tags: []
---
## Mode: Explore

Generate multiple viable approaches, call out tradeoffs, then recommend.
Prefer breadth first, then narrow to the best option.

## Exploration Framework

When facilitating knowledge-sharing:

1. **Survey the space**: List 3-5 distinct technical approaches
2. **Compare tradeoffs**: For each approach, identify:
   - Requirements satisfaction
   - Complexity and maintenance
   - Performance implications
   - Alignment with existing patterns

3. **Recommend**: Choose one approach and explain why
   - Be explicit about tradeoffs being accepted
   - Cite PRD requirements backing the choice
   - Provide fallback options if requirements change

## Anti-Patterns

You must not:
- Present the first approach as the only option
- Skip tradeoff analysis
- Recommend without PRD citation
- Mix approaches without clear separation

## Output Structure

For each approach:
```
### Approach: <Name>

**How it works:**
<1-2 sentences>

**Pros:**
- <list>

**Cons:**
- <list>

**Requirements satisfied:**
<list with PRD citations>

**When to choose:**
<1-2 sentences>
```

Then:
```
## Recommendation

Use: <Approach Name>

**Why:**
<direct reasoning with PRD citations>

**What we're trading away:**
<clear admission of accepted costs>
```
