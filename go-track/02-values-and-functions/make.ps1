# Windows equivalent of the Makefile: .\make.ps1 run | test | break
param([ValidateSet('run', 'test', 'break')][string]$Task = 'test')
Set-Location $PSScriptRoot
switch ($Task) {
    'run'   { go run . }
    'test'  { go test -count=1 -v -skip '^TestBreak' . }
    'break' { go test -count=1 -v -run '^TestBreak' . }
}
