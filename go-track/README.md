# The Go Track

Learn Go from zero, one small concept at a time. Every lesson follows the same cycle:

1. **Learn**: read the lesson's `README.md`. It explains the concept with comparisons to Java, C and JS, and shows small examples.
2. **Build**: fill in the `TODO`s in the lesson's starter file.
3. **Test**: run `make test` until everything passes.
4. **Break**: run `make break`. These tests attack your code with nasty inputs (edge cases, huge numbers, unicode, slow paths). Fix what breaks.

When both pass, tell Claude `lesson NN done`. Claude runs the tests itself, then asks you a few questions to check that the ideas stuck. After the quiz you write a short recap and move on.

## Lessons

| # | Lesson | What you learn |
|---|---|---|
| 01 | [Hello, Go](01-hello/) | running a program, packages, imports, functions, tests |
| 02 | [Values and functions](02-values-and-functions/) | variables, types, zero values, conversions, multiple returns |
| 03 | [Control flow](03-control-flow/) | `if`, `for`, `switch`, looping over strings |
| 04–15 | coming next | see [CURRICULUM.md](../CURRICULUM.md) |

Start with [01-hello](01-hello/). Don't skip ahead: each lesson assumes the ones before it.

## Commands

Run these from inside a lesson folder, e.g. `cd go-track/01-hello`.

| Command | What it does |
|---|---|
| `make run` | runs the lesson's `main.go` so you can see your code's output |
| `make test` | runs the normal tests |
| `make break` | runs the attack tests (from lesson 02 on) |

No `make` (native Windows)? Use `.\make.ps1 run`, `.\make.ps1 test`, `.\make.ps1 break`.

In VS Code you can also click the small **run test** link that appears above each `func Test...` in a `_test.go` file.

## Reading test output

All output appears in the terminal where you ran the command.

```text
=== RUN   TestGreeting
    greet_test.go:9: Greeting("Gopher") = "", want "Hello, Gopher!"
--- FAIL: TestGreeting (0.00s)
```

- `=== RUN` means the test started.
- The indented line is the test's message: file, line number, then what your code returned (`""`) and what it should have returned (`"Hello, Gopher!"`).
- `--- PASS` or `--- FAIL` is the verdict.
- `ok` at the very end means every test passed. `FAIL` means at least one did not.

## Stuck?

Ask Claude anything: syntax, an error message, why a test fails, how something compares to Java. Claude explains and gives hints, but won't write the exercise for you. That part is yours.
