# Lesson 01: Hello, Go

**You'll learn:** how to run a Go program, what packages and imports are, how to write a function, and how Go tests work.
**You'll build:** three small string functions.
**Time:** about an hour.

---

## Learn

### 1. Your first program

Open [main.go](main.go):

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello from Go!")
	fmt.Println(Greeting("Gopher"))
	fmt.Println(Shout("go"))
	fmt.Println(Introduce("Gopher", 13))
}
```

Line by line:

- `package main`: every Go file starts by saying which **package** it belongs to. A package is a folder of `.go` files that are compiled together. The package named `main` is special: it builds into a program you can run.
- `import "fmt"`: pulls in the standard library's `fmt` package (short for "format"). It handles printing and string formatting.
- `func main() { ... }`: the entry point, where the program starts.
- `fmt.Println(...)`: calls the `Println` function from the `fmt` package. It prints its arguments, then a newline.

| | Entry point |
|---|---|
| Java | `public static void main(String[] args)` inside a class |
| C | `int main(int argc, char **argv)` |
| JS / Node | none; the file runs top to bottom |
| **Go** | `func main()` in `package main`. No class, no arguments, no return value |

There are no semicolons. Go adds them for you at the end of each line. One consequence: an opening `{` must be on the **same line** as `func main()`, not the next one.

### 2. Run it

From this folder:

```sh
make run        # or: go run .
```

`go run .` compiles every `.go` file in the current folder (`.`) into a temporary program and runs it. Go is compiled like C, not interpreted like JS. The compiler is just fast enough that it feels instant.

You should see `Hello from Go!` followed by **three empty lines**. Those are the return values of `Greeting`, `Shout` and `Introduce`, which currently return an empty string `""`. Fixing that is your job.

> Related: `go build` writes the compiled program to a file instead of running it. Try `go build -o hello . && ./hello`. The result is a single binary with no runtime to install, unlike Java's JVM or Node.

### 3. Functions

Open [greet.go](greet.go):

```go
func Greeting(name string) string {
	return ""
}
```

Read it as: function `Greeting` takes a parameter `name` of type `string` and returns a `string`.

**Types come after names.** That is backwards from Java and C, and the same order as TypeScript:

| Language | Same function |
|---|---|
| Java | `static String greeting(String name)` |
| C | `char *greeting(char *name)` |
| TypeScript | `function greeting(name: string): string` |
| **Go** | `func Greeting(name string) string` |

Two files, one package: `main.go` calls `Greeting`, which lives in `greet.go`, with no import between them. Every file in the same folder and package sees everything in the others.

### 4. Strings

- String literals use double quotes: `"hello"`. Backticks make a raw multi-line string: `` `C:\no\escapes` ``.
- `+` joins strings: `"Hello, " + name`.
- Go **never** converts types for you. `"age: " + 13` is a compile error, while Java and JS would quietly turn `13` into `"13"`. To mix a number into a string, use `fmt.Sprintf`:

```go
s := fmt.Sprintf("%s has %d legs", "cat", 4) // "cat has 4 legs"
```

`Sprintf` works like C's `sprintf` (or Java's `String.format`) but returns the string. `%s` is a string, `%d` an integer, and `%v` means "any value, default format", which is handy when you're unsure.

`s := ...` creates a new variable `s` and works out its type from the right-hand side. Lesson 02 covers this properly.

### 5. Imports and the standard library

To use another package, import it and prefix calls with its name:

```go
import (
	"fmt"
	"strings"
)

func example() string {
	return strings.Repeat("ab", 3) // "ababab"
}
```

Documentation for every standard package is at [pkg.go.dev/std](https://pkg.go.dev/std). For example, [pkg.go.dev/strings](https://pkg.go.dev/strings) lists `ToUpper`, `Repeat`, `Contains` and the rest. You can also run `go doc strings.ToUpper` in the terminal.

**Why `fmt.Println` and not `fmt.println`?** In Go, a name that starts with a **capital letter** is visible outside its package, like Java's `public`. A lowercase name is private to its package. That's why every function you call from a standard package is capitalized.

Two strict rules that will bite you:

- **Importing a package you don't use is a compile error.** Not a warning, an error.
- **Declaring a local variable you don't use is a compile error.**

Go does this on purpose, so dead code can't pile up. It's annoying for the first week, and then you stop noticing.

### 6. Tests

Open [greet_test.go](greet_test.go):

```go
func TestGreeting(t *testing.T) {
	got := Greeting("Gopher")
	want := "Hello, Gopher!"
	if got != want {
		t.Errorf("Greeting(%q) = %q, want %q", "Gopher", got, want)
	}
}
```

- Test files end in `_test.go`. `go run` ignores them; `go test` runs them.
- A test is any function named `TestXxx` that takes `t *testing.T` (a pointer to a test helper, which lesson 04 explains). There's no JUnit annotation or Jest `describe`; the name alone is enough.
- There's no `assertEquals`. You compare with a plain `if`, and call `t.Errorf` to report a failure. `t.Errorf` uses the same `%` verbs as `Sprintf`. `%q` prints a value in quotes, so an empty string shows up as `""` instead of nothing.

Run them:

```sh
make test       # or: go test -v .
```

Right now every test fails:

```text
=== RUN   TestGreeting
    greet_test.go:9: Greeting("Gopher") = "", want "Hello, Gopher!"
--- FAIL: TestGreeting (0.00s)
...
FAIL
make: *** [Makefile:9: test] Error 1
```

`greet_test.go:9` is the file and line that reported the failure. Then you get what your code returned and what it should have returned. The final `make: *** ... Error 1` line only means "a test failed"; it's not a separate problem.

To run a single test, pass its name: `go test -v -run TestShout .`

---

## Build

Fill in the three `TODO`s in [greet.go](greet.go). The comment above each function says exactly what it must return.

1. **`Greeting`**: join strings with `+`.
2. **`Shout`**: import `strings` and use `strings.ToUpper`.
3. **`Introduce`**: first try `name + " is " + age + " years old"` and read the compiler error. Then fix it with `fmt.Sprintf`. Remember to import `fmt` in `greet.go`; imports belong to a single file.

After each one, run `make run` to see the output, then `make test`.

## Test

```sh
make test
```

You're done with this step when you see `--- PASS` three times and `ok  gotrack/01-hello` at the end.

## Break

Lesson 01 has no attack tests. Instead, **break the code yourself on purpose** and get to know Go's error messages while the stakes are zero. Make each change, run `make test`, read the error, then undo the change.

1. Delete the closing `"` of a string in `greet.go`.
2. Add `x := 5` as a line inside `Greeting` and never use `x`.
3. Add `"os"` to an import block without using it.
4. Move `main`'s opening `{` to its own line.
5. Change `Greeting`'s return to `return 42`.
6. Rename `Greeting` to `greeting` in `greet.go` only.
7. Change the expected `want` in a test to something wrong, and read how a failing test reports itself.

For each one, notice **which file and line** the error points to, and whether the message tells you how to fix it. You'll meet every one of these errors again by accident, so it pays to recognize them now.

---

## Done?

- [ ] `make test` passes.
- [ ] `make run` prints the three lines.
- [ ] You did all 7 deliberate breaks.

Tell Claude **`lesson 01 done`**. You'll get a short quiz (3–5 questions, no trick questions), then write a recap in `LEARNING_LOGS/go-track/01-hello.md` using [the template](../../LEARNING_LOGS/TEMPLATE_RECAP.md).
