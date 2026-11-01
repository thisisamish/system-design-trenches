# Lesson 03: Control Flow

**You'll learn:** `if`, Go's single loop keyword `for`, `switch`, and how to loop over a string correctly.
**You'll build:** a letter-grade function, a sum, a string reverser, and a prime checker.
**Time:** 1–2 hours.

---

## Learn

### 1. `if`

```go
if x > 10 {
	fmt.Println("big")
} else if x > 5 {
	fmt.Println("medium")
} else {
	fmt.Println("small")
}
```

Compared to Java, C and JS:

- **No parentheses** around the condition. `if (x > 10)` still compiles, but `gofmt` removes the parentheses.
- **Braces are required**, even for one line. There's no `if (x) doThing();` shorthand.
- **The condition must be a `bool`.** There's no "truthiness": `if count { }` and `if name { }` are compile errors. Write `if count != 0` and `if name != ""`. This rules out a whole family of C and JS bugs.

Operators are the usual ones: `==`, `!=`, `<`, `<=`, `>`, `>=`, `&&`, `||`, `!`. There's no `===`, because Go never converts types, so `==` is always strict.

**`if` with a short statement.** You can run one statement before the condition. Variables declared there exist only inside the `if`/`else`:

```go
if n := len(name); n > 10 {
	fmt.Println("long name:", n)
}
// n doesn't exist here
```

You'll see this pattern everywhere once errors arrive in lesson 07.

### 2. `for` is the only loop

Go has no `while` and no `do`. `for` does every job:

```go
// Classic three-part loop, like C/Java (no parentheses)
for i := 0; i < 5; i++ {
	fmt.Println(i)
}

// "while" loop: just a condition
for n > 1 {
	n = n / 2
}

// Infinite loop: leave a way out with break
for {
	if done() {
		break
	}
}
```

`break` exits the loop and `continue` skips to the next round, same as in Java/C/JS.

`i++` is a **statement**, not an expression: `x := i++` doesn't compile, and there's no `++i`. `+=`, `-=`, `*=` work as usual.

### 3. `for range`

`range` loops over a collection and gives you an index and a value each round:

```go
for i, ch := range "abc" {
	fmt.Println(i, ch)   // 0 97, 1 98, 2 99
}
```

Use `_` to ignore the part you don't need: `for _, ch := range s`. Go 1.22 can also range over an integer: `for i := range 5` gives 0 to 4.

You'll range over slices and maps in lessons 05 and 06. Strings first, because they have a twist.

### 4. Strings are bytes, characters are runes

A Go `string` is a sequence of **bytes** in UTF-8 encoding. English letters take 1 byte each, but most other characters take more:

| Character | Bytes in UTF-8 |
|---|---|
| `a` | 1 |
| `é` | 2 |
| `♥` | 3 |
| `语` | 3 |

That gives two ways to walk through a string, and they behave differently:

```go
s := "héllo"

len(s)           // 6, not 5: len counts bytes

for i := 0; i < len(s); i++ {
	s[i]         // a byte (uint8): one piece of a character, not always a whole one
}

for i, r := range s {
	r            // a rune (int32): one whole character
	i            // the byte position where that character starts: 0, 1, 3, 4, 5
}
```

Indexing with `s[i]` gives you **bytes**. Ranging with `for ... range s` decodes UTF-8 and gives you **runes**. For plain English text they look identical. They only differ once real-world text shows up.

To turn a rune back into a string, use `string(r)`. Character literals use single quotes: `'a'` is a rune, `"a"` is a string.

| | One character | Length |
|---|---|---|
| Java | `char` (UTF-16, 2 bytes; emoji need two) | `s.length()` counts UTF-16 units |
| C | `char` (1 byte, no idea about Unicode) | `strlen` counts bytes |
| JS | string of length 1 (UTF-16 again) | `s.length` counts UTF-16 units |
| **Go** | `rune` (a full Unicode code point) | `len(s)` counts **bytes**; `utf8.RuneCountInString(s)` counts characters |

Strings are **immutable**, as in Java and JS: you can't do `s[0] = 'H'`. You build new strings instead, for example with `+`.

### 5. `switch`

```go
switch day {
case "Sat", "Sun":
	fmt.Println("weekend")
case "Fri":
	fmt.Println("almost")
default:
	fmt.Println("weekday")
}
```

Compared to Java/C/JS:

- **No fall-through, so no `break` needed.** Only the matching case runs. (If you really want fall-through, the `fallthrough` keyword exists, but you'll almost never use it.)
- One case can match **several values**: `case "Sat", "Sun":`.
- Cases can be any type that supports `==`, not just ints and enums.

**A switch with no value** checks each case as a `bool`, top to bottom, and runs the first one that's true. It's a cleaner `if / else if` chain:

```go
switch {
case temp < 0:
	fmt.Println("freezing")
case temp < 20:
	fmt.Println("cool")
default:
	fmt.Println("warm")
}
```

Order matters here: the first true case wins.

---

## Build

Fill in the four `TODO`s in [flow.go](flow.go). Read each doc comment carefully: it is the full spec.

1. **`Grade`**: use a switch. Check the doc comment for exactly where each grade starts and ends.
2. **`SumTo`**: use a `for` loop and a running total.
3. **`Reverse`**: build a new string with the characters in reverse order.
4. **`IsPrime`**: try dividing `n` by smaller numbers. `%` gives the remainder; a remainder of 0 means "divides evenly".

`make run` shows your functions' output; `make test` checks them.

## Test

```sh
make test
```

## Break

```sh
make break
```

The attack tests in [break_test.go](break_test.go) hit four weak spots that real code fails on all the time:

- **Boundaries**: exactly 90, exactly 60, 0, 100, just outside the valid range.
- **Nothing**: zero and negative inputs.
- **Real-world text**: accented letters, Chinese, symbols.
- **Big inputs**: one prime check on a 10-digit number, with a 1-second limit.

For each failure, read the input that broke your code, then work through what your code does with that exact input, line by line.

<details>
<summary>Hint for Reverse (open only if stuck)</summary>

Look at what the broken output looks like, then reread section 4. Which kind of loop hands you whole characters?

</details>

<details>
<summary>Hint for IsPrime speed (open only if stuck)</summary>

If `n` has a divisor bigger than its square root, it must also have one smaller than its square root (they come in pairs: 36 = 4 × 9). So how far do you actually need to check? You can write the loop condition without importing anything: compare `i * i` with `n`.

</details>

---

## Done?

- [ ] `make test` passes.
- [ ] `make break` passes.
- [ ] You can explain what `len("héllo")` returns and why.

Tell Claude **`lesson 03 done`** for the quiz. Then write `LEARNING_LOGS/go-track/03-control-flow.md` from [the template](../../LEARNING_LOGS/TEMPLATE_RECAP.md).
