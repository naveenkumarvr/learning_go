# Section 07: Testing, Debugging and Code Quality

**Stage:** Stage 2 — Professional Go Engineering

## Purpose
Build confidence through layered tests, reproducible debugging, and automated quality gates.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Unit tests and table-driven tests** — Test behavior across representative cases with useful failure messages.
2. **Test doubles and interface seams** — Use fakes and small interfaces only where they improve isolation.
3. **Integration tests and test fixtures** — Exercise real boundaries such as files, databases, and HTTP servers with controlled fixtures.
4. **Fuzzing and property-based thinking** — Use Go fuzz tests to discover edge cases and define invariants.
5. **Race detection and coverage interpretation** — Use go test -race and coverage as signals, not as proof of correctness.
6. **Debugging, profiling and quality gates** — Reproduce bugs, inspect state, run static checks, and build CI gates.

## Expected outcomes
- Explain and apply unit tests and table-driven tests in idiomatic Go.
- Explain and apply test doubles and interface seams in idiomatic Go.
- Explain and apply integration tests and test fixtures in idiomatic Go.
- Explain and apply fuzzing and property-based thinking in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-07/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Unit Tests And Table Driven Tests](../projects/section-07/topics/topic-01-unit-tests-and-table-driven-tests.md)
- [Topic 2: Test Doubles And Interface Seams](../projects/section-07/topics/topic-02-test-doubles-and-interface-seams.md)
- [Topic 3: Integration Tests And Test Fixtures](../projects/section-07/topics/topic-03-integration-tests-and-test-fixtures.md)
- [Topic 4: Fuzzing And Property Based Thinking](../projects/section-07/topics/topic-04-fuzzing-and-property-based-thinking.md)
- [Topic 5: Race Detection And Coverage Interpretation](../projects/section-07/topics/topic-05-race-detection-and-coverage-interpretation.md)
- [Topic 6: Debugging Profiling And Quality Gates](../projects/section-07/topics/topic-06-debugging-profiling-and-quality-gates.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-07/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-07/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-07/capstones/capstone-3-independent-engineering-challenge.md)
