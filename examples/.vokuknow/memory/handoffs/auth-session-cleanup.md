# Handoff: auth session cleanup

## Current state
Root cause is understood, but cleanup and call-site review are still pending.

## Must know
- Tenant overrides exist.
- `RequestContext` drives secure-cookie behavior.
- Do not simplify this flow to global config only.

## Evidence
- src/Auth/SessionBootstrap.php
- src/Http/RequestContext.php
- tests/Auth/SessionBootstrapTest.php

## Next recommended step
Review cookie policy call sites before refactoring session initialization.

## Open questions
- Which remaining cookie policy call sites bypass tenant-aware context?
