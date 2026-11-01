# Windows equivalent of the Makefile: .\make.ps1 run | test
param([ValidateSet('run', 'test')][string]$Task = 'test')
Set-Location $PSScriptRoot
switch ($Task) {
    'run'  { go run . }
    'test' { go test -count=1 -v . }
}
