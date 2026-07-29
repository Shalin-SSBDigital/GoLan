# Go Garbage Collector — Tri-Color Mark-Sweep

> Complete learning guide following the Go-by-Python teaching method.

---

## 1. What is it?

Go's **Garbage Collector (GC)** is an automatic memory cleaner that finds and frees memory that your program no longer needs. Go uses a **Non-Generational, Tri-Color Mark-Sweep** collector:

- **Non-Generational** — treats all memory as one space (no young/old separation)
- **Tri-Color** — uses three colors (white, gray, black) to track which objects are live
- **Mark-Sweep** — two phases: mark live objects, then sweep dead ones
- **Concurrent** — runs alongside your program (minimal pauses)

## 2. Why do we need it?

**Problem:** Every time you create a value (`x := &MyStruct{}`, `make([]int, 100)`), it uses memory. If you never free memory, your program grows forever and crashes. In C, you manually call `free()`. In Go, the GC does it automatically.

**Go's solution:** A concurrent GC that:
- Runs **without stopping your entire program** (most of the time)
- Pauses are typically **< 500 microseconds**
- Is **tunable** via the `GOGC` environment variable
- Works **automatically** — you never think about it

## 3. Python Comparison

| Feature | Go | Python |
|---|---|---|
| GC type | Non-generational, concurrent mark-sweep | Generational, reference counting + mark-sweep |
| Collection trigger | Heap size doubles (default GOGC=100) | Reference count drops to 0, or threshold |
| Pause time | ~500 μs (very short) | Variable (can be long with many objects) |
| User control | `GOGC`, `runtime.GC()` | `gc.collect()` |
| Detect collection | `GODEBUG=gctrace=1` | `gc.get_stats()` |
| Concurrency | Most phases concurrent with program | Reference counting is immediate, mark-sweep can pause |
| Generations | One (non-generational) | Three (young, middle, old) |

## 4. Syntax

No syntax for GC — it runs **automatically** in the background. But you can interact with it:

```go
import "runtime"

runtime.GC()                    // Force a GC cycle (blocking)
runtime.GOMEMLIMIT()            // Soft memory limit (Go 1.19+)

// Debug: print GC stats
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Println("Alloc:", m.Alloc)             // currently allocated
fmt.Println("TotalAlloc:", m.TotalAlloc)   // total allocated so far
fmt.Println("NumGC:", m.NumGC)             // number of GC cycles
fmt.Println("PauseTotal:", m.PauseTotalNs) // total pause time

// Environment variables:
// GOGC=100      default: start GC when heap doubles
// GOGC=off      disable GC (dangerous!)
// GODEBUG=gctrace=1  print GC stats every cycle
```

## 5. Simple Example

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    var m runtime.MemStats

    // Allocate lots of data
    for i := 0; i < 10; i++ {
        data := make([]byte, 10*1024*1024) // 10 MB
        _ = data
        time.Sleep(100 * time.Millisecond)

        runtime.ReadMemStats(&m)
        fmt.Printf("Alloc: %d MB, NumGC: %d\n",
            m.Alloc/1024/1024, m.NumGC)
    }
}

// Run with: GODEBUG=gctrace=1 go run main.go
```

**Line-by-line:**
1. `runtime.ReadMemStats(&m)` — reads current memory statistics
2. `make([]byte, 10*1024*1024)` — allocates 10 MB (triggers GC as heap grows)
3. `m.Alloc` — currently allocated bytes on heap
4. `m.NumGC` — how many GC cycles have completed

## 6. Python Equivalent

```python
import gc
import time

# Python GC is different — reference counting + generational
print("Thresholds:", gc.get_threshold())  # (700, 10, 10)
print("Garbage:", gc.garbage)

# Force collection
gc.collect()  # same as runtime.GC()

# Disable GC
gc.disable()  # same as GOGC=off (dangerous!)

# Get stats
for i in range(10):
    data = bytearray(10 * 1024 * 1024)  # 10 MB
    time.sleep(0.1)
    print("Garbage collector stats:", gc.get_stats())

# Run with: python3 -v to see GC traces
```

**Key difference:** Python uses **reference counting** — objects are freed IMMEDIATELY when their reference count reaches 0. Go's GC runs **periodically** in the background. Both have trade-offs.

## 7. Step-by-Step Execution

```
GC CYCLE — PHASE 1: MARK (Concurrent)

Before mark:                         After mark:
Heap: [A, B, C, D, E, F, G]         White = C, E, G (dead)
                                    Gray = B (needs scanning)
Roots: [A, D]                       Black = A, D, F (live)

Step 1: Color roots gray
  [A(Gray), D(Gray)] + heap

Step 2: Scan gray objects
  Scan A → finds B → color B gray, color A black
  Scan D → finds F → color F gray, color D black

Step 3: Continue scanning
  Scan B → no children → color B black
  Scan F → finds nothing → color F black

Step 4: All done — only white objects remain
  White: [C, E, G] → these are DEAD (unreachable)
  Black: [A, B, D, F] → these are LIVE (keep)

GC CYCLE — PHASE 2: SWEEP (Concurrent)

Walk entire heap:
  Found C: not in black set → FREE
  Found E: not in black set → FREE
  Found G: not in black set → FREE
  Found A: in black set → KEEP
  Found B: in black set → KEEP
  ...

Memory is returned to the system (or kept for reuse).
```

## 8. Visual Explanation

```
THE TRI-COLOR SYSTEM:

  ⚪ WHITE:  Not visited yet. At end of mark, = DEAD (trash)
  🔘 GRAY:   Found alive, but haven't scanned its children yet
  ⚫ BLACK:  Fully scanned. ALL alive = definitely KEEP

MARK PHASE FLOW:

  Start:  [⚪A, ⚪B, ⚪C, ⚪D, ⚪E, ⚪F]
          roots = [A, D]
          │
          ▼
  Step 1: [🔘A, ⚪B, ⚪C, 🔘D, ⚪E, ⚪F]
          roots colored gray
          │
          ▼
  Step 2: [⚫A, 🔘B, ⚪C, ⚫D, ⚪E, 🔘F]
          A → B, D → F scanned. A, D now black
          │
          ▼
  Step 3: [⚫A, ⚫B, ⚪C, ⚫D, ⚪E, ⚫F]
          All gray scanned → all black or white
          │
          ▼
  Result: ⚪C, ⚪E = DEAD → sweep will free them
          ⚫A, ⚫B, ⚫D, ⚫F = ALIVE → keep

WRITE BARRIER (prevents "hide a white behind black"):

  During concurrent GC:
    if you do: blackObj.field = whiteObj
    Write barrier fires → whiteObj is colored GRAY
    So the GC will still scan it!

  Without write barrier:
    GC scans A → done (black)
    You put C inside A: A.field = C
    GC never re-scans A → C stays white → FREED → BUG!

GC ASSISTS ("If you mess it up, you clean it up"):

  Program allocates fast: 🚀🚀🚀
  GC marks slowly:       🐢
  Program hits threshold → GC forces program to HELP mark
  This slows the program down → rate of allocation decreases
  GC catches up → normal operation resumes

GOGC TUNING:

  GOGC = 100 (default):
    heap_before = 10 MB
    heap_limit = 20 MB  ← GC starts when heap reaches 20 MB
    heap grows 2× before collection

  GOGC = 200:
    heap_before = 10 MB
    heap_limit = 30 MB  ← less frequent GC, more memory usage
    heap grows 3× before collection

  GOGC = 50:
    heap_before = 10 MB
    heap_limit = 15 MB  ← more frequent GC, less memory usage
    heap grows 1.5× before collection
```

## 9. Real-World Analogy

**Cleaning Robot in a Messy Playroom (from the PDF):**

| GC Concept | Playroom Analogy |
|---|---|
| **White objects** | Toys the robot hasn't looked at yet. Marked for trash |
| **Gray objects** | Toys the robot knows you're using, but hasn't checked inside yet |
| **Black objects** | Toys fully scanned — you're definitely using them |
| **Roots** | Toys in your hands right now (definitely keep) |
| **Mark phase** | Robot walks around, tags toys as keep or trash |
| **Sweep phase** | Robot throws away all white (trash) toys |
| **Write barrier** | Sensor on the floor. If you pick up a white toy and move it, sensor turns it gray |
| **GC Assist** | If you're making a mess faster than robot cleans, robot makes YOU help clean |
| **GOGC** | How full should the room get before robot starts cleaning? |
| **GODEBUG=gctrace=1** | Watch the robot's cleaning log |

## 10. Real-World Use Cases

| Use Case | GC Impact |
|---|---|
| **High-throughput APIs** | GC pauses cause latency spikes. Tune GOGC or set GOMEMLIMIT |
| **Real-time systems** | Pre-allocate everything to avoid GC during critical paths |
| **Data pipelines** | Large heap → longer GC cycles. Monitor with gctrace |
| **Microservices** | Small heaps → fast GC. But watch for GC Assist under load |
| **Game servers** | Frame pacing must account for GC pauses |
| **CLI tools** | GC doesn't matter for short-lived programs |

## 11. Common Beginner Mistakes

**Mistake 1: Disabling GC**
```go
// DEBUG ONLY! NEVER in production
debug.SetGCPercent(-1)  // GOGC=off
```
**Why it hurts:** Your program grows until it runs out of memory and crashes.

**Mistake 2: Calling runtime.GC() manually**
```go
for {
    runtime.GC()  // Forces STW pause! Hurts performance
}
```
**Why it hurts:** `runtime.GC()` forces a blocking, synchronous collection. The concurrent GC is better.

**Mistake 3: Ignoring GC pressure**
```go
func handleRequest() {
    for i := 0; i < 10000; i++ {
        x := &SomeStruct{}  // lots of heap allocations
    }
}
```
**Fix:** Pool objects with `sync.Pool` or reduce allocations:
```go
var pool = sync.Pool{New: func() any { return &SomeStruct{} }}
func handleRequest() {
    for i := 0; i < 10000; i++ {
        x := pool.Get().(*SomeStruct)
        // use x
        pool.Put(x)
    }
}
```

**Mistake 4: Not monitoring GC in production**
```go
// Run with: GODEBUG=gctrace=1
// Output: gc 1 @0.008s 4%: 0.024+0.30+0.012 ms clock, ...
// You can SEE pause times and GC CPU usage
```

## 12. Best Practices

1. **Don't optimize GC prematurely** — the default GC is excellent
2. **Use `GODEBUG=gctrace=1`** to observe GC behavior in staging
3. **Set `GOMEMLIMIT`** (Go 1.19+) for predictable memory usage
4. **`sync.Pool`** for frequently allocated, temporary objects
5. **Pre-allocate slices**: `make([]T, 0, expectedCap)` reduces reallocations
6. **Use value receivers** for small structs (fewer heap allocations)
7. **Check escape analysis**: `go build -gcflags="-m"` before optimizing GC
8. **`GOGC=off`** is NEVER safe in production
9. **GC pauses are NOT the enemy** — GC Assist (program helping GC) is usually the bigger cost
10. **Profile first** — use `pprof` to find allocation hot spots before GC-tuning

## 13. Summary Table

| Python | Go | Notes |
|---|---|---|
| Reference counting | Concurrent mark-sweep | Different approaches |
| `gc.collect()` | `runtime.GC()` | Force collection |
| `gc.disable()` | `GOGC=off` | Disable (dangerous) |
| `gc.get_stats()` | `runtime.ReadMemStats()` | Get GC statistics |
| `gc.set_threshold()` | `debug.SetGCPercent()` | Tune collection frequency |
| `-X` no equivalent | `GOMEMLIMIT` (Go 1.19+) | Soft memory limit |
| 3 generations | 1 generation (non-generational) | Heap organization |
| Generational threshold | `GOGC` (double heap size) | Collection trigger |

## 14. Key Takeaways

1. Go's GC is **Non-Generational, Tri-Color, Concurrent Mark-Sweep**
2. **Mark** = find live objects (color white→gray→black)
3. **Sweep** = free white (dead) objects
4. **Write barrier** prevents losing objects modified during GC
5. **GC Assist** slows your program if you allocate faster than GC can mark
6. **GOGC=100** = start GC when heap doubles (default)
7. Use `GODEBUG=gctrace=1` to see GC logs
8. GC pauses are typically **< 500 μs**
9. Don't call `runtime.GC()` manually — let the auto GC work
10. Profile before optimizing — most GC is fast enough

---

## Practice Exercises

### Easy: Watch GC in Action
Write a program that allocates a 100 MB slice, then sets it to nil. Run with `GODEBUG=gctrace=1 go run main.go`. Observe the GC cycle: mark phase time, sweep phase time, and how much was freed.

### Medium: GOGC Tuning
Write a program that allocates memory in a loop (allocate 1 MB, sleep 10ms, repeat 100 times). Run 3 times with `GOGC=100`, `GOGC=50`, and `GOGC=200`. Measure total runtime, number of GC cycles, and max heap size using `runtime.ReadMemStats`.

### Challenging: sync.Pool Performance
Create a benchmark that:
1. Allocates 1,000,000 small structs directly (no pooling)
2. Allocates the same number using `sync.Pool`
Run with `go test -bench=. -benchmem`. Compare allocations per operation and nanoseconds per operation. Report the difference.
