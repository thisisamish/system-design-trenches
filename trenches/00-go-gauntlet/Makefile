.PHONY: attack attack-local

# Stages 0.1 + 0.2 inside the 128MB / 0.5 CPU container, with the race detector.
attack:
	docker compose run --rm --build gauntlet

# No Docker: no memory cap, and -race needs gcc on your PATH.
attack-local:
	go test -count=1 -v ./...
