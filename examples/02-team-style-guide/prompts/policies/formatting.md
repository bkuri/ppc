---
id: policies/formatting
desc: Output formatting standards.
priority: 11
tags: []
---
## Formatting Rules

All output must follow these formatting standards.

## Code Blocks

Code blocks must:
- Specify language: \`\`\`go, \`\`\`typescript, \`\`\`bash
- Be copy-paste runnable (no placeholders like `// your code here`)
- Include imports when showing full files
- Show the complete working example, not snippets

Example:
\`\`\`typescript
import { useState } from 'react';

function Counter() {
  const [count, setCount] = useState(0);
  return <button onClick={() => setCount(count + 1)}>{count}</button>;
}
\`\`\`

## Headings

- Use sentence case (not Title Case)
- H1: Main title (one per document)
- H2: Major sections
- H3: Subsections
- H4+: Rare — if needed, consider restructuring

Incorrect:
> # How To Configure The System
> ## Set Up Your Environment
> ### Installation Steps

Correct:
> # How to configure the system
> ## Set up your environment
> ### Installation steps

## Lists

**Unordered lists:** Use hyphens (\`-\`)
- Item one
- Item two
  - Nested item

**Ordered lists:** Use \`1.\`
1. First step
2. Second step
3. Third step

Do not mix formats in the same list.

## Inline Code

Use backticks for:
- Function names: \`useEffect\`
- Variable names: \`isAuthenticated\`
- File paths: \`src/utils/auth.ts\`
- CLI commands: \`npm install\`
- Config keys: \`mode: explore\`
- CSS properties: \`display: flex\`

Do not use backticks for emphasis.

## Links

- Use descriptive anchor text: [React documentation](https://react.dev)
- Do not use raw URLs: https://react.dev
- Use \`<code>\` styling for URLs in monospaced contexts: \`https://react.dev\`

## Blockquotes

Use blockquotes only for:
- Critical warnings: > **Warning:** This operation cannot be undone.
- Security considerations: > This endpoint requires authentication.
- Non-obvious edge cases: > Note: This only works in Node.js 18+.

Do not use for:
- Emphasis
- Quoting people
- Formatting decorations

## Tables

Use only when presenting structured data that doesn't fit as lists:

| Component | Props | Status |
|-----------|-------|--------|
| Button | \`label\`, \`onClick\` | Stable |
| Modal | \`isOpen\`, \`onClose\` | Beta |

Avoid tables for:
- Step-by-step instructions (use ordered lists)
- Feature comparisons with complex criteria (use prose)
- Code examples (use code blocks)

## Horizontal Rules

Use to separate major sections:
- Before "Related resources" section
- Between distinct, unrelated topics
- After code examples to separate from following text

Do not use for:
- Decoration
- Frequent section breaks (reorganize instead)

## Whitespace

- Single blank line between paragraphs
- Two blank lines before H2, one blank line before H3
- No trailing whitespace
- No consecutive blank lines
