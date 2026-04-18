# Architecture

vokuknow is organized as a thin CLI on top of testable internal packages.

- `cmd/vokuknow`: entrypoint
- `internal/cli`: Cobra command wiring
- `internal/app`: orchestration for init/build/lint/doctor/show/generate
- `internal/schema`: schema loading + validation
- `internal/generator`: deterministic markdown/json artifact generation
- `internal/lint`: structure/policy/drift checks
- `internal/doctor`: health-check wrapper around lint
- `internal/fs`: deterministic file write helpers

The schema in `.vokuknow/schema/` is the source of truth.
