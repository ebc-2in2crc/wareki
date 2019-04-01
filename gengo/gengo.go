package gengo

import (
	"errors"
	"time"
)

type Gengo struct {
	name      string
	shortName string
	kanjiName string
	from      time.Time
	to        time.Time
}

func MEIJI() *Gengo {
	return meiji
}

func TAISHO() *Gengo {
	return taisho
}

func SHOWA() *Gengo {
	return showa
}

func HEISEI() *Gengo {
	return heisei
}

func REIWA() *Gengo {
	return reiwa
}

func Now() (*Gengo, error) {
	return Date(time.Now())
}

func Date(datetime time.Time) (*Gengo, error) {
	for _, g := range values {
		if g.Between(datetime) {
			return g, nil
		}
	}
	return nil, errors.New("range error. must specify date: greater than 1868/01/24")
}

func (g *Gengo) Name() string {
	return g.name
}

func (g *Gengo) ShortName() string {
	return g.shortName
}

func (g *Gengo) KanjiName() string {
	return g.kanjiName
}
func (g *Gengo) Between(t time.Time) bool {
	return t.After(g.from) && t.Before(g.to)
}

func (g *Gengo) Convert(t time.Time) int {
	return t.Year() - g.from.Year() + 1
}

func (g *Gengo) ToAC(year int) int {
	return g.from.Year() + year - 1
}

func Values() []*Gengo {
	return values
}

var values = []*Gengo{meiji, taisho, showa, heisei, reiwa}

var jst, _ = time.LoadLocation("Asia/Tokyo")

var (
	meiji = &Gengo{
		"meiji",
		"M",
		"明治",
		time.Date(1868, 1, 24, 23, 59, 59, 0, jst),
		time.Date(1912, 7, 30, 0, 0, 0, 0, jst),
	}
	taisho = &Gengo{
		"taisho",
		"T",
		"大正",
		time.Date(1912, 7, 29, 23, 59, 59, 0, jst),
		time.Date(1926, 12, 25, 00, 00, 00, 0, jst),
	}
	showa = &Gengo{
		"showa",
		"S",
		"昭和",
		time.Date(1926, 12, 24, 23, 59, 59, 0, jst),
		time.Date(1989, 1, 8, 0, 0, 0, 0, jst),
	}
	heisei = &Gengo{
		"heisei",
		"H",
		"平成",
		time.Date(1989, 1, 7, 23, 59, 59, 0, jst),
		time.Date(2019, 5, 1, 0, 0, 0, 0, jst),
	}
	reiwa = &Gengo{
		"reiwa",
		"R",
		"令和",
		time.Date(2019, 4, 30, 23, 59, 59, 0, jst),
		time.Date(9999, 12, 31, 23, 59, 59, 0, jst),
	}
)
