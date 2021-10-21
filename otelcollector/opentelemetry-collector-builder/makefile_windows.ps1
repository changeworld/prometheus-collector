# Temporary method to build the OTL collector and fluent-bit for windows

# building otelcollector
Write-Output "building otelcollector"
Remove-Item .\otelcollector
go get
go build -o otelcollector .
Move-Item .\otelcollector .\otelcollector.exe

Write-Output "FINISHED building otelcollector"

# building fluent-bit plugin

Write-Output "building fluent-bit plugin"

Set-Location ..
Set-Location fluent-bit
Set-Location src

.\makefile_windows.ps1

Set-Location ..
Set-Location ..
Set-Location opentelemetry-collector-builder

Write-Output "FINISHED building fluent-bit plugin"