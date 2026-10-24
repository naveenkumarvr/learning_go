# Final Capstone 1: Cloud-Native Platform Controller

## Purpose
Build a production-shaped Go controller that reconciles a declared platform resource into actual state. It should include a clear API model, idempotent reconciliation, bounded retries, status reporting, structured telemetry, secure defaults, graceful shutdown, and tests for transient and permanent failures.

## Concepts integrated
Go controllers, Kubernetes APIs, reconciliation, queues, retries, observability, security, release engineering.

## Operating constraints
- Default to local development and open-source tooling.
- No paid cloud resources or third-party API usage without explicit approval.
- Use mocks or local services when they provide sufficient evidence.
- Never use real production credentials, private customer data, or sensitive prompts.
- State non-goals and avoid implementing infrastructure beyond the learning objectives.

## Required deliverables
1. Working source code and automated tests.
2. `README.md` with architecture, prerequisites, build/run/test commands, troubleshooting, and cleanup.
3. `DESIGN.md` with requirements, assumptions, architecture/data flow, component responsibilities, alternatives, trade-offs, failure modes, and rollout/recovery plan.
4. `OPERATIONS.md` with health signals, useful logs/metrics, failure diagnosis, safe shutdown, and a concise runbook.
5. Test evidence and a short demo script.
6. Security/dependency review appropriate to the scope.
7. Performance or load evidence where relevant, with workload and limitations stated.
8. A retrospective explaining what would change for a real production deployment.

## Architecture expectations
Create a diagram specific to your implementation. It must identify trust boundaries, external dependencies, state ownership, failure paths, and telemetry flow. Explain why the chosen boundaries fit the requirements and compare at least two alternatives.

## Milestones
1. Requirements and architecture proposal.
2. API/data contracts and a minimal vertical slice.
3. Core behavior and automated tests.
4. Failure handling, security, observability, and lifecycle management.
5. Integration and contract tests.
6. Performance/reliability exercise.
7. Operational runbook and final design defense.

## Validation criteria
- [ ] Functional behavior matches stated requirements.
- [ ] Important normal, boundary, failure, timeout, and cancellation paths are tested.
- [ ] Interfaces and component boundaries are proportionate and explained.
- [ ] Error handling, resource ownership, and shutdown behavior are correct.
- [ ] Security and data-handling assumptions are explicit.
- [ ] Observability supports diagnosis without leaking secrets or sensitive payloads.
- [ ] Build/test/run instructions are reproducible.
- [ ] Performance claims have measurements and reproducible workload details, or are explicitly marked out of scope with justification.
- [ ] No critical/high-severity issue remains.
- [ ] The learner can defend trade-offs and explain what would be required before production adoption.

## Submission
Create your own folder under `practice/final-capstone-1/`. The repository does not include a starter solution. Ask the Validator to assess the exact submission path and this specification.

## Resources
- Go documentation: https://go.dev/doc/
- Go standard library: https://pkg.go.dev/std
- Kubernetes documentation (for controller work): https://kubernetes.io/docs/
- OpenTelemetry documentation (for telemetry work): https://opentelemetry.io/docs/
