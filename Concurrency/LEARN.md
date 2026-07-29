# Concurrency in Go

> Complete learning guide following the Go-by-Python teaching method.

---

## 1. What is it?

**Concurrency** is the ability of a program to **deal with multiple tasks at the same time**. It's about **structuring** your program to handle multiple things that are in progress simultaneously.

Go has concurrency built into the language with three primitives:
- **Goroutines** — lightweight concurrent functions
- **Channels** — typed pipes for communication between goroutines
- **Select** — wait on multiple channels at once

**Important distinction:**
- **Concurrency** = structure (dealing with many things)
- **Parallelism** = execution (doing many things at once)
- Go enables concurrency. Parallelism happens automatically if hardware supports it.

## 2. Why do we need it?

**Problem:** Modern programs need to handle many things at once — thousands of network requests, multiple database queries, file operations, API calls. Without concurrency, everything runs sequentially, which is slow.

**Go's solution:** Go makes concurrency simple and safe:
- `go keyword` starts a goroutine — no complex thread management
- Channels provide safe communication — no shared memory bugs
- `select` handles multiple channels — no callback hell
- Goroutines are lightweight — you can have millions

**Python comparison:** In Python, you need:
- `threading` (heavy threads, GIL limits)
- `asyncio` (requires `async`/`await` everywhere)
- `multiprocessing` (separate processes, complex communication)
Go's approach is simpler: just `go fn()`.

## 3. Python Comparison

| Feature | Go | Python |
|---|---|---|
| Concurrent task | `go fn()` | `threading.Thread`, `asyncio.create_task` |
| Communication | `ch <- val` / `<-ch` | `queue.Queue`, `asyncio.Queue` |
| Multi-wait | `select { case ... }` | `asyncio.wait(FIRST_COMPLETED)` |
| Mutual exclusion | `sync.Mutex` | `threading.Lock` |
| Wait for group | `sync.WaitGroup` | `concurrent.futures.wait(ALL_COMPLETED)` |
| Per-task stack | ~2 KB | ~8 MB (thread) or ~1 KB (coroutine) |
| Async keyword needed? | No | Yes (`async def`, `await`) |
| Parallel CPU work | Automatic (if cores) | Blocked by GIL |

## 4. Syntax

```go
go myFunction()              // Start a goroutine

ch := make(chan int)         // Create a channel
ch <- 42                     // Send
value := <-ch                // Receive

select {                     // Wait on multiple channels
case v1 := <-ch1:
    // handle v1
case v2 := <-ch2:
    // handle v2
case <-time.After(1 * time.Second):
    // timeout after 1 second
default:
    // non-blocking: run if no channel is ready
}

var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

## 5. Simple Example

```go
package main

import (
    "fmt"
    "time"
)

func doWork(id int, done chan bool) {
    fmt.Printf("Worker %d: starting...\n", id)
    time.Sleep(1 * time.Second)
    fmt.Printf("Worker %d: done!\n", id)
    done <- true  // signal completion
}

func main() {
    done := make(chan bool)

    // Start 3 concurrent workers
    for i := 1; i <= 3; i++ {
        go doWork(i, done)
    }

    // Wait for all 3 to complete
    for i := 1; i <= 3; i++ {
        <-done  // receive 3 completion signals
    }

    fmt.Println("All done!")
}
```

**Line-by-line explanation:**
1. `make(chan bool)` — creates a channel for completion signals
2. `go doWork(i, done)` — starts worker i concurrently
3. `<-done` — blocks until a worker signals completion
4. Loop 3 times — receives from the channel 3 times (once per worker)
5. After all 3 receives, main continues — all workers are done

## 6. Python Equivalent

```python
import threading
import time

def do_work(worker_id, results):
    print(f"Worker {worker_id}: starting...")
    time.sleep(1)
    print(f"Worker {worker_id}: done!")
    results.append(worker_id)  # signal completion

results = []  # shared list — needs lock!
lock = threading.Lock()

def safe_work(worker_id):
    do_work(worker_id, results)

threads = []
for i in range(1, 4):
    t = threading.Thread(target=safe_work, args=(i,))
    threads.append(t)
    t.start()

for t in threads:
    t.join()

print("All done!")
```

**Key differences:**
- Go uses **channel** for signaling (`done <- true`) — no shared variable needed
- Python uses **shared list** `results` — needs synchronization
- Go's channel is **built-in** and **type-safe**
- Go: `<-done` blocks like `.join()`, but one channel can signal many goroutines

## 7. Step-by-Step Execution

### Sequential (without concurrency):

```
main()                     Time: 0ms
  doWork(1) ──► sleep 1s ──► print ──► Time: 1000ms
  doWork(2) ──► sleep 1s ──► print ──► Time: 2000ms
  doWork(3) ──► sleep 1s ──► print ──► Time: 3000ms
Total: 3 seconds
```

### Concurrent (with goroutines):

```
main()                     Time: 0ms
  │
  ├── go doWork(1) ──► sleep 1s ──► done ───┐
  ├── go doWork(2) ──► sleep 1s ──► done ───┤  Time: ~1 second
  ├── go doWork(3) ──► sleep 1s ──► done ───┘
  │
  └── <-done (3 times) ──► All done!
Total: ~1 second
```

## 8. Visual Explanation

```
CONCURRENCY vs PARALLELISM:

CONCURRENCY (structure — one CPU, switching tasks):
  ┌─────┐   ┌─────┐   ┌─────┐
  │Task1│──►│Task2│──►│Task3│   One person juggling 3 tasks
  └─────┘   └─────┘   └─────┘   Switches between them
     │         │         │
     ▼         ▼         ▼
  [Goroutine 1] [Goroutine 2] [Goroutine 3]
  ───┬──────┬──────┬──────┬──────┬──────► Time
     │ T1   │ T2   │ T3   │ T1   │
     └──────┴──────┴──────┴──────┘
  (Only one runs at any instant, but they overlap in time)

PARALLELISM (execution — multiple CPUs, truly simultaneous):
  Core 0: ── T1 ──────────────────────────────────
  Core 1: ── T2 ──────────────────────────────────
  Core 2: ── T3 ──────────────────────────────────
  (All run at the exact same time)

FAN-OUT / FAN-IN PATTERN:

  Jobs                     Workers                  Results
  ┌─────┐          ┌──────────────┐          ┌──────────────┐
  │Job 1│──┬──────►│  Worker 1    │──┬──────►│  Result 1    │
  ├─────┤  │       └──────────────┘  │       ├──────────────┤
  │Job 2│──┼──────►│  Worker 2    │──┼──────►│  Result 2    │
  ├─────┤  │       └──────────────┘  │       ├──────────────┤
  │Job 3│──┼──────►│  Worker 3    │──┼──────►│  Result 3    │
  ├─────┤  │       └──────────────┘  │       ├──────────────┤
  │Job 4│──┤                         │       │  Result 4    │
  ├─────┤  │                         │       ├──────────────┤
  │Job 5│──┘                         └──────►│  Result 5    │
  └─────┘                                   └──────────────┘

  Fan-out: jobs distributed across workers
  Fan-in:  results collected from workers
```

## 9. Real-World Analogy

**Restaurant Kitchen:**

| Concept | Kitchen Analogy |
|---|---|
| **Sequential** | One chef cooks dish by dish. Start-to-finish. Slow. |
| **Concurrent** | Chef starts rice, while rice cooks chops veggies, while veggies cook plates. Uses time efficiently. |
| **Parallel** | Two chefs, each cooking their own order. Needs more stoves (CPUs). |
| **Goroutine** | "Hey, start chopping these onions!" — chef delegates without waiting |
| **Channel** | Passthrough window between chefs. "Here are the chopped onions." |
| **Select** | Waiter stands at two windows: "Whichever order is ready first." |
| **Mutex** | Only one chef can use the stove at a time. |

## 10. Real-World Use Cases

| Pattern | Use Case | Go Implementation |
|---|---|---|
| **Fan-out/Fan-in** | Process 100 files concurrently | Send file paths to jobs channel, workers process, send results back |
| **Pipeline** | ETL data processing | Extract → ch1 → Transform → ch2 → Load |
| **Select + timeout** | API calls with timeout | Make API call on one channel, timeout on another |
| **Worker pool** | Database migration tool | Fixed number of workers process migration jobs |
| **Pub/Sub** | Event broadcasting | Multiple subscribers listen on same channel |
| **Throttling** | Rate-limited API client | Buffered channel as token bucket |

## 11. Common Beginner Mistakes

**Mistake 1: Deadlock — main goroutine blocked waiting**
```go
func main() {
    ch := make(chan int)
    ch <- 42          // ❌ BLOCKS: no receiver yet!
    fmt.Println(<-ch) // never reached
}
```
**Fix:** Send in a goroutine, or use a buffered channel

**Mistake 2: Data race — shared variable without mutex**
```go
var counter int
for i := 0; i < 1000; i++ {
    go func() { counter++ }()  // ❌ RACE: multiple goroutines write counter
}
```
**Fix:** Use mutex or channel
```go
var mu sync.Mutex
go func() { mu.Lock(); counter++; mu.Unlock() }()
```

**Mistake 3: Select without default blocks forever**
```go
select {
case v := <-ch:
    fmt.Println(v)
// no default — blocks if ch is empty (may be intentional)
}
```
**Fix:** Add `default:` for non-blocking behavior

**Mistake 4: Mixing concurrency and parallelism assumptions**
```go
// DON'T assume goroutines will run in any order
for i := 0; i < 3; i++ {
    go fmt.Println(i)  // prints ??? (random order)
}
```
**Fix:** Always coordinate output or use channels for ordering

**Mistake 5: Starting more goroutines than needed**
```go
for _, item := range hugeList {
    go process(item)  // ❌ 100,000 goroutines! Memory disaster!
}
```
**Fix:** Use a bounded worker pool with a fixed number of goroutines

## 12. Best Practices

1. **Prefer channels over mutexes** for communication between goroutines
2. **Use mutexes** for protecting simple shared state (counters, caches)
3. **Use `select`** for any function that waits on multiple channels
4. **Always set timeouts** on channel operations that might block
5. **Use `go run -race`** to detect data races in development
6. **Limit goroutine count** with worker pools — unbounded goroutines can exhaust memory
7. **Close channels from the sender side only**
8. **Name channels clearly** — `jobsCh`, `resultsCh`, `doneCh` helps readability
9. **Don't launch goroutines without a plan to stop them**
10. **Concurrency is not parallelism** — design for concurrency, measure parallelism

## 13. Summary Table

| Python | Go | Notes |
|---|---|---|
| `threading.Thread.start()` | `go fn()` | Launch concurrent task |
| `queue.Queue` | `chan T` | Typed communication pipe |
| `q.put(x)` | `ch <- x` | Send to channel |
| `q.get()` | `<-ch` | Receive from channel |
| `threading.Lock()` | `sync.Mutex` | Mutual exclusion |
| `asyncio.wait(FIRST_COMPLETED)` | `select { case ... }` | Wait on multiple |
| `thread.join()` | `sync.WaitGroup.Wait()` | Wait for completion |
| `concurrent.futures.ThreadPoolExecutor` | Worker pool pattern | Bounded concurrent tasks |

## 14. Key Takeaways

1. **Concurrency** = structure (dealing with many things). **Parallelism** = execution (doing many things).
2. Go enables concurrency with **goroutines, channels, and select**
3. Goroutines are **cheap** (~2 KB) — launch thousands without worry
4. Channels are **typed pipes** — safe communication between goroutines
5. `select` lets you **wait on multiple channels** — first ready wins
6. Mutex protects **shared state** — use when channels don't fit
7. Cancelling goroutines can be done with a **done channel**
8. Always use **WaitGroup** or **channels** to wait, never `time.Sleep`
9. Test with **`go run -race`** to catch data races
10. Go motto: "Do not communicate by sharing memory; share memory by communicating."

---

## Practice Exercises

### Easy: Three Concurrent Printers
Write a program that starts 3 goroutines, each printing a different message 5 times. Use a WaitGroup to wait for all to finish. The output should be interleaved (showing concurrency).

### Medium: Select with Timeout
Write a program with two channels. One goroutine sends a message after 2 seconds, another after 500ms. Use `select` to print whichever arrives first. Also add a timeout case that fires at 3 seconds.

### Challenging: Worker Pool Pipeline
Create a 3-stage pipeline:
1. **Generator** goroutine: sends numbers 1-10 to channel A
2. **Worker pool** (3 goroutines): receives from channel A, squares the number, sends to channel B
3. **Collector** goroutine: receives from channel B, sums the results, sends final sum to channel C
Main goroutine prints the final sum. Expected: 1² + 2² + ... + 10² = 385
