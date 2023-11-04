# a2fa

a2fa is a command line tool for generating and validating one-time password.

<!-- TOC -->

  * [Introduction](#introduction)
  * [Installation instructions](#installation-instructions)
    + [Linux](#linux)
    + [macOS](#macOS)
    + [Windows](#windows)
  * [Usage](#usage)
  * [Examples](#examples)
    + [Generate code](#generate-code)
    + [Add account](#add-account)
    + [Remove accounts](#remove-accounts)
    + [Update acccount](#update-acccount)
    + [List acccount](#list-acccount)
  * [Reporting issues](#reporting-issues)
  * [Contributing](#contributing)
  * [License](#license)

<!-- /TOC -->

## Introduction

a2fa means annoying two-factor authentication. Its purpose is to get rid of phones and be able to authenticate easily. It keeps synced with Google Authenticator, Microsoft Authenticator.

**Description**:

* An easy-to-use substitute for 2FA apps like TOTP Google authenticator.
* Supports the OATH algorithms, such TOTP and HOTP.
* No need for network connection.
* No need for phone.

## Installation instructions

### Linux

Download precompiled binary from [release](https://github.com/csyezheng/a2fa/releases/) page. 

Unzip the download and cd to the extracted folder.

```
tar -zxf a2fa_Linux_x86_64.tar.gz
cd a2fa_Linux_x86_64
```

Copy binary file

```
sudo cp a2fa /usr/bin/
sudo chown root:root /usr/bin/a2fa
sudo chmod 755 /usr/bin/a2fa
```

### macOS

Download precompiled binary from [release](https://github.com/csyezheng/a2fa/releases/) page. 

Unzip the download and cd to the extracted folder.

```
tar -zxf a2fa_Darwin_x86_64.tar.gz
cd a2fa_Darwin_x86_64
```

Move a2fa to your $PATH.

```
sudo mkdir -p /usr/local/bin
sudo mv a2fa  /usr/local/bin/
```

### Windows

Download precompiled binary from [release](https://github.com/csyezheng/a2fa/releases/) page. Open this file in the Explorer and extract `a2fa.exe`.  a2fa is a portable executable so you can place it wherever is convenient. Open a CMD window (or powershell) and run the binary. 

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
a2fa remove [flags] <account name> <account name>...
a2fa update [flags] <account name> <secret key>
a2fa list [flags] [account name] [account name]...
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
a2fa generate [flags] <secret key>
```

**Examples**:

```
a2fa generate ADOO3MCCCVO5AVD6
```

Output: Code: 488821

**Show flags and usage**:

````
a2fa generate -h
````

### Add account

```
a2fa add [flags] <account name> <secret key>
```

**Examples**:

```
a2fa add GitHub ADOO3MCCCVO5AVD6
```

Output: account added successfully

**Show flags and usage**:

```
a2fa add -h
```

### List acccount

```shell
a2fa list [flags] [account name] [account name]...
```

**Examples**:

```
a2fa list
```

Output:

```
0. GitHUb 414033
1. Google 337590
2. Microoft 54936
3. Apple 70362
```

### Remove accounts

```
a2fa remove [flags] <account name> <account name>...
```

**Examples**:

```
a2fa remove GitHub
```

Output: accounts deleted successfully

### Update acccount

```
a2fa update [flags] <account name> <secret key>
```

**Examples**:

```
a2fa update GitHub 5BRSSSBJUWBQBOXE
```

Output: account updated successfully

**Show flags and usage**:

```
a2fa update -h
```

## Reporting issues

If you encounter any problems, you can open an issue in our bug tracker, please fill the issue template with *as much information as possible*.

## Contributing

Refer to [CONTRIBUTING.md](CONTRIBUTING.md)

## License

Apache License 2.0, see [LICENSE](LICENSE).
