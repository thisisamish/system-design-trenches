# System Design Trenches

A multi-month proof-of-learning, incident archive, and forkable curriculum. Learn Go and distributed systems by **breaking things under pressure**, not by reading about them.

> Constrain -> Crash -> Diagnose -> Re-Architect. See [PHILOSOPHY.md](PHILOSOPHY.md).

## Scoreboard

| Trench | Arena | Baseline RPS / p99 | Hardened RPS / p99 | Crashes Survived | Post-Mortem |
|---|---|---|---|---|---|
| 0 | Go Gauntlet | n/a | n/a | 0 | - |
| 1 | Connection & I/O Wall | - | - | 0 | - |
| 2 | Race Conditions | - | - | 0 | - |
| 3 | Cascading Failures | - | - | 0 | - |
| 4 | Storage Ceiling | - | - | 0 | - |
| 5 | Network Chaos | - | - | 0 | - |

Roadmap: [CURRICULUM.md](CURRICULUM.md). Incident reports: [LEARNING_LOGS/](LEARNING_LOGS/).

## How to Fork & Fight

1. Clone the repo.
2. Pick a trench in `trenches/`. Each is self-contained.
3. Run its attack (`make attack`, or `.\attack.ps1` on Windows without make). Watch it die.
4. Diagnose with logs, `-race`, `pprof`, `docker stats`. Fix it yourself.
5. Write a post-mortem from `LEARNING_LOGS/TEMPLATE_POSTMORTEM.md`.

### Reset to Naive

Each trench's broken baseline is preserved in a Git tag:

```sh
git checkout trench-0-naive   # broken baseline
git checkout trench-0-survived  # one engineer's fix (spoilers)
```

(Tags are created once this directory is a git repo; run `git init` and commit first.)

## Dev Environment

Use the dev container; it works the same on Windows, macOS, and Linux.

1. Install Docker (Docker Desktop with WSL2 integration on Windows) and VS Code with the Dev Containers extension.
2. **Windows: clone inside WSL** (`~/code/...`), not under `/mnt/c`. It is much faster.
3. Open the folder in VS Code and choose **Reopen in Container**.

You get Go 1.22, gcc (so `go test -race` works), `make`, graphviz for `pprof`, and nested Docker for compose-based trenches. Then run `make attack` from a trench directory.

Without the container you need Go 1.22+, gcc, make, and Docker. `attack.ps1` is a fallback for native Windows.
