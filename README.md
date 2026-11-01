# System Design Trenches

A multi-month learning log and forkable curriculum: learn Go from zero, then distributed-systems design, by **building things and then breaking them**.

> Learn → Build → Test → Break. See [PHILOSOPHY.md](PHILOSOPHY.md).

## Start here

```sh
cd go-track/01-hello
make run
make test
```

Then open [go-track/01-hello/README.md](go-track/01-hello/README.md) and follow it. How lessons work: [go-track/README.md](go-track/README.md).

## Progress

### Phase 1: Go Track

| Lesson | Topic | Status | Recap |
|---|---|---|---|
| 01 | Hello, Go | not started | - |
| 02 | Values and functions | not started | - |
| 03 | Control flow | not started | - |
| 04–15 | see [CURRICULUM.md](CURRICULUM.md) | not written yet | - |

### Phase 2: System-Design Trenches

| Trench | Arena | Baseline RPS / p99 | Hardened RPS / p99 | Post-Mortem |
|---|---|---|---|---|
| 1 | Connection & I/O Wall | - | - | - |
| 2 | Race Conditions | - | - | - |
| 3 | Cascading Failures | - | - | - |
| 4 | Storage Ceiling | - | - | - |
| 5 | Network Chaos | - | - | - |

Roadmap: [CURRICULUM.md](CURRICULUM.md). Recaps and post-mortems: [LEARNING_LOGS/](LEARNING_LOGS/).

## How to Fork & Learn

1. Clone the repo and open it in the dev container (below).
2. Start at `go-track/01-hello` and work through the lessons in order.
3. For each lesson: read the README, fill in the TODOs, `make test`, `make break`, fix what breaks.
4. Close each lesson with a short recap from `LEARNING_LOGS/TEMPLATE_RECAP.md`.

### Git tags

- `go-NN-done`: one learner's state after finishing lesson NN (contains solutions, so spoilers).
- `original-gauntlet`: the repo's first, much harsher design, kept for reference.

## Dev Environment

Use the dev container; it works the same on Windows, macOS, and Linux.

1. Install Docker (Docker Desktop with WSL2 integration on Windows) and VS Code with the Dev Containers extension.
2. **Windows: clone inside WSL** (`~/code/...`), not under `/mnt/c`. It is much faster.
3. Open the folder in VS Code and choose **Reopen in Container**.

You get Go 1.22, gcc (so `go test -race` works), `make`, graphviz for `pprof`, and nested Docker for the later trenches.

Without the container you need Go 1.22+ (plus gcc, make, and Docker for later lessons). Every `Makefile` has a `make.ps1` beside it for native Windows.
