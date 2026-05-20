# Skill: crystallizer

## Purpose
Distill completed work into durable digests and reusable lessons.

## Repo locations
- Schema: .vokuknow/schema/
- Memory discoveries: .vokuknow/memory/discoveries/
- Memory claims: .vokuknow/memory/claims/
- Memory private claims: .vokuknow/memory/claims/private/
- Memory digests: .vokuknow/memory/digests/
- Memory handoffs: .vokuknow/memory/handoffs/
- Sources raw: .vokuknow/sources/raw/
- Sources normalized: .vokuknow/sources/normalized/
- Audit: .vokuknow/audit/

## Required workflow
1. Inspect existing memory artifacts and sources before writing new claims or digests.
2. Preserve provenance fields: captured_by, confidence, observed_at, source_ref.
3. Follow contradiction strategy: prefer newer authoritative claims with stronger observation support.
4. Enforce private/shared boundary using private paths: .vokuknow/memory/claims/private.
5. Apply redaction rules before storing: api_key, password, token.
6. Log self-directed decisions and missing guidance in .vokuknow/audit/decision-log.md.
7. Append auditable entries for edits, promotions, deletions, and resolutions.
8. Crystallize finished work into digests when complete.

## Quality rules
- Use deterministic language.
- Prefer authoritative, recent, and multi-supported claims.
- Never leave unresolved placeholders.

## Anti-drift constraints
- Do not invent schema fields.
- Do not skip provenance.
- Keep shared/private policy consistent.
