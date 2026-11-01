# Post-Mortem: <Trench / Stage - Incident Title>

Date: <YYYY-MM-DD>   Author: <name>

## 1. The Attack Command & Load Profile
<Exact command. Concurrency, duration, payload shape, resource limits.>

## 2. The Smoking Gun
<Raw evidence at the moment of death: error logs, `go test -race` report, `pprof` output, `EXPLAIN ANALYZE`, `docker stats`. Paste it, do not paraphrase.>

## 3. Root Cause Analysis
<Which hardware or concurrency limit broke, and why. Mechanism, not symptom.>

## 4. The Architectural Fix & Trade-offs
* **Change:** <what you changed>
* **Before:** <latency / throughput / failure rate>
* **After:** <latency / throughput / failure rate>
* **New vulnerability introduced:** <what this fix costs or exposes>

## 5. Survival Checklist
- [ ] Attack re-run survives repeatedly
- [ ] Commit tagged `trench-X-survived`
