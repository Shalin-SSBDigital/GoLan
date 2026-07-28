// ============================================================
//  CHANNELS IN GO — with Lego Brick Analogy 🧱
// ============================================================
//
//                      🧱  LEGO BRICK ANALOGY  🧱
//
//   Alex and Sam are in different rooms, building Lego.
//          ┌──────────── Room A ────────────┐
//          │  Alex builds spaceships!       │
//          │  But Alex needs RED 4x2 bricks │
//          │  that Sam has.                 │
//          └────────────────────────────────┘
//                      │
//                      │  PIPE through the wall
//                      ▼
//          ┌──────────── Room B ────────────┐
//          │  Sam has a bucket of bricks!   │
//          │  Sam pushes a red 4x2 brick    │
//          │  into the pipe → it arrives    │
//          │  in Alex's room.               │
//          │  Alex takes it and builds!     │
//          └────────────────────────────────┘
//
//   ┌─────────────────────────────────────────────────────────┐
//   │  The PIPE is a CHANNEL                                  │
//   │  The BRICK is the DATA (type-specific)                  │
//   │  Alex PUSHING a brick = SEND (ch <- brick)              │
//   │  Sam TAKING a brick = RECEIVE (brick := <-ch)           │
//   │  Only ONE type of brick fits a given pipe type          │
//   └─────────────────────────────────────────────────────────┘
//
// ┌────────────┬────────────────────────────────────────────────┐
// │    Go      │                 Python                         │
// ├────────────┼────────────────────────────────────────────────┤
// │ ch :=      │ from queue import Queue                       │
// │   make(    │ q = Queue()                                   │
// │   chan int)│ q.put(42)       # send                        │
// │ ch <- 42   │ val = q.get()   # receive                     │
// │ val := <-ch│                                                │
// │            │ BUT: goroutines + channels are BUILT-IN,       │
// │            │ not library. No import needed.                 │
// │            │ Channels BLOCK until ready.                    │
// │            │ Python Queue needs polling or put/get wait.    │
// └────────────┴────────────────────────────────────────────────┘
// ============================================================

package main

import (
	"fmt"
	"time"
)

// ============================================================
//  BASIC CHANNEL — Unbuffered
//  Unbuffered channel = a pipe that can hold ONLY ONE brick.
//  Alex puts a brick in → the brick waits until Sam takes it.
//  Sam cannot take until Alex puts. Both must be ready.
// ============================================================

func alexSendsBrick(ch chan string) {
	// Alex has a red 4x2 Lego brick
	brick := "🧱 Red 4x2 brick"
	fmt.Printf("  👦 Alex pushes a %s into the pipe...\n", brick)
	ch <- brick // Send — blocks until Sam receives
	fmt.Println("  👦 Alex: 'Brick sent! Back to building.'")
}

func samReceivesBrick(ch chan string) {
	// Sam waits at the pipe opening
	fmt.Println("  👧 Sam: 'Waiting for a brick...'")
	brick := <-ch // Receive — blocks until Alex sends
	fmt.Printf("  👧 Sam got a %s from the pipe!\n", brick)
	fmt.Println("  👧 Sam: 'Perfect, adding it to my castle!'")
}

func main() {
	fmt.Println("========================================")
	fmt.Println("  🧱  LEGO BRICK CHANNEL DEMO 🧱")
	fmt.Println("========================================")
	fmt.Println()

	// Create the pipe — a channel for LEGO bricks (strings)
	// make(chan string) creates an UNBUFFERED channel
	// The pipe has zero capacity — only 1 brick can pass at a time
	pipe := make(chan string)

	fmt.Println("─── Scenario 1: One brick through the pipe ───")
	fmt.Println()

	// Alex and Sam work at the same time (goroutines)
	go alexSendsBrick(pipe)
	go samReceivesBrick(pipe)

	// Give them time to finish
	time.Sleep(1 * time.Second)
	fmt.Println()

	// ============================================================
	//  BUFFERED CHANNEL — Pipe with a bucket
	//  Sam puts several bricks into the pipe before Alex looks.
	//  The pipe has a waiting shelf (buffer).
	//  Sam can fill the shelf and walk away.
	//  Alex picks up bricks whenever.
	// ============================================================
	fmt.Println("─── Scenario 2: Buffered pipe (shelf holds 3 bricks) ───")
	fmt.Println()

	// make(chan string, 3) creates a BUFFERED channel
	// The pipe has a shelf that holds up to 3 bricks
	shelfPipe := make(chan string, 3)

	// Sam quickly puts 3 bricks into the shelf-pipe (no one receiving yet!)
	shelfPipe <- "🧱 Blue 2x4 brick"
	shelfPipe <- "🧱 Yellow 1x2 brick"
	shelfPipe <- "🧱 Green 2x2 brick"
	fmt.Println("  👧 Sam stuffed 3 bricks into the shelf-pipe!")
	fmt.Println("  👧 Sam: 'Alex can grab them whenever.'")
	fmt.Println()

	// Alex picks them up one at a time
	fmt.Println("  👦 Alex reaches into the shelf-pipe...")
	for i := 0; i < 3; i++ {
		brick := <-shelfPipe
		fmt.Printf("  👦 Alex pulled out: %s\n", brick)
	}
	fmt.Println("  👦 Alex: 'Got all the bricks! Thanks Sam!'")
	fmt.Println()

	// ============================================================
	//  CHANNEL WITH GOROUTINE SYNC — Build Together
	//  Alex and Sam build different parts, then swap bricks.
	//  This shows how channels sync two goroutines.
	// ============================================================
	fmt.Println("─── Scenario 3: Two builders swap bricks ───")
	fmt.Println()

	pipeAtoB := make(chan string)
	pipeBtoA := make(chan string)

	go func() {
		// Alex builds a wing, then sends it
		fmt.Println("  👦 Alex: 'Building a wing panel...'")
		time.Sleep(200 * time.Millisecond)
		pipeAtoB <- "🪽 Wing panel (from Alex)"
		fmt.Println("  👦 Alex: 'Wing sent! Now waiting for Sam's wheel...'")

		// Alex waits for Sam's wheel
		wheel := <-pipeBtoA
		fmt.Printf("  👦 Alex received: %s\n", wheel)
		fmt.Println("  👦 Alex: 'Now I can finish the spaceship!'")
	}()

	go func() {
		// Sam builds a wheel, then sends it
		fmt.Println("  👧 Sam: 'Making a wheel...'")
		time.Sleep(300 * time.Millisecond)
		pipeBtoA <- "⚙️ Wheel (from Sam)"
		fmt.Println("  👧 Sam: 'Wheel sent! Now waiting for Alex's wing...'")

		// Sam waits for Alex's wing
		wing := <-pipeAtoB
		fmt.Printf("  👧 Sam received: %s\n", wing)
		fmt.Println("  👧 Sam: 'Perfect, the car chassis is ready!'")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()

	// ============================================================
	//  CLOSING A CHANNEL — No more bricks
	//  Sam says "that's all I have" by closing the pipe.
	//  Alex knows to stop waiting.
	// ============================================================
	fmt.Println("─── Scenario 4: Closing the pipe (no more bricks) ───")
	fmt.Println()

	brickBucket := make(chan string, 2)
	brickBucket <- "🧱 White 2x2 brick"
	brickBucket <- "🧱 Black 4x2 brick"
	close(brickBucket) // Sam: "That's all I've got!"
	fmt.Println("  👧 Sam closed the pipe: 'No more bricks!'")

	fmt.Println("  👦 Alex unloads the pipe...")
	for brick := range brickBucket {
		fmt.Printf("  👦 Alex got: %s\n", brick)
	}
	// After close, range loop exits automatically
	fmt.Println("  👦 Alex: 'Empty pipe. Time to build!'")

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("  🏗️  BUILD COMPLETE!  🏗️")
	fmt.Println("========================================")
}

// ============================================================
//  KEY CONCEPTS SUMMARY
// ============================================================
//
//  ┌───────────────────┬──────────────────────────────────────┐
//  │ Concept           │ Meaning                              │
//  ├───────────────────┼──────────────────────────────────────┤
//  │ make(chan T)      │ Create unbuffered pipe for type T    │
//  │                   │ (1 brick at a time, sync handoff)    │
//  ├───────────────────┼──────────────────────────────────────┤
//  │ make(chan T, N)   │ Create buffered pipe holding N bricks│
//  │                   │ (async, sender doesn't block until N)│
//  ├───────────────────┼──────────────────────────────────────┤
//  │ ch <- value       │ Push a brick into the pipe           │
//  │                   │ (blocks if pipe full)               │
//  ├───────────────────┼──────────────────────────────────────┤
//  │ value := <-ch     │ Pull a brick from the pipe           │
//  │                   │ (blocks if pipe empty)              │
//  ├───────────────────┼──────────────────────────────────────┤
//  │ close(ch)         │ Say "no more bricks"                 │
//  │                   │ Receivers can detect via `v, ok`     │
//  ├───────────────────┼──────────────────────────────────────┤
//  │ range ch          │ Keep pulling until pipe is closed    │
//  │                   │ Auto-exits when closed               │
//  ├───────────────────┼──────────────────────────────────────┤
//  │ Goroutine + chan  │ Two kids + a pipe = concurrency!     │
//  │                   │ They build simultaneously,           │
//  │                   │ syncing only when swapping bricks.   │
//  └───────────────────┴──────────────────────────────────────┘
//
//  ⚠️ Deadlock warning: If Alex pushes and nobody receives,
//     everyone freezes. Always pair send ↔ receive, or use
//     buffered channels so bricks have a shelf to sit on.
// ============================================================