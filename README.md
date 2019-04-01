# wareki

wareki は西暦と和暦を変換するプログラムです.

## Description

wareki は西暦と和暦を変換します.

西暦から和暦への変換は、和暦の元号は西暦に応じて自動的に決まります. たとえば, 1989/01/08 を指定すると H1 (平成1年) に, 1989/01/07 を指定すると S64 (昭和64年) になります.
西暦は省略でき, デフォルト値はシステム日付になります.

デフォルトでは, 元号は英大文字1文字で出力しますが (e.g. H) 漢字で出力することもできます (e.g. 平成)
また, オプションで和暦から西暦に変換することもできます.

令和・平成・昭和・大正・明治に対応しています.

## Usage

```sh
$ date "+%Y/%m/%d"
2018/05/10

$ wareki
H30

$ wareki 2017
H29

$ wareki --kanji
平成30

$ wareki 1989
H1

$ wareki 1989/01/08
H1

$ wareki 1989/01/07
S64

$ wareki --heisei 30
2018

$ wareki --heisei 1
1989

$ wareki --showa 1
1926

$ wareki --help
# ...

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

## Contribution

1. Fork this repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Rebase your local changes against the master branch
5. Run test suite with the go test ./... command and confirm that it passes
6. Run gofmt -s
7. Create new Pull Request

## Licence

[MIT](https://github.com/ebc-2in2crc/wareki/blob/master/LICENSE)

## Author

[ebc-2in2crc](https://github.com/ebc-2in2crc)
