> [中文版本](principles.zh-CN.md)
# Principles for Layered AGENTS.md Files

Layered AI instruction files work best when they are treated as context architecture, not as dumping grounds for every preference a team has ever had.

## 1. Put Context Near the Work

The root file should explain the project. A module file should explain the module. A test folder file should explain how tests should be written.

Good local instructions include:

- Domain invariants.
- Naming rules that are not obvious from code.
- Security or correctness traps.
- Local test commands.
- Patterns the team has explicitly rejected.

## 2. Keep the Root Small

The root `AGENTS.md` should answer:

- What is this project?
- What stack does it use?
- What should never be changed casually?
- How should an agent verify work?
- Where should deeper rules be found?

If the root file grows too large, move material into `docs/` or a local `AGENTS.md`.

## 3. Prefer Constraints Over Preferences

"Use clear names" is weak because it does not change behavior.

"All external API errors must be converted into `DomainError` before leaving the service layer" is strong because it tells the agent how to act.

## 4. Record Review Lessons

When a human repeatedly corrects an AI-generated pattern, update the nearest relevant `AGENTS.md`.

Example:

```md
## Change Notes

- Do not mock the repository interface in integration tests; use the in-memory store from `tests/fixtures`.
```

This turns review feedback into project memory.

## 5. Avoid Token Inflation

Every instruction has a cost. Remove rules that are:

- Already enforced by code or tests.
- Too generic to affect output.
- Duplicated across many folders.
- No longer true after a refactor.

## 6. Make Conflicts Explicit

If a child folder intentionally overrides a parent rule, say so directly.

```md
## Local Override

The root project prefers JSON logs. This CLI folder uses plain text logs because the output is consumed by humans in terminals.
```

## 7. Review Instruction Changes Like Code

Instruction files shape future implementation work. Pull requests that change `AGENTS.md` should explain:

- What recurring problem the change solves.
- Which folder or workflow it affects.
- How the team can tell whether it helped.
