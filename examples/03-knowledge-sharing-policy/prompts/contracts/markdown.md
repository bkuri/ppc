---
id: contracts/markdown
desc: Output valid Markdown.
priority: 100
tags: [output:markdown]
---
## Output Contract

Output valid Markdown without extra commentary or formatting artifacts.

## Markdown Requirements

All output must be valid Markdown that can be:
- Rendered by standard Markdown parsers
- Copy-pasted into documentation systems
- Read as plain text without formatting tools

## Formatting Rules

- Use standard Markdown syntax (CommonMark compatible)
- Headers: \`#\`, \`##\`, \`###\` (ATX style preferred)
- Lists: \`-\` for unordered, \`1.\` for ordered
- Code blocks: \`\`\`language for syntax highlighting
- Inline code: \`\` for short references
- Blockquotes: \`>\` for callouts or warnings

## Prohibited Output

You must not include:
- Debug text or meta-commentary
- "Here is the output:" prefixes
- "Done" or completion messages
- Line number markers
- Tool-specific formatting

## Minimal Guarantees

At minimum, ensure:
- All headers are properly closed
- Code blocks have language identifiers
- Lists are properly indented
- Links use \`[text](url)\` format
- Escaped characters are handled correctly

## Validation

If you are unsure your output is valid Markdown, prefer:
- Simpler formatting (fewer nested lists)
- Code blocks over inline code for complex snippets
- Plain text over HTML entities
- Standard syntax over extensions
