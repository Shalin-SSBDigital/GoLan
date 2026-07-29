# Channels in Go

> Complete learning guide following the Go-by-Python teaching method.

---

## 1. What is it?

A **channel** is a typed pipe that connects goroutines. One goroutine sends a value into the channel, and another goroutine receives that value out of the channel.

Think of a channel like a **pipe through a wall** between two rooms. Alex pushes a Lego brick into the pipe from Room A, and Sam pulls that same brick out in Room B.

Key properties:
- Channels are **typed** — a `chan int` only carries integers
- Channels **block** — sending blocks until someone receives, and vice versa
- Channels are **safe** — multiple goroutines can use them without locks

## 2. Why do we need it?

**Problem:** When two goroutines need to share data or coordinate, you could use shared variables with locks. But locks are error-prone — they cause deadlocks, races, and complicated code.

**Go's solution:** Channels provide a **safe, simple communication mechanism** between goroutines. The Go mantra:

> "Do not communicate by sharing memory; instead, share memory by communicating."

Instead of "goroutine A writes to variable X, goroutine B reads X (protected by mutex)", you say "goroutine A sends value to channel, goroutine B receives from channel."

## 3. Python Comparison

| Feature | Go | Python |
|---|---|---|
| Create | `ch := make(chan T)` | `q = queue.Queue()` |
| Send | `ch <- value` | `q.put(value)` |
| Receive | `value := <-ch` | `q.get()` |
| Buffered | `make(chan T, N)` | `Queue(maxsize=N)` |
| Close | `close(ch)` | No built-in (sentinel value pattern) |
| Iterate | `for v := range ch` | Manual loop with `q.get()` |
| Block on send | Yes (if buffer full) | No (unless `block=True`) |
| Block on receive | Yes (if empty) | Yes (unless `block=False`) |
| Type-safe | Yes (compile-time) | No (any Python object) |

**Key difference:** Channels are **built into Go** — no imports needed. Python needs `from queue import Queue`.

## 4. Syntax

```go
ch := make(chan int)          // Unbuffered channel — holds 0 items
ch := make(chan int, 5)       // Buffered channel — holds 5 items

ch <- 42                      // Send: push 42 into the channel
value := <-ch                 // Receive: pull a value out of the channel

close(ch)                     // Close: signal that no more sends will happen

for v := range ch {           // Iterate: receive until channel is closed
    fmt.Println(v)
}

v, ok := <-ch                 // Check if channel is open
// ok == true  → value received
// ok == false → channel is closed, v is zero value
```

**Operators explained:**
- `make(chan T)` — creates a channel of type T (`chan T`)
- `ch <-` — **arrow pointing into** the channel = **send**
- `<-ch` — **arrow pointing out of** the channel = **receive**
- `close(ch)` — marks the channel as closed (sends will panic)
- `range ch` — receives until the channel is closed

## 5. Simple Example

```go
package main

import "fmt"

func main() {
    // Create a channel for strings (like a pipe for Lego bricks)
    pipe := make(chan string)

    // Send in a goroutine (sends block until received)
    go func() {
        pipe <- "Hello from the pipe!"  // send
    }()

    // Receive in main goroutine (blocks until something is sent)
    msg := <-pipe                        // receive
    fmt.Println(msg)
}
```

**Line-by-line explanation:**
1. `pipe := make(chan string)` — creates an unbuffered channel for strings
2. `go func() { pipe <- "..." }()` — goroutine sends a string into the pipe
3. `msg := <-pipe` — main receives from the pipe (blocks until the goroutine sends)
4. Both happen simultaneously — the send and receive are synchronized

## 6. Python Equivalent

```python
from queue import Queue
import threading

q = Queue()

def sender():
    q.put("Hello from the queue!")

t = threading.Thread(target=sender)
t.start()

msg = q.get()
print(msg)
t.join()
```

**Line-by-line comparison:**

| Go | Python |
|---|---|
| `pipe := make(chan string)` | `q = Queue()` |
| `go func() { pipe <- msg }()` | `t = Thread(target=sender); t.start()` with `q.put(msg)` |
| `msg := <-pipe` | `msg = q.get()` |
| No close needed here | No close needed here |

## 7. Step-by-Step Execution

### Unbuffered Channel

```
main() creates pipe (chan string)
    │
    ├── go func() launches goroutine
    │      └── pipe <- "Hello"  ──► BLOCKS (no receiver yet)
    │                                    │
    └── msg := <-pipe  ──► BLOCKS (nothing in pipe yet)
         │                     │
         │              ┌──────┘
         │              ▼
         │    SEND AND RECEIVE HAPPEN AT THE SAME TIME
         │    "Hello" is transferred directly from sender to receiver
         │
         ├── msg = "Hello"
         └── goroutine unblocks and exits
```

### Buffered Channel (buffer = 2)

```
ch := make(chan string, 2)

ch <- "A"     ──► stored in buffer [_, _]  →  [A, _]
ch <- "B"     ──► stored in buffer [A, _]  →  [A, B]
ch <- "C"     ──► BLOCKS! Buffer is full [A, B]

msg := <-ch   ──► msg = "A", buffer [_, B]  →  now ch <- "C" can proceed
```

## 8. Visual Explanation

```
UNBUFFERED CHANNEL (synchronous handoff):

  Sender                          Receiver
  ┌────────┐                     ┌────────┐
  │  "Hi!"  │─ ─ ─ ┬ ─ ─ ─ ─ ►│  "Hi!"  │
  └────────┘       │           └────────┘
                   │
              Both must be
              ready at the
              exact same time

BUFFERED CHANNEL (buffer = 3):

  Sender places up to 3 items without receiver

  ┌────────┐    ┌────┬────┬────┐    ┌────────┐
  │  "A"   │──►│ A  │ B  │ C  │──►│  "A"   │
  │  "B"   │──►│    │    │    │──►│  "B"   │
  │  "C"   │──►└────┴────┴────┘──►│  "C"   │
  └────────┘    Buffer shelf      └────────┘

  Sender blocks ONLY when buffer is full (4th item)
  Receiver blocks ONLY when buffer is empty

CLOSE + RANGE:

  close(ch)   ──► channel is marked "no more sends"

  for v := range ch:   keeps pulling until channel is closed
                       then exits automatically

  v, ok := <-ch:       ok == false when channel is closed AND empty
```

## 9. Real-World Analogy

**Lego Brick Through a Pipe (Alex & Sam):**

Alex and Sam are in different rooms, building Lego. There's a pipe through the wall.

| Go Concept | Lego Analogy |
|---|---|
| `make(chan string)` | Installing a pipe between rooms |
| `ch <- "red brick"` | Alex pushes a red brick into the pipe |
| `brick := <-ch` | Sam pulls a brick out of the pipe |
| Unbuffered channel | Pipe holds exactly 1 brick. Alex must wait until Sam takes it |
| Buffered channel | Pipe has a shelf holding N bricks. Alex can put bricks on the shelf and walk away |
| `close(ch)` | Sam tapes a note on the pipe: "No more bricks needed" |
| `range ch` | Sam keeps pulling bricks until the "no more" note is seen |

## 10. Real-World Use Cases

| Use Case | How Channels Help |
|---|---|
| **Worker pools** | Jobs channel distributes work to goroutine workers |
| **Pipelines** | Stage 1 sends to ch1, Stage 2 receives from ch1 and sends to ch2 |
| **Fan-out** | One channel feeds multiple goroutine workers |
| **Fan-in** | Multiple goroutines send results to one channel |
| **Task cancellation** | Close a done channel to signal all goroutines to stop |
| **Rate limiting** | Buffered channel as a token bucket |
| **Request/response** | Each request gets its own response channel |

## 11. Common Beginner Mistakes

**Mistake 1: Deadlock on unbuffered channel (send without receive)**
```go
ch := make(chan int)
ch <- 42       // ❌ BLOCKS FOREVER — no one is receiving!
fmt.Println(<-ch)  // never reached
```
**Fix:** Always ensure there's a receiver, use a goroutine, or use a buffered channel

**Mistake 2: Sending on a closed channel**
```go
ch := make(chan int)
close(ch)
ch <- 42       // ❌ PANIC: send on closed channel
```
**Fix:** Only the sender should close. Never send after close.

**Mistake 3: Range on a channel that's never closed**
```go
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    // forgot close(ch)
}()
for v := range ch {   // ❌ LEAKS — waits forever for more values
    fmt.Println(v)
}
```
**Fix:** Always `close(ch)` when done sending.

**Mistake 4: Using channels where a mutex is simpler**
```go
// Over-engineered — a mutex is simpler for protecting a counter
type Counter struct {
    value int
    ch     chan int
}
```
**Fix:** Use channels for communication, mutexes for shared state.

## 12. Best Practices

1. **Unbuffered channels** for synchronization (handoff guarantees both are ready)
2. **Buffered channels** for decoupling (sender and receiver work at different speeds)
3. **Only the sender should close** a channel — never the receiver
4. **Close is a signal**, not a cleanup — receivers detect it via `v, ok := <-ch`
5. **Don't close channels you don't own** — leads to panics
6. **Prefer `range` for reading** until closed — cleaner than manual `v, ok` loops
7. **Use buffered channels for known capacities** to avoid goroutine leaks
8. **Use `select` for multiple channels** — clean way to wait on many at once
9. **Nil channels block forever** — sometimes useful for disabling cases in select
10. **Channel of struct{}** is a zero-size channel — used purely for signaling

## 13. Summary Table

| Python | Go | Notes |
|---|---|---|
| `from queue import Queue` | Built-in (`make(chan T)`) | No import needed |
| `q = Queue()` | `ch := make(chan T)` | Unbuffered |
| `q = Queue(maxsize=5)` | `ch := make(chan T, 5)` | Buffered |
| `q.put(x)` | `ch <- x` | Send |
| `x = q.get()` | `x := <-ch` | Receive |
| No equivalent | `close(ch)` | Signal no more sends |
| No equivalent | `for x := range ch` | Iterate until closed |
| `q.empty()` | `len(ch)` | Current buffer length |
| `q.maxsize` | `cap(ch)` | Buffer capacity |

## 14. Key Takeaways

1. Channels are **typed pipes** for goroutine communication
2. **Unbuffered** channels sync sender and receiver (both must be ready)
3. **Buffered** channels decouple sender and receiver (buffer holds N items)
4. `ch <- value` = **send**, `value := <-ch` = **receive**
5. `close(ch)` signals "no more sends" — only the sender closes
6. `for v := range ch` iterates until the channel is closed
7. Send on closed channel = **panic**
8. Channels **block** — sends block until received, receives block until sent
9. Use buffered channels to avoid deadlocks in simple scenarios
10. Go motto: "Share memory by communicating"

---

## Practice Exercises

### Easy: Simple Message Pass
Write a program where the main goroutine sends "hello" to a goroutine via an unbuffered channel, and the goroutine prints it. Use WaitGroup to wait for the goroutine.

### Medium: Number Pipeline
Create a pipeline: goroutine 1 generates numbers 1-5 and sends them to a channel. Goroutine 2 receives, doubles each number, and sends to another channel. Main goroutine prints the results.

### Challenging: Fan-Out Worker Pool
Create a buffered channel of jobs (5 jobs: "job 1" through "job 5"). Launch 3 worker goroutines that each read from the jobs channel, process (print the job name + "processed"), and send results to a results channel. Use close and range properly. Ensure no deadlocks and all jobs are processed exactly once.
