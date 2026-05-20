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
- Log every material self-directed decision in .vokuknow/audit/decision-log.md when docs, skills, or schema guidance is missing.
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
		".vokuknow/skills/vokuknow/SKILL.md": `---
name: vokuknow
description: Capture durable repo-local memory after meaningful discovery, debugging, refactoring, or implementation work.
---

# vokuknow memory skill

Use this skill when work created reusable learning that future agents should load before editing the same area, or when the current task exposes missing guidance that forced the agent to decide for itself.

## Core rule

Treat task complete + memory captured as the real done condition for non-trivial work.

If you must choose because the docs, skills, prompts, or schema do not say enough, log that blind spot immediately in .vokuknow/audit/decision-log.md.

## Workflow: log blind spots during work

1. Append an entry each time the agent must make a material decision without clear repo guidance.
2. Record the code area, the decision made, why the decision was needed, and what guidance was missing.
3. Add repo-local evidence paths so the gap can be reviewed later.
4. Convert repeated or durable lessons from the log into memory artifacts when the task is complete.

## Workflow: capture discovery memory

1. Review existing memory in the same code area before writing anything new.
2. Extract durable learnings only (ownership, invariants, coupling, failed assumptions, warning signs).
3. Select artifact type:
   - discovery for exploration findings
   - claim for durable evidence-backed statements
4. Update an existing artifact if the area is already covered; only create a new file when scope is materially different.
5. Record evidence as repo-local paths and separate facts from hypotheses/open questions.
6. Validate with the checklist before finalizing.

## Workflow: capture debugging or implementation lessons

1. Capture the symptom/change, root cause or invariant, failed prior assumption, and reusable lesson.
2. Save to a digest artifact when the lesson should shape future edits.
3. Include touched files as evidence and mark confidence.
4. If the lesson invalidates old memory, update or supersede the older artifact.
5. Validate with the checklist before finalizing.

## Workflow: handoff incomplete work

1. Create or update a handoff artifact when continuation is likely.
2. Record current state, must-know constraints, evidence, next recommended step, and open questions.
3. Keep instructions specific enough that the next agent can resume without re-discovery.
4. Validate with the checklist before finalizing.

## Validation checklist

- Memory content is durable and reusable, not scratch notes.
- Existing artifact reuse was considered before creating a new file.
- Every durable claim includes repo-local evidence paths.
- Every material self-directed decision caused by missing guidance is logged in .vokuknow/audit/decision-log.md.
- Facts, hypotheses, and open questions are clearly separated.
- Confidence is explicit.
- No secrets, credentials, tokens, personal data, or sensitive literals are stored.

## Privacy and anti-drift constraints

- Keep shared/private boundaries intact (.vokuknow/memory/claims/private for private claims).
- Redact or summarize sensitive source content instead of copying it.
- Do not invent undocumented schema fields or placeholder sections.
- Keep artifacts concise and code-area scoped so they remain maintainable.

## References

- Templates: .vokuknow/templates/
- Policy: .vokuknow/policy/
- Audit log: .vokuknow/audit/decision-log.md
- Memory store: .vokuknow/memory/
`,
		".vokuknow/templates/decision-entry.md": `# Decision entry: <short title>

## Area
<primary code area or subsystem>

## Decision made
<what the agent decided to do>

## Why a decision was needed
<what was ambiguous, undocumented, or missing>

## Missing guidance
<which doc, skill, policy, or schema rule should exist or be clarified>

## Evidence
- <repo-local path>
- <repo-local path>

## Follow-up
<what should be documented, automated, or taught later>
`,
		".vokuknow/audit/decision-log.md": `# Decision log

Append a new entry whenever the agent must make a material decision because repo guidance is missing or ambiguous.

## Entry: <short title>

### Area
<primary code area or subsystem>

### Decision made
<what the agent decided to do>

### Why a decision was needed
<what was ambiguous, undocumented, or missing>

### Missing guidance
<which doc, skill, policy, or schema rule should be added or clarified>

### Evidence
- <repo-local path>
- <repo-local path>

### Follow-up
<what should be documented, automated, or taught later>
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

Memory artifacts exist to reduce repeated discovery, expose documentation blind spots, and prevent future agents from making avoidable wrong assumptions.

## Log blind spots during work

- if the agent must make a material decision because guidance is missing, append it to .vokuknow/audit/decision-log.md immediately
- record what was decided, why the decision was needed, and which doc, skill, or policy should be improved
- turn repeated or durable blind spots into updated docs, skills, or memory artifacts once the task is complete

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
		".vokuknow/prompts/examples/decision-log.md": `# Decision log

Use this workflow when the agent has to choose an approach because the current docs, skills, or schema do not provide enough guidance.

## Fill these inputs

- CODE_AREA: src/Auth/SessionBootstrap.php
- DECISION_TITLE: choose tenant-aware cookie source
- DECISION_MADE: Used request context instead of global config as the source of truth for secure-cookie behavior.
- WHY_DECISION_WAS_NEEDED: Existing docs explained secure cookies but did not define which component owned the final decision.
- MISSING_GUIDANCE: Add a short ownership note to the auth/session documentation and the vokuknow skill checklist.
- EVIDENCE_PATHS: src/Auth/SessionBootstrap.php, src/Http/RequestContext.php, tests/Auth/SessionBootstrapTest.php
- AUDIT_TARGET_FILE: .vokuknow/audit/decision-log.md

## Expected outcome

Append a structured decision entry that makes the blind spot visible before the task context is lost.
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
