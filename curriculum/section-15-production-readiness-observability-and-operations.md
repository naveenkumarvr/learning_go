# Section 15: Production Readiness, Observability and Operations

**Stage:** Stage 4 — Cloud-Native and AI Platform Specialization

## Purpose
Operate Go applications with measurable health, actionable telemetry, and safe release practices.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Structured logs, metrics and traces** — Emit correlated telemetry with stable labels and useful operational context.
2. **Health, readiness and diagnostics** — Separate liveness from readiness and expose safe diagnostic endpoints.
3. **SLOs, SLIs and error budgets** — Translate service objectives into actionable monitoring and release decisions.
4. **Graceful shutdown and lifecycle management** — Drain requests, stop background work, and close dependencies within deadlines.
5. **Configuration, feature flags and rollout** — Validate configuration and support safe, reversible release strategies.
6. **Incident response and runbooks** — Diagnose a simulated incident and create a concise, actionable runbook.

## Expected outcomes
- Explain and apply structured logs, metrics and traces in idiomatic Go.
- Explain and apply health, readiness and diagnostics in idiomatic Go.
- Explain and apply slos, slis and error budgets in idiomatic Go.
- Explain and apply graceful shutdown and lifecycle management in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-15/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.
