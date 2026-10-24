# Section 11: Databases, Caching and Messaging

**Stage:** Stage 3 — Production Systems Engineering

## Purpose
Persist data and communicate between services while managing consistency, failure, and operational behavior.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **SQL and database access** — Use database/sql, connection pools, context, and explicit query boundaries.
2. **Transactions and consistency** — Design transactions, understand isolation trade-offs, and avoid long-running locks.
3. **Migrations and schema evolution** — Apply versioned migrations and maintain compatibility during deployment.
4. **Caching and invalidation** — Use TTLs, cache-aside patterns, and explicit stale-data assumptions.
5. **Queues and event-driven processing** — Consume messages with retries, idempotency, and dead-letter handling.
6. **Outbox, deduplication and recovery** — Build reliable event publication patterns and recover from partial failures.

## Expected outcomes
- Explain and apply sql and database access in idiomatic Go.
- Explain and apply transactions and consistency in idiomatic Go.
- Explain and apply migrations and schema evolution in idiomatic Go.
- Explain and apply caching and invalidation in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-11/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Sql And Database Access](../projects/section-11/topics/topic-01-sql-and-database-access.md)
- [Topic 2: Transactions And Consistency](../projects/section-11/topics/topic-02-transactions-and-consistency.md)
- [Topic 3: Migrations And Schema Evolution](../projects/section-11/topics/topic-03-migrations-and-schema-evolution.md)
- [Topic 4: Caching And Invalidation](../projects/section-11/topics/topic-04-caching-and-invalidation.md)
- [Topic 5: Queues And Event Driven Processing](../projects/section-11/topics/topic-05-queues-and-event-driven-processing.md)
- [Topic 6: Outbox Deduplication And Recovery](../projects/section-11/topics/topic-06-outbox-deduplication-and-recovery.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-11/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-11/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-11/capstones/capstone-3-independent-engineering-challenge.md)
