# Scalable Concurrency Pipelines

> Complete learning guide following the Go-by-Python teaching method.

---

## 1. What is it?

A **concurrency pipeline** is a series of stages connected by channels. Data flows through the pipeline one stage at a time, where each stage runs concurrently in its own goroutine. Like a factory assembly line — raw materials go in, finished products come out, and multiple workers process different parts simultaneously.

**Key patterns:**
- **Fan-out** — distribute work across multiple goroutine workers
- **Fan-in** — combine results from multiple goroutines into one channel
- **Pipeline** — chain stages: Stage1 → ch1 → Stage2 → ch2 → Stage3

## 2. Why do we need it?

**Problem:** Processing data sequentially is slow. If each step takes 100ms, processing 1000 items takes 100 seconds. But if you pipeline and parallelize, the same work can finish in seconds.

**Go's solution:** Channels + goroutines = natural pipeline construction. Each stage is a function that takes an input channel and returns an output channel. Stages connect with `ch1 := stage1(input)` then `ch2 := stage2(ch1)`.

## 3. Python Comparison

| Feature | Go | Python |
|---|---|---|
| Pipeline stage | Goroutine + channel | Generator / asyncio queue |
| Fan-out | Multiple goroutines reading one channel | `concurrent.futures.ThreadPoolExecutor` |
| Fan-in | Merge multiple channels into one | `asyncio.gather()` |
| Worker pool | Fixed goroutines reading jobs channel | `ThreadPoolExecutor(max_workers=N)` |
| Cancel pipeline | `context.Context` | `asyncio.CancelledError` |
| Rate limit | Token bucket + channel | `asyncio.Semaphore` |
| Pipeline stage function | `func(in <-chan T) <-chan R` | Generator function with `yield` |

## 4. Syntax

```go
// Pipeline stage signature (idiomatic Go):
func stageName(input <-chan InputType) <-chan OutputType {
    output := make(chan OutputType)
    go func() {
        defer close(output)
        for v := range input {
            // process v
            output <- processedV
        }
    }()
    return output
}

// Fan-out (multiple workers reading one channel):
func fanOut(input <-chan Job, numWorkers int) []<-chan Result {
    workers := make([]<-chan Result, numWorkers)
    for i := 0; i < numWorkers; i++ {
        workers[i] = worker(input)
    }
    return workers
}

// Fan-in (merge multiple channels into one):
func fanIn(channels ...<-chan Result) <-chan Result {
    output := make(chan Result)
    var wg sync.WaitGroup
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan Result) {
            defer wg.Done()
            for v := range c {
                output <- v
            }
        }(ch)
    }
    go func() {
        wg.Wait()
        close(output)
    }()
    return output
}
```

## 5. Simple Example

```go
package main

import (
    "fmt"
    "sync"
)

// Stage 1: Generate numbers
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Stage 3: Print results
func printResults(in <-chan int) {
    for r := range in {
        fmt.Println("Result:", r)
    }
}

func main() {
    // Pipeline: generate → square → print
    numbers := generate(1, 2, 3, 4, 5)
    squares := square(numbers)
    printResults(squares)
}
```

**Line-by-line:**
1. `generate()` — returns a channel, launches goroutine to send values
2. `square()` — takes input channel, returns output channel, processes concurrently
3. `printResults()` — reads from channel until closed
4. Pipeline flows: `1,2,3,4,5` → `1,4,9,16,25` → printed

## 6. Python Equivalent

```python
import asyncio

# Python equivalent using async generators
async def generate(*nums):
    for n in nums:
        yield n

async def square(input_gen):
    async for n in input_gen:
        yield n * n

async def main():
    gen = generate(1, 2, 3, 4, 5)
    sq = square(gen)
    async for r in sq:
        print(f"Result: {r}")

asyncio.run(main())
```

**Key differences:**
- Go stages run in **separate goroutines** (true concurrency). Python async generators run in **one event loop** (cooperative multitasking).
- Go channels are **typed** and **buffered/unbuffered**. Python generators are **untyped** and **unbuffered**.
- Go fan-out is **trivial** (multiple goroutines, one channel). Python needs thread pool executors.

## 7. Step-by-Step Execution

```
PIPELINE FLOW:

  generate()              square()              printResults()
  ┌──────────┐    ch1     ┌──────────┐    ch2    ┌──────────────┐
  │          │  ───────►  │          │  ───────►  │              │
  │ 1,2,3,4,5│           │ 1,4,9,16,│           │ prints each  │
  │          │           │ 25       │           │ result       │
  └──────────┘           └──────────┘           └──────────────┘

  Time:
  ch1: [1]─[2]─[3]─[4]─[5]─────►
  ch2: ────[1]─[4]─[9]─[16]─[25]►
  print: ────[1]─[4]─[9]─[16]─[25]

  The pipeline processes concurrently!
  While square() works on 2, print() is printing 1.
  Total time ≈ max(stage time) × count, NOT sum(stage times) × count

FAN-OUT / FAN-IN FOR LOG PROCESSING:

  Log Generator              Analysis Workers              Aggregator
  (Kafka/File)               (Fan-out)                    (Fan-in)
  ┌────────┐    logsCh       ┌──────────┐    resultsCh     ┌──────────┐
  │        │  ────────────►  │ Worker 1 │  ────────────►  │          │
  │ Millions│                ├──────────┤                 │ Combined │
  │ of logs │                │ Worker 2 │                  │ output   │
  │        │                 ├──────────┤                  │          │
  └────────┘                 │ Worker 3 │                  └──────────┘
                             └──────────┘

  Worker pool size = runtime.NumCPU() (typically)
  Each worker processes different log line simultaneously
  Results are all merged into one aggregator
```

## 8. Visual Explanation

```
BASIC PIPELINE (3 stages):

  ┌──────────┐   ch1   ┌──────────┐   ch2   ┌──────────┐
  │ Stage 1  │ ──────► │ Stage 2  │ ──────► │ Stage 3  │
  │ Generate │         │ Process  │         │ Collect  │
  └──────────┘         └──────────┘         └──────────┘

FAN-OUT / FAN-IN (distributed log processing):

            ┌──────────┐
            │  Source   │
            │ (Channel) │
            └────┬─────┘
                 │
         ┌───────┼───────────┐
         │       │           │
         ▼       ▼           ▼
     ┌────────┐ ┌────────┐ ┌────────┐
     │Worker 1│ │Worker 2│ │Worker 3│  ← Fan-out
     └───┬────┘ └───┬────┘ └───┬────┘
         │          │          │
         └──────────┼──────────┘
                    ▼
             ┌──────────┐
             │  Merge   │
             │ (Fan-in) │
             └────┬─────┘
                  │
                  ▼
             ┌──────────┐
             │Collector │
             └──────────┘

PIPELINE WITH CONTEXT CANCELLATION:

  ctx, cancel := context.WithCancel(context.Background())

  stage1(ctx, input) → stage2(ctx, ch1) → stage3(ctx, ch2)

  If any stage fails:
    cancel()  ← cascades through all stages via ctx.Done()
    Goroutines exit cleanly
    No goroutine leaks

RATE LIMITED PIPELINE:

  Token bucket channel:
  ticker := time.NewTicker(rate)
  for range ticker.C {
      bucket <- struct{}{}
  }

  Worker waits for token before processing:
  for job := range jobs {
      <-bucket  // blocks if rate exceeded
      process(job)
  }
```

## 9. Real-World Analogy

**Factory Assembly Line (Distributed Log Processing):**

| Pipeline Concept | Factory Analogy |
|---|---|
| **Pipeline stage** | A station on the assembly line (e.g., "Paint Station") |
| **Channel** | Conveyor belt between stations |
| **Fan-out** | Multiple identical stations processing items in parallel |
| **Fan-in** | Multiple conveyor belts merging into one output belt |
| **Worker pool** | 5 painters, each painting a different item simultaneously |
| **Generator** | Raw materials arrive on a truck |
| **Bounded parallelism** | Only N painters — beyond that, the line backs up |
| **Context cancellation** | Emergency stop button — everything shuts down |
| **Rate limiting** | Output belt can only handle 10 items/minute — throttle input |

**Distributed Log Processing Pipeline:**
1. **Generator** reads millions of log lines from Kafka
2. **Fan-out** to 8 workers (one per CPU core)
3. Each worker parses logs, runs regex for security patterns
4. **Fan-in** results into single aggregator
5. Aggregator sends alerts to dashboard/database

## 10. Real-World Use Cases

| Use Case | Pipeline Architecture |
|---|---|
| **ETL pipelines** | Extract → Transform → Load (fan-out transformations) |
| **Log processing** | Read logs → Parse → Analyze → Alert → Store |
| **Image processing** | Load images → Resize → Filter → Compress → Save |
| **API aggregation** | Fan-out to 5 services → Fan-in results → Respond |
| **Video transcoding** | Split video → Transcode chunks → Merge → Upload |
| **Data validation** | Read records → Validate (fan-out) → Collect errors → Report |

## 11. Common Beginner Mistakes

**Mistake 1: Not closing output channels**
```go
func stage(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for v := range in {
            out <- v * 2
        }
        // forgot close(out) — next stage hangs forever!
    }()
    return out
}
```
**Fix:** Always `defer close(out)` or `close(out)` when the goroutine finishes.

**Mistake 2: Unbounded goroutine creation**
```go
for _, log := range millionsOfLogs {
    go process(log)  // MILLIONS of goroutines — crash!
}
```
**Fix:** Use a bounded worker pool.
```go
workers := runtime.NumCPU()
for i := 0; i < workers; i++ {
    go worker(jobsCh)
}
```

**Mistake 3: Deadlock in fan-in (forgetting wg.Wait)**
```go
func fanIn(chs ...<-chan int) <-chan int {
    out := make(chan int)
    for _, ch := range chs {
        go func(c <-chan int) {
            for v := range c {
                out <- v
            }
        }(ch)
    }
    // Never closed! Reader hangs forever
    return out
}
```
**Fix:** Use `sync.WaitGroup` to close after all inputs are done.

**Mistake 4: No context cancellation**
```go
func longPipeline(ctx context.Context, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for v := range in {
            // No way to cancel mid-pipeline!
            result := doHeavyWork(v)
            out <- result
        }
    }()
    return out
}
```
**Fix:** Always select on `ctx.Done()` in long-running stages.

## 12. Best Practices

1. **Follow the pipeline pattern**: `func(in <-chan T) <-chan R` — consistent signatures
2. **Always close output channels** when the goroutine finishes
3. **Use `sync.WaitGroup`** in fan-in to know when to close the merged channel
4. **Use `context.Context`** as first parameter for cancelable pipelines
5. **Bounded parallelism**: use a fixed-size worker pool, never unbounded goroutines
6. **Error handling**: send errors through a "result struct" channel, not as separate error channels
7. **Rate limit** upstream if downstream can't keep up
8. **Buffer channels** when stages work at different speeds (but not too much)
9. **Test with `go run -race`** — data races are common in pipelines
10. **Monitor goroutine count**: `runtime.NumGoroutine()` to detect leaks

## 13. Summary Table

| Python | Go | Notes |
|---|---|---|
| Generator (`yield`) | Goroutine + channel | Pipeline stage |
| `asyncio.gather()` | Fan-in pattern | Merge results |
| `ThreadPoolExecutor` | Worker pool (fan-out) | Bounded goroutines |
| `asyncio.CancelledError` | `context.Context.Done()` | Pipeline cancellation |
| `asyncio.Semaphore` | Token bucket channel | Rate limiting |
| `asyncio.Queue(maxsize)` | `make(chan T, N)` | Bounded channel |
| Generator exit | `close(ch)` | Signal pipeline end |

## 14. Key Takeaways

1. **Pipeline** = stages connected by channels (each stage is a goroutine)
2. **Fan-out** = distribute work across N workers (one input channel, N goroutines)
3. **Fan-in** = merge N output channels into one (use WaitGroup)
4. **Always close output channels** — or downstream stages hang forever
5. **Always use bounded parallelism** — worker pool, not unbounded goroutines
6. **Use `context.Context`** for clean cancellation of pipelines
7. **Buffer channels** to handle speed differences between stages
8. **Send errors through data channels** — use result structs with error fields
9. **Rate limit** to protect downstream services
10. **Test with `go run -race`** to catch race conditions

---

## Practice Exercises

### Easy: 3-Stage Pipeline
Build a pipeline: generate numbers 1-100 → double each → print. Each stage should be its own goroutine with `func(in <-chan int) <-chan int` signature.

### Medium: Fan-Out Prime Finder
Build a worker pool pipeline:
1. Generator sends numbers 1-100 to a job channel
2. 4 workers each receive from jobs, check if the number is prime, send to results
3. Fan-in merges results, collector prints primes
Use WaitGroup in the fan-in to ensure proper channel closing.

### Challenging: Log Processing Pipeline with Context Cancellation
Build a simulated log processing pipeline:
1. Generate 1000 log entries (random strings with "ERROR", "WARN", "INFO" levels)
2. Fan-out to 4 workers that filter for "ERROR" entries
3. Fan-in filtered results
4. Add `context.Context` — cancel after 10 errors found (simulate "too many errors, stop processing"). All goroutines should exit cleanly when cancelled.
