# Contributing to PPC

Thanks for your interest in contributing to the Prompt Policy Compiler!

This guide will help you understand how to contribute effectively.

## Philosophy

PPC is built on these principles:

- **Simplicity over features**: Prefer boring solutions
- **Correctness over cleverness**: Explicit rules over implicit behavior
- **Determinism**: Identical inputs must always produce identical outputs
- **Fail loudly**: No silent resolutions or implicit overrides

If a feature violates these principles, it doesn't belong in PPC.

## Types of Contributions

### Bug Reports

Found a bug? Please open an issue with:

1. **What you expected to happen**
2. **What actually happened**
3. **Steps to reproduce**
4. **Your environment** (OS, Go version, PPC version)

Example:

```
Title: ppc doctor fails on circular requires

Expected: Error message indicating circular dependency
Actual: Panic with nil pointer dereference

Steps:
1. Create base.md that requires itself
2. Run: ppc doctor --strict

Environment: Linux, Go 1.22, PPC v0.2.0
```

### Feature Requests

Have an idea? Check first:

1. Does it violate PPC's philosophy?
2. Is it already in the ROADMAP?
3. Would it add hidden behavior or implicit resolution?

If your idea passes these tests, open an issue describing:

- **Problem it solves**
- **Why PPC is the right place for it**
- **How it affects output determinism**

### Code Contributions

Want to contribute code? Start here:

#### 1. Check the ROADMAP

Look at [ROADMAP.md](ROADMAP.md) to see what's planned and prioritized.

#### 2. Discuss first (optional but recommended)

For anything beyond bugfixes, open an issue first so we can discuss approach before you invest time.

#### 3. Fork and branch

```bash
git clone https://github.com/yourusername/ppc.git
cd ppc
git checkout -b feature/your-feature-name
```

#### 4. Make changes

Keep these principles in mind:

- **Minimal changes**: Only change what's necessary for the feature
- **Readable code**: Prefer clarity over cleverness
- **No magic**: Explicit over implicit
- **Preserve determinism**: Test that identical inputs → identical outputs

#### 5. Run tests

```bash
go test ./...
```

Make sure all tests pass.

#### 6. Test with examples

If you changed compilation logic, verify examples still work:

```bash
for dir in examples/0*; do
  echo "Testing $dir..."
  cd $dir
  ppc doctor
  ppc explore --profile explore
  cd ../..
done
```

#### 7. Commit with clear messages

```bash
git commit -m "Fix: Doctor reports false positive on reachable modules

The reachability check was not handling transitive requires correctly.
This fix walks the require graph to the full closure before marking
modules as unreachable.

Fixes #123"
```

#### 8. Push and open a pull request

```bash
git push origin feature/your-feature-name
```

Open a PR against `main` with a clear description of your changes.

## Code Style

### Go

- Follow `gofmt` conventions
- Use meaningful variable names (avoid single letters except in loops)
- Keep functions small and focused
- Add comments for public functions and non-obvious logic

### Markdown

- Use `gfm` (GitHub-flavored Markdown)
- Headers: sentence case (not Title Case)
- Lists: `-` for unordered, `1.` for ordered
- Code blocks: specify language (e.g., \`\`\`go)

### YAML

- Use 2-space indentation
- Keep frontmatter minimal
- Document all fields with comments when non-obvious

## Testing Standards

### Unit Tests

- Test public APIs and error cases
- Use table-driven tests for multiple scenarios
- Keep test names descriptive: `TestFunctionName_Scenario`

### Golden Tests

- Use for snapshot testing (exact output comparison)
- Store fixtures in `tests/testdata/`
- Update fixtures only when output changes are intentional

### Example Tests

- Run `ppc doctor --strict` on all examples
- Verify profiles compile without errors
- Check that variable substitution works correctly

## Documentation

When adding features:

1. Update relevant `.md` files
2. Update README if user-facing
3. Add examples if it's a major feature
4. Keep examples in the `examples/` directory
5. Test that examples compile and validate

## Reporting Security Issues

**Do not** open public issues for security vulnerabilities.

Email security concerns to the maintainer privately.

## Release Process

Releases follow semantic versioning: `vX.Y.Z`

- **X (Major)**: Breaking changes
- **Y (Minor)**: New features, backward compatible
- **Z (Patch)**: Bug fixes

The maintainer will:

1. Update CHANGELOG.md
2. Update version in code
3. Tag release with `git tag`
4. Push to GitHub
5. Automated CI builds and publishes artifacts

## Getting Help

- **Questions?** Open a discussion or issue
- **Want to discuss design?** Open an issue with `[RFC]` prefix
- **Need clarification?** Comment on existing issues

## Recognition

Contributors will be recognized in:

- Release notes
- CONTRIBUTORS file (if added)
- GitHub's contributors page

Thank you for helping make PPC better! 🎉
