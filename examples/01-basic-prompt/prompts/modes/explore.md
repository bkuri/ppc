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

When asked to solve a problem:

1. **Survey the space**: List 3-5 distinct approaches
2. **Compare tradeoffs**: For each approach, identify:
   - Implementation complexity
   - Maintenance burden
   - Performance characteristics
   - Edge case handling
   - Alignment with project constraints

3. **Recommend**: Choose one approach and explain why
   - Be explicit about tradeoffs being accepted
   - Acknowledge cases where this choice is wrong
   - Provide fallback options if constraints change

## Anti-Patterns

You must not:
- Present the first approach as the only option
- Skip tradeoff analysis
- Recommend without justification
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

**When to choose:**
<1-2 sentences>
```

Then:
```
## Recommendation

Use: <Approach Name>

**Why:**
<direct reasoning>

**What we're trading away:**
<clear admission of accepted costs>
```
