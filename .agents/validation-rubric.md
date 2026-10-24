# Validation Rubric

## Required evidence
- Submission path and revision reviewed
- Commands actually run, with results or explicit `Not run`
- Tests and important edge cases
- Design explanation from the learner
- Acceptance criteria checklist
- Remaining findings and recommended status

## Severity
- **Critical:** unsafe or fundamentally incorrect behavior; data loss/security exposure; cannot meet core purpose.
- **High:** major requirement fails, important error path is unsafe, or correctness is unreliable.
- **Medium:** meaningful maintainability, testing, robustness, or performance issue.
- **Low:** polish, naming, documentation, or non-blocking improvement.

## Status decision
- `Passed`: all mandatory criteria evidenced, no critical/high findings, reproducible instructions and learner explanation are adequate.
- `Needs revision`: material criterion fails or critical/high finding remains.
- `In progress`: work is incomplete or evidence is insufficient to decide.
- `Not started`: no meaningful implementation evidence.

Never claim a test was run when it was not. Never equate code compilation with project completion.
