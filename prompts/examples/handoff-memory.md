# Handoff memory

Use the `handoff_memory` category when work is incomplete and the next agent needs durable context before continuing.

## Fill these inputs

- `CODE_AREA`: `src/Auth/SessionBootstrap.php`
- `HANDOFF_SCOPE`: `auth session cleanup`
- `FILES_TOUCHED`: `src/Auth/SessionBootstrap.php, src/Http/RequestContext.php`
- `EVIDENCE_PATHS`: `src/Auth/SessionBootstrap.php, src/Http/RequestContext.php, tests/Auth/SessionBootstrapTest.php`
- `MEMORY_KIND`: `handoff`
- `MEMORY_TARGET_FILE`: `.vokuknow/memory/handoffs/auth-session-cleanup.md`
- `OPEN_QUESTIONS`: `Which cookie policy call sites still bypass tenant overrides?`
- `CONFIDENCE_LEVEL`: `high`

## Expected outcome

Write a handoff artifact that tells the next agent what is done, what must be preserved, and the next recommended step.
