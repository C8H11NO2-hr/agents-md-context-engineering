# AGENTS.md

## Purpose

This folder demonstrates local AI instructions for a small JavaScript package.

## Local Rules

- Use modern ESM syntax.
- Keep the task store dependency-free.
- Throw `TaskNotFoundError` for missing task IDs.
- Tests use Node's built-in test runner.

## Verification

```bash
npm test
npm start
```

## Change Notes

- Do not add a bundler or transpiler; the example should stay readable.
