# Memory rules

## Purpose

Memory artifacts exist to reduce repeated discovery and prevent future agents from making avoidable wrong assumptions.

## Save memory when

- exploration revealed durable structure or coupling
- debugging established a root cause or disproved an assumption
- implementation clarified invariants, ownership, or dangerous call paths
- incomplete work needs a clean handoff

## Do not save memory when

- the note is only temporary scratch work
- the content is already captured in an existing artifact
- the statement has no supporting evidence
- the information is sensitive and should not be persisted

## Required qualities

- repo-local and file-based
- specific to a code area or lesson
- evidence-backed
- explicit about confidence
- clear about facts vs. hypotheses vs. open questions
- updated in place when a nearby artifact already exists
