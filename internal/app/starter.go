package app

func starterFiles() map[string]string {
	return map[string]string{
		".vokuknow/schema/AGENTS.md": `# AGENTS Contract

## Purpose
This repository uses a local-first wiki contract to preserve durable, auditable engineering knowledge.

## Operating rules
- Use sources, observations, claims, and digests as separate knowledge layers.
- Every claim update must include provenance.
- Shared and private knowledge must stay separated by policy.
- Apply secret redaction and sensitive-data filtering before storage.
- Keep an audit trail for edits, deletions, promotions, and contradiction resolutions.
- Resolve contradiction by preferring source recency, source authority, and supporting observations.
- Human override is allowed, but default computed resolution must remain visible.
- Crystallization is required after completed work to capture reusable lessons.
- Anti-drift: follow schema and policy files exactly; do not invent undocumented fields.
`,
		".vokuknow/schema/entity_types.yaml": `entities:
  - name: source
    description: Raw or normalized input evidence
  - name: observation
    description: Structured notes derived from sources
  - name: claim
    description: Testable statement backed by observations
  - name: digest
    description: Crystallized lesson from completed work
`,
		".vokuknow/schema/relation_types.yaml": `relations:
  - name: supports
    description: Source or observation supports a claim
  - name: contradicts
    description: Two claims conflict
  - name: derived_from
    description: Artifact was produced from an upstream artifact
  - name: supersedes
    description: New claim replaces previous claim version
`,
		".vokuknow/schema/claim_policies.yaml": `required_provenance_fields:
  - source_ref
  - observed_at
  - captured_by
  - confidence
contradiction_strategy: prefer newer authoritative claims with stronger observation support
update_rule: update existing claim when scope and entity key match; create new claim when scope or key differs
crystallization_required: true
`,
		".vokuknow/schema/source_policies.yaml": `authorities:
  - name: primary_repo
    weight: 100
  - name: official_docs
    weight: 80
  - name: issue_discussion
    weight: 60
`,
		".vokuknow/schema/privacy_policies.yaml": `redaction_patterns:
  - api_key
  - password
  - token
private_paths:
  - .vokuknow/claims/private
promotion_rule: private claims may be promoted to shared after explicit review and redaction
`,
		".vokuknow/schema/prompt_profiles.yaml": `profiles:
  default:
    goal_template: "Generate deterministic workflow output for %s."
    done_template: "%s updates are applied deterministically with provenance, privacy, and auditability preserved."
`,
	}
}
