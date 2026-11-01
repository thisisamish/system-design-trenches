# Curriculum

Tick a box only after the lesson gate: tests pass, break tests pass, quiz done, recap written.

## Phase 1: The Go Track (`go-track/`)

Tutor-led. Every lesson is Learn → Build → Test → Break.

- [ ] 01 Hello, Go: `go run`, packages, imports, functions, tests (break: read compiler and test errors on purpose)
- [ ] 02 Values and functions: variables, types, zero values, conversions, multiple returns (break: integer overflow)
- [ ] 03 Control flow: `if`, `for`, `switch`, strings vs runes (break: boundaries, unicode, speed)
- [ ] 04 Structs, methods, pointers vs values (break: a mutation that silently does nothing) *(not yet written)*
- [ ] 05 Arrays and slices: `len`, `cap`, `append` (break: shared backing arrays) *(not yet written)*
- [ ] 06 Maps (break: nil maps, missing keys) *(not yet written)*
- [ ] 07 Errors: `(val, err)`, wrapping, `errors.Is` (break: flood of malformed input) *(not yet written)*
- [ ] 08 Interfaces (break: a hostile implementation) *(not yet written)*
- [ ] 09 Goroutines and `WaitGroup` (break: race detector on shared state) *(not yet written)*
- [ ] 10 Mutexes (break: 10,000 concurrent callers) *(not yet written)*
- [ ] 11 Channels and `select` (break: deadlocks, goroutine leaks) *(not yet written)*
- [ ] 12 `net/http` and JSON (break: load test) *(not yet written)*
- [ ] 13 `context` and timeouts (break: slow upstream, disconnecting clients) *(not yet written)*
- [ ] 14 `pprof` (break: find a planted leak) *(not yet written)*
- [ ] 15 Checkpoint project: bounded worker pool and graceful shutdown (break: flood plus `SIGTERM` mid-flight) *(not yet written)*

## Phase 2: System-Design Trenches (`trenches/`, designed after Phase 1)

Same cycle with bigger systems. Support fades: Trench 1 is tutor-led, Trenches 4–5 are close to adversary mode. Each trench ends with a post-mortem.

### Trench 1: Connection Pool & I/O Wall (support: tutor)
- [ ] `sql.DB` pool tuning / PgBouncer
- [ ] Indexing vs write amplification
- [ ] Redis read-through cache
- [ ] Boss: cache stampede
- [ ] Post-mortem

### Trench 2: Race Conditions & Double-Booking (support: concepts on request)
- [ ] Isolation levels
- [ ] `FOR UPDATE` / `SKIP LOCKED` vs optimistic versioning
- [ ] Redis Lua atomic decrement
- [ ] Idempotency keys
- [ ] Post-mortem

### Trench 3: Cascading Failures & Backpressure (support: hints on request)
- [ ] Timeouts, backoff + jitter, circuit breaker
- [ ] Queue decoupling
- [ ] Load shedding, DLQ
- [ ] Post-mortem

### Trench 4: Storage Ceiling (support: adversary)
- [ ] Keyset pagination
- [ ] Partitioning
- [ ] Replica reads + replication lag
- [ ] Sharding with consistent hashing
- [ ] Post-mortem

### Trench 5: Split-Brain & Network Chaos (support: adversary)
- [ ] CAP trade-offs in practice
- [ ] Graceful degradation
- [ ] Leases, fencing tokens, failover
- [ ] Post-mortem
