# Philosophy: Learn, Build, Test, Break

## Why reading alone fails

Tutorials teach a mechanism and stop. You memorize `sync.Mutex`, pass the quiz, and forget it, because nothing you built ever corrupted its data.

## Why "thrown in the deep end" fails too

This repo first tried the opposite: broken code, an attack, no explanations. Without the basics, that isn't learning; it's guessing. You can't diagnose a nil-map panic in a language whose syntax you can't read yet.

## The cycle

1. **Learn**: one concept, explained with comparisons to languages you already know.
2. **Build**: write it yourself, from a stub.
3. **Test**: prove it does what the spec says.
4. **Break**: attack tests hit the edges: boundaries, huge numbers, real-world text, concurrency, load. Find out where your code fails, and fix it.

Then repeat with a slightly harder concept. Each lesson stands on the previous one.

## Rules

* Start from the barebones. Complexity grows only once the current level feels solid.
* The AI is a tutor who explains anything but never hands over the solution. Writing it is the learning.
* Breaking things is part of every lesson, in small doses, not a final exam.
* Support fades gradually. In the later system-design trenches you lead the diagnosis yourself.
* Understanding is checked, not assumed: a short quiz and a written recap close every lesson.
