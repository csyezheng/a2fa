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
    download_link="https://github.com/csyezheng/a2fa/releases/latest/download/a2fa_Linux_x86_64.tar.gz"
    a2fa_archive="a2fa_Linux_x86_64.tar.gz"
    ;;
  Darwin)
    OS='osx'
    download_link="https://github.com/csyezheng/a2fa/releases/latest/download/a2fa_Darwin_x86_64.tar.gz"
    a2fa_archive="a2fa_Darwin_x86_64.tar.gz"
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


printf "Downloading package, please wait\n"
curl -LO "$download_link"

printf "Extracting archive...\n"
decompressed_dir="/tmp/a2fa"
if [ ! -d "$decompressed_dir" ]
then
  mkdir "$decompressed_dir"
fi

tar -xzf "$a2fa_archive" --directory "$decompressed_dir"

printf "Successfully extracted archive\n"

cd "$decompressed_dir"

printf "Starting package install...\n"

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

printf "Successfully installed\n"

cd ..

if [ -d "$decompressed_dir" ]
then
  rm -r "$decompressed_dir"
fi