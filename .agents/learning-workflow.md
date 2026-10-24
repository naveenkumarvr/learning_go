# Learning Workflow and Validation Contract

## Standard loop
1. **Orient** — check `progress.md`; identify one topic and its prerequisites.
2. **Teach** — define outcomes, explain the mental model, connect to practical systems, and ask a small reasoning question.
3. **Practice** — choose one project from the topic spec. Do not assign all three at once unless requested.
4. **Review** — ask the learner to explain the design and identify assumptions.
5. **Validate** — use the exact acceptance criteria; run checks where available.
6. **Remediate** — map each failure to a missing concept and targeted exercise.
7. **Revalidate** — rerun failed checks and inspect the changed evidence.
8. **Integrate** — connect topics through section capstones.
9. **Reflect** — record key learning and update `progress.md` only when evidence supports a status change.

## Project progression
- **Project 1 — Focused:** isolates the core concept with a narrow implementation.
- **Project 2 — Applied:** integrates the concept with realistic I/O, failure modes, tests, or neighboring concepts.
- **Project 3 — Independent:** requires the learner to make and defend design decisions under stated constraints.

The three projects must not be trivial renamings of the same exercise.

## Validation rubric
Score each relevant dimension as `Pass`, `Partial`, `Fail`, or `Not applicable`:
- Functional requirements and acceptance tests
- Correct Go semantics and idiomatic design
- Error handling and edge cases
- Automated tests and test quality
- Concurrency safety when relevant
- Complexity and performance reasoning when relevant
- Security and data handling when relevant
- Observability, lifecycle, and operational readiness when relevant
- Design rationale, alternatives, and trade-offs
- Reproducible build/run instructions

A project passes only if all mandatory acceptance criteria pass, no critical/high-severity issue remains, and the learner can explain key decisions. Do not turn this into a fake precision score; explain evidence and gaps.

## Remediation protocol
For each gap:
1. Name the concept and observed evidence.
2. Explain the smallest useful mental model.
3. Point to a targeted reference.
4. Give one small exercise that isolates the gap.
5. Re-check the original acceptance criterion.
6. Revalidate the project after the fix.

## Reports
Reports are created on demand, using a concise format: scope, submission/version reviewed, checks run, results, findings, remediation, final recommendation, and progress change if any. Do not create a report when a short chat response is enough.
