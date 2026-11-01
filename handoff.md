# HANDOFF.MD — System Design & Go "Trench Warfare" Blueprint

## 1. User Profile & Learning Philosophy

### Engineer Profile
* **Experience:** Software Engineer (2 YOE).
* **Existing Stack:** Production experience in **Java, C, and JavaScript**; extensive project experience in **C++, TypeScript, and Python**.
* **Target Language:** **Go (Golang)** — starting from zero in Go, leveraging strong existing mental models of pointers/memory (from C/C++), concurrency/threads/interfaces (from Java), and event loops/async I/O (from JS/TS).

### Core Learning Philosophy
The user learns exclusively **by doing, failing under pressure, and reverse-engineering**.
* **Zero Academic Fluff:** Never lecture on syntax or theory first. Do not explain Go's scheduler, CAP theorem, caching strategies, or locking mechanisms until the user's system has actively crashed, leaked memory, deadlocked, or corrupted data because those mechanisms were missing or misused.
* **High Agency, High Pressure:** The brain works best when thrown into the trenches. Your job as an AI coding agent is not to be a tutor—**you are the adversary, the chaos engineer, and the war-room pair programmer.**

---

## 2. Living Repository & Artifact Architecture

This repository is designed to be a **permanent, multi-month proof-of-learning, incident archive, and forkable curriculum** that another engineer can later clone to run their own trench warfare journey.

When initializing the repository, scaffold and maintain the following structure:

```text
system-design-trenches/
├── README.md                  # Public-facing overview, scoreboard, and "How to Fork & Fight" guide
├── PHILOSOPHY.md              # Manifesto on Trench Warfare learning (Agency > Academia)
├── CURRICULUM.md              # Living roadmap & checklist from Trench 0 (Go) to Trench 5
├── AGENTS.md                  # Universal instructions for AI coding agents (Cursor, Windsurf, Copilot, etc.)
├── CLAUDE.md                  # Specific operational rules & slash-workflows for Claude Code
├── LEARNING_LOGS/             # Post-mortems, flamegraphs, metrics snapshots, & architectural decisions
│   ├── TEMPLATE_POSTMORTEM.md # Standardized incident report template for every crash
│   ├── trench-0-go/           # Logs & post-mortems from Trench 0
│   ├── trench-1-io-wall/      # Logs & post-mortems from Trench 1
│   ├── trench-2-contention/   # ...
│   ├── trench-3-backpressure/ # ...
│   ├── trench-4-storage/      # ...
│   └── trench-5-chaos/        # ...
├── trenches/                  # Self-contained combat arenas (separate Docker Compose per trench)
│   ├── 00-go-gauntlet/        # Trench 0: Raw Go concurrency, net/http, context, & pprof
│   ├── 01-connection-io-wall/ # Trench 1: High-throughput read/write & cache stampedes
│   ├── 02-race-conditions/    # Trench 2: Flash sale locking & idempotency
│   ├── 03-cascading-failures/ # Trench 3: Async queues, backpressure, & circuit breakers
│   ├── 04-storage-ceiling/    # Trench 4: 50M rows, indexing, partitioning, & sharding
│   └── 05-network-chaos/      # Trench 5: Split-brain, Toxiproxy, & consensus
└── shared/                    # Reusable k6 helpers, Toxiproxy wrappers, & Prometheus/Grafana configs
```

### Purpose of Key Root Artifacts

1. **`README.md` (The Public Proof-of-Work & Forkable Guide):**
   * Tracks the user's live progress table across all Trenches (Baseline RPS/Latency vs. Hardened RPS/Latency, Crashes Survived, Post-Mortems Linked).
   * Contains a **"Reset to Naive"** guide (via Git tags like `trench-1-naive`, `trench-1-hardened`) so any future human cloning the repo can checkout the broken state and fight the battles themselves.
2. **`PHILOSOPHY.md` (The Manifesto):**
   * Codifies why "Training -> Exam -> Practicality" fails and how "Constrain -> Crash -> Diagnose -> Re-Architect" builds visceral engineering intuition.
3. **`CURRICULUM.md` (The War Map):**
   * Tracks every sub-boss, failure mode, and survival skill unlocked across Trenches 0–5.
4. **`LEARNING_LOGS/` (Post-Mortem Artifacts for Posterity):**
   * **Rule:** A Trench is **not complete** until the user writes a post-mortem in `LEARNING_LOGS/trench-X/` using `TEMPLATE_POSTMORTEM.md`.
   * Every post-mortem must capture:
     1. **The Attack Command & Load Profile** (what hit the system).
     2. **The Smoking Gun** (exact error logs, `pprof` output, `EXPLAIN ANALYZE`, or `docker stats` at the moment of death).
     3. **Root Cause Analysis** (what hardware or concurrency limit broke and why).
     4. **The Architectural Fix & Trade-offs** (what changed, before/after latency & throughput numbers, and what *new* vulnerability that fix introduced).
5. **`AGENTS.md` & `CLAUDE.md` (The AI Guardrails):**
   * Enforce the strict "Never write the fix for the user; build the naive target, write the attack script, and demand diagnosis" protocol across every session.

---

## 3. Rules of Engagement for Claude Code (CRITICAL)

1. **You Build the Naive Target & The Weapon:**
   * When starting a Trench, scaffold the **deliberately naive, unoptimized baseline** (or buggy Go concurrency harness in Trench 0) and constrain its hardware in `docker-compose.yml` (`cpus: '0.5'`, `mem_limit: 256m`).
   * Create a Git tag (e.g., `trench-X-naive`) so the broken baseline is preserved forever for future learners.
   * Write the **attack script** (`k6`, `go test -race`, `wrk`, or `Toxiproxy` chaos script) that will overwhelm or corrupt the naive build.
2. **Never Write the Architectural Fix First:**
   * **DO NOT** pre-optimize the code. **DO NOT** write the worker pool, the `context.WithTimeout`, the Redis cache, or the `SELECT FOR UPDATE` lock for the user.
3. **Bridge to Existing C / Java / JS Mental Models (For Go):**
   * Since the user knows C, Java, and JS/TS, explain Go mechanics *only when asked or during debugging* by contrasting them briefly with what they already know:
     * Structs & pointers -> like C, but with escape analysis and garbage collection (no pointer arithmetic).
     * Interfaces -> implicit/structural typing (like TypeScript shapes), unlike Java's explicit `implements`.
     * Goroutines -> M:N green threads multiplexed over OS threads (lighter than Java platform threads, preemptive unlike JS single-threaded event loop).
     * Error handling -> explicit C-style return values `(val, err)` instead of Java/JS `try/catch` exceptions.
4. **Force Diagnosis Through Metrics & Logs:**
   * After the system crashes, ask: *"What broke, and what do `pprof`, `-race`, `docker stats`, or the logs say?"*
5. **Enforce the Post-Mortem Before Moving On:**
   * Once the user's fix survives the attack and the Boss Fight, prompt them to document the incident in `LEARNING_LOGS/` and tag the commit (`trench-X-survived`) before unlocking the next Trench.

---

## 4. The Curriculum: Trench 0 to Trench 5

### Trench 0: Learning Go in the Trenches (The Go Gauntlet)
* **Objective:** Go from 0 Go knowledge to writing production-grade, race-free, profiled Go network services by breaking Go's runtime primitives in 4 rapid micro-combats.
* **Location:** `trenches/00-go-gauntlet/`
* **How It Works:** Claude scaffolds a partially broken or naive CLI/HTTP harness with failing stress tests (`go test -race -bench=.`) and `net/http/pprof` enabled. The user must write Go to fix and harden each stage.

* **Stage 0.1: Structs, Slices, Pointers, & Escape Analysis (The Memory Trap)**
  * **The Build:** Write a high-frequency in-memory log parser and ring buffer in Go (`structs`, slices, maps, pointer vs. value receivers, `error` interfaces).
  * **The Trap:** Slice capacity reallocations, mutating shared backing arrays across functions, and nil-pointer / nil-map panics under load.
  * **Survival Lesson:** Value vs. pointer semantics (C mental model vs. Go GC), how slice headers (`ptr, len, cap`) actually work under the hood, implicit interfaces, and idiomatic `if err != nil` error wrapping (`fmt.Errorf("%w", err)`).

* **Stage 0.2: Goroutines, Channels, & The Data Race Minefield**
  * **The Build:** Build a concurrent aggregator that fans out 1,000 tasks across goroutines and merges results into an in-memory map/state struct.
  * **The Attack:** Run `go test -race` and hit it with 10,000 concurrent calls.
  * **Expected Failure:** `fatal error: concurrent map writes`, lost updates, and permanently blocked goroutines (goroutine leaks) waiting on unbuffered channels that nobody reads from.
  * **Survival Lesson:** `sync.Mutex` vs. `sync.RWMutex`, `sync.WaitGroup`, buffered vs. unbuffered channels, `select` statements, and why "communicate by sharing memory vs. share memory by communicating" has real performance trade-offs.

* **Stage 0.3: `context.Context`, Zombie Goroutines, & `pprof` Flamegraphs**
  * **The Build:** Build a raw `net/http` proxy server that fetches data from a slow upstream mock server and returns JSON (`encoding/json`).
  * **The Attack:** Fire 5,000 requests where clients disconnect after 100ms while upstream takes 2,000ms, inside a `128MB RAM` container.
  * **Expected Failure:** Goroutine count explodes to 50,000+, file descriptors exhaust (`too many open files` because `resp.Body.Close()` was missed or `http.Client` had no timeout), and the container OOM-kills.
  * **Survival Lesson:** Propagating `context.Context` cancellation across HTTP requests and DB calls, `defer` mechanics (and loop pitfalls), tuning `http.Transport` connection pools, and reading live heap/goroutine dumps via `go tool pprof`.

* **Stage 0.4: Boss Fight — TheBounded Worker Pool & Graceful Shutdown**
  * **The Build:** Combine Stages 0.1–0.3 into a bounded concurrent job ingestion HTTP server in Go with `database/sql` (or `pgx`).
  * **The Attack:** Flood the server with requests while sending `SIGTERM` (`docker stop`) mid-flight.
  * **Survival Lesson:** Implementing a backpressured worker pool with channels, catching OS signals (`os/signal`), draining in-flight requests with `srv.Shutdown(ctx)` without dropping data, and logging a clean exit.

---

### Trench 1: The Connection Pool & I/O Wall
* **Domain:** High-Throughput Read/Write Service in Go (URL Shortener or Telemetry Ingest).
* **Naive Setup:** Single Go API container (`0.5 CPU`, `128MB RAM`, unconfigured `sql.DB` pool) + Single PostgreSQL container (`max_connections=40`, `256MB RAM`, no indexes beyond PK).
* **The Attack:** `k6` script firing 2,000 concurrent VUs with a 90/10 read/write split and Zipfian (hot-key) distribution.
* **Expected Failure:** `pq: sorry, too many clients already`, sky-high CPU wait on Postgres disk I/O, `500 Internal Server Error` cascades.
* **Target Survival Lessons:**
  * Tuning Go's `db.SetMaxOpenConns`, `db.SetMaxIdleConns`, and external pooling (`PgBouncer`).
  * B-Tree indexing vs. write amplification (`EXPLAIN (ANALYZE, BUFFERS)`).
  * Read-through caching with Redis (`go-redis`).
  * **Boss Fight:** Cache stampede (thundering herd) when a hot key expires under 2,000 RPS -> user must discover and implement request coalescing (`golang.org/x/sync/singleflight`) or mutex locks.

---

### Trench 2: Race Conditions & Double-Booking
* **Domain:** High-Contention Flash Sale / Concert Ticket Booking API in Go.
* **Naive Setup:** Go endpoint that runs `SELECT stock`, checks `if stock > 0`, runs `UPDATE stock = stock - 1`, and inserts an order row.
* **The Attack:** `k6` script syncing 5,000 concurrent requests via a barrier to buy the final 100 tickets at the exact same millisecond, plus duplicate client retries.
* **Expected Failure:** Massive inventory overselling (stock hits `-842`) or database deadlock freezes (`ERROR: deadlock detected`).
* **Target Survival Lessons:**
  * Database isolation levels (`READ COMMITTED` vs. `SERIALIZABLE`).
  * Pessimistic locking (`SELECT ... FOR UPDATE` / `SKIP LOCKED`) vs. Optimistic concurrency control (version columns).
  * Atomic Redis Lua scripts for inventory decrement.
  * Idempotency keys to survive client retries without double-charging.

---

### Trench 3: Cascading Failures & Backpressure
* **Domain:** Synchronous Order Checkout / Webhook Processing Pipeline in Go.
* **Naive Setup:** Go API receives payload and makes a synchronous HTTP call to a mock third-party service (configured to take 500ms and fail 15% of the time).
* **The Attack:** Spike traffic 10x for 90 seconds while degrading the mock third-party service to a 3,000ms response time.
* **Expected Failure:** Goroutine buildup, memory exhaustion, saturation of downstream connections, and complete unavailability of unrelated `/health` and read endpoints.
* **Target Survival Lessons:**
  * Context timeouts, exponential backoff with jitter, and state-machine Circuit Breakers.
  * Decoupling ingestion from processing via Message Queues (RabbitMQ / Kafka / Redis Streams).
  * Go worker pool concurrency tuning, load shedding (`429 Too Many Requests`), and Dead Letter Queues (DLQs) for poison-pill messages.

---

### Trench 4: The Single-Node Storage Ceiling
* **Domain:** High-Volume Time-Series / Chat Message Store in Go + Postgres.
* **Naive Setup:** Single Postgres container (`256MB RAM`) pre-seeded via a fast Go `COPY` seeder with 25–50 million rows so indexes exceed `shared_buffers` and available RAM.
* **The Attack:** Concurrent paginated range queries (`OFFSET 100000 LIMIT 50`) alongside continuous bulk inserts.
* **Expected Failure:** Disk swap thrashing, query latency jumping from 5ms to 5,000ms+, write locks blocking reads.
* **Target Survival Lessons:**
  * Cursor/keyset pagination vs. `OFFSET` performance collapse.
  * Postgres declarative table partitioning (time-based or range-based).
  * Read/Write splitting with a Primary-Replica setup in Go—and handling **replication lag** when a user writes data and immediately reads stale state.
  * Application-level sharding in Go using consistent hashing across multiple Postgres containers.

---

### Trench 5: Split-Brain & Network Chaos
* **Domain:** Multi-Node Distributed Stateful System (3 Go API nodes, Nginx LB, Redis Sentinel/Cluster, Postgres Primary-Replica).
* **Naive Setup:** Distributed cluster assuming reliable local Docker networking.
* **The Attack:** Automated chaos script using `Toxiproxy` and `docker kill` to inject 400ms latency, half-open TCP connections (blackholes), 25% packet loss, or killing the DB primary mid-transaction.
* **Expected Failure:** Split-brain writes, hung Go HTTP/DB connections, inconsistent cache vs. DB state, total cluster outage.
* **Target Survival Lessons:**
  * Real-time CAP theorem trade-offs (Availability vs. Consistency during partition).
  * Graceful degradation (serving stale cache or shedding load when dependencies die).
  * Distributed leases, fencing tokens, heartbeats, and automated failover.

---

## 5. Initialization Protocol (What Claude Code Should Do First)

When the user starts a session in an empty or freshly initialized repository and references `handoff.md`:

1. **Scaffold the Living Repository Skeleton:**
   * Create `README.md`, `PHILOSOPHY.md`, `CURRICULUM.md`, `AGENTS.md`, `CLAUDE.md`, and `LEARNING_LOGS/TEMPLATE_POSTMORTEM.md`.
2. **Initialize Trench 0 (`trenches/00-go-gauntlet/`):**
   * Scaffold **Stage 0.1 & Stage 0.2** with a `go.mod`, a deliberately buggy/naive Go concurrency & memory harness, a `Dockerfile` / `docker-compose.yml` (`128MB RAM`), and a brutal test/attack suite (`Makefile` with `make attack`).
3. **Brief the User & Pull the Pin:**
   * Show the user the repository layout, give them the single command to run the first Go gauntlet failure (`make attack`), and wait for the crash. **Do not explain the fix.**