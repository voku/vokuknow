# Skill: wiki-operator

## Purpose
Operate the repo-local wiki with deterministic claim/query/update behavior.

## Repo locations
- Schema: .vokuknow/schema/
- Claims shared: .vokuknow/claims/shared/
- Claims private: .vokuknow/claims/private/
- Sources raw: .vokuknow/sources/raw/
- Sources normalized: .vokuknow/sources/normalized/
- Digests: .vokuknow/digests/
- Audit: .vokuknow/audit/

## Required workflow
1. Inspect existing claims and sources before writing new claims.
2. Preserve provenance fields: source_ref, observed_at.
3. Follow contradiction strategy: prefer recent and authoritative.
4. Enforce private/shared boundary using private paths: .vokuknow/claims/private.
5. Apply redaction rules before storing: token.
6. Append auditable entries for edits, promotions, deletions, and resolutions.
7. Crystallize finished work into digests when complete.

## Quality rules
- Use deterministic language.
- Prefer authoritative, recent, and multi-supported claims.
- Never leave unresolved placeholders.

## Anti-drift constraints
- Do not invent schema fields.
- Do not skip provenance.
- Keep shared/private policy consistent.
