# Goroutines in Go

> Complete learning guide following the Go-by-Python teaching method.

---

## 1. What is it?

A **goroutine** is a lightweight concurrent function execution in Go. You add the keyword `go` before any function call, and Go runs it in the background — without blocking the rest of your program.

Goroutines are **NOT** operating system threads. They are managed by Go's runtime scheduler, which multiplexes thousands (or millions) of goroutines onto a small pool of OS threads.

## 2. Why do we need it?

**Problem:** In most languages, running multiple tasks at once requires threads. Threads are expensive:
- Each thread takes ~8 MB of stack memory
- Creating a thread takes ~1 millisecond
- Context switching between threads is slow

**Go's solution:** Goroutines fix all of this:
- Start with ~2 KB stack (grows as needed)
- Creation takes ~0.1 microseconds
- Go's scheduler handles switching efficiently

This means you can have **thousands of concurrent tasks** in Go — something impractical with OS threads.

## 3. Python Comparison

| Feature | Go | Python |
|---|---|---|
| Launch concurrent task | `go fn()` | `threading.Thread(target=fn).start()` |
| Stack size | ~2 KB (grows) | ~8 MB (fixed per thread) |
| Max practical count | Millions | Thousands |
| Wait for completion | `sync.WaitGroup` | `.join()` (one at a time) |
| Async needed? | No — just `go` | Yes — `asyncio` + `await` |
| Scheduler | Go runtime (userspace) | OS kernel (threads) |
| CPU parallelism | Automatic (if cores available) | GIL blocks parallel CPU work |

## 4. Syntax

```go
go functionName()         // Start a named function as goroutine

go func() {               // Start an anonymous (inline) goroutine
    // code runs concurrently
}()

go func(x int) {          // With parameter
    fmt.Println(x)
}(42)
```

**Keywords explained:**
- `go` — tells Go to run this function in a new goroutine
- `func` — same function syntax, nothing special
- `()` at end — invokes the anonymous function immediately

## 5. Simple Example

```go
package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 0; i < 3; i++ {
        fmt.Println(msg)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go printMessage("Hello from goroutine!")  // runs concurrently
    printMessage("Hello from main!")          // runs in main goroutine

    time.Sleep(500 * time.Millisecond)        // wait for goroutine
}
```

**Line-by-line explanation:**
- `go printMessage(...)` — starts a goroutine, returns immediately
- `printMessage("Hello from main!")` — runs normally in the main goroutine
- `time.Sleep(500ms)` — gives the goroutine time to finish (otherwise main exits and kills it)

## 6. Python Equivalent

```python
import threading
import time

def print_message(msg):
    for i in range(3):
        print(msg)
        time.sleep(0.1)

# Python way
t = threading.Thread(target=print_message, args=("Hello from thread!",))
t.start()                     # equivalent to "go printMessage(...)"
print_message("Hello from main!")
t.join()                      # wait for thread to finish
```

**Key differences:**
- Go: `go fn()` — just 1 word. Python: `threading.Thread(target=fn, args=()).start()` — 3-4 lines
- Go: No `.join()` needed — use `WaitGroup` instead
- Python: `.join()` blocks until thread finishes (one at a time)

## 7. Step-by-Step Execution

```
main() starts
    │
    ├── go printMessage("goroutine")  ──►  Go scheduler queues this
    │      │                                  ┌─ GOROUTINE QUEUE ──┐
    │      │                                  │ printMessage()     │
    │      │                                  └────────────────────┘
    │      │      Go runtime will run it on an available OS thread
    │
    ├── printMessage("main")  ──►  runs immediately in main goroutine
    │      ├─ prints "Hello from main!"   (i=0)
    │      ├─ sleeps 100ms  ──►  scheduler may run the other goroutine now!
    │      │                       ├─ prints "Hello from goroutine!" (i=0)
    │      │                       └─ sleeps 100ms
    │      ├─ prints "Hello from main!"   (i=1)
    │      └─ sleeps 100ms
    │
    └── main() exits ──► ALL goroutines die (if unfinished)
```

## 8. Visual Explanation

```
WITHOUT WaitGroup (goroutine may never run):

  main ──► go fn() ──► [background] ──► main exits
                             │               │
                          goroutine        KILLED!
                          queued            (never runs)

WITH WaitGroup (proper waiting):

  main ──► go fn() ──► [background] ──► wg.Wait() ──► main exits
                             │               │             │
                       goroutine runs    goroutine      main waits
                       calls wg.Done()   done signal    for done

CONCURRENT vs SEQUENTIAL timing:

  SEQUENTIAL:   [Task 1: 200ms] ──► [Task 2: 200ms] ──► [Task 3: 200ms] = 600ms

  CONCURRENT:   [Task 1: 200ms] ─────────────────────────────┐
                [Task 2: 200ms] ───────────────────────────┐ │
                [Task 3: 200ms] ─────────────────────────┐ │ │ ≈ 200ms!
                                                         ▼ ▼ ▼
                                                  All finish together!
```

## 9. Real-World Analogy

**Factory Assembly Line:**

- **Main goroutine** = Foreman/Manager. Assigns tasks, doesn't do them.
- `go fn()` = "Hey Worker #1, start assembling part A" — foreman doesn't wait.
- **Goroutines** = Workers on the floor. Each works independently.
- `sync.WaitGroup` = Clipboard with checkboxes. Foreman adds checkboxes (`wg.Add(1)`), workers check them off (`wg.Done()`), foreman waits for all checked (`wg.Wait()`).
- **OS Thread** = Actual physical table. One worker per table. But Go's scheduler can put multiple workers at one table, switching between them rapidly.

**Key insight:** You don't manage tables (threads). You just assign workers (goroutines), and the scheduler finds tables for them.

## 10. Real-World Use Cases

| Use Case | Why Goroutines? |
|---|---|
| **Web servers** | One goroutine per incoming request. Handle 10k+ connections easily |
| **File processing** | Process 100 files concurrently — goroutine per file |
| **Background workers** | Worker pool consuming jobs from a channel |
| **API calls** | Fan-out: call 5 external APIs concurrently, collect all results |
| **Real-time data** | One goroutine per WebSocket connection |
| **Batch jobs** | Process chunks of data in parallel goroutines |
| **Database queries** | Run multiple queries concurrently, wait for first result |
| **Health checks** | Ping 50 services in parallel instead of sequentially |

## 11. Common Beginner Mistakes

**Mistake 1: Main exits before goroutine runs**
```go
func main() {
    go fmt.Println("Never runs!")
}   // ❌ main exits immediately, goroutine never executes
```
**Fix:** Use WaitGroup (not sleep)
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("Runs!")
}()
wg.Wait()
```

**Mistake 2: Forgetting to call wg.Done()**
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    // forgot wg.Done() — program hangs forever!
}()
wg.Wait()  // waits forever — DEADLOCK
```
**Fix:** Always use `defer wg.Done()` at the start of the goroutine

**Mistake 3: Loop variable capture bug (Go < 1.22)**
```go
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)  // prints 3, 3, 3 — all see the LAST value!
    }()
}
```
**Fix:** Pass i as argument (creates a copy per iteration)
```go
for i := 0; i < 3; i++ {
    go func(id int) {
        fmt.Println(id)  // prints 0, 1, 2
    }(i)
}
```

**Mistake 4: Using time.Sleep instead of WaitGroup**
```go
go doWork()
time.Sleep(1 * time.Second)  // fragile — what if work takes 2 seconds?
```
**Fix:** Always use WaitGroup for proper synchronization

## 12. Best Practices

1. **Always use `sync.WaitGroup`** to wait for goroutines (never `time.Sleep`)
2. **Start with `defer wg.Done()`** as the first line of your goroutine
3. **Pass loop variables as arguments**, don't capture them by closure
4. **Limit goroutine count** with worker pools — unbounded goroutines can exhaust memory
5. **Handle panics inside goroutines** — an unhandled panic crashes the whole program
6. **Use channels to communicate** between goroutines (not shared variables)
7. **Test with `go run -race`** to detect data races
8. **Prefer `go` for I/O-bound work** (files, network, databases) — CPU-bound work may need a worker pool
9. **Name your goroutines** with clear function names so stack traces are readable
10. **Don't start goroutines you can't stop** — use context cancellation or done channels

## 13. Summary Table

| Python | Go | Notes |
|---|---|---|
| `threading.Thread(target=fn).start()` | `go fn()` | Launch concurrent task |
| `.join()` | `sync.WaitGroup.Wait()` | Wait for completion |
| Not needed | `wg.Add(1)` + `wg.Done()` | Track completion count |
| `threading.active_count()` | No built-in | Goroutine count |
| `threading.current_thread()` | No built-in | Current goroutine |
| `threading.Lock()` | `sync.Mutex` | Mutual exclusion |
| OS thread (~8 MB stack) | Goroutine (~2 KB stack) | Memory per concurrent task |
| GIL-limited parallelism | True parallelism (if cores) | CPU utilization |

## 14. Key Takeaways

1. `go fn()` launches a goroutine — Go's simplest concurrency primitive
2. Goroutines are **lightweight** (~2 KB stack, millions possible)
3. **Not OS threads** — Go's scheduler multiplexes them onto threads
4. **WaitGroup** is the standard way to wait: `Add → Done → Wait`
5. Always pass loop variables as arguments (closure gotcha)
6. Use `defer wg.Done()` at the top of every goroutine
7. Main exiting **kills all goroutines** — always wait for them
8. Goroutines + channels = Go's concurrency model
9. Test with `go run -race` for data race detection
10. Goroutines are cheap, but not free — use worker pools for large numbers

---

## Practice Exercises

### Easy: Count to 5 in Parallel
Write a program that launches 5 goroutines, each printing a number 1-5. Use a WaitGroup to ensure all finish before main exits.

### Medium: Concurrent URL Fetcher (Simulated)
Write a program that launches 3 goroutines, each simulating an API call by sleeping a different amount of time (200ms, 400ms, 600ms). Use a WaitGroup and print when each one completes.

### Challenging: Closure Fix
Given this buggy code:
```go
func main() {
    names := []string{"Alice", "Bob", "Charlie"}
    for _, name := range names {
        go func() {
            fmt.Println("Hello,", name)
        }()
    }
    time.Sleep(100 * time.Millisecond)
}
```
Fix it so it prints "Hello, Alice", "Hello, Bob", "Hello, Charlie" (any order). Then rewrite it to use a WaitGroup instead of time.Sleep.
