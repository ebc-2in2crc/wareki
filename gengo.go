package wareki

import (
	"errors"
	"time"
)

// Gengo 日本の元号を表す
type Gengo struct {
	name      string
	shortName string
	kanjiName string
	from      time.Time
	to        time.Time
}

// MEIJI 明治を表す Gengo を取得する
func MEIJI() *Gengo {
	return meiji
}

// TAISHO 大正を表す Gengo を取得する
func TAISHO() *Gengo {
	return taisho
}

// SHOWA 昭和を表す Gengo を取得する
func SHOWA() *Gengo {
	return showa
}

// HEISEI 平成を表す Gengo を取得する
func HEISEI() *Gengo {
	return heisei
}

// REIWA 令和を表す Gengo を取得する
func REIWA() *Gengo {
	return reiwa
}

// Now 現在日付の Gengo を取得する
func Now() (*Gengo, error) {
	return Date(time.Now())
}

// Date datetime に対応する Gengo を取得する
func Date(datetime time.Time) (*Gengo, error) {
	for _, g := range values {
		if g.Between(datetime) {
			return g, nil
		}
	}
	return nil, errors.New("range error. must specify date: greater than 1868/01/24")
}

// Name 元号のローマ字表現
func (g *Gengo) Name() string {
	return g.name
}

// ShortName 元号の短縮形のローマ字表現
func (g *Gengo) ShortName() string {
	return g.shortName
}

// KanjiName 元号の漢字表現
func (g *Gengo) KanjiName() string {
	return g.kanjiName
}

// Between 元号の範囲かどうか判定する。
// t が元号の範囲内のとき true を返す。
func (g *Gengo) Between(t time.Time) bool {
	return t.After(g.from) && t.Before(g.to)
}

// Convert 西暦を和暦にする
func (g *Gengo) Convert(t time.Time) int {
	return t.Year() - g.from.Year() + 1
}

// ToAC 和暦を西暦にする
func (g *Gengo) ToAC(year int) int {
	return g.from.Year() + year - 1
}

// Values Gengo の配列を返す
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
