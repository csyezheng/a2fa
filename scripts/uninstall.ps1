Write-Host "Uninstalling a2fa" -ForegroundColor Green

$destination = "$env:LOCALAPPDATA\a2fa"

Remove-Item -Path $destination -Recurse -Force -ErrorAction SilentlyContinue

Write-Host "Uninstall a2fa successfully" -ForegroundColor Green

$separator = [System.IO.Path]::PathSeparator
$modifiedPath = $env:Path -split $separator -ne $destination -join $separator
[Environment]::SetEnvironmentVariable("Path", $modifiedPath, [EnvironmentVariableTarget]::User)

Write-Host "The a2fa entry in the environment variable Path has been deleted" -ForegroundColor Green