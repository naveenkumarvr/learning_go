# Section 13: Security and Supply Chain Integrity

**Stage:** Stage 3 — Production Systems Engineering

## Purpose
Secure Go services, dependencies, build artifacts, and operational interfaces.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Threat modeling and secure defaults** — Identify assets, trust boundaries, abuse cases, and security requirements.
2. **Authentication, authorization and secrets** — Protect credentials and enforce least privilege at service boundaries.
3. **Input validation and injection defenses** — Handle untrusted inputs safely in HTTP, SQL, files, and command execution.
4. **TLS, certificates and transport security** — Configure secure transport and reason about certificate lifecycle.
5. **Dependency and build supply chain** — Audit dependencies, pin/review updates, and produce traceable build artifacts.
6. **Secure testing and vulnerability response** — Add security checks, reproduce findings, patch safely, and document residual risk.

## Expected outcomes
- Explain and apply threat modeling and secure defaults in idiomatic Go.
- Explain and apply authentication, authorization and secrets in idiomatic Go.
- Explain and apply input validation and injection defenses in idiomatic Go.
- Explain and apply tls, certificates and transport security in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-13/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Threat Modeling And Secure Defaults](../projects/section-13/topics/topic-01-threat-modeling-and-secure-defaults.md)
- [Topic 2: Authentication Authorization And Secrets](../projects/section-13/topics/topic-02-authentication-authorization-and-secrets.md)
- [Topic 3: Input Validation And Injection Defenses](../projects/section-13/topics/topic-03-input-validation-and-injection-defenses.md)
- [Topic 4: Tls Certificates And Transport Security](../projects/section-13/topics/topic-04-tls-certificates-and-transport-security.md)
- [Topic 5: Dependency And Build Supply Chain](../projects/section-13/topics/topic-05-dependency-and-build-supply-chain.md)
- [Topic 6: Secure Testing And Vulnerability Response](../projects/section-13/topics/topic-06-secure-testing-and-vulnerability-response.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-13/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-13/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-13/capstones/capstone-3-independent-engineering-challenge.md)
