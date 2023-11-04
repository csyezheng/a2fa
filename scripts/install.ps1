$downloadURL = "https://github.com/csyezheng/a2fa/releases/latest/download/a2fa_Windows_x86_64.zip"
$packageName = "a2fa_Windows_x86_64.zip"

$tempDir = Join-Path ([System.IO.Path]::GetTempPath()) ([System.IO.Path]::GetRandomFileName())
New-Item -ItemType Directory -Path $tempDir -Force -ErrorAction SilentlyContinue

$packagePath = Join-Path -Path $tempDir -ChildPath $packageName

Write-Host "Downloading package, please wait" -ForegroundColor Green
Invoke-WebRequest -Uri $downloadURL -OutFile $packagePath

$contentPath = Join-Path -Path $tempDir -ChildPath "new"
New-Item -ItemType Directory -Path $contentPath -ErrorAction SilentlyContinue

Write-Host "Extracting archive..." -ForegroundColor Green
Expand-Archive -Path $packagePath -DestinationPath $contentPath
Write-Host "Successfully extracted archive" -ForegroundColor Green

Write-Host "Starting package install..." -ForegroundColor Green

# $Destination = "$env:USERPROFILE\Packages\a2fa"
$Destination = "$env:LOCALAPPDATA\a2fa"
New-Item -ItemType Directory -Path $Destination -Force -ErrorAction SilentlyContinue

Get-ChildItem -Recurse -Path "$contentPath" -File | ForEach-Object {
    $DestinationFilePath = Join-Path $Destination $_.fullname.replace($contentPath, "")
    Copy-Item $_.fullname -Destination $DestinationFilePath
}

Write-Host "a2fa has been installed at $Destination" -ForegroundColor Green

$newPath = $env:Path + [System.IO.Path]::PathSeparator + $Destination
[Environment]::SetEnvironmentVariable("PATH", $newPath, [EnvironmentVariableTarget]::User)

Write-Host "Path environment variable modified; restart your shell to use the new value." -ForegroundColor Green
Write-Host "Successfully installed." -ForegroundColor Green

Remove-Item -Path $tempDir -Recurse -Force -ErrorAction SilentlyContinue