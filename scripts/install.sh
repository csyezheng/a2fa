#!/usr/bin/env bash

# exit codes
# 0 - exited without problems
# 1 - OS not supported by this script

set -e

#detect the platform
OS="$(uname)"
case $OS in
  Linux)
    OS='linux'
    ;;
  Darwin)
    OS='osx'
    binTgtDir=/usr/local/bin
    ;;
  *)
    echo 'OS not supported'
    exit 1
    ;;
esac

OS_type="$(uname -m)"
case "$OS_type" in
  x86_64|amd64)
    OS_type='amd64'
    ;;
  *)
    echo 'OS type not supported'
    exit 1
    ;;
esac

version=$(eval "curl https://api.github.com/repos/csyezheng/a2fa/releases/latest -s | jq .name -r")

download_link="https://github.com/csyezheng/a2fa/releases/download/$version/a2fa_Linux_x86_64.tar.gz"
a2fa_archive="a2fa_Linux_x86_64.tar.gz"

curl -OS "$download_link"

decompressed_dir="tmp_a2fa"
mkdir "$decompressed_dir"
tar -xzf "$a2fa_archive" --directory "$decompressed_dir"

cd "$decompressed_dir"

case "$OS" in
  'linux')
    cp a2fa /usr/bin/a2fa.new
    chmod 755 /usr/bin/a2fa.new
    chown root:root /usr/bin/a2fa.new
    mv /usr/bin/a2fa.new /usr/bin/a2fa
    ;;
  'osx')
    mkdir -m 0555 -p ${binTgtDir}
    cp a2fa ${binTgtDir}/a2fa.new
    mv ${binTgtDir}/a2fa.new ${binTgtDir}/a2fa
    chmod a=x ${binTgtDir}/a2fa
    ;;
  *)
    echo 'OS not supported'
    exit 2
esac

version=$(a2fa --version 2>>errors | head -n 1)
printf "\n${version} has successfully installed."