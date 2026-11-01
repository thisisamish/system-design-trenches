.PHONY: run test break

# Run main.go so you can see your functions' output.
run:
	go run .

# Normal tests. -v prints every test, -count=1 skips Go's test cache.
test:
	go test -count=1 -v -skip '^TestBreak' .

# Attack tests, once `make test` passes.
break:
	go test -count=1 -v -run '^TestBreak' .
