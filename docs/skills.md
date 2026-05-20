# Skills

Generated schema-driven skills live in `.vokuknow/skills/`.

Default generated skills:

- `wiki-operator.md`
- `crystallizer.md`
- `contradiction-resolver.md`

Bootstrapped memory skill:

- `.vokuknow/skills/vokuknow/SKILL.md`

The bootstrapped `vokuknow` skill now uses a workflow-first format inspired by agent skill packs such as `kepano/obsidian-skills`:

- explicit workflow sections for discovery, debugging/implementation, and handoff
- explicit blind-spot logging into `.vokuknow/audit/decision-log.md` when the agent has to decide for itself
- explicit review of new decision-log entries so guidance is updated or a concrete follow-up is left behind
- a validation checklist before finalizing write-back
- clear privacy and anti-drift constraints
- references to templates, policies, and memory store paths

Generated skills remain deterministic and schema-derived; the `vokuknow` skill provides the practical operator workflow used during memory capture.
