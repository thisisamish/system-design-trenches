FROM golang:1.22-bookworm
WORKDIR /app
COPY go.mod ./
COPY internal ./internal
# -race needs cgo (gcc is in this image)
ENV CGO_ENABLED=1
CMD ["go", "test", "-race", "-count=1", "-timeout=120s", "-v", "./..."]
