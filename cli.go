package main

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"

	"github.com/ebc-2in2crc/golang-suburi/gengo"
	"github.com/urfave/cli"
)

const (
	ExitCodeOK = iota
	ExitCodeError
)

type CLO struct {
	outStream, errStream io.Writer
}

var clo *CLO

func (c *CLO) Run(args []string) int {
	clo = c

	app := cli.NewApp()
	app.Name = AppName
	app.Usage = "西暦を和暦に変換する"
	app.Version = Version
	app.HideHelp = true
	app.HideVersion = true
	app.Description = description()
	app.Flags = flags()
	cli.AppHelpTemplate = appHelpTemplate()
	app.Action = action()
	app.Writer = c.outStream
	app.ErrWriter = c.errStream

	err := app.Run(args)
	if err == nil {
		return ExitCodeOK
	} else {
		fmt.Fprintf(clo.errStream, "%v\n", err)
		return ExitCodeError
	}
}

func description() string {
	return `AC に指定した西暦を和暦に変換します.
  和暦の元号は西暦に応じて自動的に決まります. たとえば, 1989/01/08 を指
  定すると H1 (平成1年) に, 1989/01/07 を指定すると S64 (昭和64年) に
  なります.
  AC は省略でき, デフォルト値はシステム日付になります.
  デフォルトでは, 元号は英大文字1文字で出力しますが (e.g. H) --kanji オ
  プションを指定することにより漢字で出力することもできます (e.g. 平成)
  また, --meiji, --taisho, --showa, --heisei オプションにより, 和暦か
  ら西暦に変換することもできます.`
}

func flags() []cli.Flag {
	return []cli.Flag{
		cli.IntFlag{
			Name:  "meiji, M",
			Usage: "明治から西暦に変換します",
		},
		cli.IntFlag{
			Name:  "taisho, T",
			Usage: "大正から西暦に変換します",
		},
		cli.IntFlag{
			Name:  "showa, S",
			Usage: "昭和から西暦に変換します",
		},
		cli.IntFlag{
			Name:  "heisei, H",
			Usage: "平成から西暦に変換します",
		},
		cli.BoolFlag{
			Name:  "kanji, K",
			Usage: "元号を漢字で出力します",
		},
		cli.BoolFlag{
			Name:  "help, h",
			Usage: "このヘルプを表示します",
		},
		cli.BoolFlag{
			Name:  "version, v",
			Usage: "バージョンを表示します",
		},
	}
}

func appHelpTemplate() string {
	return `NAME:
  {{.Name}} - {{.Usage}}
	
USAGE:
  {{.Name}} [options] [AC]
	
DESCRIPTION:
  {{.Description}}
	
OPTIONS:
  {{range .Flags}}{{.}}
  {{end}}
`
}

func action() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if c.Bool("h") == true {
			cli.ShowAppHelp(c)
			return nil
		}
		if c.Bool("v") == true {
			cli.ShowVersion(c)
			return nil
		}
		if mustWarekiToAC(c) == true {
			warekiToAC(c)
			return nil
		} else {
			return acToWareki(c)
		}
	}
}

func mustWarekiToAC(c *cli.Context) bool {
	return c.Int("M") != 0 ||
		c.Int("T") != 0 ||
		c.Int("S") != 0 ||
		c.Int("H") != 0
}

func warekiToAC(c *cli.Context) {
	switch {
	case c.Int("M") != 0:
		fmt.Fprintf(clo.outStream, "%d\n", gengo.MEIJI.ToAC(c.Int("M")))
	case c.Int("T") != 0:
		fmt.Fprintf(clo.outStream, "%d\n", gengo.TAISHO.ToAC(c.Int("T")))
	case c.Int("S") != 0:
		fmt.Fprintf(clo.outStream, "%d\n", gengo.SHOWA.ToAC(c.Int("S")))
	case c.Int("H") != 0:
		fmt.Fprintf(clo.outStream, "%d\n", gengo.HEISEI.ToAC(c.Int("H")))
	}
}

func acToWareki(c *cli.Context) error {
	f := func(t time.Time, kanji bool) (string, error) {
		g, err := gengo.Date(t)
		if err != nil {
			return "", err
		}

		year := g.Convert(t)
		if kanji {
			return g.KanjiName() + strconv.Itoa(year), nil
		} else {
			return g.ShortName() + strconv.Itoa(year), nil
		}
	}

	// 西暦から和暦に変換
	// 引数がないときはシステム日付を和暦に変換
	if c.NArg() == 0 {
		str, err := f(time.Now(), c.Bool("K"))
		if err == nil {
			fmt.Fprintf(clo.outStream, "%s\n", str)
			return nil
		} else {
			return err
		}
	}

	// 引数があるときは日付にパースして和暦に変換
	s := c.Args()[0]
	match, err := regexp.MatchString("^\\d{4}(/\\d{2}(/\\d{2})?)?$", s)
	if err != nil {
		return err
	}
	if match == false {
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

	str, err := f(t, c.Bool("K"))
	if err == nil {
		fmt.Fprintf(clo.outStream, "%s\n", str)
		return nil
	} else {
		return err
	}
}
