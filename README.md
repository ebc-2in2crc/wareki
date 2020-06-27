[English](README.md) | [日本語](README_ja.md)

# wareki

[![Build Status](https://travis-ci.com/ebc-2in2crc/wareki.svg?branch=master)](https://travis-ci.com/ebc-2in2crc/wareki)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/ebc-2in2crc/wareki?status.svg)](https://godoc.org/github.com/ebc-2in2crc/wareki)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/wareki)](https://goreportcard.com/report/github.com/ebc-2in2crc/wareki)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/wareki.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/wareki.svg?label=version)

`wareki` converts between Japanese calendar and western calendar.

## Description

`wareki` converts between Japanese calendar and western calendar.

When converting from the Western calendar to the Japanese calendar, the era of the Japanese calendar is automatically determined according to the Western calendar.
The year of Western calendar can be omitted and the default value is the system date.

By default, the Japanese era is output with one uppercase letter (e.g. `R`) but can also be output with kanji (e.g. `令和`)
You can also convert from Japanese calendar to Western calendar by specifying an option.

`wareki` support Reiwa, Heisei, Taisho, Meiji.

## Usage

```sh
$ date "+%Y/%m/%d"
2019/05/01

$ wareki
R1

$ wareki --kanji
令和1

$ wareki 2018
H30

$ wareki 2019/05/01
R1

$ wareki 2019/04/30
H31

$ wareki --reiwa 1
2019

$ wareki --heisei 1
1989

$ wareki --heisei 1
1989

$ wareki --showa 1
1926

$ wareki --help
# ...
```

Or, you can use Docker.

```sh
$ date "+%Y/%m/%d"
2019/05/01

$ docker container run --rm ebc2in2crc/wareki
R1

$ docker container run --rm ebc2in2crc/wareki --kanji
令和1
```

## Installation

### Developer

```
$ go get -u github.com/ebc-2in2crc/wareki/...
```

### User

Download from the following url.

- [https://github.com/ebc-2in2crc/wareki/releases](https://github.com/ebc-2in2crc/wareki/releases)

Or, you can use Homebrew (Only macOS).

```sh
$ brew tap ebc-2in2crc/wareki
$ brew install wareki
```

Or, you can use Docker.

```sh
$ docker pull ebc2in2crc/wareki
```

## Contribution

1. Fork this repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Rebase your local changes against the master branch
5. Run test suite with the go test ./... command and confirm that it passes
6. Run gofmt -s
7. Create new Pull Request

## License

[MIT](https://github.com/ebc-2in2crc/wareki/blob/master/LICENSE)

## Author

[ebc-2in2crc](https://github.com/ebc-2in2crc)
