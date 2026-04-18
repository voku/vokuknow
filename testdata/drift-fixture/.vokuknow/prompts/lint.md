# Prompt: lint

## Task
Lint the knowledge contract and report policy or structure violations.

## Goal
Generate deterministic workflow output for lint.

## Required steps
1. Read .vokuknow/schema/AGENTS.md and relevant policy YAML.
2. Inspect existing wiki claims, sources, and digests before changes.
3. Apply provenance requirements: captured_by, confidence, observed_at, source_ref.
4. Enforce contradiction resolution: prefer newer authoritative claims with stronger observation support.
5. Enforce privacy boundary and redaction.
6. Write auditable updates with deterministic formatting.
7. If work is complete, crystallize reusable lessons into digests.

## Constraints
- Deterministic output only.
- No unresolved placeholders.
- Keep shared/private boundaries intact.
- Prefer recent authoritative sources with multiple observations.

## Done condition
lint updates are applied deterministically with provenance, privacy, and auditability preserved.

## Output expectations
- Updated files under .vokuknow/ with clear diffs.
- Audit-safe edits with provenance.
