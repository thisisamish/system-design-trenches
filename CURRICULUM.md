# Curriculum: The War Map

Mark a box only when the fix survives the attack. A trench is done when its post-mortem exists.

## Trench 0: Go Gauntlet (`trenches/00-go-gauntlet/`)
- [ ] 0.1 Memory Trap: slices, pointers, nil maps, error wrapping
- [ ] 0.2 Data Race Minefield: goroutines, channels, mutexes, leaks
- [ ] 0.3 Zombie Goroutines: `context`, `http.Transport`, `pprof` (not yet scaffolded)
- [ ] 0.4 Boss: bounded worker pool + graceful shutdown (not yet scaffolded)
- [ ] Post-mortem in `LEARNING_LOGS/trench-0-go/`

## Trench 1: Connection Pool & I/O Wall
- [ ] `sql.DB` pool tuning / PgBouncer
- [ ] Indexing vs write amplification
- [ ] Redis read-through cache
- [ ] Boss: cache stampede
- [ ] Post-mortem

## Trench 2: Race Conditions & Double-Booking
- [ ] Isolation levels
- [ ] `FOR UPDATE` / `SKIP LOCKED` vs optimistic versioning
- [ ] Redis Lua atomic decrement
- [ ] Idempotency keys
- [ ] Post-mortem

## Trench 3: Cascading Failures & Backpressure
- [ ] Timeouts, backoff + jitter, circuit breaker
- [ ] Queue decoupling
- [ ] Load shedding, DLQ
- [ ] Post-mortem

## Trench 4: Storage Ceiling
- [ ] Keyset pagination
- [ ] Partitioning
- [ ] Replica reads + replication lag
- [ ] Sharding with consistent hashing
- [ ] Post-mortem

## Trench 5: Split-Brain & Network Chaos
- [ ] CAP trade-offs in practice
- [ ] Graceful degradation
- [ ] Leases, fencing tokens, failover
- [ ] Post-mortem
