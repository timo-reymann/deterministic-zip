deterministic-zip
===
[![GitHub Release](https://img.shields.io/github/v/release/timo-reymann/deterministic-zip?label=version)](https://github.com/timo-reymann/deterministic-zip/releases/latest)
[![CircleCI Build Status](https://circleci.com/gh/timo-reymann/deterministic-zip.svg?style=shield)](https://app.circleci.com/pipelines/github/timo-reymann/deterministic-zip)
[![codecov](https://codecov.io/gh/timo-reymann/deterministic-zip/branch/main/graph/badge.svg?token=6O7X0VO5L6)](https://codecov.io/gh/timo-reymann/deterministic-zip)
[![Renovate](https://img.shields.io/badge/renovate-enabled-green?logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAzNjkgMzY5Ij48Y2lyY2xlIGN4PSIxODkuOSIgY3k9IjE5MC4yIiByPSIxODQuNSIgZmlsbD0iI2ZmZTQyZSIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTUgLTYpIi8+PHBhdGggZmlsbD0iIzhiYjViNSIgZD0iTTI1MSAyNTZsLTM4LTM4YTE3IDE3IDAgMDEwLTI0bDU2LTU2YzItMiAyLTYgMC03bC0yMC0yMWE1IDUgMCAwMC03IDBsLTEzIDEyLTktOCAxMy0xM2ExNyAxNyAwIDAxMjQgMGwyMSAyMWM3IDcgNyAxNyAwIDI0bC01NiA1N2E1IDUgMCAwMDAgN2wzOCAzOHoiLz48cGF0aCBmaWxsPSIjZDk1NjEyIiBkPSJNMzAwIDI4OGwtOCA4Yy00IDQtMTEgNC0xNiAwbC00Ni00NmMtNS01LTUtMTIgMC0xNmw4LThjNC00IDExLTQgMTUgMGw0NyA0N2M0IDQgNCAxMSAwIDE1eiIvPjxwYXRoIGZpbGw9IiMyNGJmYmUiIGQ9Ik04MSAxODVsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzI1YzRjMyIgZD0iTTIyMCAxMDBsMjMgMjNjNCA0IDQgMTEgMCAxNkwxNDIgMjQwYy00IDQtMTEgNC0xNSAwbC0yNC0yNGMtNC00LTQtMTEgMC0xNWwxMDEtMTAxYzUtNSAxMi01IDE2IDB6Ii8+PHBhdGggZmlsbD0iIzFkZGVkZCIgZD0iTTk5IDE2N2wxOC0xOCAxOCAxOC0xOCAxOHoiLz48cGF0aCBmaWxsPSIjMDBhZmIzIiBkPSJNMjMwIDExMGwxMyAxM2M0IDQgNCAxMSAwIDE2TDE0MiAyNDBjLTQgNC0xMSA0LTE1IDBsLTEzLTEzYzQgNCAxMSA0IDE1IDBsMTAxLTEwMWM1LTUgNS0xMSAwLTE2eiIvPjxwYXRoIGZpbGw9IiMyNGJmYmUiIGQ9Ik0xMTYgMTQ5bDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMxZGRlZGQiIGQ9Ik0xMzQgMTMxbDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMxYmNmY2UiIGQ9Ik0xNTIgMTEzbDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMyNGJmYmUiIGQ9Ik0xNzAgOTVsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzFiY2ZjZSIgZD0iTTYzIDE2N2wxOC0xOCAxOCAxOC0xOCAxOHpNOTggMTMxbDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMzNGVkZWIiIGQ9Ik0xMzQgOTVsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzFiY2ZjZSIgZD0iTTE1MyA3OGwxOC0xOCAxOCAxOC0xOCAxOHoiLz48cGF0aCBmaWxsPSIjMzRlZGViIiBkPSJNODAgMTEzbDE4LTE3IDE4IDE3LTE4IDE4ek0xMzUgNjBsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzk4ZWRlYiIgZD0iTTI3IDEzMWwxOC0xOCAxOCAxOC0xOCAxOHoiLz48cGF0aCBmaWxsPSIjYjUzZTAyIiBkPSJNMjg1IDI1OGw3IDdjNCA0IDQgMTEgMCAxNWwtOCA4Yy00IDQtMTEgNC0xNiAwbC02LTdjNCA1IDExIDUgMTUgMGw4LTdjNC01IDQtMTIgMC0xNnoiLz48cGF0aCBmaWxsPSIjOThlZGViIiBkPSJNODEgNzhsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzAwYTNhMiIgZD0iTTIzNSAxMTVsOCA4YzQgNCA0IDExIDAgMTZMMTQyIDI0MGMtNCA0LTExIDQtMTUgMGwtOS05YzUgNSAxMiA1IDE2IDBsMTAxLTEwMWM0LTQgNC0xMSAwLTE1eiIvPjxwYXRoIGZpbGw9IiMzOWQ5ZDgiIGQ9Ik0yMjggMTA4bC04LThjLTQtNS0xMS01LTE2IDBMMTAzIDIwMWMtNCA0LTQgMTEgMCAxNWw4IDhjLTQtNC00LTExIDAtMTVsMTAxLTEwMWM1LTQgMTItNCAxNiAweiIvPjxwYXRoIGZpbGw9IiNhMzM5MDQiIGQ9Ik0yOTEgMjY0bDggOGM0IDQgNCAxMSAwIDE2bC04IDdjLTQgNS0xMSA1LTE1IDBsLTktOGM1IDUgMTIgNSAxNiAwbDgtOGM0LTQgNC0xMSAwLTE1eiIvPjxwYXRoIGZpbGw9IiNlYjZlMmQiIGQ9Ik0yNjAgMjMzbC00LTRjLTYtNi0xNy02LTIzIDAtNyA3LTcgMTcgMCAyNGw0IDRjLTQtNS00LTExIDAtMTZsOC04YzQtNCAxMS00IDE1IDB6Ii8+PHBhdGggZmlsbD0iIzEzYWNiZCIgZD0iTTEzNCAyNDhjLTQgMC04LTItMTEtNWwtMjMtMjNhMTYgMTYgMCAwMTAtMjNMMjAxIDk2YTE2IDE2IDAgMDEyMiAwbDI0IDI0YzYgNiA2IDE2IDAgMjJMMTQ2IDI0M2MtMyAzLTcgNS0xMiA1em03OC0xNDdsLTQgMi0xMDEgMTAxYTYgNiAwIDAwMCA5bDIzIDIzYTYgNiAwIDAwOSAwbDEwMS0xMDFhNiA2IDAgMDAwLTlsLTI0LTIzLTQtMnoiLz48cGF0aCBmaWxsPSIjYmY0NDA0IiBkPSJNMjg0IDMwNGMtNCAwLTgtMS0xMS00bC00Ny00N2MtNi02LTYtMTYgMC0yMmw4LThjNi02IDE2LTYgMjIgMGw0NyA0NmM2IDcgNiAxNyAwIDIzbC04IDhjLTMgMy03IDQtMTEgNHptLTM5LTc2Yy0xIDAtMyAwLTQgMmwtOCA3Yy0yIDMtMiA3IDAgOWw0NyA0N2E2IDYgMCAwMDkgMGw3LThjMy0yIDMtNiAwLTlsLTQ2LTQ2Yy0yLTItMy0yLTUtMnoiLz48L3N2Zz4=)](https://renovatebot.com)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=timo-reymann_deterministic-zip&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=timo-reymann_deterministic-zip)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=timo-reymann_deterministic-zip&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=timo-reymann_deterministic-zip)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=timo-reymann_deterministic-zip&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=timo-reymann_deterministic-zip)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ftimo-reymann%2Fdeterministic-zip.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Ftimo-reymann%2Fdeterministic-zip?ref=badge_shield)

Simple (almost drop-in) replacement for zip that produces deterministic files.

## Installation

### Automatic install

```bash
bash <(curl -sS https://raw.githubusercontent.com/timo-reymann/deterministic-zip/main/installer)
```

### Manual

#### Linux (64-bit)

```bash
curl -LO https://github.com/timo-reymann/deterministic-zip/releases/download/$(curl -Lso /dev/null -w %{url_effective} https://github.com/timo-reymann/deterministic-zip/releases/latest | grep -o '[^/]*$')/deterministic-zip_linux-amd64 && \
chmod +x deterministic-zip_linux-amd64 && \
sudo mv deterministic-zip_linux-amd64 /usr/local/bin/deterministic-zip
```

#### Darwin (Intel)

##### brew

```bash
brew tap timo-reymann/deterministic-zip
brew install deterministic-zip
```

##### manual

```bash
curl -LO https://github.com/timo-reymann/deterministic-zip/releases/download/$(curl -Lso /dev/null -w %{url_effective} https://github.com/timo-reymann/deterministic-zip/releases/latest | grep -o '[^/]*$')/deterministic-zip_darwin-amd64 && \
chmod +x deterministic-zip_darwin-amd64 && \
sudo mv deterministic-zip_darwin-amd64 /usr/local/bin/deterministic-zip
```

### Install with go

```bash
go get -u github.com/timo-reymann/deterministic-zip
```

### Supported platforms

The following platforms are supported (and have prebuilt binaries):

- Linux (gcc)
    - 32-bit
    - 64-bit
    - ARM 64-bit
    - ARM 32-bit
- Darwin
    - 64-bit
    - ARM (M1)
- Windows
    - ARM
    - 32-bit
    - 64-bit
- FreeBSD
  - 32-bit
  - 64-bit
  - ARM 64-bit
  - ARM 32-bit
- OpenBSD
  - 32-bit
  - 64-bit

Binaries for all of these can be found on
the [latest release page](https://github.com/timo-reymann/deterministic-zip/releases/latest).

## Usage

```sh
deterministic-zip -h
```

## FAQ

### Why?!

Why another zip-tool? What is this deterministic stuff?!

When we are talking about deterministic it means that the hash of the zip file won't change unless the contents of the
zip file changes.

This means only the content, no metadata. You can achieve this with zip, yes.

The problem that still remains is that the order is almost unpredictable and zip is very platform specific, so you will
end up with a bunch of crazy shell pipelines. And I am not even talking about windows at this point.

So this is where this tool comes in, it is intended to be a drop-in replacement for zip in your build process.

The use cases for this are primary:

- Zipping serverless code
- Backups or other files that get rsynced

#### Want to know more about the topic of deterministic/reproducible builds?

I can recommend the following resources:

- [reproducible-builds.org](https://reproducible-builds.org/)
- [Debian Wiki](https://wiki.debian.org/ReproducibleBuilds/About)

### How reliable is it?

Of course, it is not as reliable as the battle-proven and billions of times executed zip.

Even though I am heavily relying on the go stdlib this software can of course have bugs. And you are welcome to report
them and help make this even more stable. Of course there will be tests to cover most use cases but at the end this is
still starting from scratch, so if you need advanced features or just dont feel comfortable about using this tool don't
do it!

### Differences between zip and deterministic-zip

Please see [docs/differences](./docs/differences)

### Why not just using another project out there?

As far as I know the following (GitHub) projects exist:

- [bboe/deterministic_zip (Python)](https://github.com/bboe/deterministic_zip)
    - You must list files explicitly
    - Changed order -> changed zip
    - You will need to install Python (no problem on Linux/Mac) and the package
- [bitgenics/deterministic-zip (NodeJS/JavaScript)](https://github.com/bitgenics/deterministic-zip#readme)
    - Support for globs and ignores order
    - You need to install node.js, the package, and it has no cli interface
- [orf/deterministic-zip (Rust)](https://github.com/orf/deterministic-zip)
    - has prebuilt binaries for all relevant platforms (and other can be built easily)
    - very basic, but you can customize compression (nice feature)

All in all they are just simply not what I needed. My favourite is Rust, because its just simply dropping in a binary.
Something that's very convenient especially when it comes to Docker builds.

The main problem that all these solutions share is that it in my opinion cool things like excluding patterns, that I
regularly use are simply not implemented, and i REALLY love glob patterns.


## Libraries

This whole project wouldnt be possible with the great work of the
following libraries:

- [glob by gobwas](https://github.com/gobwas/glob)
- [pflag by spf13](https://github.com/spf13/pflag)
- [go stdlib](https://github.com/golang/go)



## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ftimo-reymann%2Fdeterministic-zip.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Ftimo-reymann%2Fdeterministic-zip?ref=badge_large)