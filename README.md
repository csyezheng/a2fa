# a2fa

a2fa *("annoying two-factor authentication")* is a command line tool for generating and validating one-time password.

<!-- TOC -->

  * [Introduction](#introduction)
  * [Installation](#installation)
    + [macOS](#macOS)
    + [Linux](#linux)
    + [Windows](#windows)
    + [Alternate ways](#Alternate-ways)
  * [Usage](#usage)
  * [Examples](#examples)
    + [Generate code](#generate-code)
    + [Add account](#add-account)
    + [List acccount](#list-acccount)
    + [Remove accounts](#remove-accounts)
    + [Update acccount](#update-acccount)
  * [Reporting issues](#reporting-issues)
  * [Contributing](#contributing)
  * [License](#license)

<!-- /TOC -->

## Introduction

a2fa means annoying two-factor authentication. Its purpose is to get rid of phones and be able to authenticate easily. It keeps synced with Google Authenticator, Microsoft Authenticator.

**Description**:

* An easy-to-use substitute for 2FA apps like TOTP Google authenticator.
* Supports the OATH algorithms, such as TOTP and HOTP.
* No need for network connection.
* No need for phone.

## Installation

### macOS

```
brew install csyezheng/tap/a2fa
```

### Linux

**Archlinux**

```
paru -S a2fa
```

```
yay -S a2fa
```

**Debian-based Linux**

```
echo 'deb [trusted=yes] https://apt.fury.io/csyezheng/ /' | sudo tee /etc/apt/sources.list.d/fury.list
sudo apt update
sudo apt install a2fa
```

### Windows

```
winget install -e --id csyezheng.a2fa
```

### Alternate ways

Please see the [installation ](docs/installation.md)

## Usage

```
a2fa [command] [flags] [args]
```

```
Available Commands:
  add         Add account and its secret key
  completion  Generate the autocompletion script for the specified shell
  generate    Generate one-time password from secret key
  help        Help about any command
  list        List all added accounts and password code
  remove      Remove account and its secret key
  update      Add account and its secret key
  version     show version
```

```
a2fa generate [flags] <secret key>
a2fa add [flags] <account name> <secret key>
a2fa remove <account name> [user name]
a2fa update [flags] <account name> <secret key>
a2fa list [account name]
```

Commonly used flags

```
Flags:
  -b, --base32         use base32 encoding of KEY instead of hex (default true)
  -c, --counter int    used for HOTP, A counter C, which counts the number of iterations
  -e, --epoch int      used for TOTP, epoch (T0) which is the Unix time from which to start counting time steps
  -H, --hash string    A cryptographic hash method H (SHA1, SHA256, SHA512) (default "SHA1")
  -h, --help           help for generate
  -i, --interval int   used for TOTP, an interval (Tx) which will be used to calculate the value of the counter CT (default 30)
  -l, --length int     A HOTP value length d (default 6)
  -m, --mode string    use use time-variant TOTP mode or use event-based HOTP mode (default "totp")
```

## Examples

### Generate code

Generate a **time-based** one-time password but do not save the secret key

```
a2fa generate ADOO3MCCCVO5AVD6
```

Generate a **counter-based** one-time password with counter 1

```
a2fa generate -m hotp -c 1 ADOO3MCCCVO5AVD6
```

### Add account

Add an account named GitHub

```
a2fa add GitHub ADOO3MCCCVO5AVD6
```

Add an account, the account name is GitHub, the user name is csyezheng

```
a2fa add GitHub:csyezheng ADOO3MCCCVO5AVD6
```

### List acccount

List all accounts

```shell
a2fa list 
```

List all accounts named GitHub

```
a2fa list GitHub
```

List accounts whose account name is GitHub and whose username is csyezheng

```
a2fa list GitHub:csyezheng
```

List accounts whose account name is GitHub and whose username is csyezheng

```
a2fa list GitHub csyezheng
```

### Remove accounts

Remove all accounts named GitHub

```
a2fa remove GitHub
```

Delete accounts  whose account name is GitHub and whose username is csyezheng

```
a2fa remove GitHub csyezheng
```

### Update acccount

Update the secret key of accounts which account name is GitHub

```
a2fa update GitHub 5BRSSSBJUWBQBOXE
```

Update the secret key of accounts which account name is GitHub and the username is csyezheng

```
a2fa update GitHub:csyezheng 5BRSSSBJUWBQBOXE
```

## Reporting issues

If you encounter any problems, you can open an issue in our bug tracker, please fill the issue template with *as much information as possible*.

## Contributing

Refer to [CONTRIBUTING.md](CONTRIBUTING.md)

## License

Apache License 2.0, see [LICENSE](LICENSE).
