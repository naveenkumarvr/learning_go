# Section 05: Error Handling, Resource Management and Reliability

**Stage:** Stage 1 — Go Foundations

## Purpose
Make failures explicit, preserve useful context, and manage resources safely.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Errors as values** — Return errors explicitly and avoid panic for expected operational failures.
2. **Wrapping and classifying errors** — Use errors.Is/errors.As, sentinel errors, typed errors, and contextual wrapping.
3. **Defer and cleanup** — Release files, locks, and other resources reliably; understand defer argument evaluation.
4. **Panic and recover boundaries** — Know when panic is appropriate and recover only at deliberate boundaries.
5. **Validation and defensive programming** — Validate inputs at boundaries and make invalid states difficult to represent.
6. **Timeouts, cancellation and reliability basics** — Use context cancellation, deadlines, and bounded operations in preparation for production services.

## Expected outcomes
- Explain and apply errors as values in idiomatic Go.
- Explain and apply wrapping and classifying errors in idiomatic Go.
- Explain and apply defer and cleanup in idiomatic Go.
- Explain and apply panic and recover boundaries in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-05/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Errors As Values](../projects/section-05/topics/topic-01-errors-as-values.md)
- [Topic 2: Wrapping And Classifying Errors](../projects/section-05/topics/topic-02-wrapping-and-classifying-errors.md)
- [Topic 3: Defer And Cleanup](../projects/section-05/topics/topic-03-defer-and-cleanup.md)
- [Topic 4: Panic And Recover Boundaries](../projects/section-05/topics/topic-04-panic-and-recover-boundaries.md)
- [Topic 5: Validation And Defensive Programming](../projects/section-05/topics/topic-05-validation-and-defensive-programming.md)
- [Topic 6: Timeouts Cancellation And Reliability Basics](../projects/section-05/topics/topic-06-timeouts-cancellation-and-reliability-basics.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-05/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-05/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-05/capstones/capstone-3-independent-engineering-challenge.md)
