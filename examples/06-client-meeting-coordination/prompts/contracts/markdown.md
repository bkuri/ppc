---
id: markdown-contract
title: Markdown Output Contract
description: All output must be valid, well-formatted Markdown
requires: []
tags: [contract:markdown]
---

All output from this policy must be valid CommonMark Markdown with consistent formatting.

## Markdown Requirements

### Headings

- Use `#` for the document title.
- Use `##` for major sections.
- Use `###` for subsections.
- Do not skip heading levels (e.g., don't go from `##` to `####`).

### Code Blocks

- Use triple backticks for code blocks: ```code```
- Include language tags when applicable: ```go, ```yaml, ```md
- Do not indent code with spaces.

### Lists

- Use `-` for unordered lists.
- Use `1.` for ordered lists.
- Start list items on a new line.

### Emphasis and Links

- Use `**bold**` for bold text.
- Use `*italic*` for italic text.
- Use `[text](url)` for links.
- Do not use HTML tags.

### Blockquotes

- Use `>` for blockquotes.
- Place `>` at the start of each line.

### Whitespace

- No trailing whitespace at the end of lines.
- One blank line between paragraphs.
- One blank line between sections.

### Special Characters

- Escape backticks within inline code: \`code\`
- Escape asterisks when not used for emphasis: \*
- Escape underscores when not used for emphasis: \_

## Validation

The output must:

- Parse as valid CommonMark Markdown.
- Render correctly in standard markdown viewers.
- Not contain malformed or ambiguous syntax.
