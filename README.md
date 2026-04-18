# vokuknow

`vokuknow` is a memory-forcing toolkit for coding agents.

It is designed to work with `vokuprompt`, which owns deterministic prompt compilation. `vokuknow` supplies the memory-specific skill, templates, policies, examples, and repo-local artifact conventions that force durable write-back after meaningful work.

## Product boundary

### `vokuprompt` owns

- deterministic prompt compilation
- category registry
- pattern composition
- placeholder manifests
- the generic execution contract

### `vokuknow` owns

- the memory-specific skill
- memory capture templates
- memory policy guidance
- example prompt workflows
- repo-local memory artifact conventions
- optional thin helper tooling later if it proves necessary

`vokuknow` is not a hosted memory system, vector database, graph platform, or new prompt compiler.

## Goal

After code discovery, debugging, refactoring, or implementation, the agent must write back the durable learnings that would help the next agent understand the codebase faster and make fewer wrong assumptions.

The toolkit enforces four things:

1. **Capture**: save durable learnings after meaningful work.
2. **Structure**: store them in consistent repo-local artifacts.
3. **Evidence**: separate facts from guesses and attach proof.
4. **Reuse**: make future agents load those learnings before they start hacking.

## Core rule

> After meaningful discovery or implementation work, the agent must save durable learnings before considering the task fully complete.

## Repository contents

- `skills/vokuknow/SKILL.md`: memory write-back rules for agents
- `templates/`: structured templates for discoveries, claims, digests, and handoffs
- `policy/`: guidance for memory quality, evidence, privacy, and retention
- `prompts/examples/`: example memory-oriented prompt inputs
- `examples/.vokuknow/memory/`: example repo-local memory artifacts

## Repo-local memory layout

```text
.vokuknow/
  memory/
    discoveries/
    claims/
    digests/
    handoffs/
    open-questions/
```

Keep the storage model file-based and local to the repository. Start with Markdown artifacts before adding any helper automation.

## Recommended workflow

1. Use `vokuprompt` to choose a task category.
2. Do the actual discovery, debugging, refactoring, or implementation work.
3. Detect when durable learning happened.
4. Run a memory-oriented prompt workflow.
5. Save a structured artifact into `.vokuknow/memory/`.
6. Load relevant artifacts before future work in the same area.

## Artifact types

- **Discovery**: what exploration revealed and why it matters later
- **Claim**: a durable, evidence-backed statement worth reusing
- **Digest**: the reusable lesson from a debugging or implementation effort
- **Handoff**: what the next agent must know before continuing

## Using this toolkit with `vokuprompt`

The intended pairing is:

- `vokuprompt` compiles the memory-oriented contract
- `vokuknow` provides the skill, templates, policies, and file conventions used to write the result back

See `prompts/examples/` for example prompt inputs and `examples/.vokuknow/memory/` for example outputs.

## Future scope

If lightweight automation proves useful, `vokuknow` may later grow thin helper tooling to:

- scaffold memory files
- lint required sections
- verify evidence presence
- detect likely duplicates
- flag secret-looking strings

That tooling is secondary. The primary product is the content pack that forces durable, evidence-backed write-back.
