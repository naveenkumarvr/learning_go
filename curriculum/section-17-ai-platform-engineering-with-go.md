# Section 17: AI Platform Engineering with Go

**Stage:** Stage 4 — Cloud-Native and AI Platform Specialization

## Purpose
Use Go for reliable AI platform services around inference, retrieval, orchestration, and governance.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Inference gateway fundamentals** — Build a provider-neutral gateway with timeouts, request validation, and usage metadata.
2. **Streaming responses and cancellation** — Proxy streaming output while preserving cancellation and backpressure.
3. **Model/provider adapters and routing** — Route across providers or local inference backends using explicit policies.
4. **RAG service integration** — Integrate ingestion/retrieval APIs, citations, metadata filters, and access-control checks.
5. **Rate limits, quotas and cost controls** — Enforce per-tenant limits and track token or request usage with clear accounting assumptions.
6. **Evaluation, audit and AI observability** — Record safe traces and evaluation outcomes without leaking sensitive prompts or secrets.

## Expected outcomes
- Explain and apply inference gateway fundamentals in idiomatic Go.
- Explain and apply streaming responses and cancellation in idiomatic Go.
- Explain and apply model/provider adapters and routing in idiomatic Go.
- Explain and apply rag service integration in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-17/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.
