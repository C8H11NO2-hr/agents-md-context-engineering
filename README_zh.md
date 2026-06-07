# AGENTS.md 上下文工程

[English](README.md)

这是一个实用的开源指南，介绍如何使用分层 `AGENTS.md` 文件，让 AI 辅助编程更可靠、成本更低，也更容易被团队审查和维护。

本仓库把一份关于 AI 编程上下文文件的研究型文档，整理成一个可复用的开源项目。它面向希望提升 Vibe Coding 质量的开发者：减少提示词漂移，减少 AI 凭空发明项目约定，让 AI 在真实代码库中更像一个理解上下文的协作者。

## 为什么需要这个项目

AI 编程智能体需要的不只是“指令”，而是在正确项目层级上的正确指令。

很多团队会从一个根目录的大型提示词文件开始。这对小项目可行，但当代码库变大后，单个巨大文件会变得嘈杂，AI 在处理局部任务时也更容易被无关信息干扰。分层 `AGENTS.md` 架构把上下文放到代码附近：

- 根目录 `AGENTS.md`：项目目标、技术选型、安全规则、贡献流程。
- 模块级 `AGENTS.md`：领域架构、命名规则、不变量、测试要求。
- 示例级 `AGENTS.md`：人类和 AI 都容易检查的局部约定。

它本质上是一种轻量级上下文工程：全局指导保持简洁，局部知识靠近代码，并让 Git 保存这些团队决策。

## 仓库内容

```text
.
├── AGENTS.md                         # 给 AI 编程智能体的根目录指令
├── README.md                         # 英文 README
├── README_zh.md                      # 中文 README
├── docs/
│   ├── analysis.zh-CN.md             # 从原始中文文档转换而来，并统一为 AGENTS.md
│   ├── principles.md                 # 分层上下文文件的设计原则
│   └── templates/
│       ├── root-AGENTS.md            # 根目录指令模板
│       └── module-AGENTS.md          # 模块目录指令模板
├── examples/
│   ├── go-task-service/              # Go 示例，包含局部 AGENTS.md
│   ├── python-task-service/          # Python 示例，包含局部 AGENTS.md
│   └── js-task-service/              # JavaScript 示例，包含局部 AGENTS.md
├── llms.txt                          # 方便 LLM 阅读的仓库地图
├── CONTRIBUTING.md
├── CODE_OF_CONDUCT.md
├── SECURITY.md
└── LICENSE
```

## 核心思想

1. **渐进式披露**
   根目录文件保持简短，把详细规则放到真正需要它们的目录中。

2. **空间局部性**
   如果 AI 正在编辑 `services/`，服务层规则就应该靠近 `services/`，而不是埋在根目录长文档里。

3. **只记录非显而易见的知识**
   不要重复代码已经表达出来的信息。应记录决策背景、不变量、失败模式和团队偏好。

4. **版本化的项目记忆**
   把 `AGENTS.md` 当成可维护的项目记忆。当人类反复纠正某个 AI 错误模式时，把经验写进最近的相关 `AGENTS.md`。

5. **可审查的 AI 协作**
   指令文件会影响未来代码生成，因此对 `AGENTS.md` 的修改也应该像代码一样被审查。

## 快速开始

把模板复制到你自己的项目中：

```bash
cp docs/templates/root-AGENTS.md AGENTS.md
cp docs/templates/module-AGENTS.md src/your-module/AGENTS.md
```

然后围绕三个问题修改它们：

- AI 在这里绝对不应该做什么？
- 哪些局部约定无法从代码本身直接看出来？
- 哪些命令可以证明修改是正确的？

## 示例项目

每个示例都故意保持很小。重点不是任务服务本身，而是展示 Go、Python、JavaScript 项目中局部 `AGENTS.md` 如何指导 AI 做出不同决策。

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

## 推荐的 `AGENTS.md` 结构

建议使用短小、具体、可执行的章节：

```md
# AGENTS.md

## Purpose
这个目录负责什么。

## Invariants
必须始终保持成立的规则。

## Do Not
已知失败模式和禁止走的捷径。

## Commands
本地格式化、测试和验证命令。

## Change Notes
来自过往 AI / 人类审查循环的小经验。
```

## 开源项目健康度

本项目包含了让开源仓库更容易被采用和贡献的基础文件：

- `README.md`：项目定位和上手说明。
- `README_zh.md`：中文说明。
- `LICENSE`：明确复用权利。
- `CONTRIBUTING.md`：贡献流程。
- `CODE_OF_CONDUCT.md`：社区行为预期。
- `SECURITY.md`：安全问题报告方式。
- `CITATION.cff`：研究型引用信息。
- `.github` issue 和 pull request 模板。
- `llms.txt`：AI 可读的仓库导航。

这个结构参考了 GitHub 和 Open Source Guides 对公开项目的建议：项目应提供基础文档、贡献指南、开源许可证和社区协作预期。参考资料见 [docs/references.md](docs/references.md)。

## 源文档

原始中文分析文档已经转换为 Markdown：[docs/analysis.zh-CN.md](docs/analysis.zh-CN.md)。原始 Word 文档保存在 `docs/source/`。其中早期文件名已经全部统一为 `AGENTS.md`。

## 许可证

MIT。见 [LICENSE](LICENSE)。
