# AGENTS.md Context Engineering

[中文说明](README_zh.md)

A practical open-source guide for using layered `AGENTS.md` files to make AI-assisted programming more reliable, cheaper, and easier to review.

This repository turns a research-style document about AI coding context files into a reusable project for developers who want better vibe coding: less prompt drift, fewer invented conventions, and more useful AI collaboration inside real codebases.

## Why This Exists

AI coding agents do not only need instructions. They need the right instructions at the right layer of the project.

Many teams start with one large root prompt file. That works for small projects, but it becomes noisy as the codebase grows. A layered `AGENTS.md` architecture gives the agent local context near the code it is changing:

- Root `AGENTS.md`: project-wide goals, technology choices, safety rules, and contribution workflow.
- Module `AGENTS.md`: domain-specific architecture, naming rules, invariants, and testing expectations.
- Example `AGENTS.md`: concrete local conventions that are easy for both humans and agents to inspect.

The result is a lightweight form of context engineering: keep global guidance small, move local knowledge close to the code, and let version control preserve the reasoning.

## What Is Inside

```text
.
├── AGENTS.md                         # Root instructions for AI coding agents
├── docs/
│   ├── analysis.zh-CN.md             # Converted source analysis, normalized to AGENTS.md
│   ├── principles.md                 # Design principles for layered context files
│   └── templates/
│       ├── root-AGENTS.md            # Starter root instruction file
│       └── module-AGENTS.md          # Starter module instruction file
├── examples/
│   ├── go-task-service/              # Go example with local AGENTS.md
│   ├── python-task-service/          # Python example with local AGENTS.md
│   └── js-task-service/              # JavaScript example with local AGENTS.md
├── llms.txt                          # LLM-friendly map of this repository
├── CONTRIBUTING.md
├── CODE_OF_CONDUCT.md
├── SECURITY.md
└── LICENSE
```

## Core Ideas

1. **Progressive disclosure**
   Keep the root file short. Put detailed rules in the folder where they matter.

2. **Spatial locality**
   If the agent is editing `services/`, the service-layer rules should be physically close to `services/`.

3. **Non-obvious knowledge only**
   Do not restate what code already shows. Capture decisions, invariants, failure modes, and team preferences.

4. **Versioned project memory**
   Treat `AGENTS.md` files as maintainable project memory. When a human corrects a recurring AI mistake, update the nearest relevant file.

5. **Reviewable AI collaboration**
   Changes to instructions should be reviewed like code because they influence future code.

## Quick Start

Copy the templates into your own project:

```bash
cp docs/templates/root-AGENTS.md AGENTS.md
cp docs/templates/module-AGENTS.md src/your-module/AGENTS.md
```

Then edit them with three questions in mind:

- What should the agent never do here?
- What local conventions are not obvious from the code?
- What commands prove the work is correct?

## Example Projects

Each example is intentionally small. The point is not the task service itself; the point is how local `AGENTS.md` files guide an agent differently in Go, Python, and JavaScript.

### Go

```bash
cd examples/go-task-service
go test ./...
go run ./cmd/server
```

### Python

```bash
cd examples/python-task-service
PYTHONPATH=src python -m unittest discover -s tests
PYTHONPATH=src python -m task_service
```

### JavaScript

```bash
cd examples/js-task-service
npm test
npm start
```

## Recommended `AGENTS.md` Shape

Use short sections with concrete constraints:

```md
# AGENTS.md

## Purpose
What this folder owns.

## Invariants
Rules that must remain true.

## Do Not
Known failure modes and forbidden shortcuts.

## Commands
Local commands for formatting, testing, and verification.

## Change Notes
Small lessons learned from previous AI/human review cycles.
```

## Repository Health

This project includes the files that make an open-source repository easier to adopt and contribute to:

- `README.md` for project framing and onboarding.
- `LICENSE` for clear reuse rights.
- `CONTRIBUTING.md` for contribution workflow.
- `CODE_OF_CONDUCT.md` for community expectations.
- `SECURITY.md` for responsible reporting.
- `CITATION.cff` for research-style attribution.
- `.github` issue and pull request templates.
- `llms.txt` for AI-readable repository navigation.

This structure follows current GitHub and Open Source Guides recommendations that public projects should include basic documentation, contribution guidance, licensing, and community expectations. See [docs/references.md](docs/references.md).

## Source Document

The original Chinese analysis has been converted into Markdown at [docs/analysis.zh-CN.md](docs/analysis.zh-CN.md). The source Word document is preserved under `docs/source/`. All references to the earlier name have been normalized to `AGENTS.md`.

## License

MIT. See [LICENSE](LICENSE).
