# 🏗️ The DevOps Architect: Go-Lang Mastery Syllabus

This roadmap is designed to transform 10 years of DevOps experience into **Go-Lang Architectural Mastery**. Every language feature is treated as a direct solution to an infrastructure problem.

---

## 🗺️ Curriculum Overview
**Goal:** Reach "Principal" status by building 20 production-grade tools.

### Module 1: Configuration, State & Error Handling
**The Goal:** Moving from loose YAML/JSON to type-safe infrastructure definitions.  
**Concepts:** Basic types, Structs, Pointers (Memory addresses), Custom Error types, and the `io` package.

#### 🛠️ Scenarios (Build Together)
1. **Inventory Parser:** Build a tool that validates a fleet of 500+ server definitions and flags hardware mismatches.
2. **Secret Redactor:** Create a stream processor that reads logs and masks sensitive environment variables.
3. **State Serializer:** Write a "Mini-Terraform" staterunner that saves and loads resource states to local disk.

> 🎓 **Certification Challenge:** Build a **Linting Engine** for Kubernetes manifests that enforces custom security policies.

---

### Module 2: Networking & The Infrastructure API
**The Goal:** Building robust, production-grade clients for Cloud APIs (AWS/GCP).  
**Concepts:** `net/http`, JSON Marshalling, `context` (timeouts/cancellation), and Interfaces (for mocking APIs).

#### 🛠️ Scenarios (Build Together)
1. **Health Check Mesh:** Build a tool that probes 50 endpoints and reports latency/uptime via JSON.
2. **S3/Blob Wrapper:** Create a storage-agnostic interface that can switch between AWS S3 and Local Minio.
3. **API Rate Limiter:** Build a middleware that protects your internal tools from being overwhelmed by script requests.

> 🎓 **Certification Challenge:** Build a **Multi-Cloud Status Dashboard CLI** that aggregates health data from three different cloud providers concurrently.

---

### Module 3: Concurrency & Parallel Execution
**The Goal:** Mastering Go's "Killer Feature"—running thousands of tasks without a massive memory footprint.  
**Concepts:** Goroutines, Channels, `sync` package (WaitGroups/Mutexes), and Worker Pools.

#### 🛠️ Scenarios (Build Together)
1. **Parallel Log Grep:** Scan 10GB of distributed logs across 100 files in parallel to find a specific Trace ID.
2. **Concurrent IAM Rotator:** Rotate 1,000 IAM keys across 50 sub-accounts simultaneously with a 10-worker limit.
3. **Real-time Event Bus:** Build a channel-based system that listens for "Deployment" events and triggers "Slack" notifications.

> 🎓 **Certification Challenge:** Build a **High-Performance Port Scanner** that can scan a /24 network range for open ports in under 5 seconds.

---

### Module 4: Professional CLI Engineering
**The Goal:** Building tools that feel as polished as `kubectl` or `terraform`.  
**Concepts:** Cobra (Commands/Flags), Viper (Config files/Env vars), and `slog` (Structured Logging).

#### 🛠️ Scenarios (Build Together)
1. **The "Ops" Swiss Army Knife:** Build a multi-command CLI (e.g., `ops cluster get`, `ops node drain`).
2. **Dynamic Config Loader:** A tool that merges flags, environment variables, and `.yaml` files with a specific hierarchy.
3. **Progressive Log Watcher:** Create a CLI that streams logs from multiple pods with color-coded, structured output.

> 🎓 **Certification Challenge:** Build **gosh (Go-Shell)**: A CLI tool that executes remote commands on a fleet of servers via SSH and returns a unified report.

---

### Module 5: The "Principal" Level (K8s & Providers)
**The Goal:** Extending the tools we use every day.  
**Concepts:** Building Custom Terraform Providers and Kubernetes Operators (using `controller-runtime`).

#### 🛠️ Scenarios (Build Together)
1. **Custom Terraform Resource:** Build a provider that manages a simple "Local File" resource as if it were a cloud resource.
2. **K8s Sidecar Injector:** Write a controller that automatically adds a "Logging Agent" container to any new Pod.
3. **Resource Reaper:** Build an operator that deletes "Orphaned" cloud resources based on specific labels.

> 🎓 **Final Certification Challenge:** Build an **"Auto-Scaler" Operator** that monitors a custom external metric (like a queue length) and scales Kubernetes Deployments accordingly.

---
🚀 *Ready to begin? Let's start with Module 1: Configuration & State.*