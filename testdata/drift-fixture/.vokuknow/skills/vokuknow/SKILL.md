# vokuknow

Use this skill to capture durable repo-local memory after non-trivial discovery, debugging, refactoring, or implementation work.

## Required behavior

- Save durable learnings after non-trivial work.
- Prefer updating an existing memory artifact over creating a duplicate.
- Separate facts, hypotheses, and open questions.
- Always attach evidence paths for every durable claim.
- Never save secrets, credentials, tokens, personal data, or other sensitive material.
- Create a handoff when work is incomplete or another agent is likely to continue.
- Treat task complete + memory captured as the real done condition.

## What to capture

Capture only learnings that are likely to help the next agent:

- ownership boundaries
- invariants and constraints
- coupling between files or systems
- root causes and failed assumptions
- reusable debugging warnings
- implementation lessons that change how future edits should be made

Do not save transient scratch notes, speculative summaries without evidence, or redundant copies of existing memory.

## Artifact selection

- Use a discovery artifact after exploration or code reading.
- Use a claim artifact for a durable statement that is supported by evidence.
- Use a digest artifact after debugging or implementation clarifies a reusable lesson.
- Use a handoff artifact when work is incomplete or continuation is likely.

## Write-back rules

- Update existing memory in the same area or logical subsystem before creating a new file.
- Use clear titles tied to the code area or lesson.
- Record evidence as repo-local paths.
- Mark confidence explicitly.
- Distinguish confirmed facts from hypotheses or unresolved questions.
- Keep artifacts concise, specific, and reusable by future agents.
