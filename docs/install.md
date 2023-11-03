---
title: "Install"
description: "a2fa Installation"
---

a2fa is a command line tool for generating and validating one-time password. Its purpose is to get rid of phones and be able to authenticate easily. It keeps synced with Google Authenticator, Microsoft Authenticator.

# Install

## Linux installation {#linux}

### Script installation {#linux-script-installation}

    sudo -v ; curl https://raw.githubusercontent.com/csyezheng/a2fa/main/scripts/install.sh | sudo bash

```
a2fa -h
```

### Manual installation {#linux-manual-installation}

Fetch and unpack

    version=$(eval "curl https://api.github.com/repos/csyezheng/a2fa/releases/latest -s | jq .name -r")
    curl -O "https://github.com/csyezheng/a2fa/releases/download/$version/a2fa_Linux_x86_64.tar.gz"
    cd a2fa_Linux_x86_64.tar.gz

Copy binary file

    sudo cp a2fa /usr/bin/
    sudo chown root:root /usr/bin/a2fa
    sudo chmod 755 /usr/bin/a2fa
    
    a2fa -h

## macOS installation {#macos}

### Installation with brew {#macos-brew}

    brew install a2fa

### Manual installation {#macos-manual-installation}

To avoid problems with macOS gatekeeper enforcing the binary to be signed and
notarized it is enough to download with `curl`.

Download the latest version of a2fa.

    version=$(eval "curl https://api.github.com/repos/csyezheng/a2fa/releases/latest -s | jq .name -r")
    cd && curl -O https://github.com/csyezheng/a2fa/releases/download/$version/a2fa_Darwin_x86_64.tar.gz

Unzip the download and cd to the extracted folder.

    unzip -a a2fa_Darwin_x86_64.tar.gz && cd a2fa_Darwin_x86_64

Move a2fa to your $PATH. You will be prompted for your password.

    sudo mkdir -p /usr/local/bin
    sudo mv a2fa /usr/local/bin/

(the `mkdir` command is safe to run, even if the directory already exists).

Remove the leftover files.

    cd .. && rm -rf a2fa_Darwin_x86_64 a2fa_Darwin_x86_64.tar.gz
    
    a2fa -h

## Windows installation {#windows}

### Windows package manager (Winget) {#windows-winget}

To install a2fa
```
winget install -e --id csyezheng.a2fa
```
To uninstall a2fa
```
winget uninstall -e --id csyezheng.a2fa
```

### Manual download{#windows-manual-download}

```
version=$(eval "curl https://api.github.com/repos/csyezheng/a2fa/releases/latest -s | jq .name -r")
wget https://github.com/csyezheng/a2fa/releases/download/$version/a2fa_Windows_x86_64.zip
```

```
unzip -a a2fa_Windows_x86_64.zip && cd a2fa_Windows_x86_64
```

```
a2fa -h
```

