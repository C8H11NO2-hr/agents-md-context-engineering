# AGENTS.md

## Purpose

This repository documents and demonstrates layered `AGENTS.md` files for AI-assisted software development.

## Project Rules

- Use `AGENTS.md` as the canonical AI instruction filename.
- Do not introduce alternative names unless documenting compatibility through symlinks or tool adapters.
- Keep examples small, runnable, and focused on context engineering rather than framework complexity.
- Prefer clear Markdown over clever formatting.
- When editing generated or converted text, preserve the author's argument while improving consistency.

## Documentation Style

- Write for working programmers.
- Favor practical rules, examples, and review checklists.
- Distinguish evidence, recommendation, and speculation.
- Avoid claiming universal performance gains unless backed by cited research or clearly framed as a hypothesis.

## Example Project Expectations

- Each example must have a local `AGENTS.md`.
- Each example should include at least one command that verifies behavior.
- Keep dependencies minimal.
- Show language-specific conventions without making the examples large.

## Verification

Run what is available on the local machine:

```bash
go test ./examples/go-task-service/...
cd examples/python-task-service && PYTHONPATH=src python -m unittest discover -s tests
cd examples/js-task-service && npm test
```
