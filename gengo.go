package wareki

import (
	"errors"
	"time"
)

// Gengo represents Japanese era.
type Gengo struct {
	name      string
	shortName string
	kanjiName string
	from      time.Time
	to        time.Time
}

// MEIJI returns a Gengo that represents MEIJI era.
func MEIJI() *Gengo {
	return meiji
}

// TAISHO returns a Gengo that represents TAISHO era.
func TAISHO() *Gengo {
	return taisho
}

// SHOWA returns a Gengo that represents SHOWA era.
func SHOWA() *Gengo {
	return showa
}

// HEISEI returns a Gengo that represents HEISEI era.
func HEISEI() *Gengo {
	return heisei
}

// REIWA returns a Gengo that represents REIWA era.
func REIWA() *Gengo {
	return reiwa
}

// Now returns a current Gengo.
func Now() (*Gengo, error) {
	return Date(time.Now())
}

// Date returns a Gengo that always uses the given time.Time.
func Date(t time.Time) (*Gengo, error) {
	for _, g := range values {
		if g.Between(t) {
			return g, nil
		}
	}
	return nil, errors.New("range error. must specify date: greater than 1868/01/24")
}

// Name returns a romanized representation of the Gengo.
func (g *Gengo) Name() string {
	return g.name
}

// ShortName returns a short romanized representation of the Gengo.
func (g *Gengo) ShortName() string {
	return g.shortName
}

// KanjiName returns a kanji representation of the Gengo.
func (g *Gengo) KanjiName() string {
	return g.kanjiName
}

// Between reports whether t is between Gengo.
func (g *Gengo) Between(t time.Time) bool {
	return t.After(g.from) && t.Before(g.to)
}

// Convert converts western calendar into Japanese calendar.
func (g *Gengo) Convert(t time.Time) int {
	return t.Year() - g.from.Year() + 1
}

// ToAC converts Japanese calendar into western calendar.
func (g *Gengo) ToAC(year int) int {
	return g.from.Year() + year - 1
}

// Values returns an array That includes all Gengo.
func Values() []*Gengo {
	return values
}

var values = []*Gengo{meiji, taisho, showa, heisei, reiwa}

var jst = time.FixedZone("JST", 9*60*60)

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
