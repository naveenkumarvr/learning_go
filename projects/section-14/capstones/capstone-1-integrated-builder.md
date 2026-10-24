# Section 14 Capstone 1: Integrated Builder

**Section:** Performance, Runtime and Advanced Go Internals  
**Level:** Integrated implementation

## Scenario
Diagnose performance using measurements and understand enough runtime behavior to make informed trade-offs. Your task is to deliver a coherent system rather than a set of disconnected exercises. This capstone integrates multiple topics from the section:

- Benchmarking and profiling: Create representative benchmarks and use CPU, heap, block, and mutex profiles.
- Garbage collection and allocation: Interpret allocation patterns and GC behavior without premature tuning.
- Scheduler and goroutine behavior: Understand runtime scheduling, blocking, and concurrency limits at a practical level.
- Memory layout and compiler diagnostics: Use escape analysis and compiler output to explain observed behavior.
- Latency, throughput and load testing: Model workloads, identify bottlenecks, and distinguish saturation from downstream latency.
- Optimization with evidence: Make one change at a time, compare baselines, and protect correctness with tests.

## Problem statement
Combine the section's first half topics into a small but complete tool or subsystem. Create a runnable, documented implementation that solves a realistic engineering problem. Choose a scope that can be completed locally, state assumptions, and avoid adding infrastructure that is not needed to prove the learning outcomes.

## Requirements
- Define user-facing behavior, constraints, non-goals, and measurable acceptance criteria.
- Provide an architecture diagram and explain component responsibilities and data/control flow.
- Implement idiomatic Go with clear error handling and lifecycle ownership.
- Include unit tests and integration tests appropriate to the design.
- Exercise invalid inputs, dependency failures, and recovery behavior.
- Run `gofmt`, `go vet ./...`, and `go test ./...`; use race detection and benchmarks when relevant.
- Document how to build, run, test, troubleshoot, and safely clean up.
- Discuss at least two design alternatives and why the selected design fits.
- Identify one real-world use case and one situation where this design would be unnecessary or harmful.

## Suggested architecture
```mermaid
flowchart LR
    A[CLI or API boundary] --> B[Validation and orchestration]
    B --> C[Core domain logic]
    C --> D[Storage / network / runtime adapter]
    B --> E[Errors and telemetry]
    D --> F[Tests and failure simulation]
```
Adapt the diagram to the actual design. Do not force these components if the project is simpler.

## Milestones
1. Requirements and design note.
2. Thin vertical slice with a reproducible local run.
3. Core implementation and tests.
4. Failure handling and quality gates.
5. Design review, alternatives, and final demonstration.

## Acceptance criteria
- [ ] The stated user-facing requirements work.
- [ ] Design and component boundaries are explained and proportionate.
- [ ] Tests cover normal, boundary, and failure behavior.
- [ ] Build/test commands are reproducible and documented.
- [ ] No critical/high-severity issue remains.
- [ ] The learner can explain key choices and trade-offs without relying on a memorized script.

## Submission
Create your own folder under `practice/section-14/capstone-1/` containing source, tests, README, and `DESIGN.md`. This repository intentionally contains no sample implementation.

## Validation
The Validator should inspect the submitted implementation against this document, run available checks, report unrun checks honestly, and recommend a status based on evidence. Passing this capstone does not automatically pass the whole section unless all required topic criteria are also met.

## Resources
- Go documentation: https://go.dev/doc/
- Go standard library: https://pkg.go.dev/std
- Go blog: https://go.dev/blog/
