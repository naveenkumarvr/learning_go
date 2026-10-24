# Section 10: Networking, APIs and Communication Protocols

**Stage:** Stage 2 — Professional Go Engineering

## Purpose
Build reliable HTTP services and clients with explicit API contracts and resilient network behavior.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **TCP, sockets and network fundamentals** — Understand connections, ports, DNS, TLS, and network failure modes relevant to Go services.
2. **HTTP handlers and routing** — Build composable handlers with request validation and clear response contracts.
3. **HTTP clients and resilience** — Set timeouts, reuse transports, handle retries carefully, and respect idempotency.
4. **Middleware, authentication and authorization** — Implement cross-cutting concerns without coupling business logic to transport details.
5. **API schemas and versioning** — Define stable request/response models and compatibility expectations.
6. **gRPC and protocol trade-offs** — Build a small RPC service and compare REST, gRPC, and messaging for concrete use cases.

## Expected outcomes
- Explain and apply tcp, sockets and network fundamentals in idiomatic Go.
- Explain and apply http handlers and routing in idiomatic Go.
- Explain and apply http clients and resilience in idiomatic Go.
- Explain and apply middleware, authentication and authorization in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-10/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.
