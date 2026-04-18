# Debugging digest

Use the `debugging_digest` category after debugging establishes a symptom, root cause, failed assumption, and fix that future agents should reuse.

## Fill these inputs

- `CODE_AREA`: `src/Auth/SessionBootstrap.php`
- `FILES_TOUCHED`: `src/Auth/SessionBootstrap.php, src/Http/RequestContext.php, tests/Auth/SessionBootstrapTest.php`
- `PRIOR_ASSUMPTION`: `proxy changes should not affect session cookies`
- `UPDATED_ASSUMPTION`: `proxy headers can change effective scheme handling and cookie behavior`
- `EVIDENCE_PATHS`: `src/Auth/SessionBootstrap.php, src/Http/RequestContext.php, tests/Auth/SessionBootstrapTest.php`
- `MEMORY_KIND`: `digest`
- `MEMORY_TARGET_FILE`: `.vokuknow/memory/digests/secure-cookies-after-reverse-proxy.md`
- `CONFIDENCE_LEVEL`: `high`

## Expected outcome

Write a digest artifact that records the root cause, failed assumption, fix, and future warning signs.
