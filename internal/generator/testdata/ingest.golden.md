# Prompt: ingest

## Task
Ingest new source material into normalized wiki knowledge artifacts.

## Goal
Goal for ingest

## Required steps
1. Read .vokuknow/schema/AGENTS.md and relevant policy YAML.
2. Inspect existing wiki claims, sources, and digests before changes.
3. Apply provenance requirements: source_ref, observed_at.
4. Enforce contradiction resolution: prefer recent and authoritative.
5. Enforce privacy boundary and redaction.
6. Write auditable updates with deterministic formatting.
7. If work is complete, crystallize reusable lessons into digests.

## Constraints
- Deterministic output only.
- No unresolved placeholders.
- Keep shared/private boundaries intact.
- Prefer recent authoritative sources with multiple observations.

## Done condition
Done for ingest

## Output expectations
- Updated files under .vokuknow/ with clear diffs.
- Audit-safe edits with provenance.
