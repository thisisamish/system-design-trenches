# Lesson 02: Values and Functions

**You'll learn:** declaring variables, Go's basic types, zero values, constants, type conversions, and functions that return more than one value.
**You'll build:** five small number functions: a temperature converter, a money splitter, a percentage, a time converter, and an average.
**Time:** 1–2 hours.

---

## Learn

### 1. Declaring variables

Go has two ways to declare a variable:

```go
var count int = 10   // full form: name, type, value
var count = 10       // type inferred from the value
var count int        // no value: count is 0 (see "zero values" below)

count := 10          // short form: declare + assign, type inferred
```

The short form `:=` is what you'll use most. It only works **inside functions**, and only for **new** variables. To change an existing variable, use plain `=`:

```go
x := 1
x = 2    // fine: changes x
x := 3   // compile error: no new variables on left side of :=
```

| | Declare an int |
|---|---|
| Java | `int count = 10;` or `var count = 10;` |
| C | `int count = 10;` |
| JS / TS | `let count = 10;` / `let count: number = 10;` |
| **Go** | `count := 10` or `var count int = 10` |

Inferred types are still **static**: once `count` is an `int`, it stays an `int`, as in Java and TS and unlike JS.

You can declare several at once: `a, b := 1, "two"`.

### 2. Basic types

| Go | What it is | Closest Java / C |
|---|---|---|
| `int` | signed integer, 64 bits on any modern machine | Java `long`, C `long` (on 64-bit Linux) |
| `int8`, `int16`, `int32`, `int64` | signed integer of exactly that many bits | Java `byte`/`short`/`int`/`long`, C `int32_t` etc. |
| `uint`, `uint8` … `uint64` | unsigned integer | C `unsigned`. Java has no unsigned types |
| `float64` | 64-bit floating point | `double` |
| `float32` | 32-bit floating point | `float` |
| `bool` | `true` / `false` | `boolean` / `bool` |
| `string` | immutable text (UTF-8 bytes) | Java `String`. Not a C `char*` |
| `byte` | another name for `uint8` | C `unsigned char` |
| `rune` | another name for `int32`; one Unicode character | Java `int` code point |

Use `int` for integers and `float64` for decimals unless you have a reason not to.

**Integers have fixed sizes and wrap around silently.** An `int8` holds −128 to 127. Add 1 to 127 and you get −128: no error, no exception, no panic. Java's `int` behaves the same way. JS numbers don't, because every JS number is a float. The largest values are available as constants: `math.MaxInt8`, `math.MaxInt32`, `math.MaxInt`, and so on.

### 3. Zero values

Every variable in Go always has a value. If you don't give it one, it gets its type's **zero value**:

| Type | Zero value |
|---|---|
| any number | `0` |
| `bool` | `false` |
| `string` | `""` (empty string) |

```go
var n int       // 0
var ok bool     // false
var name string // ""
```

There's no "uninitialized garbage" like C locals, no `undefined` like JS, and no compiler complaint like Java's "might not have been initialized". This matters a lot later: Go code often relies on zero values being useful.

### 4. Constants and integer division

```go
const Pi = 3.14159
const SecondsPerDay = 24 * 60 * 60
```

Constants are computed by the compiler. A number written straight into code, like `9` or `5`, is an **untyped constant**: it takes on whatever type it's combined with.

Integer division throws away the remainder, as in Java and C:

```go
7 / 2      // 3   (both are integers)
7.0 / 2    // 3.5 (one is a float)
7 % 2      // 1   (remainder)
```

So in an expression like `x * (9 / 5)`, the `9 / 5` part is worked out first, **as integers**, and gives `1`. Keep that in mind when you write `CelsiusToFahrenheit`.

### 5. Conversions are always explicit

Go never converts between number types for you, not even from `int` to `float64`, which Java does automatically:

```go
a := 10      // int
b := 2.5     // float64
c := a * b   // compile error: mismatched types int and float64

c := float64(a) * b   // 25.0: you convert explicitly with Type(value)
n := int(b)           // 2: converting float to int drops the decimals
```

Conversion syntax is `Type(value)`, like a function call. Compare Java's `(double) a` cast.

**Where you put the conversion matters.** `float64(a / b)` divides first, as integers, then converts the already-truncated result. `float64(a) / float64(b)` converts first, then divides as floats. These give different answers. The same applies to overflow: the arithmetic happens in whatever type the operands have **at that moment**.

### 6. Functions, in more depth

```go
func add(a int, b int) int {
	return a + b
}

func add(a, b int) int {   // same thing: consecutive params of one type share it
	return a + b
}
```

**Functions can return several values.** This is one of Go's most-used features:

```go
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

q, r := divmod(17, 5)   // q = 3, r = 2
```

Java would need a small class or an array for this, C would use out-pointers, and JS would return an array or object. In Go you just list the return types in parentheses.

You can also **name** the return values. They start at their zero value, and they document what each one means:

```go
func divmod(a, b int) (quotient int, remainder int) {
	return a / b, a % b
}
```

If you only want some of the returned values, throw the rest away with `_` (the **blank identifier**):

```go
q, _ := divmod(17, 5)   // ignore the remainder
```

You need `_` because an unused variable is a compile error, as you saw in lesson 01.

### 7. Reading the tests in this lesson

The tests here use a common Go pattern called a **table-driven test**:

```go
tests := []struct {
	c, want float64
}{
	{100, 212},
	{0, 32},
}
for _, tt := range tests {
	if got := CelsiusToFahrenheit(tt.c); got != tt.want { ... }
}
```

You don't need to write this yet; lessons 04 and 05 cover structs and slices. To read it: `tests` is a list of rows, each row has fields `c` and `want`, and the `for` loop checks your function against every row. When a row fails, the message tells you which input it was.

Floats are compared with a small tolerance (`closeTo`), because `0.1 + 0.2` is not exactly `0.3` in any language.

---

## Build

Fill in the five `TODO`s in [values.go](values.go). Each function's comment gives its exact behavior and an example.

1. **`CelsiusToFahrenheit`**: one line of arithmetic. Watch the integer division from section 4.
2. **`SplitCents`**: returns two values. Use `/` and `%`.
3. **`PercentOf`**: takes `int`s, returns a `float64`. Think about where the conversion goes (section 5).
4. **`MillisFromSeconds`**: takes an `int32`, returns an `int64`.
5. **`Average`**: the average of two numbers, rounded down.

Use `make run` to see your functions in action, and `make test` to check them.

## Test

```sh
make test
```

When all five tests pass, move on. If one fails, read the message: it shows the input, what you returned, and what was expected.

## Break

```sh
make break
```

The attack tests in [break_test.go](break_test.go) feed your functions much bigger numbers: a 30-day timeout in seconds, and averages of the largest possible `int`. These come up in real systems (timeouts, file sizes, IDs, timestamps), and code that passes the normal tests often breaks here.

When a break test fails:

1. Read the failure. Is your answer a little off, or wildly wrong (negative, for example)?
2. Work out what type each piece of your expression has **at each step** of the calculation, and what the largest value that type can hold is.
3. Fix it without changing the function signatures.

<details>
<summary>Hint for MillisFromSeconds (open only if stuck)</summary>

Your function returns an `int64`, but which type is the multiplication actually done in? Section 5 has the answer.

</details>

<details>
<summary>Hint for Average (open only if stuck)</summary>

`a + b` can be bigger than the largest `int` even though `a`, `b` and their average all fit. Converting to `float64` doesn't help either: a `float64` has only 53 bits for digits, so very large integers lose their last digits. Can you reach the average without ever computing `a + b`?

</details>

One break test probably passes already. Work out why that function can't overflow when the others can.

---

## Done?

- [ ] `make test` passes.
- [ ] `make break` passes.
- [ ] You can explain *why* each failing break test failed, not just what you changed.

Tell Claude **`lesson 02 done`** for the quiz. Then write `LEARNING_LOGS/go-track/02-values-and-functions.md` from [the template](../../LEARNING_LOGS/TEMPLATE_RECAP.md).
