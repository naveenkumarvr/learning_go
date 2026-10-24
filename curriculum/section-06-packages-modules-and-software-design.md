# Section 06: Packages, Modules and Software Design

**Stage:** Stage 2 — Professional Go Engineering

## Purpose
Organize maintainable Go code and evolve APIs without creating unnecessary layers.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Package boundaries and naming** — Organize by responsibility, keep packages cohesive, and avoid vague utility packages.
2. **Module versioning and dependency management** — Manage module requirements, upgrades, replace directives, and reproducible dependency resolution.
3. **Configuration and dependency injection** — Separate configuration parsing from runtime wiring and inject dependencies explicitly.
4. **Layering and ports/adapters** — Keep domain logic testable and isolate I/O where the complexity warrants it.
5. **Public API and compatibility** — Design exported symbols carefully and understand semver implications.
6. **Refactoring and code review** — Refactor incrementally, preserve behavior with tests, and review for clarity and maintainability.

## Expected outcomes
- Explain and apply package boundaries and naming in idiomatic Go.
- Explain and apply module versioning and dependency management in idiomatic Go.
- Explain and apply configuration and dependency injection in idiomatic Go.
- Explain and apply layering and ports/adapters in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-06/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Package Boundaries And Naming](../projects/section-06/topics/topic-01-package-boundaries-and-naming.md)
- [Topic 2: Module Versioning And Dependency Management](../projects/section-06/topics/topic-02-module-versioning-and-dependency-management.md)
- [Topic 3: Configuration And Dependency Injection](../projects/section-06/topics/topic-03-configuration-and-dependency-injection.md)
- [Topic 4: Layering And Ports Adapters](../projects/section-06/topics/topic-04-layering-and-ports-adapters.md)
- [Topic 5: Public Api And Compatibility](../projects/section-06/topics/topic-05-public-api-and-compatibility.md)
- [Topic 6: Refactoring And Code Review](../projects/section-06/topics/topic-06-refactoring-and-code-review.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-06/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-06/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-06/capstones/capstone-3-independent-engineering-challenge.md)
