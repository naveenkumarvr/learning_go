# Section 01: Development Environment and Go Workflow

**Stage:** Stage 1 — Go Foundations

## Purpose
Become productive with the Go toolchain and understand the compile, run, test, format, and dependency workflow.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Install and inspect Go** — Install a supported Go release, inspect GOROOT/GOPATH, and understand the toolchain and environment.
2. **Go workspace and module basics** — Create modules, understand go.mod/go.sum, and work with package paths and module boundaries.
3. **Build, run, format and vet** — Use go run, go build, go test, gofmt, go vet, and explain what each command catches.
4. **Editor and debugging workflow** — Configure VS Code/Kiro Go support, use gopls, debugger breakpoints, and inspect runtime values.
5. **CLI workflow and reproducibility** — Create repeatable scripts/Make targets and document commands so another engineer can reproduce a build.
6. **Go documentation and dependency hygiene** — Navigate go doc/pkg.go.dev, choose dependencies carefully, and inspect module versions.

## Expected outcomes
- Explain and apply install and inspect go in idiomatic Go.
- Explain and apply go workspace and module basics in idiomatic Go.
- Explain and apply build, run, format and vet in idiomatic Go.
- Explain and apply editor and debugging workflow in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-01/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Install And Inspect Go](../projects/section-01/topics/topic-01-install-and-inspect-go.md)
- [Topic 2: Go Workspace And Module Basics](../projects/section-01/topics/topic-02-go-workspace-and-module-basics.md)
- [Topic 3: Build Run Format And Vet](../projects/section-01/topics/topic-03-build-run-format-and-vet.md)
- [Topic 4: Editor And Debugging Workflow](../projects/section-01/topics/topic-04-editor-and-debugging-workflow.md)
- [Topic 5: Cli Workflow And Reproducibility](../projects/section-01/topics/topic-05-cli-workflow-and-reproducibility.md)
- [Topic 6: Go Documentation And Dependency Hygiene](../projects/section-01/topics/topic-06-go-documentation-and-dependency-hygiene.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-01/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-01/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-01/capstones/capstone-3-independent-engineering-challenge.md)
