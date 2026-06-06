# Module 1 Summary

## Module 1 Reference: Configuration, State, and Error Handling

This architecture reference log captures foundational mechanics, operational trade-offs, and compiler behaviors of Go 1.26.2.
Each language primitive is framed through the lens of cloud-native systems engineering (Kubernetes, Terraform, and low-latency microservices).

---

## 1. Variables, Constants, and Zero Values

### Explicit vs Short Declarations

Go handles memory layout and variable tracking through strict initialization patterns.

- **Explicit Declaration (var)**: Best when declaring a variable without an immediate value. It safely initializes to the type-safe default.
- **Short Declaration (:=)**: Declares and infers type in one step. Limited to function scope and cannot be used at package level.
- **Reassignment (=)**: Updates data in an existing allocation without creating a new one.

Using := inside an inner block when a variable already exists in an outer block causes **variable shadowing**.
That creates a temporary local duplicate and leaves the outer source of truth unchanged.

### Production Risk of Zero Values

If a variable is not explicitly initialized, Go assigns its type-specific zero value:

- integers -> 0
- strings -> empty string
- booleans -> false
- pointers/slices/maps -> nil

**Infrastructure risk**:
If a Port field is not initialized, it defaults to 0. If that value is propagated into deployment configuration without validation, it can trigger critical failures.

### Syntax Reference

```go
package main

const DefaultTimeoutSeconds = 30

func main() {
    // Explicit declaration (zero value: "")
    var clusterName string

    // Short declaration (type inference)
    nodeCount := 5

    // Reassignment
    clusterName = "prod-us-east"

    _ = nodeCount
}
```

---

## 2. Arrays, Slices, and Memory Allocation

### Slices Under the Hood

- Arrays are fixed-size, and length is part of the type.
- Slices are dynamic views over arrays.
- A slice header typically contains:
  - pointer to backing array
  - length
  - capacity

When append exceeds capacity, Go allocates a new backing array (often larger), copies data, updates the pointer, and old storage becomes GC-eligible.

### Optimization Strategy with make

To reduce repeated allocations and copies, pre-allocate capacity when limits are known.

### Syntax Reference

```go
// Nil slice (len=0, cap=0)
var targetIPs []string

// Slice literal
regions := []string{"us-west-1", "us-east-1"}

// Pre-allocated slice (len=0, cap=500)
metricsBuffer := make([]float64, 0, 500)

_, _, _ = targetIPs, regions, metricsBuffer
```

---

## 3. Maps (Strict Key-Value Infrastructure Maps)

### Mechanics

Maps are hash tables with average O(1) lookup.
Go maps are strictly typed for both keys and values.

- Uninitialized maps are nil.
- Writing to a nil map causes panic.
- Always initialize mutable maps using make.

### Syntax Reference

```go
environmentTags := make(map[string]string)

environmentTags["env"] = "production"
environmentTags["tier"] = "data-layer"

val, exists := environmentTags["secure_mode"]
if !exists {
    // Handle fallback
}
_ = val
```

---

## 4. Control Flow and Switch Architectures

### Inline Variable Initialization

Variables declared in if initialization are scoped to that if/else chain.

### Conditionless Switch

A switch without a target evaluates each case as a boolean expression.
This often replaces deep if/else ladders with cleaner validation logic.

### Syntax Reference

```go
if err := runHealthCheck(); err != nil {
    return err
}

switch {
case targetPort > 65535 || targetPort < 1:
    log.Fatal("Port boundary failure")
case targetPort < 1024:
    log.Printf("Warning: privileged port allocation")
default:
    log.Printf("Port parameters cleared")
}
```

---

## 5. Structs vs Instances

### Structural Blueprints

A struct defines a composite typed layout.

- **Struct type**: Blueprint (fields and types)
- **Instance**: Concrete memory allocation of that blueprint

### Syntax Reference

```go
type DatabaseConfig struct {
    HostName string
    DBPort   int
    Username string
}

func main() {
    prodDB := DatabaseConfig{
        HostName: "rds-cluster-01.internal",
        DBPort:   5432,
        Username: "sys_admin",
    }
    _ = prodDB
}
```

---

## 6. Pointers and Memory Efficiency

### Core Operators

- **&**: Get address of a value
- **\*Type**: Pointer type declaration
- **\*ptr**: Dereference pointer value

Go also supports implicit dereferencing with field access on pointers: ptr.Field.

### When to Use Pointers

- Persistent mutation across function boundaries
- Avoid copying large structs
- Share central mutable state (connections, pools, caches)

### Syntax Reference

```go
var originalConfig ServerConfig = ServerConfig{Host: "k8s-node-01"}

var cfgPointer *ServerConfig = &originalConfig
var detachedCopy ServerConfig = *cfgPointer

_, _ = cfgPointer, detachedCopy
```

---

## 7. Methods vs Functions (Pointer Receivers)

### Architectural Binding

- A function is standalone.
- A method is attached to a type through a receiver.
- Pointer receivers operate on the original object state.

### Syntax Reference

```go
func validatePortStandalone(cfg ServerConfig) bool {
    return cfg.Port <= 65535
}

func (cfg *ServerConfig) validatePort() bool {
    return cfg.Port <= 65535
}

func main() {
    server := ServerConfig{Port: 8080}
    validatePortStandalone(server)
    server.validatePort()
}
```

---

## 8. Data Ingestion Mechanics (Bytes, Tags, and Unmarshal)

### Why []byte Dominates I/O

At the OS and network layer, data is transmitted and stored as bytes.
Go parsers consume byte slices efficiently for transport and decoding.

### json.Unmarshal Flow

1. Scan bytes and parse JSON structure.
2. Match keys to struct field mappings.
3. Convert values to typed fields.
4. Write decoded output into the destination via pointer.

### Syntax Reference

```go
type TargetManifest struct {
    TargetHost string // mapped from target_host
    SecureMode bool   // mapped from secure_mode
}
```

---

## 9. Capstone: Inventory Parser Pipeline

This blueprint demonstrates:

- JSON ingestion
- Struct-based decoding
- Pointer receiver validation
- Error reporting through the error interface

```go
package main

import (
    "encoding/json"
    "fmt"
)

type ServerConfig struct {
    Host  string // json key: host_name
    Port  int    // json key: port_num
    IsSSL bool   // json key: ssl_enabled
}

func (cfg *ServerConfig) checkPort() error {
    switch {
    case cfg.Port > 65535 || cfg.Port < 1:
        return fmt.Errorf("invalid port %d for host %s", cfg.Port, cfg.Host)
    case cfg.Port < 1024:
        return fmt.Errorf("privileged port %d for host %s", cfg.Port, cfg.Host)
    default:
        return nil
    }
}

func (cfg *ServerConfig) checkSecurity() error {
    if !cfg.IsSSL {
        return fmt.Errorf("insecure deployment posture on host %s", cfg.Host)
    }
    return nil
}

func main() {
    rawFleetData := `[
        {"host_name":"prod-web-01","port_num":80,"ssl_enabled":true},
        {"host_name":"prod-db-01","port_num":5432,"ssl_enabled":false}
    ]`

    var fleet []*ServerConfig

    if err := json.Unmarshal([]byte(rawFleetData), &fleet); err != nil {
        fmt.Printf("[FATAL] parse failure: %v\n", err)
        return
    }

    for i := range fleet {
        fmt.Printf("\n==== Host: %s ====\n", fleet[i].Host)

        if err := fleet[i].checkPort(); err != nil {
            fmt.Printf("[ROUTING WARNING] %v\n", err)
        } else {
            fmt.Printf("[ROUTING CLEARED] Port %d\n", fleet[i].Port)
        }

        if err := fleet[i].checkSecurity(); err != nil {
            fmt.Printf("[SECURITY WARNING] %v\n", err)
        } else {
            fmt.Printf("[SECURITY CLEARED] SSL enabled\n")
        }
    }
}
```
