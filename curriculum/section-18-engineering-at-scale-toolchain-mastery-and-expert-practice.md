# Section 18: Engineering at Scale, Toolchain Mastery and Expert Practice

**Stage:** Stage 4 — Cloud-Native and AI Platform Specialization

## Purpose
Demonstrate independent engineering judgment across design, implementation, release, operation, and long-term maintenance.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Technical design documents and RFCs** — Write concise proposals with constraints, alternatives, risks, rollout, and decision records.
2. **Large-scale refactoring and migrations** — Change APIs and data models incrementally with compatibility and rollback plans.
3. **Multi-service integration and contract testing** — Validate cross-service assumptions and failure behavior at system boundaries.
4. **Release engineering and reproducible builds** — Automate versioning, provenance, artifact validation, and release promotion.
5. **Operational game days and failure injection** — Exercise failure modes, recovery objectives, and runbook quality in a safe environment.
6. **Independent architecture defense** — Defend an end-to-end design with evidence, trade-offs, operational plans, and post-implementation review.

## Expected outcomes
- Explain and apply technical design documents and rfcs in idiomatic Go.
- Explain and apply large-scale refactoring and migrations in idiomatic Go.
- Explain and apply multi-service integration and contract testing in idiomatic Go.
- Explain and apply release engineering and reproducible builds in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-18/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.
