# Section 03: Data Structures, Types and Memory Semantics

**Stage:** Stage 1 — Go Foundations

## Purpose
Understand Go's core data structures, value semantics, pointers, and common correctness traps.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Arrays and slices** — Understand length/capacity, append, reslicing, copying, and backing-array aliasing.
2. **Maps and sets** — Use maps safely, handle missing keys, avoid relying on iteration order, and model set operations.
3. **Structs and composition** — Design structs, embed types when appropriate, and prefer composition over inheritance-like designs.
4. **Pointers and value semantics** — Understand addressability, copying, mutation, pointer receivers, and nil pointers.
5. **Allocation and escape intuition** — Use compiler diagnostics and benchmarks to form evidence-based allocation hypotheses.
6. **Choosing data structures** — Compare slices, maps, heaps, queues, and sorted data using complexity and workload constraints.

## Expected outcomes
- Explain and apply arrays and slices in idiomatic Go.
- Explain and apply maps and sets in idiomatic Go.
- Explain and apply structs and composition in idiomatic Go.
- Explain and apply pointers and value semantics in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-03/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Arrays And Slices](../projects/section-03/topics/topic-01-arrays-and-slices.md)
- [Topic 2: Maps And Sets](../projects/section-03/topics/topic-02-maps-and-sets.md)
- [Topic 3: Structs And Composition](../projects/section-03/topics/topic-03-structs-and-composition.md)
- [Topic 4: Pointers And Value Semantics](../projects/section-03/topics/topic-04-pointers-and-value-semantics.md)
- [Topic 5: Allocation And Escape Intuition](../projects/section-03/topics/topic-05-allocation-and-escape-intuition.md)
- [Topic 6: Choosing Data Structures](../projects/section-03/topics/topic-06-choosing-data-structures.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-03/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-03/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-03/capstones/capstone-3-independent-engineering-challenge.md)
