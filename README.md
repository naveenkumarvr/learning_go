# Go Learning: Foundations to Cloud-Native and AI Platform Engineering

A practical, evidence-based Go curriculum for an experienced DevOps engineer. The goal is not to collect completed lessons; it is to demonstrate the ability to design, implement, test, secure, deploy, debug, and operate production-style Go systems.

## Start here
1. Read `curriculum/roadmap.md`.
2. Open `progress.md` and choose the first topic marked `Not started`.
3. Ask the **Go Learning Instructor** agent to teach that topic and select Project 1.
4. Create your own solution under `practice/` using the topic project's submission requirements.
5. Ask the **Go Project Validator** to review the exact submission path and project spec.
6. Fix findings and request revalidation. Only evidence-based validation changes progress.
7. Reports are generated only when requested or useful, under `reports/`.

## Learning rules
- Start from your existing DevOps, AWS, Kubernetes, Terraform, CI/CD, and observability experience. Teach Go-specific differences and only revisit prerequisites when evidence shows a gap.
- Use Socratic questions and Feynman explanations. Ask for reasoning; don't immediately give away complete solutions.
- Each topic has three progressive projects. Project 1 isolates the concept, Project 2 applies it to a realistic subsystem, and Project 3 requires independent design.
- Compilation is necessary but never sufficient for a pass.
- If a concept is unclear or validation fails, diagnose the gap, teach the missing prerequisite, assign focused remediation, then revalidate.
- Keep the system lean: `progress.md` is the single progress source of truth. No reports or practice submissions are pre-created.
- Prefer local-first tools and open-source components. Use cloud resources only when they materially improve the lesson. Avoid creating paid resources without explicit approval.

## Tools
- The shared rules live in `AGENTS.md` and `.agents/`.
- Kiro entry points live in `.kiro/agents/`; Kiro steering files live in `.kiro/steering/`.
- VS Code / GitHub Copilot agent entry points live in `.github/agents/`.
- Keep the shared instructions authoritative; platform-specific files should point back to them.

## Status meanings
- `Not started`: no meaningful work submitted.
- `In progress`: currently learning or implementing.
- `Needs revision`: a validation found material gaps.
- `Passed`: the specified acceptance criteria are met with evidence.

A conversation, explanation, or code that merely compiles does not itself justify `Passed`.
