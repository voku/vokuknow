# Prompts

Generated prompts live in `.vokuknow/prompts/`.

Default prompts:

- `ingest.md`
- `query.md`
- `update-claim.md`
- `resolve-claim.md`
- `crystallize.md`
- `handoff.md`
- `lint.md`

Each prompt includes task, goal, required steps, constraints, done condition, and output expectations.

Generated prompts also tell agents to append `.vokuknow/audit/decision-log.md` entries when they have to make material decisions without enough repo guidance, then close the guidance gap in the same task or leave a concrete follow-up.
