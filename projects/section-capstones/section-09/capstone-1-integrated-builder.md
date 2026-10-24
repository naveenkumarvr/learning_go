# Section 09 Capstone 1: Integrated Builder

**Section:** Concurrency and Parallelism  
**Level:** Integrated implementation

## Scenario
Use goroutines and synchronization safely, with bounded work and explicit cancellation. Your task is to deliver a coherent system rather than a set of disconnected exercises. This capstone integrates multiple topics from the section:

- Goroutines and lifecycle ownership: Start concurrent work with clear ownership and shutdown behavior.
- Channels and communication patterns: Use channels for coordination and data flow without overcomplicating simple code.
- Mutexes, atomics and shared state: Protect invariants and choose synchronization based on the access pattern.
- Worker pools and backpressure: Bound concurrency, queue sizes, and memory under load.
- Cancellation, timeouts and leak prevention: Propagate context, unblock workers, and test shutdown behavior.
- Race conditions and concurrency testing: Reproduce interleavings, run race detection, and reason about safety and liveness.

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
Create your own folder under `practice/section-09/capstone-1/` containing source, tests, README, and `DESIGN.md`. This repository intentionally contains no sample implementation.

## Validation
The Validator should inspect the submitted implementation against this document, run available checks, report unrun checks honestly, and recommend a status based on evidence. Passing this capstone does not automatically pass the whole section unless all required topic criteria are also met.

## Resources
- Go documentation: https://go.dev/doc/
- Go standard library: https://pkg.go.dev/std
- Go blog: https://go.dev/blog/
