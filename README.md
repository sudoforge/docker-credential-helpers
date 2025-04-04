[![GitHub release](https://img.shields.io/github/release/docker/docker-credential-helpers.svg?style=flat-square)](https://github.com/docker/docker-credential-helpers/releases/latest)
[![PkgGoDev](https://img.shields.io/badge/go.dev-docs-007d9c?style=flat-square&logo=go&logoColor=white)](https://pkg.go.dev/github.com/docker/docker-credential-helpers)
[![Build Status](https://img.shields.io/github/actions/workflow/status/docker/docker-credential-helpers/build.yml?label=build&logo=github&style=flat-square)](https://github.com/docker/docker-credential-helpers/actions?query=workflow%3Abuild)
[![Codecov](https://img.shields.io/codecov/c/github/docker/docker-credential-helpers?logo=codecov&style=flat-square)](https://codecov.io/gh/docker/docker-credential-helpers)
[![Go Report Card](https://goreportcard.com/badge/github.com/docker/docker-credential-helpers?style=flat-square)](https://goreportcard.com/report/github.com/docker/docker-credential-helpers)

## Introduction

docker-credential-helpers is a suite of programs to use native stores to keep Docker credentials safe.

## Installation

Go to the [Releases](https://github.com/docker/docker-credential-helpers/releases) page and download the binary that works better for you. Put that binary in your `$PATH`, so Docker can find it.

## Building

You can build the credential helpers using Docker:

```shell
# install emulators
$ docker run --privileged --rm tonistiigi/binfmt --install all

# create builder
$ docker buildx create --use

# build credential helpers from remote repository and output to ./bin/build
$ docker buildx bake "https://github.com/docker/docker-credential-helpers.git"

# or from local source
$ git clone https://github.com/docker/docker-credential-helpers.git
$ cd docker-credential-helpers
$ docker buildx bake
```

Or if the toolchain is already installed on your machine:

1. Download the source.

```shell
$ git clone https://github.com/docker/docker-credential-helpers.git
$ cd docker-credential-helpers
```

2.  Use `make` to build the program you want. That will leave an executable in the `bin` directory inside the repository.

```shell
$ make osxkeychain
```

3.  Put that binary in your `$PATH`, so Docker can find it.

```shell
$ cp bin/build/docker-credential-osxkeychain /usr/local/bin/
```

## Usage

### With the Docker Engine

Set the `credsStore` option in your `~/.docker/config.json` file with the suffix of the program you want to use. For instance, set it to `osxkeychain` if you want to use `docker-credential-osxkeychain`.

```json
{
  "credsStore": "osxkeychain"
}
```

### With other command line applications

The sub-package [client](https://godoc.org/github.com/docker/docker-credential-helpers/client) includes
functions to call external programs from your own command line applications.

There are three things you need to know if you need to interact with a helper:

1. The name of the program to execute, for instance `docker-credential-osxkeychain`.
2. The server address to identify the credentials, for instance `https://example.com`.
3. The username and secret to store, when you want to store credentials.

You can see examples of each function in the [client](https://godoc.org/github.com/docker/docker-credential-helpers/client) documentation.

### Available programs

- gopass: Provides a helper to use `gopass` as credentials store.
- osxkeychain: Provides a helper to use the OS X keychain as credentials store.
- pass: Provides a helper to use `pass` as credentials store.
- secretservice: Provides a helper to use the D-Bus secret service as credentials store.
- wincred: Provides a helper to use Windows credentials manager as store.

#### Note regarding `gopass`

`gopass` requires manual intervention in order for `docker-credential-gopass` to
work properly: a password store must be initialized. Please ensure to review the
upstream [quick start guide][gopass-quick-start] for more information.

[gopass-quick-start]: https://github.com/gopasspw/gopass#quick-start-guide

#### Note regarding `pass`

`pass` requires manual interview in order for `docker-credential-pass` to
work properly. It must be initialized with a `gpg2` key ID. Make sure your GPG
key exists is in `gpg2` keyring as `pass` uses `gpg2` instead of the regular
`gpg`.

## Development

A credential helper can be any program that can read values from the standard input. We use the first argument in the command line to differentiate the kind of command to execute. There are four valid values:

- `store`: Adds credentials to the keychain. The payload in the standard input is a JSON document with `ServerURL`, `Username` and `Secret`.
- `get`: Retrieves credentials from the keychain. The payload in the standard input is the raw value for the `ServerURL`.
- `erase`: Removes credentials from the keychain. The payload in the standard input is the raw value for the `ServerURL`.
- `list`: Lists stored credentials. There is no standard input payload.

This repository also includes libraries to implement new credentials programs in Go. Adding a new helper program is pretty easy. You can see how the OS X keychain helper works in the [osxkeychain](osxkeychain) directory.

1. Implement the interface `credentials.Helper` in `YOUR_PACKAGE/`
2. Create a main program in `YOUR_PACKAGE/cmd/`.
3. Add make tasks to build your program and run tests.

## License

MIT. See [LICENSE](LICENSE) for more information.
