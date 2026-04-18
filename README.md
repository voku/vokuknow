# vokuknow

`vokuknow` is a local-first Go CLI that generates deterministic skills, prompts, and instruction manifests for a repo-local knowledge wiki.

## What it is

- A schema-driven control layer for coding agents
- A deterministic generator for `.vokuknow/` artifacts
- A validator for contract drift and policy issues

## What it is not

- Not a memory server
- Not a vector database platform
- Not a graph engine or hosted SaaS product

## Quick start

```bash
go run ./cmd/vokuknow init
go run ./cmd/vokuknow build
go run ./cmd/vokuknow lint
go run ./cmd/vokuknow doctor
```

## Commands

- `vokuknow init`
- `vokuknow build`
- `vokuknow lint`
- `vokuknow doctor`
- `vokuknow show skill <name>`
- `vokuknow show prompt <name>`
- `vokuknow skill generate <name>`
- `vokuknow prompt generate <name>`

## Why deterministic generation matters

Deterministic prompt and skill generation prevents silent drift, keeps outputs diff-friendly, and makes agent behavior reproducible across reruns.

## Development

See:

- `docs/architecture.md`
- `docs/schema.md`
- `docs/skills.md`
- `docs/prompts.md`
- `docs/development.md`
