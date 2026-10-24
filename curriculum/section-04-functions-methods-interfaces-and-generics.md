# Section 04: Functions, Methods, Interfaces and Generics

**Stage:** Stage 1 — Go Foundations

## Purpose
Design reusable Go APIs with explicit contracts, small interfaces, and appropriate abstraction.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Functions and multiple returns** — Use parameters, return values, named returns sparingly, and multiple return values for result/error.
2. **Methods and receiver choices** — Choose value vs pointer receivers consistently and understand method sets.
3. **Interfaces and implicit implementation** — Define small consumer-side interfaces and use compile-time assertions when useful.
4. **Function values and closures** — Use callbacks, closures, and function types without hiding control flow.
5. **Generics and type constraints** — Use type parameters where they simplify reusable algorithms or data structures.
6. **API design and abstraction judgment** — Evaluate whether an abstraction improves clarity, testability, or flexibility enough to justify its cost.

## Expected outcomes
- Explain and apply functions and multiple returns in idiomatic Go.
- Explain and apply methods and receiver choices in idiomatic Go.
- Explain and apply interfaces and implicit implementation in idiomatic Go.
- Explain and apply function values and closures in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-04/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.
