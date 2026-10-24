# Section 12: Architecture and Distributed Systems

**Stage:** Stage 3 — Production Systems Engineering

## Purpose
Design systems around failure boundaries, consistency needs, and operational constraints.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Requirements and quality attributes** — Translate user needs into measurable latency, availability, throughput, and recovery targets.
2. **Service boundaries and modular monoliths** — Choose boundaries based on ownership and change patterns rather than fashion.
3. **Distributed failure and timeouts** — Design for partial failure, retries, deadlines, and overload.
4. **Idempotency and workflow state** — Make retried operations safe and represent workflow progress explicitly.
5. **Consistency, coordination and leader concepts** — Understand practical consistency choices and the costs of coordination.
6. **Resilience patterns and graceful degradation** — Use bulkheads, circuit breakers, rate limits, and fallbacks with clear failure semantics.

## Expected outcomes
- Explain and apply requirements and quality attributes in idiomatic Go.
- Explain and apply service boundaries and modular monoliths in idiomatic Go.
- Explain and apply distributed failure and timeouts in idiomatic Go.
- Explain and apply idempotency and workflow state in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-12/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Requirements And Quality Attributes](../projects/section-12/topics/topic-01-requirements-and-quality-attributes.md)
- [Topic 2: Service Boundaries And Modular Monoliths](../projects/section-12/topics/topic-02-service-boundaries-and-modular-monoliths.md)
- [Topic 3: Distributed Failure And Timeouts](../projects/section-12/topics/topic-03-distributed-failure-and-timeouts.md)
- [Topic 4: Idempotency And Workflow State](../projects/section-12/topics/topic-04-idempotency-and-workflow-state.md)
- [Topic 5: Consistency Coordination And Leader Concepts](../projects/section-12/topics/topic-05-consistency-coordination-and-leader-concepts.md)
- [Topic 6: Resilience Patterns And Graceful Degradation](../projects/section-12/topics/topic-06-resilience-patterns-and-graceful-degradation.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-12/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-12/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-12/capstones/capstone-3-independent-engineering-challenge.md)
