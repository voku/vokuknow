# Prompt: query

## Task
Query the existing wiki and return evidence-backed answers.

## Goal
Generate deterministic workflow output for query.

## Required steps
1. Read .vokuknow/schema/AGENTS.md and relevant policy YAML.
2. Inspect existing memory artifacts, sources, and digests before changes.
3. Apply provenance requirements: captured_by, confidence, observed_at, source_ref.
4. Enforce contradiction resolution: prefer newer authoritative claims with stronger observation support.
5. Enforce privacy boundary and redaction.
6. If you must decide because docs or skills are insufficient, append an entry to .vokuknow/audit/decision-log.md and either update the missing guidance in the same task or leave a concrete follow-up in that entry.
7. Write auditable updates with deterministic formatting.
8. If work is complete, crystallize reusable lessons into digests.

## Constraints
- Deterministic output only.
- No unresolved placeholders.
- Keep shared/private boundaries intact.
- Prefer recent authoritative sources with multiple observations.

## Done condition
query updates are applied deterministically with provenance, privacy, and auditability preserved.

## Output expectations
- Updated files under .vokuknow/ with clear diffs.
- Decision-log updates when missing guidance forced a material choice, plus either same-task guidance updates or explicit follow-up notes.
- Audit-safe edits with provenance.
