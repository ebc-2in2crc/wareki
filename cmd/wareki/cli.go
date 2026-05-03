package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"

	"github.com/ebc-2in2crc/wareki"
	"github.com/urfave/cli/v3"
)

const (
	exitCodeOK = iota
	exitCodeError
)

// CLO 標準入出力を入れ替えるための構造体
type CLO struct {
	inputStream          io.Reader
	outStream, errStream io.Writer
}

var clo *CLO

// Run エントリーポイント
func (c *CLO) Run(ctx context.Context, args []string) int {
	clo = c

	app := &cli.Command{
		Name:               appName,
		Usage:              "西暦を和暦に変換する",
		Version:            version,
		HideHelp:           true,
		HideVersion:        true,
		Description:        description(),
		Flags:              flags(),
		CustomHelpTemplate: appHelpTemplate(),
		Action:             action(),
		Writer:             c.outStream,
		ErrWriter:          c.errStream,
	}

	err := app.Run(ctx, args)
	if err != nil {
		_, _ = fmt.Fprintf(clo.errStream, "%v\n", err)
		return exitCodeError
	}
	return exitCodeOK
}

func description() string {
	return `AD に指定した西暦を和暦に変換します.
  和暦の元号は西暦に応じて自動的に決まります. たとえば, 1989/01/08 を指
  定すると H1 (平成1年) に, 1989/01/07 を指定すると S64 (昭和64年) に
  なります.
  AD は省略でき, デフォルト値はシステム日付になります.
  デフォルトでは, 元号は英大文字1文字で出力しますが (e.g. H) --kanji オ
  プションを指定することにより漢字で出力することもできます (e.g. 平成)
  また, --meiji, --taisho, --showa, --heisei, --reiwa オプションによ
  り, 和暦から西暦に変換することもできます.`
}

func flags() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:    "meiji",
			Aliases: []string{"M"},
			Usage:   "明治から西暦に変換します",
		},
		&cli.IntFlag{
			Name:    "taisho",
			Aliases: []string{"T"},
			Usage:   "大正から西暦に変換します",
		},
		&cli.IntFlag{
			Name:    "showa",
			Aliases: []string{"S"},
			Usage:   "昭和から西暦に変換します",
		},
		&cli.IntFlag{
			Name:    "heisei",
			Aliases: []string{"H"},
			Usage:   "平成から西暦に変換します",
		},
		&cli.IntFlag{
			Name:    "reiwa",
			Aliases: []string{"R"},
			Usage:   "令和から西暦に変換します",
		},
		&cli.BoolFlag{
			Name:    "kanji",
			Aliases: []string{"K"},
			Usage:   "元号を漢字で出力します",
		},
		&cli.BoolFlag{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "このヘルプを表示します",
		},
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "バージョンを表示します",
		},
	}
}

func appHelpTemplate() string {
	return `NAME:
  {{.Name}} - {{.Usage}}
	
USAGE:
  {{.Name}} [options] [AD]
	
DESCRIPTION:
  {{.Description}}
	
OPTIONS:
  {{range .Flags}}{{.}}
  {{end}}
`
}

func action() cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) error {
		if cmd.Bool("help") {
			return cli.ShowAppHelp(cmd)
		}
		if cmd.Bool("version") {
			cli.ShowVersion(cmd)
			return nil
		}
		if mustWarekiToAD(cmd) {
			return warekiToAD(cmd)
		}
		return acToWareki(cmd)
	}
}

func mustWarekiToAD(cmd *cli.Command) bool {
	return cmd.Int("meiji") != 0 ||
		cmd.Int("taisho") != 0 ||
		cmd.Int("showa") != 0 ||
		cmd.Int("heisei") != 0 ||
		cmd.Int("reiwa") != 0
}

func warekiToAD(cmd *cli.Command) error {
	switch {
	case cmd.Int("meiji") != 0:
		_, err := fmt.Fprintf(clo.outStream, "%d\n", wareki.MEIJI().ToAD(int(cmd.Int("meiji"))))
		return err
	case cmd.Int("taisho") != 0:
		_, err := fmt.Fprintf(clo.outStream, "%d\n", wareki.TAISHO().ToAD(int(cmd.Int("taisho"))))
		return err
	case cmd.Int("showa") != 0:
		_, err := fmt.Fprintf(clo.outStream, "%d\n", wareki.SHOWA().ToAD(int(cmd.Int("showa"))))
		return err
	case cmd.Int("heisei") != 0:
		_, err := fmt.Fprintf(clo.outStream, "%d\n", wareki.HEISEI().ToAD(int(cmd.Int("heisei"))))
		return err
	case cmd.Int("reiwa") != 0:
		_, err := fmt.Fprintf(clo.outStream, "%d\n", wareki.REIWA().ToAD(int(cmd.Int("reiwa"))))
		return err
	}
	return nil
}

func acToWareki(cmd *cli.Command) error {
	// 西暦から和暦に変換
	// 引数がないときはシステム日付を和暦に変換
	if cmd.Args().Len() == 0 {
		str, err := _acToWareki(time.Now(), cmd.Bool("kanji"))
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(clo.outStream, "%s\n", str)
		return err
	}

	// 引数があるときは日付にパースして和暦に変換
	s := cmd.Args().Get(0)
	if s != "-" {
		return printWareki(cmd, s)
	}

	scanner := bufio.NewScanner(clo.inputStream)
	for scanner.Scan() {
		text := scanner.Text()
		if err := printWareki(cmd, text); err != nil {
			return err
		}
	}

	return nil
}

func _acToWareki(t time.Time, kanji bool) (string, error) {
	g, err := wareki.Date(t)
	if err != nil {
		return "", err
	}

	year := g.Convert(t)
	if kanji {
		return g.KanjiName() + strconv.Itoa(year), nil
	}
	return g.ShortName() + strconv.Itoa(year), nil
}

func printWareki(cmd *cli.Command, s string) error {
	match, err := regexp.MatchString("^\\d{4}(/\\d{2}(/\\d{2})?)?$", s)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("invalid date format. must specify date: e.g.) 2018 or 2018/01 or 2018/01/01")
	}

	// タイムゾーンは JST 固定
	switch len(s) {
	case 4:
		s = s + "/01/01 JST"
	case 7:
		s = s + "/01 JST"
	case 10:
		s = s + " JST"
	}

	t, err := time.Parse("2006/01/02 MST", s)
	if err != nil {
		return err
	}

	str, err := _acToWareki(t, cmd.Bool("kanji"))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(clo.outStream, "%s\n", str)
	return err
}
