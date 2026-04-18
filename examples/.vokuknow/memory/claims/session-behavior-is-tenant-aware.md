# Claim: Session behavior is tenant-aware

## Statement
Session cookie policy is tenant-aware and must not be changed only at global config level.

## Evidence
- src/Auth/SessionBootstrap.php
- tests/Auth/SessionBootstrapTest.php

## Confidence
high

## Status
accepted

## Notes
Any refactor should verify how tenant overrides are resolved before simplifying configuration inputs.
