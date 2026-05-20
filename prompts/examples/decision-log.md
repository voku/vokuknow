# Decision log

Use this workflow when the agent has to choose an approach because the current docs, skills, or schema do not provide enough guidance.

Log ownership, invariant, safety-boundary, fallback, or workflow choices. Skip trivial style-only choices unless they expose a broader missing rule.

## Fill these inputs

- `CODE_AREA`: `src/Auth/SessionBootstrap.php`
- `DECISION_TITLE`: `choose tenant-aware cookie source`
- `DECISION_MADE`: `Used request context instead of global config as the source of truth for secure-cookie behavior.`
- `WHY_DECISION_WAS_NEEDED`: `Existing docs explained secure cookies but did not define which component owned the final decision.`
- `MISSING_GUIDANCE`: `Add a short ownership note to the auth/session documentation and the vokuknow skill checklist.`
- `EVIDENCE_PATHS`: `src/Auth/SessionBootstrap.php, src/Http/RequestContext.php, tests/Auth/SessionBootstrapTest.php`
- `AUDIT_TARGET_FILE`: `.vokuknow/audit/decision-log.md`

## Expected outcome

Append a structured decision entry that makes the blind spot visible before the task context is lost.
