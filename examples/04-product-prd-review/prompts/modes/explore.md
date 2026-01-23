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

When reviewing PRDs for {{product_name}}:

1. **Survey the space**: List 3-5 distinct product approaches
2. **Compare tradeoffs**: For each approach, identify:
   - Problem fit and user value
   - Technical feasibility
   - Timeline and resource requirements
   - Risk level ({{risk_level}})

3. **Recommend**: Choose one approach and explain why
   - Be explicit about tradeoffs being accepted
   - Link to success criteria
   - Provide fallback options if scope changes

## Anti-Patterns

You must not:
- Present the first approach as the only option
- Skip tradeoff analysis
- Recommend without linking to success criteria
- Mix approaches without clear separation

## Output Structure

For each approach:
```
### Approach: <Name>

**How it solves the problem:**
<1-2 sentences>

**Pros:**
- <list>

**Cons:**
- <list>

**Success criteria addressed:**
<list linking to PRD acceptance criteria>

**When to choose:**
<1-2 sentences>
```

Then:
```
## Recommendation

Use: <Approach Name>

**Why:**
<direct reasoning linking to success criteria>

**What we're trading away:**
<clear admission of accepted costs>

**Risk assessment:**
<how this approach mitigates {{risk_level}} risk>
```
