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


## Project files

- [Topic 1: Functions And Multiple Returns](../projects/section-04/topics/topic-01-functions-and-multiple-returns.md)
- [Topic 2: Methods And Receiver Choices](../projects/section-04/topics/topic-02-methods-and-receiver-choices.md)
- [Topic 3: Interfaces And Implicit Implementation](../projects/section-04/topics/topic-03-interfaces-and-implicit-implementation.md)
- [Topic 4: Function Values And Closures](../projects/section-04/topics/topic-04-function-values-and-closures.md)
- [Topic 5: Generics And Type Constraints](../projects/section-04/topics/topic-05-generics-and-type-constraints.md)
- [Topic 6: Api Design And Abstraction Judgment](../projects/section-04/topics/topic-06-api-design-and-abstraction-judgment.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-04/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-04/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-04/capstones/capstone-3-independent-engineering-challenge.md)
