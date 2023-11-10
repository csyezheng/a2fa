# a2fa

a2fa *("annoying two-factor authentication")* is a command line tool for generating and validating one-time password.

<!-- TOC -->

  * [Introduction](#introduction)
  * [Installation instructions](#installation-instructions)
    + [Linux](#linux)
    + [macOS](#macOS)
    + [Windows](#windows)
    + [Manual installation and more](#Manual-installation-and-more)
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

## Installation instructions

### Linux

One-liner bash script to install a2fa

```
sudo -v ; curl https://raw.githubusercontent.com/csyezheng/a2fa/main/scripts/install.sh | sudo bash
```

### macOS

One-liner bash script to install a2fa

```
sudo -v ; curl https://raw.githubusercontent.com/csyezheng/a2fa/main/scripts/install.sh | sudo bash
```

### Windows

One-liner PowerShell script to install a2fa

```
Invoke-Expression "& { $(Invoke-RestMethod 'https://raw.githubusercontent.com/csyezheng/a2fa/main/scripts/install.ps1') }"
```

### Manual installation and more

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
  -H, --hash string    A cryptographic hash method H (default "SHA1")
  -h, --help           help for add
      --hotp           use event-based HOTP mode
  -i, --interval int   used for TOTP, an interval (Tx) which will be used to calculate the value of the counter CT (default 30)
  -l, --length int     A HOTP value length d (default 6)
      --totp           use use time-variant TOTP mode (default true)
```

## Examples

### Generate code

```
a2fa generate ADOO3MCCCVO5AVD6
```

### Add account

```
a2fa add AccountName ADOO3MCCCVO5AVD6
```

```
a2fa add AccountName:username ADOO3MCCCVO5AVD6
```

### List acccount

```shell
a2fa list 
```

```
a2fa list AccountName
```

### Remove accounts

```
a2fa remove AccountName
```

```
a2fa remove AccountName username
```

### Update acccount

```
a2fa update AccountName 5BRSSSBJUWBQBOXE
```

```
a2fa update AccountName:username 5BRSSSBJUWBQBOXE
```

## Reporting issues

If you encounter any problems, you can open an issue in our bug tracker, please fill the issue template with *as much information as possible*.

## Contributing

Refer to [CONTRIBUTING.md](CONTRIBUTING.md)

## License

Apache License 2.0, see [LICENSE](LICENSE).
