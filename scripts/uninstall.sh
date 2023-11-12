#!/usr/bin/env bash

# exit codes
# 0 - exited without problems
# 1 - OS not supported by this script

set -e

#detect the platform
OS="$(uname)"
case $OS in
  Linux)
    sudo rm -rf /usr/bin/a2fa
    ;;
  Darwin)
    rm -rf /usr/local/bin/a2fa
    ;;
  *)
    echo 'OS not supported'
    exit 1
    ;;
esac

printf "Successfully uninstalled a2fa\n"