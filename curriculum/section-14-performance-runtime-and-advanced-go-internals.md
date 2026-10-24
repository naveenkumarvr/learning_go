# Section 14: Performance, Runtime and Advanced Go Internals

**Stage:** Stage 3 — Production Systems Engineering

## Purpose
Diagnose performance using measurements and understand enough runtime behavior to make informed trade-offs.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Benchmarking and profiling** — Create representative benchmarks and use CPU, heap, block, and mutex profiles.
2. **Garbage collection and allocation** — Interpret allocation patterns and GC behavior without premature tuning.
3. **Scheduler and goroutine behavior** — Understand runtime scheduling, blocking, and concurrency limits at a practical level.
4. **Memory layout and compiler diagnostics** — Use escape analysis and compiler output to explain observed behavior.
5. **Latency, throughput and load testing** — Model workloads, identify bottlenecks, and distinguish saturation from downstream latency.
6. **Optimization with evidence** — Make one change at a time, compare baselines, and protect correctness with tests.

## Expected outcomes
- Explain and apply benchmarking and profiling in idiomatic Go.
- Explain and apply garbage collection and allocation in idiomatic Go.
- Explain and apply scheduler and goroutine behavior in idiomatic Go.
- Explain and apply memory layout and compiler diagnostics in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-14/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Benchmarking And Profiling](../projects/section-14/topics/topic-01-benchmarking-and-profiling.md)
- [Topic 2: Garbage Collection And Allocation](../projects/section-14/topics/topic-02-garbage-collection-and-allocation.md)
- [Topic 3: Scheduler And Goroutine Behavior](../projects/section-14/topics/topic-03-scheduler-and-goroutine-behavior.md)
- [Topic 4: Memory Layout And Compiler Diagnostics](../projects/section-14/topics/topic-04-memory-layout-and-compiler-diagnostics.md)
- [Topic 5: Latency Throughput And Load Testing](../projects/section-14/topics/topic-05-latency-throughput-and-load-testing.md)
- [Topic 6: Optimization With Evidence](../projects/section-14/topics/topic-06-optimization-with-evidence.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-14/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-14/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-14/capstones/capstone-3-independent-engineering-challenge.md)
