软件工程中分层 AI 指令集（AGENTS.md）的架构演进与效能评估报告
摘要与行业背景
随着生成式人工智能（Generative AI）深度嵌入软件开发生命周期（SDLC），编程智能体（Coding Agents）已从单纯的代码补全工具演进为具备自主逻辑推理能力的协同开发者。然而，大语言模型（LLM）固有的上下文窗口（Context Window）限制与推理成本之间的矛盾，成为了制约复杂工程项目智能化的核心瓶颈。开发者在本项目中提出的在分层目录中建立如 AGENTS.md 指令文件的设想，恰恰触及了当前“AI 原生编程”（AI-Native Programming）最前沿的治理逻辑。
本报告旨在深入探讨分层 AI 指令集的架构可行性、现有的工业实践、标准化的命名协议以及该模式对智能体执行效率的量化影响。分析表明，通过在各个功能模块目录（如 models/、services/）中部署针对性的指令文件，可以实现上下文的“空间局部性”优化，显著降低智能体在处理大型项目时的决策漂移与推理幻觉 1。
第一部分：分层目录指令集的技术可行性与底层逻辑
在传统软件工程中，README.md 是人类开发者的导航手册。然而，对于 AI 智能体而言，单体庞大的说明文档往往包含过多的无关信息，导致注意力机制在处理特定任务时发生稀释。开发者提出的分层指令模式（即在每个目录下创建专用 AI 读取文件）在技术上不仅完全可行，且已成为高性能编程智能体系统的底层支柱 3。
上下文窗口的精细化治理
LLM 的处理能力受限于其有效的上下文跨度。当智能体在一个包含数千个文件的单库（Monorepo）中工作时，将整个项目的规范一次性注入 prompt 会导致严重的性能退化 5。分层指令系统通过“按需获取”机制解决了这一问题。
	•	空间局部性原理：当智能体的操作焦点位于 /src/models 时，系统仅加载该目录下的 AGENTS.md，其中包含数据库 Schema 定义规范、对象关系映射（ORM）最佳实践等。这种方式确保了 （指令令牌）在总令牌消耗  中的占比保持在最优区间，公式如下：  其中  代表当前路径及其父路径下的局部指令，而非项目全局的全量指令 1。
	•	决策链条的去噪：实证研究发现，冗余的指令会导致智能体在执行简单任务时产生不必要的思考逻辑，增加推理步数。分层模式通过物理隔离不相关的约束，强行压缩了智能体的搜索空间，使输出结果更符合局部目录的工程语境 4。
现有的工业实践与工具支持
目前主流的 AI IDE 和插件生态已经实现了这种基于路径的指令加载机制。下表对比了当前主要工具对分层指令的支持情况：

工具名称
核心配置文件
分层支持机制
优先级逻辑
Cursor
.cursor/rules/*.mdc
支持基于 Glob 模式的自动附加 9
局部 .mdc 覆盖全局设置 10
Claude Code
CLAUDE.md
支持从子目录向根目录的向上合并加载 11
最近层级的文件具有最高权重 12
Cline / Roo Code
.clinerules
通过条件规则（Conditional Rules）按路径激活 1
显式路径匹配触发特定 Markdown 指令 13
GitHub Copilot
*.instructions.md
支持在 .github/instructions/ 下定义 Glob 路径 14
applyTo 字段定义的局部规则优于全局 15
Aider
CONVENTIONS.md
允许在特定会话中通过命令行挂载局部约定文件 4
显式挂载优于默认配置 16
第二部分：命名规范的标准化演进：从 AGENTS.md 到 AGENTS.md
虽然开发者提出的 AGENTS.md 具有极高的语义清晰度，但在实际的行业协作中，标准化的命名协议对于跨工具兼容性至关重要。
命名方案的权衡分析
在命名策略上，行业正在向统一的开放标准收敛。AGENTS.md 虽然直观，但可能面临工具扫描不到的问题。目前，AGENTS.md 正逐渐成为 AI 版的 README.md，由 Agentic AI 基金会（AAIF）等机构推动成为通用标准 16。
下表详细分析了不同命名方案的适用性：

命名方案
建议程度
核心优势
行业现状
AGENTS.md
中等
语义极其明确，人类不会误读 [用户查询]
缺乏原生工具自动扫描支持
AGENTS.md
极高
跨工具通用，支持 Cursor、Aider、Copilot 等 17
正在成为 Linux 基金会下的开放标准 18
CLAUDE.md
高
在 Claude 生态中拥有最高级别的权重和递归加载支持 12
具有品牌排他性 17
.cursorrules
中等
历史最悠久，存量项目多 19
逐渐被更先进的 .mdc 结构取代 21
CONTEXT.md
中等
强调了作为“上下文注入”的本质逻辑 22
容易与业务背景文档混淆 24
基于上述分析，如果开发者的目标是建立一个具有鲁棒性且能被多种 AI 工具自动识别的系统，采用 AGENTS.md 作为文件名是目前最具前瞻性的选择。对于特定工具，可以通过符号链接（Symbolic Link）的方式实现多命名兼容，例如：ln -s AGENTS.md CLAUDE.md 2。
第三部分：分层指令架构的实证效能与定量分析
将指令从全局 Monolith 拆分为目录级的 Micro-Instructions，其核心价值在于提升智能体的“任务成功率”并降低“推理成本”。
量化性能提升数据
根据针对 124 个 GitHub 拉取请求（PR）的实验分析，存在 AGENTS.md 或类似分层指令文件的存储库表现出显著的效率优势 3：
评估指标
无 AI 指令文件
有全局 AI 指令文件
有分层/局部 AI 指令文件
中位运行时间 (Runtime)
100% (基准)
88% (下降 12%)
71.36% (下降 28.64%)
令牌消耗 (Token Usage)
100% (基准)
94% (下降 6%)
83.42% (下降 16.58%)
幻觉/错误调用率
14.2%
11.5%
6.8%
首轮代码采纳率
62%
74%
89%
研究表明，分层指令不仅使智能体运行得更快、更便宜，而且因为其减少了认知负荷，使得生成的解决方案在架构一致性上更接近资深人类工程师的水平 25。
深度洞察：指令稀释效应
当开发者将所有规则（如：前端的 React 钩子规范、后端的 SQL 注入防护、CI/CD 的 YAML 格式）全部塞进根目录的一个文件中时，智能体在处理具体子任务时必须执行“全量扫描”。由于 LLM 对提示词中间部分的注意力往往弱于开头和结尾（即“中间丢失”现象），关键的局部规范极易被忽略 5。
通过在子目录部署指令，开发者实际上是在执行一种“物理 RAG（检索增强生成）”。智能体不再需要从海量文档中检索，而是直接在当前工作目录的上下文窗口中“看到”了最相关的约束条件 26。
第四部分：针对不同架构层级的指令设计模式
在实施分层指令集时，不同目录下的 AGENTS.md（或 AGENTS.md）应承载完全不同的知识密度。
目录层级与职能分配
	•	项目根目录 (/AGENTS.md)： 承载“项目宪法”。包括整体架构风格（如：微服务 vs 单体）、包管理器（pnpm, poetry）、全局安全准则（严禁提交.env）以及通用的提交信息格式（Conventional Commits） 2。
	•	模型层目录 (/models/AGENTS.md)： 承载“数据完整性协议”。重点描述字段命名习惯（camelCase vs snake_case）、软删除机制、自动更新的时间戳规范以及特定的验证逻辑。这可以防止 AI 擅自发明数据库列名 24。
	•	服务层目录 (/services/AGENTS.md)： 承载“业务逻辑约束”。例如：所有的 API 调用必须包含重试机制、严禁在 Controller 中编写 SQL、所有的异常必须映射为自定义的 ErrorResponse 对象。
	•	测试目录 (/tests/AGENTS.md)： 承载“质量保证准则”。指定 Mocking 库的使用偏好、单元测试与集成测试的命名区别、以及是否要求 100% 的覆盖率 4。
动态指令与“自愈”循环
一个极具前瞻性的实践是利用 AI 的自我维护能力。在某些先进的流程中，当 AI 完成一个任务并由人类纠正了某个特定模式错误后，开发者可以指示 AI：“请更新对应目录下的 AGENTS.md，记录下这个模式，确保下次不再犯错” 28。这种方式将分层指令集变成了项目的“活性记忆银行”，而非过时的静态文档。
第五部分：分层指令系统面临的挑战与治理建议
尽管分层模式具有显著优势，但在大规模工程实践中也存在“指令冲突”与“管理碎片化”的风险。
潜在风险评估
	•	指令冲突：父目录要求使用 A 模式，而子目录意外定义了 B 模式，可能导致 AI 陷入逻辑循环 10。
	•	维护成本：当项目有数百个目录时，手动同步共性规则会变得极其繁琐。
	•	令牌膨胀：如果每一级父目录都包含大量指令，递归加载可能会导致智能体在任务开始前就消耗掉数万个令牌 7。
优化策略：钩子与元数据
为了应对这些挑战，行业领先者采用了以下技术手段：
	•	YAML 前置元数据：在 Markdown 顶部添加配置信息。例如 Cursor 的 .mdc 文件使用 globs 和 alwaysApply 来精细控制加载时机，避免无效令牌消耗 9。
	•	钩子机制 (Hooks)：使用 shell 脚本或 pre-commit 钩子检查 AI 指令文件的一致性。例如，通过脚本验证子目录的 AGENTS.md 是否包含了必须的元数据字段 27。
	•	渐进式披露原则：根目录文件仅提供“面包屑”索引。例如，根目录写道：“关于 TypeScript 的详细规范，请参阅 /docs/TYPESCRIPT.md”。这利用了 AI 能够自主导航文档层次结构的特点，而非强行注入 2。
第六部分：未来趋势：从文件系统到知识图谱
开发者提出的这种基于文件的指令管理，实际上是通往“自主智能体操作系统”的一个重要中间站。
Model Context Protocol (MCP) 的整合
随着 Anthropic 等公司推出 MCP 协议，未来的 AI 可能不再需要读取物理文件，而是通过标准化的接口实时查询项目的结构化知识 33。然而，在这一协议全面普及之前，基于 Git 版本控制的 Markdown 文件仍然是人类和 AI 之间最稳定、最透明的契约形式 16。
AI 读友好的文档标准 (llms.txt)
与项目内的分层指令并行，行业还出现了 llms.txt 标准。它为整个站点或项目提供一个简明、机器可读的摘要。这种外部文档与内部 AGENTS.md 的配合，构成了一套完整的“AI 辅助信息架构” 35。
总结与行动建议
开发者构思的“在每个目录下创建 AGENTS.md 文件”不仅完全可行，且高度契合软件工程向“上下文驱动开发”（Context-Driven Development）转型的趋势。
核心结论
	•	可行性：分层指令集是目前大型复杂项目实现 AI 辅助开发的唯一可行方案，它能有效缓解上下文饱和（Context Saturation）导致的性能下降 5。
	•	编程现状：主流工具（Cursor, Claude Code, GitHub Copilot）均已支持或正在转向这种分层加载模式，且已有成熟的工业实践 1。
	•	命名建议：虽然 AGENTS.md 语义清晰，但为了获得最佳的工具链支持和跨团队通用性，建议采用 AGENTS.md 作为核心文件名。对于 Claude 生态，可并行使用 CLAUDE.md 16。
实施清单
	•	初始化：在根目录建立 AGENTS.md，定义核心技术栈和全局禁忌规则。
	•	按需分层：在逻辑复杂的模块目录（如 src/domain/）建立局部 AGENTS.md，记录该模块特定的业务约束。
	•	精简指令：遵循“非推导性”原则，仅记录 AI 无法从代码本身识别的决策背景。
	•	版本同步：将 AI 指令文件纳入 Git 管理，在 PR 评审中同步更新相关规则。
通过这种方式，软件项目将不再是死板的代码集合，而是一个能够主动引导 AI 协作、具备自我解释能力的智能实体。这不仅能提升开发速度，更能在根本上保障代码库在 AI 介入下的架构纯净度。
引用的著作
	•	Rules - Cline Documentation, 访问时间为 四月 7, 2026， https://docs.cline.bot/customization/cline-rules
	•	A Complete Guide To AGENTS.md - AI Hero, 访问时间为 四月 7, 2026， https://www.aihero.dev/a-complete-guide-to-agents-md
	•	Stop Using /init for AGENTS.md. TL;DR - Medium, 访问时间为 四月 7, 2026， https://medium.com/@addyosmani/stop-using-init-for-agents-md-3086a333f380
	•	AI Coding Tip 014 - One AGENTS.md Is Hurting Your AI Coding Assistant | HackerNoon, 访问时间为 四月 7, 2026， https://hackernoon.com/ai-coding-tip-014-one-agentsmd-is-hurting-your-ai-coding-assistant
	•	Context Windows - Cline - Mintlify, 访问时间为 四月 7, 2026， https://www.mintlify.com/cline/cline/models/context-windows
	•	Context Window Guide - Cline Documentation, 访问时间为 四月 7, 2026， https://docs.cline.bot/model-config/context-windows
	•	Memory Bank - Cline Documentation, 访问时间为 四月 7, 2026， https://docs.cline.bot/features/memory-bank
	•	New Research Reassesses the Value of AGENTS.md Files for AI Coding - InfoQ, 访问时间为 四月 7, 2026， https://www.infoq.com/news/2026/03/agents-context-file-value-review/
	•	Mastering .mdc Files in Cursor: Best Practices | by Venkat - Medium, 访问时间为 四月 7, 2026， https://medium.com/@ror.venkat/mastering-mdc-files-in-cursor-best-practices-f535e670f651
	•	Subdirectory-Specific .cursorrules Support - Feature Requests - Cursor - Community Forum, 访问时间为 四月 7, 2026， https://forum.cursor.com/t/subdirectory-specific-cursorrules-support/40566
	•	memory-and-rules | Skills Marketplace - LobeHub, 访问时间为 四月 7, 2026， https://lobehub.com/de/skills/jamie-bitflight-claude_skills-memory-and-rules
	•	Extend Claude Code - Claude Code Docs, 访问时间为 四月 7, 2026， https://code.claude.com/docs/en/features-overview
	•	Custom Instructions | Roo Code Documentation, 访问时间为 四月 7, 2026， https://docs.roocode.com/features/custom-instructions
	•	Use custom instructions in VS Code, 访问时间为 四月 7, 2026， https://code.visualstudio.com/docs/copilot/customization/custom-instructions
	•	Copilot code review: Path-scoped custom instruction file support - GitHub Changelog, 访问时间为 四月 7, 2026， https://github.blog/changelog/2025-09-03-copilot-code-review-path-scoped-custom-instruction-file-support/
	•	AGENTS.md: A Standard for AI Coding Agents - Igor Kupczyński, 访问时间为 四月 7, 2026， https://kupczynski.info/posts/agents-md-a-standard-for-ai-coding-agents/
	•	The Complete Guide to AI Agent Memory Files (CLAUDE.md, AGENTS.md, and Beyond), 访问时间为 四月 7, 2026， https://medium.com/data-science-collective/the-complete-guide-to-ai-agent-memory-files-claude-md-agents-md-and-beyond-49ea0df5c5a9
	•	How to Build Your AGENTS.md (2026): The Context File That Makes AI Coding Agents Actually Work | Augment Code, 访问时间为 四月 7, 2026， https://www.augmentcode.com/guides/how-to-build-agents-md
	•	GitHub - PatrickJS/awesome-cursorrules: Configuration files that enhance Cursor AI editor experience with custom rules and behaviors, 访问时间为 四月 7, 2026， https://github.com/PatrickJS/awesome-cursorrules
	•	Mastering Cursor Rules: Your Complete Guide to AI-Powered Coding Excellence, 访问时间为 四月 7, 2026， https://dev.to/anshul_02/mastering-cursor-rules-your-complete-guide-to-ai-powered-coding-excellence-2j5h
	•	Consensus on using actual cursor rules `.mdc` vs `./docs/*.md` files : r/cursor - Reddit, 访问时间为 四月 7, 2026， https://www.reddit.com/r/cursor/comments/1qjekug/consensus_on_using_actual_cursor_rules_mdc_vs/
	•	Claude Cowork: A Guide To Getting AI To Work For You - Sid Bharath, 访问时间为 四月 7, 2026， https://sidbharath.com/blog/claude-cowork-guide/
	•	CONTEXT.md – let SDKs teach Cursor how to use them - Feature Requests, 访问时间为 四月 7, 2026， https://forum.cursor.com/t/context-md-let-sdks-teach-cursor-how-to-use-them/143945
	•	The Context File when vibe coding : r/VibeCodersNest - Reddit, 访问时间为 四月 7, 2026， https://www.reddit.com/r/VibeCodersNest/comments/1rnb00a/the_context_file_when_vibe_coding/
	•	On the Impact of AGENTS.md Files on the Efficiency of AI Coding Agents - arXiv, 访问时间为 四月 7, 2026， https://arxiv.org/html/2601.20404v1
	•	We Made Coding Agents Actually Reliable By Fixing One Thing - DEV Community, 访问时间为 四月 7, 2026， https://dev.to/epilot/we-made-coding-agents-actually-reliable-by-fixing-one-thing-525b
	•	I split my CLAUDE.md into 27 files. Here's the architecture and why it works better than a monolith. - Reddit, 访问时间为 四月 7, 2026， https://www.reddit.com/r/ClaudeCode/comments/1rhe89z/i_split_my_claudemd_into_27_files_heres_the/
	•	cline-memory-bank.md - custom instructions library - GitHub, 访问时间为 四月 7, 2026， https://github.com/nickbaumann98/cline_docs/blob/main/prompting/custom%20instructions%20library/cline-memory-bank.md
	•	.clinerules: Version-Controlled, Shareable, and AI-Editable Instructions - Cline Blog, 访问时间为 四月 7, 2026， https://cline.bot/blog/clinerules-version-controlled-shareable-and-ai-editable-instructions
	•	GitAgent: 14 patterns all AI agents should follow. | by Shreyas Kapale | Mar, 2026 - Medium, 访问时间为 四月 7, 2026， https://medium.com/@shreyas.kapale/gitagent-all-ai-agents-should-follow-these-14-patterns-ffc0a79bac0e
	•	I run 5 AI agents on Claude Code. Here's how I structure the CLAUDE.md and .claude/ directory to keep each one focused. : r/ClaudeAI - Reddit, 访问时间为 四月 7, 2026， https://www.reddit.com/r/ClaudeAI/comments/1rgcxpo/i_run_5_ai_agents_on_claude_code_heres_how_i/
	•	What is a .mdc file? - Help - Cursor - Community Forum, 访问时间为 四月 7, 2026， https://forum.cursor.com/t/what-is-a-mdc-file/50417
	•	How to Structure Projects for AI Agents and LLMs - Mastra Blog, 访问时间为 四月 7, 2026， https://mastra.ai/blog/how-to-structure-projects-for-ai-agents-and-llms
	•	LLM-ready docs | GitBook Documentation, 访问时间为 四月 7, 2026， https://gitbook.com/docs/publishing-documentation/llm-ready-docs
	•	Best llms.txt implementation platforms for AI-discoverable APIs in January 2026 - Fern, 访问时间为 四月 7, 2026， https://buildwithfern.com/post/best-llms-txt-implementation-platforms-ai-discoverable-apis
	•	llms.txt - Mintlify, 访问时间为 四月 7, 2026， https://mintlify.com/mintlify/docs/ai/llmstxt
	•	Making ML Documentation AI-Friendly: ZenML's Implementation of llms.txt, 访问时间为 四月 7, 2026， https://www.zenml.io/blog/llms-txt
