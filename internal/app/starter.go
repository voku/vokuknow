package app

func starterFiles() map[string]string {
	return map[string]string{
		".vokuknow/schema/AGENTS.md": `# AGENTS Contract

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
`,
		".vokuknow/schema/entity_types.yaml": `entities:
  - name: source
    description: Raw or normalized input evidence
  - name: observation
    description: Structured notes derived from sources
  - name: claim
    description: Testable statement backed by observations
  - name: discovery
    description: Durable learning from exploration that future agents should reuse
  - name: digest
    description: Crystallized lesson from completed work
  - name: handoff
    description: Continuation context for follow-up work
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
update_rule: update existing memory when the code area or logical subsystem already has coverage; create new memory when the scope differs materially
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
  - .vokuknow/memory/claims/private
promotion_rule: private claims may be promoted to shared after explicit review and redaction
`,
		".vokuknow/schema/prompt_profiles.yaml": `profiles:
  default:
    goal_template: "Generate deterministic workflow output for %s."
    done_template: "%s updates are applied deterministically with provenance, privacy, and auditability preserved."
`,
		".vokuknow/skills/vokuknow/SKILL.md": `# vokuknow

Use this skill to capture durable repo-local memory after non-trivial discovery, debugging, refactoring, or implementation work.

## Required behavior

- Save durable learnings after non-trivial work.
- Prefer updating an existing memory artifact over creating a duplicate.
- Separate facts, hypotheses, and open questions.
- Always attach evidence paths for every durable claim.
- Never save secrets, credentials, tokens, personal data, or other sensitive material.
- Create a handoff when work is incomplete or another agent is likely to continue.
- Treat task complete + memory captured as the real done condition.

## What to capture

Capture only learnings that are likely to help the next agent:

- ownership boundaries
- invariants and constraints
- coupling between files or systems
- root causes and failed assumptions
- reusable debugging warnings
- implementation lessons that change how future edits should be made

Do not save transient scratch notes, speculative summaries without evidence, or redundant copies of existing memory.

## Artifact selection

- Use a discovery artifact after exploration or code reading.
- Use a claim artifact for a durable statement that is supported by evidence.
- Use a digest artifact after debugging or implementation clarifies a reusable lesson.
- Use a handoff artifact when work is incomplete or continuation is likely.

## Write-back rules

- Update existing memory in the same area or logical subsystem before creating a new file.
- Use clear titles tied to the code area or lesson.
- Record evidence as repo-local paths.
- Mark confidence explicitly.
- Distinguish confirmed facts from hypotheses or unresolved questions.
- Keep artifacts concise, specific, and reusable by future agents.
`,
		".vokuknow/templates/discovery.md": `# Discovery: <short title>

## Area
<primary code area or subsystem>

## Learned
- <durable learning 1>
- <durable learning 2>

## Evidence
- <repo-local path>
- <repo-local path>

## Confidence
<high|medium|low>

## Future warning
<what the next agent should check before changing this area>

## Open questions
- <question, if any>
`,
		".vokuknow/templates/claim.md": `# Claim: <short title>

## Statement
<durable evidence-backed statement>

## Evidence
- <repo-local path>
- <repo-local path>

## Confidence
<high|medium|low>

## Status
<accepted|provisional|superseded>

## Notes
<important scope limits, caveats, or follow-up>
`,
		".vokuknow/templates/digest.md": `# Digest: <short title>

## Symptom
<what went wrong or what changed>

## Root cause
<confirmed cause>

## Failed assumption
<what the team or agent assumed incorrectly>

## Fix
<what resolved the problem or clarified the implementation>

## Evidence
- <repo-local path>
- <repo-local path>

## Reusable lesson
<what future agents should remember>

## Open questions
- <question, if any>
`,
		".vokuknow/templates/handoff.md": `# Handoff: <short title>

## Current state
<what is done and what remains>

## Must know
- <critical fact 1>
- <critical fact 2>

## Evidence
- <repo-local path>
- <repo-local path>

## Next recommended step
<best next action for the next agent>

## Open questions
- <question, if any>
`,
		".vokuknow/policy/memory_rules.md": `# Memory rules

## Purpose

Memory artifacts exist to reduce repeated discovery and prevent future agents from making avoidable wrong assumptions.

## Save memory when

- exploration revealed durable structure or coupling
- debugging established a root cause or disproved an assumption
- implementation clarified invariants, ownership, or dangerous call paths
- incomplete work needs a clean handoff

## Do not save memory when

- the note is only temporary scratch work
- the content is already captured in an existing artifact
- the statement has no supporting evidence
- the information is sensitive and should not be persisted

## Required qualities

- repo-local and file-based
- specific to a code area or lesson
- evidence-backed
- explicit about confidence
- clear about facts vs. hypotheses vs. open questions
- updated in place when an artifact covering the same area or logical subsystem already exists
`,
		".vokuknow/policy/claim_quality.md": `# Claim quality

A good claim is reusable, scoped, and supported by proof.

## Required

- A single durable statement.
- Evidence paths that a future agent can inspect.
- An explicit confidence level.
- A status that shows whether the claim is accepted, provisional, or superseded.

## Avoid

- compound claims that bundle unrelated facts
- wording that sounds certain without evidence
- vague statements such as "this probably matters"
- claims that duplicate another artifact without adding clarity

## Review questions

- Is the statement specific enough to guide future edits?
- Can another agent verify it from the listed evidence?
- Would updating an existing claim be better than creating a new one?
`,
		".vokuknow/policy/privacy.md": `# Privacy

Never store:

- secrets, credentials, tokens, or keys
- personal data or private user information
- internal URLs, identifiers, or values that should not persist in repo history
- copied logs or dumps that contain sensitive content

## Redaction rules

- Prefer describing the security-sensitive behavior instead of copying the value.
- Replace sensitive literals with redacted placeholders when the lesson still matters.
- If the evidence source contains sensitive information, reference the file path without reproducing the sensitive content.

## Principle

Memory should preserve durable understanding, not confidential data.
`,
		".vokuknow/policy/retention.md": `# Retention

Keep memory artifacts current and useful.

## Update instead of duplicate

When a discovery, claim, digest, or handoff already covers the same code area, update the existing artifact unless a separate file is clearly more useful.

## Supersede stale memory

- mark outdated claims as superseded
- refresh digests when the implementation has changed materially
- remove or resolve handoffs once the continuation is complete

## Periodic review

Review memory artifacts when:

- a subsystem is refactored
- an invariant changes
- an old assumption is disproved
- open questions have been resolved
`,
		".vokuknow/prompts/examples/code-discovery.md": `# Code discovery

Use the code_discovery category to explore an unfamiliar code area and extract durable learnings with evidence.

## Fill these inputs

- CODE_AREA: src/Auth/SessionBootstrap.php
- DISCOVERY_SCOPE: session bootstrap and cookie policy behavior
- EVIDENCE_PATHS: src/Auth/SessionBootstrap.php, tests/Auth/SessionBootstrapTest.php
- MEMORY_KIND: discovery
- MEMORY_TARGET_FILE: .vokuknow/memory/discoveries/session-bootstrap-flow.md
- CONFIDENCE_LEVEL: high
- OPEN_QUESTIONS: Do tenant overrides affect all cookie policy call sites?

## Expected outcome

Write a discovery artifact that records what the next agent should know before changing session bootstrap behavior.
`,
		".vokuknow/prompts/examples/implementation-learning.md": `# Implementation learning

Use the implementation_learning category after a code change reveals new invariants, ownership boundaries, or coupling.

## Fill these inputs

- CODE_AREA: src/Auth/SessionBootstrap.php
- FILES_TOUCHED: src/Auth/SessionBootstrap.php, src/Http/RequestContext.php
- LEARNING_TYPE: implementation lesson
- PRIOR_ASSUMPTION: secure-cookie behavior was controlled only by global config
- UPDATED_ASSUMPTION: secure-cookie behavior is request-context-driven and tenant-aware
- EVIDENCE_PATHS: src/Auth/SessionBootstrap.php, src/Http/RequestContext.php, tests/Auth/SessionBootstrapTest.php
- MEMORY_KIND: digest
- MEMORY_TARGET_FILE: .vokuknow/memory/digests/secure-cookies-after-reverse-proxy.md
- CONFIDENCE_LEVEL: high

## Expected outcome

Write a digest artifact that captures the reusable lesson created by the implementation work.
`,
		".vokuknow/prompts/examples/debugging-digest.md": `# Debugging digest

Use the debugging_digest category after debugging establishes a symptom, root cause, failed assumption, and fix that future agents should reuse.

## Fill these inputs

- CODE_AREA: src/Auth/SessionBootstrap.php
- FILES_TOUCHED: src/Auth/SessionBootstrap.php, src/Http/RequestContext.php, tests/Auth/SessionBootstrapTest.php
- PRIOR_ASSUMPTION: proxy changes should not affect session cookies
- UPDATED_ASSUMPTION: proxy headers can change effective scheme handling and cookie behavior
- EVIDENCE_PATHS: src/Auth/SessionBootstrap.php, src/Http/RequestContext.php, tests/Auth/SessionBootstrapTest.php
- MEMORY_KIND: digest
- MEMORY_TARGET_FILE: .vokuknow/memory/digests/secure-cookies-after-reverse-proxy.md
- CONFIDENCE_LEVEL: high

## Expected outcome

Write a digest artifact that records the root cause, failed assumption, fix, and future warning signs.
`,
		".vokuknow/prompts/examples/handoff-memory.md": `# Handoff memory

Use the handoff_memory category when work is incomplete and the next agent needs durable context before continuing.

## Fill these inputs

- CODE_AREA: src/Auth/SessionBootstrap.php
- HANDOFF_SCOPE: auth session cleanup
- FILES_TOUCHED: src/Auth/SessionBootstrap.php, src/Http/RequestContext.php
- EVIDENCE_PATHS: src/Auth/SessionBootstrap.php, src/Http/RequestContext.php, tests/Auth/SessionBootstrapTest.php
- MEMORY_KIND: handoff
- MEMORY_TARGET_FILE: .vokuknow/memory/handoffs/auth-session-cleanup.md
- OPEN_QUESTIONS: Which cookie policy call sites still bypass tenant overrides?
- CONFIDENCE_LEVEL: high

## Expected outcome

Write a handoff artifact that tells the next agent what is done, what must be preserved, and the next recommended step.
`,
	}
}
