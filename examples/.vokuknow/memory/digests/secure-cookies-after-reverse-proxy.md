# Digest: Why secure cookies broke after reverse proxy changes

## Symptom
Sessions failed after the proxy rollout.

## Root cause
Proxy headers changed effective scheme handling, which altered secure-cookie behavior.

## Failed assumption
We assumed secure-cookie logic was global-only.

## Fix
Adjusted request-context-driven scheme handling.

## Evidence
- src/Auth/SessionBootstrap.php
- src/Http/RequestContext.php
- tests/Auth/SessionBootstrapTest.php

## Reusable lesson
Proxy and environment assumptions must be captured near session initialization.

## Open questions
- Which other auth paths depend on the same request-context scheme calculation?
