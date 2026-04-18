# AGENTS Contract

## Purpose
This repository uses a local-first memory contract to preserve durable, auditable engineering knowledge.

## Operating rules
- Use sources, observations, claims, discoveries, digests, and handoffs as separate knowledge layers.
- Every claim update must include provenance.
- Shared and private knowledge must stay separated by policy.
- Apply secret redaction and sensitive-data filtering before storage.
- Keep an audit trail for edits, deletions, promotions, and contradiction resolutions.
- Resolve contradiction by preferring source recency, source authority, and supporting observations.
- Human override is allowed, but default computed resolution must remain visible.
- Crystallization is required after completed work to capture reusable lessons.
- Anti-drift: follow schema and policy files exactly; do not invent undocumented fields.
