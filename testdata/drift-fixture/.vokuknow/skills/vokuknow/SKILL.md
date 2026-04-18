---
name: vokuknow
description: Capture durable repo-local memory after meaningful discovery, debugging, refactoring, or implementation work.
---

# vokuknow memory skill

Use this skill when work created reusable learning that future agents should load before editing the same area.

## Core rule

Treat task complete + memory captured as the real done condition for non-trivial work.

## Workflow: capture discovery memory

1. Review existing memory in the same code area before writing anything new.
2. Extract durable learnings only (ownership, invariants, coupling, failed assumptions, warning signs).
3. Select artifact type:
   - discovery for exploration findings
   - claim for durable evidence-backed statements
4. Update an existing artifact if the area is already covered; only create a new file when scope is materially different.
5. Record evidence as repo-local paths and separate facts from hypotheses/open questions.
6. Validate with the checklist before finalizing.

## Workflow: capture debugging or implementation lessons

1. Capture the symptom/change, root cause or invariant, failed prior assumption, and reusable lesson.
2. Save to a digest artifact when the lesson should shape future edits.
3. Include touched files as evidence and mark confidence.
4. If the lesson invalidates old memory, update or supersede the older artifact.
5. Validate with the checklist before finalizing.

## Workflow: handoff incomplete work

1. Create or update a handoff artifact when continuation is likely.
2. Record current state, must-know constraints, evidence, next recommended step, and open questions.
3. Keep instructions specific enough that the next agent can resume without re-discovery.
4. Validate with the checklist before finalizing.

## Validation checklist

- Memory content is durable and reusable, not scratch notes.
- Existing artifact reuse was considered before creating a new file.
- Every durable claim includes repo-local evidence paths.
- Facts, hypotheses, and open questions are clearly separated.
- Confidence is explicit.
- No secrets, credentials, tokens, personal data, or sensitive literals are stored.

## Privacy and anti-drift constraints

- Keep shared/private boundaries intact (.vokuknow/memory/claims/private for private claims).
- Redact or summarize sensitive source content instead of copying it.
- Do not invent undocumented schema fields or placeholder sections.
- Keep artifacts concise and code-area scoped so they remain maintainable.

## References

- Templates: .vokuknow/templates/
- Policy: .vokuknow/policy/
- Memory store: .vokuknow/memory/
