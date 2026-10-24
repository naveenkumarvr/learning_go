# Section 09: Concurrency and Parallelism

**Stage:** Stage 2 — Professional Go Engineering

## Purpose
Use goroutines and synchronization safely, with bounded work and explicit cancellation.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Goroutines and lifecycle ownership** — Start concurrent work with clear ownership and shutdown behavior.
2. **Channels and communication patterns** — Use channels for coordination and data flow without overcomplicating simple code.
3. **Mutexes, atomics and shared state** — Protect invariants and choose synchronization based on the access pattern.
4. **Worker pools and backpressure** — Bound concurrency, queue sizes, and memory under load.
5. **Cancellation, timeouts and leak prevention** — Propagate context, unblock workers, and test shutdown behavior.
6. **Race conditions and concurrency testing** — Reproduce interleavings, run race detection, and reason about safety and liveness.

## Expected outcomes
- Explain and apply goroutines and lifecycle ownership in idiomatic Go.
- Explain and apply channels and communication patterns in idiomatic Go.
- Explain and apply mutexes, atomics and shared state in idiomatic Go.
- Explain and apply worker pools and backpressure in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-09/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Goroutines And Lifecycle Ownership](../projects/section-09/topics/topic-01-goroutines-and-lifecycle-ownership.md)
- [Topic 2: Channels And Communication Patterns](../projects/section-09/topics/topic-02-channels-and-communication-patterns.md)
- [Topic 3: Mutexes Atomics And Shared State](../projects/section-09/topics/topic-03-mutexes-atomics-and-shared-state.md)
- [Topic 4: Worker Pools And Backpressure](../projects/section-09/topics/topic-04-worker-pools-and-backpressure.md)
- [Topic 5: Cancellation Timeouts And Leak Prevention](../projects/section-09/topics/topic-05-cancellation-timeouts-and-leak-prevention.md)
- [Topic 6: Race Conditions And Concurrency Testing](../projects/section-09/topics/topic-06-race-conditions-and-concurrency-testing.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-09/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-09/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-09/capstones/capstone-3-independent-engineering-challenge.md)
