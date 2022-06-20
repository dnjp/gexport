# Google Auth Export

This is a tool that extracts individual QR codes from a Google Authenticator backup file which can be used to import accounts into a tool like [Raivo OTP](https://github.com/raivo-otp/ios-application).

## Usage

```
gexport [QR FILE] [OUTPUT DIRECTORY]
```

## Installation

### Go Install

```
$ go install github.com/dnjp/gexport@latest
```

### Build from Source

1. Clone the repository and change to the local directory

```
$ git clone https://github.com/dnjp/gexport
$ cd gexport
```

2. Build the binary

```
$ make
```

3. Copy the binary to your desired location

```
$ sudo cp bin/gexport /usr/local/bin/gexport
```

