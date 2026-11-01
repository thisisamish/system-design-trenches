# Philosophy: Trench Warfare Learning

## Why "Training -> Exam -> Practicality" fails

Tutorials teach a mechanism before you feel the pain it removes. You memorize `sync.Mutex`, pass the quiz, and forget it, because nothing ever corrupted your data.

## The loop

1. **Constrain** - tiny CPU and RAM limits, so limits are reached in seconds.
2. **Crash** - a naive system meets a hostile load.
3. **Diagnose** - logs, `-race`, `pprof`, `EXPLAIN ANALYZE`, `docker stats`. Evidence first.
4. **Re-Architect** - you write the fix, and you name the new weakness it creates.

## Rules

* Agency over academia: theory is introduced only after the system has died of its absence.
* The AI is the adversary and chaos engineer, never the tutor who hands over the fix.
* A trench is not complete until the post-mortem is written.
* Every fix has a cost. Record it.
