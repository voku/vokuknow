# Discovery: Session bootstrap flow

## Area
src/Auth/SessionBootstrap.php

## Learned
- Session bootstrap depends on `RequestContext`, not a global singleton.
- Tenant overrides affect cookie behavior.
- Secure-cookie behavior is environment-sensitive.

## Evidence
- src/Auth/SessionBootstrap.php
- tests/Auth/SessionBootstrapTest.php

## Confidence
high

## Future warning
Inspect tenant override flow before changing session or cookie logic.

## Open questions
- Do all cookie policy call sites respect the same tenant override path?
