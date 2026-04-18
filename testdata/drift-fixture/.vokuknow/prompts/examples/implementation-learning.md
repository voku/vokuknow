# Implementation learning

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
