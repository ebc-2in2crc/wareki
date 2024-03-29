[English](README.md) | [日本語](README_ja.md)

# wareki

![CI](https://github.com/ebc-2in2crc/wareki/actions/workflows/pr.yml/badge.svg)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/ebc-2in2crc/wareki?status.svg)](https://godoc.org/github.com/ebc-2in2crc/wareki)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/wareki)](https://goreportcard.com/report/github.com/ebc-2in2crc/wareki)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/wareki.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/wareki.svg?label=version)

wareki は西暦と和暦を変換するプログラムです。

## Description

wareki は西暦と和暦を変換します。

西暦から和暦への変換は、和暦の元号は西暦に応じて自動的に決まります. たとえば、2019/05/01 を指定すると R1 (令和1年) に、2019/04/30 を指定すると H31 (平成31年) になります。
西暦は省略でき、デフォルト値はシステム日付になります。

デフォルトでは、元号は英大文字1文字で出力しますが (e.g. R) 漢字で出力することもできます (e.g. 令和)
また、オプションで和暦から西暦に変換することもできます。

令和・平成・昭和・大正・明治に対応しています。

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

Docker を使うこともできます。

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

Go 1.16 or later.

```
$ go install github.com/ebc-2in2crc/wareki/cmd/wareki@latest
```

Go 1.15.

```
$ go get -u github.com/ebc-2in2crc/wareki/...
```

### User

次の URL からダウンロードします。

- [https://github.com/ebc-2in2crc/wareki/releases](https://github.com/ebc-2in2crc/wareki/releases)

Homebrew を使うこともできます (Mac のみ)

```sh
$ brew tap ebc-2in2crc/wareki
$ brew install wareki
```

Docker を使うこともできます。

```sh
$ docker pull ebc2in2crc/wareki
$ docker container run ebc2in2crc/wareki
R4
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
