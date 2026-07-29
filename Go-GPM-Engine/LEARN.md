# Go Scheduler — G-P-M Engine

> Complete learning guide following the Go-by-Python teaching method.

---

## 1. What is it?

The **G-P-M engine** is Go's runtime scheduler. It decides **which goroutine runs on which OS thread at any given moment**. The three letters stand for:

- **G (Goroutine)** — a lightweight task (~2 KB stack)
- **P (Processor)** — a logical CPU core (a "desk" to work at)
- **M (Machine)** — an OS thread (the "muscle" that does work)

The scheduler multiplexes **millions of Gs** onto **a few Ms** using **Ps** as the bridge. This is called the **M:N scheduling model** — M goroutines are scheduled onto N OS threads.

## 2. Why do we need it?

**Problem:** OS threads are expensive. Each thread takes ~2 MB of stack memory and ~1 ms to create. In Python, every time you do `threading.Thread(target=fn).start()`, you pay this cost. If you wanted 100,000 concurrent tasks, you'd crash.

**Go's solution:** Goroutines start at ~2 KB. A single OS thread can run thousands of goroutines by switching between them. The G-P-M scheduler makes this possible:

| Resource | OS Thread | Goroutine |
|---|---|---|
| Stack size | ~2 MB (fixed) | ~2 KB (grows as needed) |
| Create time | ~1 ms | ~0.1 μs |
| Max on 8 GB RAM | ~4,000 | ~4,000,000 |
| Context switch | OS kernel (slow) | Go runtime (fast) |

## 3. Python Comparison

| Feature | Go | Python |
|---|---|---|
| Execution unit | Goroutine (G) | OS Thread |
| Scheduling | Go runtime (user-mode) | OS kernel (kernel-mode) |
| Stack | ~2 KB, grows | ~8 MB, fixed |
| Number of Ps | `GOMAXPROCS` (default = CPU cores) | N/A (no scheduler control) |
| Work stealing | Yes (automatic) | No |
| Continuation stealing | Yes (Go's innovation) | No |
| Blocking syscall | M releases P, new M takes over | Thread blocks entirely |
| Debug scheduler | `GODEBUG=schedtrace=1000` | `threading.enumerate()` |

## 4. Syntax

There's no syntax for the G-P-M scheduler — it runs **automatically** in the Go runtime. But you can control and observe it:

```go
// Control: GOMAXPROCS
runtime.GOMAXPROCS(4)  // Set P count (usually = CPU cores)
numCPU := runtime.NumCPU()  // Get CPU count
numG := runtime.NumGoroutine()  // Get active goroutine count

// Observe scheduler
// GODEBUG=schedtrace=1000 go run main.go
// Prints: SCHED 0ms: gomaxprocs=8 idleMs=0 threads=4 ...

// The go keyword creates a G
go myFunction()  // scheduler assigns this G to a P+M
```

**The go keyword triggers the entire G-P-M dance.** Everything else is automatic.

## 5. Simple Example

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    // Show current scheduler state
    fmt.Println("CPU cores:", runtime.NumCPU())
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
    fmt.Println("Goroutines at start:", runtime.NumGoroutine())

    // Launch goroutines — scheduler manages them
    for i := 0; i < 5; i++ {
        go func(n int) {
            fmt.Println("Goroutine", n, "running on P")
        }(i)
    }

    fmt.Println("Goroutines after launch:", runtime.NumGoroutine())
}
```

**Line-by-line:**
1. `runtime.NumCPU()` — how many CPU cores the machine has
2. `runtime.GOMAXPROCS(0)` — current P count (usually = NumCPU())
3. `go func(n int){...}(i)` — creates 5 Gs, scheduler assigns them to available P+M slots
4. `runtime.NumGoroutine()` — count active goroutines (includes main + 5)

## 6. Python Equivalent

```python
import threading
import os
import time

# Python doesn't expose a scheduler directly
# But these give some insight:
print("CPU cores:", os.cpu_count())
print("Active threads:", threading.active_count())

def worker(n):
    print(f"Thread {n} running")

threads = []
for i in range(5):
    t = threading.Thread(target=worker, args=(i,))
    threads.append(t)
    t.start()

for t in threads:
    t.join()

# Python has NO equivalent of:
# - runtime.NumGoroutine()
# - runtime.GOMAXPROCS()
# - Work stealing or continuation stealing
```

**Key difference:** Python threads are OS threads — each one is heavy and managed by the kernel. Go goroutines are lightweight and managed by Go's runtime.

## 7. Step-by-Step Execution

```
1. Program starts:
   main() runs as G1 on M1 at P1
   P count = GOMAXPROCS (e.g., 4)

2. go fn() is called:
   ┌────────────────────────────────────────────────────┐
   │  A new G2 is created (starts in runnable state)    │
   │  G2 is placed on P1's LOCAL RUN QUEUE              │
   │  P1's queue: [G2, ...]                             │
   │  M1 continues executing G1                         │
   └────────────────────────────────────────────────────┘

3. Scheduler ticks (every ~10ms or at blocking points):
   ┌────────────────────────────────────────────────────┐
   │  M1:P1 picks G2 from local queue                   │
   │  G2 executes for a while                           │
   │  G2's turn ends → back to queue                    │
   │  M1:P1 picks next G from queue                     │
   └────────────────────────────────────────────────────┘

4. Blocking syscall (e.g., file read):
   ┌────────────────────────────────────────────────────┐
   │  G2 calls read(file) → blocks                      │
   │  M1 stays with G2 (both blocked until read done)   │
   │  P1 is DETACHED from M1 → available for new M      │
   │  A new/different M2 picks up P1                    │
   │  M2:P1 continues running other Gs from queue       │
   │  When read completes → G2 goes back to queue       │
   └────────────────────────────────────────────────────┘

5. Work stealing (P has empty queue):
   ┌────────────────────────────────────────────────────┐
   │  P2's queue is empty → P2 becomes "Work Thief"     │
   │  P2 randomly picks another P (e.g., P5)            │
   │  P2 steals HALF of P5's queued Gs                  │
   │  Now P2 can keep working → no idle CPU             │
   └────────────────────────────────────────────────────┘

6. Continuation stealing (Go's secret sauce):
   ┌────────────────────────────────────────────────────┐
   │  go fn() is called inside G1:                      │
   │  Instead of: put new G on queue, continue G1       │
   │  Go does:  jump into new G IMMEDIATELY             │
   │            put the rest of G1 on the queue          │
   │  Benefit: new G has hot CPU cache from G1          │
   │  (the data G1 was touching is what new G needs)    │
   └────────────────────────────────────────────────────┘
```

## 8. Visual Explanation

```
THE G-P-M TRIO:

  ┌──────────┐     ┌──────────┐     ┌──────────┐
  │    G     │     │    P     │     │    M     │
  │ Goroutine│ ◄──►│Processor │ ◄──►│ Machine  │
  │ (Task)   │     │ (Desk)   │     │ (Thread) │
  │ ~2 KB    │     │ GOMAXPROCS│    │ ~2 MB    │
  │ Millions │     │ = CPU cnt│    │ Few      │
  └──────────┘     └──────────┘     └──────────┘

THE M:N SCHEDULING MODEL:

  ┌────────┐  ┌────────┐  ┌────────┐  ┌────────┐
  │  G1    │  │  G2    │  │  G3    │  │  G4    │  ... millions of Gs
  └───┬────┘  └───┬────┘  └───┬────┘  └───┬────┘
      │           │           │           │
      ▼           ▼           ▼           ▼
  ┌─────────────────────────────────────────────┐
  │             MULTIPLEXER (scheduler)          │
  └─────┬───────────────┬───────────────┬────────┘
        │               │               │
        ▼               ▼               ▼
    ┌──────┐        ┌──────┐        ┌──────┐
    │  M1  │        │  M2  │        │  M3  │    ... few Ms (OS threads)
    └──┬───┘        └──┬───┘        └──┬───┘
       │               │               │
       ▼               ▼               ▼
   ┌──────┐        ┌──────┐        ┌──────┐
   │ P1   │        │ P2   │        │ P3   │    ... Ps (logical CPUs)
   └──────┘        └──────┘        └──────┘

LOCAL RUN QUEUE (each P has one):

  P1's queue:  [G5, G8, G2, G9, G3]
  P2's queue:  [G7, G1]                    ← almost empty!
  P3's queue:  [G4, G6, G10, G11]

  P2 steals from P1:
  P2's queue:  [G7, G1, G8, G9]           ← stolen half of P1's Gs
  P1's queue:  [G5, G2, G3]               ← remaining

GLOBAL RUN QUEUE (when all local queues empty):

  Global queue: [G100, G101, G102, ...]
  ↑ P steals from global when local is empty

WORK STEALING DIAGRAM:

  Before steal:          After steal:
  ┌────────────────┐     ┌────────────────┐
  │ P1: [1,2,3,4,5]│     │ P1: [1,2,3]    │
  ├────────────────┤     ├────────────────┤
  │ P2: [] (idle)  │     │ P2: [4,5]      │  ← stolen!
  └────────────────┘     └────────────────┘

CONTINUATION STEALING:

  Task stealing (old way):
    go fn()
    ├── Put fn() on queue
    └── Continue current function
    Result: fn() runs separately, cache cold

  Continuation stealing (Go way):
    go fn()
    ├── Jump INTO fn() NOW
    └── Put rest of current function on queue
    Result: fn() runs immediately, cache WARM!
```

## 9. Real-World Analogy

**Video Game Development Studio (from the PDF):**

| G-P-M Concept | Studio Analogy |
|---|---|
| **G (Goroutine)** | A **task ticket** — "Build character model", "Design level 3" |
| **P (Processor)** | A **desk/computer** in the office. Number of desks = office capacity |
| **M (Machine/Thread)** | A **developer** who sits at a desk and does the work |
| **Local run queue** | A developer's **to-do pile** on their desk |
| **Work stealing** | A developer who finishes their pile **steals half** from a busy colleague |
| **Continuation stealing** | A new urgent task arrives → developer **drops everything and starts it**, puts old work back in pile |
| **Blocking syscall** | Developer spills coffee, leaves desk. **Another developer** takes the desk |
| **Global queue** | **Manager's board** — tasks that nobody has picked up yet |

## 10. Real-World Use Cases

| Use Case | Why G-P-M Matters |
|---|---|
| **Web servers (Go net/http)** | Each request = 1 goroutine. Scheduler handles 10k+ concurrent requests on few threads |
| **Database connection pools** | Goroutines block on DB response, M releases P, other goroutines keep working |
| **gRPC services** | Many concurrent streams, scheduler multiplexes efficiently |
| **File servers** | Blocking I/O doesn't stall CPU — P detached, new M takes over |
| **High-throughput APIs** | Continuation stealing keeps CPU cache hot for new goroutines |
| **Kubernetes controllers** | Many concurrent reconciliation loops managed by scheduler |

## 11. Common Beginner Mistakes

**Mistake 1: Setting GOMAXPROCS too high**
```go
runtime.GOMAXPROCS(100)  // More than CPU cores!
```
**Why it hurts:** More Ps mean more contention for CPU time. Context switching overhead increases. Usually, `NumCPU()` is optimal.

**Mistake 2: Not knowing about work stealing**
```go
// Assumption: goroutines are evenly distributed
for i := 0; i < 100; i++ {
    go work()  // scheduler balances these via work stealing
}
```
**Reality:** Scheduler handles this. You don't need to manually distribute work.

**Mistake 3: Assuming goroutines are free**
```go
for {
    go func() { /* infinite loop */ }()  // goroutine leak!
}
```
**Why it hurts:** Even though goroutines are cheap, they're not free. Each takes memory and scheduler overhead. Unbounded goroutines exhaust RAM.

**Mistake 4: Blocking the M without noticing**
```go
func main() {
    runtime.GOMAXPROCS(1)
    go fmt.Println("Hello")
    // main is the only goroutine on P1
    // But main never blocks or yields!
    // The goroutine never runs until main exits
}
```
**Fix:** Use `runtime.Gosched()` to yield, or better, use proper synchronization.

## 12. Best Practices

1. **Don't set GOMAXPROCS** — let it default to NumCPU()
2. **Use `GODEBUG=schedtrace=1000`** to observe scheduler behavior in production
3. **Use worker pools** for CPU-bound work to limit goroutine count
4. **For I/O-bound work**, many goroutines are fine — scheduler handles blocking well
5. **Use `runtime.NumGoroutine()`** to detect goroutine leaks in tests
6. **Understand that the scheduler is cooperative** — goroutines yield at function calls, not time slices
7. **Don't use `runtime.Gosched()`** in production code — it's a red flag
8. **Blocking operations** (file I/O, network) are handled efficiently — don't try to avoid them
9. **CPU-bound loops** without function calls can starve other goroutines — insert `runtime.Gosched()` or restructure
10. **Trust the scheduler** — it's one of the most optimized parts of Go

## 13. Summary Table

| Python | Go |
|---|---|
| `threading.Thread` (OS thread) | Goroutine (G) |
| No equivalent | P (Processor / logical CPU) |
| `os.cpu_count()` | `runtime.NumCPU()` |
| No equivalent | `runtime.GOMAXPROCS(n)` |
| No equivalent | `runtime.NumGoroutine()` |
| Kernel thread scheduler | Go runtime scheduler |
| Thread blocks on syscall | M releases P, new M takes over |
| No work stealing | Automatic work stealing |
| No equivalent | `GODEBUG=schedtrace=1000` |
| ~8 MB per thread | ~2 KB per goroutine |

## 14. Key Takeaways

1. **G** = goroutine (task), **P** = processor (desk), **M** = machine (OS thread)
2. Go uses **M:N scheduling** — M goroutines on N threads
3. **P count** = `GOMAXPROCS` (default = CPU cores)
4. Each P has a **local run queue** of goroutines
5. **Work stealing** — idle Ps steal work from busy Ps (no idle CPU)
6. **Continuation stealing** — Go's innovation: jump into new goroutine immediately (cache warm)
7. Blocking syscall → M releases P → new M takes the P (no CPU idle)
8. Goroutines are **not free** — unbounded goroutines exhaust RAM
9. Use `GODEBUG=schedtrace=1000` to observe scheduler
10. **Trust the scheduler** — it's more optimized than manual tuning

---

## Practice Exercises

### Easy: Observe Scheduler
Write a program that launches 10 goroutines doing `fmt.Println` in a loop. Run it with `GODEBUG=schedtrace=1000` and observe the output. What does `gomaxprocs`, `idleMs`, and `threads` tell you?

### Medium: GOMAXPROCS Benchmark
Write a CPU-bound benchmark (e.g., calculating primes) that runs with `runtime.GOMAXPROCS(1)` vs `runtime.GOMAXPROCS(NumCPU())`. Measure the time difference using `time` package.

### Challenging: Goroutine Starvation
Write a program that:
1. Sets `GOMAXPROCS(1)`
2. Has a goroutine that runs an infinite CPU-bound loop (no function calls)
3. Has another goroutine that should print "I'm alive!" periodically
Observe that the second goroutine NEVER runs. Then fix it by adding `runtime.Gosched()` in the loop. Explain why this happens.
