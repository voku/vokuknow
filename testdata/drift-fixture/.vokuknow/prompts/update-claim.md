# Prompt: update-claim

## Task
Update or create a claim according to policy and provenance requirements.

## Goal
Generate deterministic workflow output for update-claim.

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
update-claim updates are applied deterministically with provenance, privacy, and auditability preserved.

## Output expectations
- Updated files under .vokuknow/ with clear diffs.
- Audit-safe edits with provenance.
