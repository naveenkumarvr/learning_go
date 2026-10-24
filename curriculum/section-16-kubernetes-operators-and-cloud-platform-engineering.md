# Section 16: Kubernetes, Operators and Cloud Platform Engineering

**Stage:** Stage 4 — Cloud-Native and AI Platform Specialization

## Purpose
Build Go-based cloud-native services and controllers that behave well in Kubernetes environments.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Containerizing Go services** — Build small, secure, reproducible images and run as a non-root user.
2. **Kubernetes APIs and client-go** — Read and update Kubernetes resources with explicit error and retry handling.
3. **Controllers and reconciliation loops** — Design idempotent reconcile logic that continuously moves observed state toward desired state.
4. **Custom resources and operator lifecycle** — Define a small CRD, validation rules, status fields, and lifecycle behavior.
5. **Leader election, work queues and rate limiting** — Scale controllers safely and avoid hot-looping on persistent failures.
6. **Platform integration and policy** — Integrate identity, configuration, observability, and deployment workflows into platform services.

## Expected outcomes
- Explain and apply containerizing go services in idiomatic Go.
- Explain and apply kubernetes apis and client-go in idiomatic Go.
- Explain and apply controllers and reconciliation loops in idiomatic Go.
- Explain and apply custom resources and operator lifecycle in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-16/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Containerizing Go Services](../projects/section-16/topics/topic-01-containerizing-go-services.md)
- [Topic 2: Kubernetes Apis And Client Go](../projects/section-16/topics/topic-02-kubernetes-apis-and-client-go.md)
- [Topic 3: Controllers And Reconciliation Loops](../projects/section-16/topics/topic-03-controllers-and-reconciliation-loops.md)
- [Topic 4: Custom Resources And Operator Lifecycle](../projects/section-16/topics/topic-04-custom-resources-and-operator-lifecycle.md)
- [Topic 5: Leader Election Work Queues And Rate Limiting](../projects/section-16/topics/topic-05-leader-election-work-queues-and-rate-limiting.md)
- [Topic 6: Platform Integration And Policy](../projects/section-16/topics/topic-06-platform-integration-and-policy.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-16/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-16/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-16/capstones/capstone-3-independent-engineering-challenge.md)
