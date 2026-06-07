# AGENTS.md

## Purpose

This folder demonstrates local AI instructions for a small Python package.

## Local Rules

- Keep application code under `src/task_service`.
- Use dataclasses for simple domain objects.
- Raise `TaskNotFound` for missing task IDs.
- Do not introduce web frameworks; this example is about project context, not HTTP.

## Verification

```bash
PYTHONPATH=src python -m unittest discover -s tests
PYTHONPATH=src python -m task_service
```

## Change Notes

- Tests should read like usage examples for developers and agents.
