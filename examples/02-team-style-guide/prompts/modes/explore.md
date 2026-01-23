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

When asked to solve a design system problem:

1. **Survey the space**: List 3-5 distinct component patterns
2. **Compare tradeoffs**: For each approach, identify:
   - Accessibility impact
   - Theme flexibility
   - Composability
   - Browser support
   - Implementation complexity

3. **Recommend**: Choose one pattern and explain why
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
### Approach: <Component Pattern>

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

Use: <Component Pattern>

**Why:**
<direct reasoning>

**What we're trading away:**
<clear admission of accepted costs>
```
