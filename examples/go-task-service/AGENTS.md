# AGENTS.md

## Purpose

This folder demonstrates local AI instructions for a small Go service.

## Local Rules

- Keep business logic in package `task`; keep the HTTP entry point in `cmd/server`.
- Do not add external dependencies for this example.
- Prefer explicit errors over panics.
- Tests should cover task creation, completion, and not-found behavior.

## Verification

```bash
go test ./...
go run ./cmd/server
```

## Change Notes

- This example is intentionally small so the `AGENTS.md` pattern remains visible.
