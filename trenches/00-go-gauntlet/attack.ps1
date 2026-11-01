# Windows equivalent of `make attack` (make is not installed by default).
param([switch]$Local)
Set-Location $PSScriptRoot
if ($Local) {
    go test -count=1 -v ./...
} else {
    docker compose run --rm --build gauntlet
}
