<div align="center">
<h1>dstp</h1>

[dstp](https://github.com/ycd/dstp), run common networking tests against your site.

![dstp gif](assets/dstp.png)

</div>


---

## Usage

```
Usage: dstp [OPTIONS] [ARGS]
Options:
	-a, --addr   <string>  The URL or the IP address to run tests against      [REQUIRED]
	-o, --out    <string>  The type of the output, either json or plaintext    [Default: plaintext] 
	-c           <bool>    Run all the tests concurrently.                     [Default: false]
	-h, --help             Show this message and exit.
```

## Motivation

A comment on [lobste.rs](https://lobste.rs/s/qtsklv/how_do_you_tell_if_problem_is_caused_by_dns#c_1nqkdp), in a thread
about DNS gave a great idea and thought a robust tool like that come in handy!

## Installation

### Packages

#### Homebrew

For Homebrew on macOS, install the [`ycd/tap/dstp`](https://github.com/ycd/homebrew-tap#readme) formula.

```zsh
brew install ycd/tap/dstp
```

#### Go Install

```zsh
go install github.com/ycd/dstp/cmd/dstp
```

#### NixOS

1. Add `dstp`to `/etc/nixos/configuration.nix`:

```
environment.systemPackages = with pkgs; [
  dstp
];
```

2. Run:

```zsh
sudo nixos-rebuild switch
```


### Downloads

Binary downloads of example are available from [the releases section on GitHub](https://github.com/ycd/dstp/releases/)
for 64-bit Windows, macOS, and Linux targets. They contain the compiled executable.

| platform     |
| ----------- | 
| [macOS ARM 64 Bit](https://github.com/ycd/dstp/releases/download/v0.3.0/dstp_0.3.0_darwin_arm64.tar.gz)
| [macOS 64 Bit](https://github.com/ycd/dstp/releases/download/v0.3.0/dstp_0.3.0_darwin_x86_64.tar.gz)
| [Linux 32-Bit](https://github.com/ycd/dstp/releases/download/v0.3.0/dstp_0.3.0_linux_i386.tar.gz)
| [Linux ARM 64 Bit](https://github.com/ycd/dstp/releases/download/v0.3.0/dstp_0.3.0_linux_arm64.tar.gz)
| [Linux 64 Bit](https://github.com/ycd/dstp/releases/download/v0.3.0/dstp_0.3.0_linux_x86_64.tar.gz)
| [Windows 64 Bit](https://github.com/ycd/dstp/releases/download/v0.3.0/dstp_0.3.0_windows_x86_64.zip)
| [Windows 32 Bit](https://github.com/ycd/dstp/releases/download/v0.3.0/dstp_0.3.0_windows_i386.zip)

### Installation from source

0. Verify that you have Go 1.17+ installed (The source code uses _( `//go:build` )_ conditional compilation directives that is introduced in Go 1.17.)

   ```
   $ go version
   ```

   If `go` is not installed, follow instructions on [the Go website](https://golang.org/doc/install).

1. Clone this repository

   ```
   $ git clone https://github.com/ycd/dstp 
   $ cd dstp
   ```

2. Build and install

   #### Unix/Linux
   ```
   # May require you to use sudo
   $ go build cmd/dstp/main.go
   $ cp dstp /usr/local/bin/dstp
   ```

   #### Mac/BSD
   ```
   # May require you to use sudo
   $ make
   $ cp dstp /usr/local/bin/dstp
   ```

3. Verify installation

   ```
   $ dstp -h 

   Usage: dstp [OPTIONS] [ARGS]
   Options:
   -a, --addr   <string>  The URL or the IP address to run te![img.png](img.png)sts against      [REQUIRED]
   -o, --out    <string>  The type of the output, either json or plaintext    [Default: plaintext]
   -c           <bool>    Run all the tests concurrently.                     [Default: false]
   -h, --help             Show this message and exit.
   ```

---

## Contributing

All kinds of Pull Requests and Feature Requests are welcomed!

## Licence

dstp's source code is licenced under the [MIT License](https://choosealicense.com/licenses/mit/).
