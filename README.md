# vokuknow

`vokuknow` is a DX-first CLI for bootstrapping deterministic repo-local knowledge artifacts for coding agents.

It is designed to pair with `vokuprompt`: `vokuprompt` compiles the task contract, while `vokuknow` provides the memory-specific skill, templates, policies, examples, and file conventions that force durable write-back and decision logging when the agent must choose without enough guidance.

## Quick start

### Install

- [Download the latest release assets](https://github.com/voku/vokuknow/releases/latest)
- Direct download links:
  - [Linux amd64](https://github.com/voku/vokuknow/releases/latest/download/vokuknow_linux_amd64.tar.gz)
  - [Linux arm64](https://github.com/voku/vokuknow/releases/latest/download/vokuknow_linux_arm64.tar.gz)
  - [macOS amd64](https://github.com/voku/vokuknow/releases/latest/download/vokuknow_darwin_amd64.tar.gz)
  - [macOS arm64](https://github.com/voku/vokuknow/releases/latest/download/vokuknow_darwin_arm64.tar.gz)
  - [Windows amd64](https://github.com/voku/vokuknow/releases/latest/download/vokuknow_windows_amd64.zip)
  - [Checksums](https://github.com/voku/vokuknow/releases/latest/download/checksums.txt)
- Or install from source with Go:

```bash
go install github.com/voku/vokuknow/cmd/vokuknow@latest
```

### Bootstrap a repository

Run these commands from the repository you want to prepare for agent memory write-back:

```bash
vokuknow init
vokuknow build
vokuknow lint
vokuknow doctor
```

What you get:

- `.vokuknow/skills/vokuknow/SKILL.md` with the write-back contract
- `.vokuknow/templates/` with structured artifact templates
- `.vokuknow/policy/` with evidence, privacy, and retention rules
- `.vokuknow/prompts/examples/` with memory-oriented prompt examples
- `.vokuknow/audit/decision-log.md` with append-only blind-spot entries for undocumented agent decisions
- `.vokuknow/memory/` with directories for discoveries, claims, digests, and handoffs

## Why this exists

When an agent must decide something because the docs, skills, or schema do not say enough, that gap should be logged immediately. After code discovery, debugging, refactoring, or implementation, the agent should also save the durable learning that would help the next agent move faster and make fewer wrong assumptions.

The toolkit enforces four things:

1. **Capture**: log blind spots during work and save durable learnings after meaningful work.
2. **Structure**: store them in consistent repo-local artifacts.
3. **Evidence**: separate facts from guesses and attach proof.
4. **Reuse**: load the relevant artifacts before future work in the same area.

## Core rule

> When the agent must decide for itself because guidance is missing, it must log that blind spot immediately, then save the durable learning before considering the task fully complete.

## Commands

- `vokuknow init`: scaffold the `.vokuknow/` layout and starter schema
- `vokuknow build`: generate deterministic skills, prompts, and manifests from schema
- `vokuknow lint`: verify schema and generated artifact consistency
- `vokuknow doctor`: run broad health checks across structure, schema, generated artifacts, and policy consistency

## Repo-local memory layout

```text
.vokuknow/
  audit/
  policy/
  prompts/examples/
  skills/vokuknow/
  templates/
  memory/
    discoveries/
    claims/
      private/
    digests/
    handoffs/
```

Keep the storage model file-based and local to the repository. Start with Markdown artifacts before adding helper automation.

## Recommended workflow

1. Use `vokuprompt` to choose a task category.
2. Do the actual discovery, debugging, refactoring, or implementation work.
3. Log self-directed decisions in `.vokuknow/audit/decision-log.md` while work is in flight.
4. Detect when durable learning happened.
5. Run a memory-oriented prompt workflow.
6. Save a structured artifact into `.vokuknow/memory/`.
7. Load relevant artifacts before future work in the same area.

## Artifact types

- **Discovery**: what exploration revealed and why it matters later
- **Claim**: a durable, evidence-backed statement worth reusing
- **Digest**: the reusable lesson from a debugging or implementation effort
- **Handoff**: what the next agent must know before continuing
- **Decision log**: append-only records of undocumented choices that expose missing docs, skills, or policies

## Repository contents

- `skills/vokuknow/SKILL.md`: memory write-back rules for agents
- `templates/`: structured templates for discoveries, claims, digests, handoffs, and decision entries
- `policy/`: guidance for memory quality, evidence, privacy, and retention
- `prompts/examples/`: example memory-oriented prompt inputs
- `audit/decision-log.md`: append-only record of guidance gaps that forced agent decisions
- `examples/.vokuknow/`: worked example repo-local memory and audit artifacts

## Product boundary

### `vokuprompt` owns

- deterministic prompt compilation
- category registry
- pattern composition
- placeholder manifests
- the generic execution contract
- no repo-local memory store or hosted knowledge system

### `vokuknow` owns

- the memory-specific skill
- memory capture templates
- memory policy guidance
- example prompt workflows
- repo-local memory artifact conventions
- optional thin helper tooling later if it proves necessary

`vokuknow` is not a vector database, graph platform, or new prompt compiler.

## Future scope

If lightweight automation proves useful, `vokuknow` may later grow thin helper tooling to:

- scaffold memory files
- lint required sections
- verify evidence presence
- detect likely duplicates
- flag secret-looking strings

That tooling is secondary. The primary product is the memory toolkit that forces durable, evidence-backed write-back.
