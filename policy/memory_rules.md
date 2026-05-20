# Memory rules

## Purpose

Memory artifacts exist to reduce repeated discovery, expose documentation blind spots, and prevent future agents from making avoidable wrong assumptions.

## Log blind spots during work

- if the agent must make a material decision because guidance is missing, append it to `.vokuknow/audit/decision-log.md` immediately
- record what was decided, why the decision was needed, and which doc, skill, or policy should be improved
- before finalizing, either update that missing guidance in the same task or leave a concrete follow-up in the log entry
- turn repeated or durable blind spots into updated docs, skills, or memory artifacts once the task is complete
- treat ownership, invariants, safety boundaries, fallback behavior, and workflow choices as material; ignore trivial style-only choices unless they reveal a bigger missing rule

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
- updated in place when an artifact covering the same area or logical subsystem already exists
