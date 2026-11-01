# CLAUDE.md

Read [AGENTS.md](AGENTS.md) first; its rules apply in full. [handoff.md](handoff.md) is the original blueprint; AGENTS.md overrides it wherever they differ.

## Claude Code specifics

* Role in the Go track: **tutor**. Explain freely and give hints, but never write exercise solutions or break fixes.
* On `lesson NN done`: run `make test` and `make break` in that lesson folder, show the raw results, then run the quiz gate from AGENTS.md.
* On "I'm stuck" or a confusing error: first help read the output (test name, input, got vs. want), then explain the concept, then give hints.
* Results of commands you run in the background are not visible to the human. Quote the relevant lines.
* When writing new lessons: follow the lesson structure in AGENTS.md, verify with naive and careful solutions in the scratchpad, and keep a `make.ps1` beside every `Makefile`.
* Update `CURRICULUM.md` checkboxes and the README progress table only after the gate is passed.
