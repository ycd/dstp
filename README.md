<div align="center">
<h1>dstp</h1>

[dstp](https://github.com/ycd/dstp) dstp, run common common networking tests against your site.

![dstp gif](assets/dstp.jpeg)

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

Run the tests against your site

```
$ dstp -c https://jvns.ca

Ping: 14.282ms
DNS: resolving 104.21.91.206
TLS: certificate is valid for 24 days
HTTPS: got 200 OK
```

## Installation

### Packages

#### Arch Linux

* [ ] For Arch Linux, install the [``]() package.

#### Homebrew

* [ ] For Homebrew on macOS, install the [``]() formula.

### Downloads

Binary downloads of example are available from [the releases section on GitHub](https://github.com/ycd/dstp/releases/)
for 64-bit Windows, macOS, and Linux targets. They contain the compiled executable.

| platform     |
| ----------- | 
| [macOS 64 Bit](https://github.com/ycd/toc/releases/download/v0.2.5/toc_0.2.5_darwin_x86_64.tar.gz)
| [Linux 32-Bit](https://github.com/ycd/toc/releases/download/v0.2.5/toc_0.2.5_linux_i386.tar.gz)
| [Linux ARM 64 Bit](https://github.com/ycd/toc/releases/download/v0.2.5/toc_0.2.5_linux_arm64.tar.gz)
| [Linux 64 Bit](https://github.com/ycd/toc/releases/download/v0.2.5/toc_0.2.5_linux_x86_64.tar.gz)
| [Windows 64 Bit](https://github.com/ycd/toc/releases/download/v0.2.5/toc_0.2.5_windows_x86_64.zip)
| [Windows 32 Bit](https://github.com/ycd/toc/releases/download/v0.2.5/toc_0.2.5_windows_i386.zip)

### Installation from source

0. Verify that you have Go 1.16+ installed

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
   $ go build .
   $ cp toc /usr/local/dstp
   ```

3. Verify installation

   ```
   $ dstp -h 

   Usage: dstp [OPTIONS] [ARGS]
   Options:
   -a, --addr   <string>  The URL or the IP address to run tests against      [REQUIRED]
   -o, --out    <string>  The type of the output, either json or plaintext    [Default: plaintext]
   -c           <bool>    Run all the tests concurrently.                     [Default: false]
   -h, --help             Show this message and exit.
   ```

---

## Contributing

All kinds of Pull Requests and Feature Requests are welcomed!

## Licence

dstp's source code is licenced under the [Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0.txt).
