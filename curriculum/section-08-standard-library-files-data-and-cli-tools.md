# Section 08: Standard Library, Files, Data and CLI Tools

**Stage:** Stage 2 — Professional Go Engineering

## Purpose
Use the standard library to build robust command-line tools and data-processing utilities.

## Prerequisites
- Prior sections or an equivalent demonstrated baseline.
- A working Go toolchain and ability to run tests locally.
- Relevant DevOps knowledge may be assumed when the section connects to cloud-native systems; focus on Go implementation and engineering judgment.

## Topics
1. **Filesystem and path handling** — Use os, io, filepath, temporary files, permissions, and atomic-write patterns.
2. **JSON, CSV and structured data** — Decode inputs defensively, validate schemas, and handle malformed records.
3. **CLI arguments and configuration** — Build predictable command interfaces with flags, environment variables, and config precedence.
4. **Streaming and buffered I/O** — Process large inputs without loading everything into memory.
5. **Logging and output contracts** — Separate machine-readable stdout from diagnostics and choose structured logging deliberately.
6. **Packaging and distribution** — Build reproducible binaries for target platforms and document installation and usage.

## Expected outcomes
- Explain and apply filesystem and path handling in idiomatic Go.
- Explain and apply json, csv and structured data in idiomatic Go.
- Explain and apply cli arguments and configuration in idiomatic Go.
- Explain and apply streaming and buffered i/o in idiomatic Go.
- Select and explain an appropriate design, including trade-offs and alternatives.
- Demonstrate behavior through tests and reproducible commands.
- Diagnose failures using evidence rather than guessing.
- Explain where the techniques apply in production and when a simpler approach is preferable.

## Learning workflow
For each topic, work through its three projects in order unless the instructor and validator agree that existing evidence supports skipping a project. Do not skip validation solely because the code compiles.

## Section capstones
Complete the three capstones under `projects/section-capstones/section-08/`. They integrate multiple topics from this section and at least one earlier concept where useful.

## Exit criteria
- Topic projects required by the learner's plan have objective validation evidence.
- The learner can explain the core mental models and key trade-offs without relying on memorized definitions.
- All three section capstones meet mandatory acceptance criteria or have explicitly documented exceptions approved by the learner.


## Project files

- [Topic 1: Filesystem And Path Handling](../projects/section-08/topics/topic-01-filesystem-and-path-handling.md)
- [Topic 2: Json Csv And Structured Data](../projects/section-08/topics/topic-02-json-csv-and-structured-data.md)
- [Topic 3: Cli Arguments And Configuration](../projects/section-08/topics/topic-03-cli-arguments-and-configuration.md)
- [Topic 4: Streaming And Buffered I O](../projects/section-08/topics/topic-04-streaming-and-buffered-i-o.md)
- [Topic 5: Logging And Output Contracts](../projects/section-08/topics/topic-05-logging-and-output-contracts.md)
- [Topic 6: Packaging And Distribution](../projects/section-08/topics/topic-06-packaging-and-distribution.md)
- [Section capstone: Capstone 1 Integrated Builder](../projects/section-08/capstones/capstone-1-integrated-builder.md)
- [Section capstone: Capstone 2 Production Scenario](../projects/section-08/capstones/capstone-2-production-scenario.md)
- [Section capstone: Capstone 3 Independent Engineering Challenge](../projects/section-08/capstones/capstone-3-independent-engineering-challenge.md)
