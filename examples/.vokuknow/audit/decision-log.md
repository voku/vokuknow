# Decision log

Append a new entry whenever the agent must make a material decision because repo guidance is missing or ambiguous.

## Entry: define material decision threshold for decision logging

### Area
skills/vokuknow/SKILL.md

### Decision made
Treated ownership, invariant, safety-boundary, fallback, and workflow choices as log-worthy, while excluding trivial style-only choices.

### Why a decision was needed
The feature required logging every material self-directed decision, but the threshold for "material" was not defined well enough to review the feature consistently.

### Missing guidance
The skill, policy, and example prompt should spell out what belongs in the decision log and what should stay out.

### Evidence
- skills/vokuknow/SKILL.md
- policy/memory_rules.md
- prompts/examples/decision-log.md

### Follow-up
Keep the examples aligned with the policy so future agents log blind spots consistently.

## Entry: add a worked example of review-time decision logging

### Area
examples/.vokuknow/audit/decision-log.md

### Decision made
Added a filled decision-log example under `examples/.vokuknow/audit/decision-log.md` instead of leaving the feature represented only by empty scaffolding.

### Why a decision was needed
The feature could be bootstrapped, but there was no concrete example showing what a review-driven blind-spot entry should look like in practice.

### Missing guidance
The repository examples should include audit artifacts alongside memory artifacts so the workflow can be understood end to end.

### Evidence
- README.md
- templates/decision-entry.md
- examples/.vokuknow/memory/digests/secure-cookies-after-reverse-proxy.md

### Follow-up
Add or refresh worked audit examples whenever the decision-log workflow changes materially.
