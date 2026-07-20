package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// =============================================================================
// STRINGS & RUNES IN GO
// =============================================================================
//
// This is one of the biggest conceptual shifts when coming from Python.
// Go's string handling is fundamentally different because:
//
//   ┌──────────────────────┬────────────────────────────────────────────────┐
//   │        Go            │                  Python 3                      │
//   ├──────────────────────┼────────────────────────────────────────────────┤
//   │ Strings are IMMUTABLE│ Strings are IMMUTABLE bytes internally,        │
//   │ byte SEQUENCES       │ but expose a SEQUENCE OF CODE POINTS          │
//   │ (UTF-8 encoded)      │ (Unicode code points / characters).           │
//   ├──────────────────────┼────────────────────────────────────────────────┤
//   │ len(s) = BYTES       │ len(s) = CHARACTERS (code points)             │
//   │ s[i] = RAW BYTE      │ s[i] = code point (as a 1-char string)        │
//   │ (not a character)    │                                                │
//   ├──────────────────────┼────────────────────────────────────────────────┤
//   │ rune = Unicode code  │ ord("世") → 19990                             │
//   │ point (int32)        │ chr(19990) → "世"                             │
//   │ '世' = 0x4E16       │                                                │
//   ├──────────────────────┼────────────────────────────────────────────────┤
//   │ UTF-8 everywhere     │ UTF-8 by default in files, but strings        │
//   │ by convention        │ are stored as UCS-1/UCS-2/UCS-4 internally    │
//   │                      │ (flexible representation, opaque to user)     │
//   └──────────────────────┴────────────────────────────────────────────────┘
//
// Bottom line:
//   Python 3 strings are a sequence of Unicode code points.
//   Go strings are a sequence of UTF-8-encoded bytes.
//
// Both handle Unicode. But Go forces you to think about the ENCODING,
// while Python hides it (you just see "characters").

// =============================================================================
// 1. BASIC STRING — Bytes Under the Hood
// =============================================================================

func demoBasicString() {
	s := "Hello"

	fmt.Println("=== Basic String: bytes vs characters ===")
	fmt.Printf("s = %q\n", s)

	// len() returns BYTES, not characters
	fmt.Printf("len(s) = %d bytes\n", len(s))
	// Python: len("Hello") → 5 characters
	// Same result here, but for DIFFERENT reasons.
	// Go counts bytes, Python counts code points.
	// They happen to match because ASCII characters are 1 byte in UTF-8.

	fmt.Print("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%02x ", s[i])
	}
	fmt.Println()

	// ⚠️ This is SAFE for ASCII, but DANGEROUS for non-ASCII strings.
	// Python: s[0] → "H" (a 1-character string)
	// Go:     s[0] → 0x48 (the raw BYTE, which happens to be 'H' in ASCII)
	fmt.Printf("s[0] = %v (type: %T, as char: %c)\n", s[0], s[0], s[0])
	// In Python, you get back a string. In Go, you get back a byte.
}

// =============================================================================
// 2. THE PROBLEM — Multi-byte Characters
// =============================================================================

func demoUnicodeProblem() {
	// "Hello" in Japanese katakana:
	s := "ハローワールド" // "Hello World" in Japanese

	fmt.Println("\n=== The Problem: Multi-byte Characters ===")
	fmt.Printf("s = %q\n", s)

	// len() gives BYTES, not characters
	fmt.Printf("len(s) = %d bytes\n", len(s))
	// But visually, there are only ~10 characters!
	// Python: len("ハローワールド") → 9 characters

	fmt.Println("\nNaive byte iteration (WRONG for non-ASCII):")
	for i := 0; i < len(s); i++ {
		fmt.Printf("  byte %d: %02x", i, s[i])
		if s[i] < 128 {
			fmt.Printf(" (%c)", s[i])
		}
		fmt.Println()
	}
	// This prints garbage bytes. Each katakana character is 3 bytes in UTF-8.
	// s[0] gives the FIRST BYTE of the first character, not the character itself.
	//
	// Python equivalent would be: s.encode("utf-8")[0]
	// But Python doesn't let you index bytes from a string — you have to
	// explicitly encode first. Go does the reverse: it always gives you bytes.

	fmt.Println("\nPython comparison:")
	fmt.Println("  Python:  s = 'ハローワールド'")
	fmt.Println("  Python:  len(s) → 9  (code points)")
	fmt.Println("  Python:  s[0] → 'ハ' (a string)")
	fmt.Println("  Go:      len(s) → 27 (UTF-8 bytes)")
	fmt.Println("  Go:      s[0] → 0xE3 (first BYTE)")
	fmt.Println()
	fmt.Println("  Python's string is a sequence of UNICODE CODE POINTS.")
	fmt.Println("  Go's string is a sequence of UTF-8 ENCODED BYTES.")
	fmt.Println("  They are fundamentally different data models.")
}

// =============================================================================
// 3. RUNES — Go's "Character"
// =============================================================================
//
// A `rune` is Go's term for a single Unicode code point.
// `rune` is an alias for `int32` (like `type rune = int32`).
//
// Python equivalent: the integer returned by ord()
//
//   'A'          → rune value: 65        (U+0041)
//   '世'         → rune value: 19990     (U+4E16)
//   'ハ'         → rune value: 12495     (U+30CF)
//
// Python: ord('世') → 19990
// Go:     '世'     → 19990  (rune literal)

func demoRunes() {
	fmt.Println("\n=== Runes ===")

	// A rune literal is a character in single quotes (like Python's ord)
	r1 := 'A'
	r2 := '世'
	r3 := 'ハ'

	fmt.Printf("'A'  = %d (U+%04X)\n", r1, r1)
	fmt.Printf("'世' = %d (U+%04X)\n", r2, r2)
	fmt.Printf("'ハ' = %d (U+%04X)\n", r3, r3)

	// Python equivalents:
	fmt.Println("\nPython comparison:")
	fmt.Println("  ord('世') →", r2)
	fmt.Println("  chr(19990) →", string(r2))
	fmt.Println("  Go has no single built-in; use string(r)")

	// Rune type is just int32:
	fmt.Printf("\nrune type: %s\n", reflect.TypeOf(r1).Kind())
	fmt.Printf("rune is just int32: rune=%d, int32(r)=%d, same=%v\n", r1, int32(r1), r1 == int32(r1))

	// Rune literal vs string literal
	fmt.Println("\nGo literal syntax comparison:")
	fmt.Printf("  '世' = rune literal  → value %[1]d (0x%[1]X)   like Python ord('世')\n", '世')
	fmt.Printf("  \"世\" = string literal → UTF-8 bytes: ")
	for _, b := range []byte("世") {
		fmt.Printf("%02x ", b)
	}
	fmt.Println("  (no direct Python equivalent — would be '世'.encode('utf-8'))")
}

// =============================================================================
// 4. RANGE LOOP — THE Correct Way to Iterate a String
// =============================================================================
//
// Using `for i, r := range s` decodes UTF-8 automatically.
// It gives you:
//   i = byte index (like Python's enumerate on bytes, but not)
//   r = rune (like iterating Python's string directly)
//
// Python:  for i, ch in enumerate(s):   # ch is a 1-char string
// Go:      for i, r := range s {         # r is a rune (int32)

func demoRangeLoop() {
	s := "ハローワールド"

	fmt.Println("\n=== Range Loop — The Correct Way ===")
	fmt.Printf("s = %q\n", s)

	// ✅ CORRECT — decodes UTF-8 on the fly
	fmt.Println("Range over string (decodes runes):")
	for i, r := range s {
		fmt.Printf("  byte[%d] = %c (U+%04X, value=%d)\n", i, r, r, r)
	}

	// ⚠️ Python comparison
	fmt.Println("\nPython equivalent:")
	fmt.Println("  for i, ch in enumerate('ハローワールド'):")
	fmt.Println("      print(i, ch, ord(ch))")
	fmt.Println("  # i is the CHARACTER INDEX (0, 1, 2, ...)")
	fmt.Println("  # Every iteration gives exactly one character")
	fmt.Println()
	fmt.Println("Go: i is the BYTE INDEX (0, 3, 6, 9, ...)")
	fmt.Println("Go: Not every iteration consumes the same bytes!")
	fmt.Println()

	// Show the byte indices are not sequential
	fmt.Println("Byte indices from range (notice the gaps at multi-byte chars):")
	for i, r := range s {
		fmt.Printf("  byte offset %2d → rune %c  (uses %d bytes)\n", i, r, utf8.RuneLen(r))
	}

	// What happens with invalid UTF-8?
	fmt.Println("\nRange on invalid UTF-8:")
	invalid := "Hello\xfe\xffWorld" // \xfe\xff are not valid UTF-8
	for i, r := range invalid {
		fmt.Printf("  byte[%d] = %U (replacement char? %v)\n", i, r, r == utf8.RuneError)
	}
	// Go replaces invalid bytes with U+FFFD (replacement character).
	// Python would likely raise an error on decode, depending on the context.
}

// =============================================================================
// 5. CONVERTING BETWEEN string, []byte, AND []rune
// =============================================================================
//
// These conversions are EXPLICIT in Go (unlike Python, where encoding/decoding
// is always an explicit method call).

func demoConversions() {
	s := "Hello, 世界"

	fmt.Println("\n=== Conversions: string ↔ []byte ↔ []rune ===")
	fmt.Printf("Original: %q\n", s)

	// string → []byte  (like Python's s.encode('utf-8'))
	// This gives you the RAW UTF-8 bytes
	bytes := []byte(s)
	fmt.Printf("[]byte:   %v\n", bytes)
	fmt.Printf("As hex:   ")
	for _, b := range bytes {
		fmt.Printf("%02x ", b)
	}
	fmt.Println()

	// []byte → string  (like Python's b.decode('utf-8'))
	// This INTERPRETS the bytes as UTF-8
	backToString := string(bytes)
	fmt.Printf("back to string: %q\n", backToString)

	// string → []rune  (Python: [ord(c) for c in s])
	// This DECODES the UTF-8 into individual code points
	runes := []rune(s)
	fmt.Printf("[]rune:   %v\n", runes)
	fmt.Printf("As chars: ")
	for _, r := range runes {
		fmt.Printf("%c ", r)
	}
	fmt.Println()

	// []rune → string  (Python: ''.join(chr(r) for r in runes))
	// This ENCOODES the runes into UTF-8 bytes
	fromRunes := string(runes)
	fmt.Printf("back from runes: %q\n", fromRunes)

	fmt.Println("\nPython equivalents:")
	fmt.Println("  bytes(s, 'utf-8')    ↔ []byte(s)")
	fmt.Println("  b.decode('utf-8')    ↔ string(byteSlice)")
	fmt.Println("  [ord(c) for c in s]  ↔ []rune(s)")
	fmt.Println("  ''.join(chr(r)...)   ↔ string(runeSlice)")
	fmt.Println()
	fmt.Println("Key difference: Go conversions are TYPE CASTS (fast and direct).")
	fmt.Println("Python methods are function calls (encode/decode).")
	fmt.Println("Go's approach is more efficient but less discoverable.")
}

// =============================================================================
// 6. utf8 PACKAGE — Inspecting and Validating UTF-8
// =============================================================================
//
// Python has `codecs` and error handlers for encoding/decoding.
// Go has `unicode/utf8` for low-level UTF-8 operations.

func demoUTF8Package() {
	s := "Hello, 世界"

	fmt.Println("\n=== unicode/utf8 Package ===")
	fmt.Printf("s = %q\n", s)

	// Count runes (characters) — like Python's len()
	fmt.Printf("utf8.RuneCountInString(s) = %d (actual characters)\n",
		utf8.RuneCountInString(s))
	fmt.Printf("len(s)                      = %d (bytes)\n", len(s))
	// In Python: len("Hello, 世界") → 9 characters
	// In Go:     utf8.RuneCountInString("Hello, 世界") → 9
	//           len("Hello, 世界") → 13 bytes (7 ASCII + 3 + 3 for 世界)

	// Check if string is valid UTF-8
	fmt.Printf("utf8.ValidString(s) = %v\n", utf8.ValidString(s))
	fmt.Printf("utf8.ValidString(\"\\\\xff\") = %v\n", utf8.ValidString("\xff"))
	// Python: try: s.encode('utf-8') except UnicodeEncodeError: ...

	// Decode first rune from a string
	r, size := utf8.DecodeRuneInString(s)
	fmt.Printf("\nFirst rune: %c (takes %d bytes)\n", r, size)

	// Decode last rune
	r, size = utf8.DecodeLastRuneInString(s)
	fmt.Printf("Last rune: %c (takes %d bytes)\n", r, size)

	// How many bytes does a rune need?
	fmt.Printf("\nRune sizes:\n")
	fmt.Printf("  'A'  → %d byte(s)\n", utf8.RuneLen('A'))
	fmt.Printf("  '世' → %d byte(s)\n", utf8.RuneLen('世'))
	fmt.Printf("  '😀'  → %d byte(s)\n", utf8.RuneLen('😀'))
	// ASCII: 1 byte, Basic Multilingual Plane: 2-3 bytes, Emoji: 4 bytes
}

// =============================================================================
// 7. WHY UTF-8? — Go's Design Decision
// =============================================================================
//
// Go chose UTF-8 for string storage because:
//   1. Backward compatible with ASCII (every ASCII string IS valid UTF-8)
//   2. Self-synchronizing (you can always find character boundaries)
//   3. Compact for most text (especially English/ASCII)
//   4. No endianness issues (unlike UTF-16)
//
// Python 3 chose a flexible internal representation:
//   - ASCII-only strings use 1 byte per char (compact)
//   - Strings with non-ASCII use 2 or 4 bytes per char (fast indexing)
//   - This is an implementation detail — you never see it
//
// Trade-off:
//   Go:     s[i] is O(1) but gives a byte, not a character. Indexing a rune
//           requires decoding (O(n)). This favors byte-level operations.
//
//   Python: s[i] is O(1) and gives a character. Internal storage is either
//           1, 2, or 4 bytes per character depending on the max code point.
//           This favors character-level operations but uses more memory.
//
//   There is NO "right" answer. Each language optimized for different use
//   cases. Go optimizes for network/text protocols (lots of byte slicing).
//   Python optimizes for text processing (lots of character access).

// =============================================================================
// 8. STRING IMMUTABILITY AND SLICING
// =============================================================================
//
// Both Go and Python strings are IMMUTABLE.
//
// But slicing works differently:
//   Python: s[0:5] creates a COPY of the characters (new string)
//   Go:     s[0:5] creates a "view" into the original bytes (no copy)
//
// Go slices are references to the underlying string's memory — essentially
// free (just a pointer + length). Python always allocates a new string.

func demoStringSlicing() {
	s := "Hello, World"

	fmt.Println("\n=== String Slicing ===")
	fmt.Printf("s = %q\n", s)

	// Slice — Go creates a new string header pointing to the same memory
	sub := s[:5]
	fmt.Printf("s[:5] = %q\n", sub)
	// This is like Python's s[:5], BUT:
	//   Python: COPY (O(n) time + allocation)
	//   Go:     VIEW (O(1) — just a pointer + length, no allocation)

	fmt.Println("\n⚠️ WARNING: Slicing multi-byte strings in the middle of a rune")
	greeting := "Hello, 世界"
	badSlice := greeting[:9] // Slices into the middle of "世"!

	fmt.Printf("greeting[:9] on %q = ", greeting)
	fmt.Printf("%q ", badSlice)
	fmt.Printf("— might be invalid UTF-8: ValidString=%v\n", utf8.ValidString(badSlice))
	// Python would never let you do this because Python's index is a
	// character index, not a byte index. s[:7] in Python gives exactly
	// the first 7 characters, always at character boundaries.

	// Safe slicing with rune count:
	runes := []rune(greeting)
	fmt.Printf("Safe: first 7 runes as string = %q\n", string(runes[:7]))
	// But this allocates (converting to runes is O(n))
}

// =============================================================================
// 9. STRING BUILDING — strings.Builder vs + vs join
// =============================================================================

func demoStringBuilding() {
	fmt.Println("\n=== String Building ===")

	// In Python, you'd do:  ''.join(parts)  or  ' ' + ' '.join(...)

	// Go's modern approach: strings.Builder (efficient, like join)
	// We'll demonstrate conceptually:
	fmt.Println("Go: strings.Builder  (most efficient)")
	fmt.Println("    var b strings.Builder")
	fmt.Println("    b.WriteString(\"Hello\")")
	fmt.Println("    b.WriteRune('世')")
	fmt.Println("    result := b.String()")
	fmt.Println()
	fmt.Println("Python: ''.join([...])  (most efficient)")
	fmt.Println("    result = ''.join(['Hello', '世'])")
	fmt.Println()
	fmt.Println("Both avoid creating intermediate strings.")
	fmt.Println("Go:  s += \"x\" creates a new string (copy), just like Python.")
	fmt.Println("      In a loop, this is O(n²) in BOTH languages.")
}

// =============================================================================
// 10. COMMON MISTAKES — Python → Go Migration Guide
// =============================================================================

func demoCommonMistakes() {
	fmt.Println("\n=== Common Mistakes (Python → Go) ===")

	// Mistake 1: Using len() for character count
	s := "世界"
	fmt.Printf("Mistake #1: len(%q) = %d (BYTES, not characters!)\n", s, len(s))
	fmt.Printf("  Correct: utf8.RuneCountInString(%q) = %d\n", s, utf8.RuneCountInString(s))
	fmt.Println("  Python: len('世界') = 2 ✓")

	// Mistake 2: Indexing into a string for characters
	fmt.Printf("\nMistake #2: %q[0] = %d (BYTE, not character!)\n", s, s[0])
	fmt.Printf("  Correct: string([]rune(%q)[0]) = %q\n", s, string([]rune(s)[0]))
	fmt.Println("  Python: '世界'[0] = '世' ✓")

	// Mistake 3: Slicing across multi-byte boundary
	fmt.Printf("\nMistake #3: %q[:1] = %q (invalid UTF-8!)\n", s, s[:1])
	fmt.Println("  Python can't even express this — character indices are safe by design.")

	// Mistake 4: Range over string vs indexing
	text := "abc世"
	fmt.Println("\nMistake #4: For loop with index vs range")
	fmt.Println("  WRONG:")
	fmt.Print("   ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%c ", text[i]) // prints raw bytes as ASCII
	}
	fmt.Println("  — garbage for multi-byte chars")
	fmt.Println("  CORRECT (range):")
	fmt.Print("   ")
	for _, r := range text {
		fmt.Printf("%c ", r)
	}
	fmt.Println("  — correct runes")
}

// =============================================================================
// 11. RAW STRINGS — Go's r'''...''' and b'''...'''
// =============================================================================

func demoRawStrings() {
	fmt.Println("\n=== Raw Strings ===")

	// Go has ONLY double-quoted interpreted strings and backtick raw strings.
	// Python has single quotes, double quotes, triple quotes, and prefixes.

	// Backtick raw string — like Python's r'''...''' or r"..."
	raw := `This is a raw string in Go.
It can span multiple lines.
Escapes like \n are NOT processed.
"Quotes" inside are fine too.`

	fmt.Println("Go backtick raw string:")
	fmt.Println(raw)

	fmt.Println("\nPython equivalents:")
	fmt.Println("  \"\"\"...\"\"\"  → triple-quoted string")
	fmt.Println("  r\"...\"     → raw string (no escape processing)")
	fmt.Println("  b\"...\"     → bytes literal (like Go's []byte(...))")
	fmt.Println()
	fmt.Println("Go has no single-quoted strings (those are runes).")
	fmt.Println("Go has NO byte-string prefix — use []byte(\"s\") instead.")
}

// =============================================================================
// 12. COMPLETE COMPARISON TABLE
// =============================================================================
//
// ┌───────────────────────────────┬────────────────────────────────────┬──────────────────────────────────────┐
// │          Feature              │            Go                     │              Python                   │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ String nature                 │ Immutable byte sequence,           │ Immutable code point sequence,        │
// │                               │ UTF-8 encoded                     │ flexible internal encoding            │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ len(s)                        │ Number of BYTES                    │ Number of CHARACTERS                  │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ s[i]                          │ Raw BYTE at index i                │ Character at index i (1-char string)  │
// │                               │ (byte, not char)                  │                                       │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Character type                │ rune (int32) — like ord result     │ 1-character string (like Python's str)│
// │                               │                                    │ ord() gives the int code point        │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ ord() / chr()                 │ '世' → rune value (compile-time)  │ ord('世') → 19990                     │
// │                               │ string('世') → "世" (at runtime) │ chr(19990) → "世"                     │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Rune literal                  │ '世'  (single quotes, like char)  │ No direct equivalent.                 │
// │                               │                                    │ ord('世') is a function call,          │
// │                               │                                    │ not a literal.                        │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Iteration                     │ for i, r := range s               │ for i, ch in enumerate(s):            │
// │                               │ i = BYTE index, r = rune          │ i = CHAR index, ch = 1-char string    │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Character count               │ utf8.RuneCountInString(s)         │ len(s)                                │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Encode to bytes               │ []byte(s)  (type conversion)      │ s.encode('utf-8')                     │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Decode from bytes             │ string(b)  (type conversion)      │ b.decode('utf-8')                     │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ String building               │ strings.Builder                   │ ''.join(parts)                        │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ String slicing                │ s[i:j] = VIEW (no copy, O(1))    │ s[i:j] = COPY (allocates, O(n))       │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Immutability                  │ Yes                               │ Yes                                   │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Raw strings                   │ Backtick: `raw\n` (no escapes)   │ r"raw" or r'''raw'''                  │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Unicode normalization         │ Not built-in (use golang.org/x/   │ unicodedata.normalize()               │
// │                               │ text/unicode/norm)               │                                       │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Invalid UTF-8                 │ Replaced with U+FFFD on range     │ Errors on decode by default           │
// │                               │ or use utf8.ValidString()         │ or .decode('utf-8', 'replace')        │
// ├───────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ String comparison             │ == compares BYTES (byte-wise)     │ == compares CODE POINTS (works        │
// │                               │ (works correctly for UTF-8)      │ correctly by definition)              │
// └───────────────────────────────┴────────────────────────────────────┴──────────────────────────────────────┘

// =============================================================================
// 13. PRACTICAL EXAMPLE — Processing a Multilingual String
// =============================================================================

func demoPracticalExample() {
	fmt.Println("\n=== Practical Example: Processing a Multilingual String ===")

	// A string with mixed English, Japanese, Chinese, and emoji
	text := "Go is 楽しい! 🎉 世界"

	fmt.Printf("Input: %q\n\n", text)

	fmt.Printf("Byte length (len):       %3d bytes\n", len(text))
	fmt.Printf("Rune count (characters): %3d characters\n\n", utf8.RuneCountInString(text))

	fmt.Println("Breakdown (byte offset, rune, bytes used):")
	fmt.Println("─────────────────────────────────────────────")
	var byteOffset int
	for _, r := range text {
		size := utf8.RuneLen(r)
		fmt.Printf("  offset %2d: %c (U+%04X, %d byte(s))\n", byteOffset, r, r, size)
		byteOffset += size
	}

	// Show raw UTF-8 bytes
	fmt.Println("\nRaw UTF-8 bytes (hex):")
	fmt.Printf("  ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%02x ", text[i])
	}
	fmt.Println()

	// Python equivalent:
	fmt.Println("\nPython equivalent:")
	fmt.Println(`  text = "Go is 楽しい! 🎉 世界"`)
	fmt.Println(`  print(len(text))               # → 17 characters`)
	fmt.Println(`  print(len(text.encode()))       # → 31 bytes  (UTF-8)`)
	fmt.Println(`  for i, ch in enumerate(text):   # ch = 1-char string`)
	fmt.Println(`      print(i, ch, ord(ch))`)
	fmt.Println()
	fmt.Println("  # Notice: Unicode escapes match!")
	fmt.Println(`  '世' == '世'  # → True`)
}

// =============================================================================
// 14. STRING HEADER — A Go Internals Peek
// =============================================================================
//
// In Go, a string is internally a struct:
//
//   type string struct {
//       data *byte  // pointer to the underlying bytes
//       len  int    // number of bytes
//   }
//
// This is why slicing is O(1) — it just creates a new header pointing
// to the same data pointer, with a new length.
//
// Python's string internal representation is more complex (flexible
// representation: 1-byte, 2-byte, or 4-byte per character) and is
// an implementation detail you're not supposed to touch.

func demoStringInternals() {
	fmt.Println("\n=== String Internals ===")

	s := "Hello"
	// Get the string header using unsafe (for educational purposes only!)
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))

	fmt.Printf("String: %q\n", s)
	fmt.Printf("StringHeader.Data = %v (pointer to bytes)\n", hdr.Data)
	fmt.Printf("StringHeader.Len  = %d (bytes)\n", hdr.Len)
	fmt.Println()
	fmt.Println("This is why slicing is cheap: no copy, just a new header.")
	fmt.Println("Python strings are NOT sliced this way — each slice copies.")
}

// =============================================================================
// SUMMARY: Python → Go String/Rune Translation Cheat Sheet
// =============================================================================
//
// Python:                          Go:
// ───────                          ──
// s = "Hello, 世界"                 s := "Hello, 世界"
//
// len(s)                           utf8.RuneCountInString(s)
//                                  // or: len([]rune(s))
//
// n = len(s.encode('utf-8'))       len(s)
//
// s[0]                             string([]rune(s)[0])
//                                  // s[0] gives the byte, not char!
//
// for i, ch in enumerate(s):       for i, r := range s {
//     print(i, ch, ord(ch))             // i = BYTE offset, r = rune
//                                  }
//
// s.encode('utf-8')                []byte(s)
//
// b.decode('utf-8')                string(b)
//
// [ord(c) for c in s]              []rune(s)
//
// ''.join(chr(r) for r in runes)  string(runes)
//                                  // Go: string([]rune) is a direct cast
//
// s[:5] (copy, O(n))              s[:5] (view, O(1))
//
// ─── Key takeaways ───
//
// 1. Go strings are BYTE SEQUENCES (UTF-8). Python strings are CODE POINT
//    SEQUENCES. This is the single most important difference.
//
// 2. Go's len(s) counts bytes; Python's len(s) counts characters.
//    Use utf8.RuneCountInString(s) in Go for character count.
//
// 3. Go's s[i] gives a raw byte; Python's s[i] gives a character string.
//    Use string([]rune(s)[i]) in Go to get the i-th character.
//
// 4. Go's range loop automatically decodes UTF-8 runes. Use it — never
//    iterate with a byte index for character processing.
//
// 5. Go's string → []rune conversion is like [ord(c) for c in s] in Python.
//    Go's []rune → string conversion is like ''.join(chr(r) for r in runes).
//
// 6. Go's string slicing creates a VIEW (no copy, O(1)). Python creates a
//    COPY (O(n)). But Go slices can cut into the middle of a multi-byte
//    character — Python never can.
//
// 7. Go's raw strings use backticks (`...`). Python uses r"..." or r"""...""".
//    Go has no single-quote string, no triple-quote, and no byte literal prefix.
//
// 8. Go has NO built-in Unicode normalization. Python has unicodedata.

func main() {
	demoBasicString()
	demoUnicodeProblem()
	demoRunes()
	demoRangeLoop()
	demoConversions()
	demoUTF8Package()
	demoStringSlicing()
	demoStringBuilding()
	demoCommonMistakes()
	demoRawStrings()
	demoPracticalExample()
	demoStringInternals()
}
