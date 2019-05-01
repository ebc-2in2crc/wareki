package wareki

import (
	"errors"
	"testing"
	"time"
)

func TestGengoInstance(t *testing.T) {
	params := []struct {
		g         *Gengo
		name      string
		shortName string
		kanjiName string
	}{
		{g: meiji, name: "meiji", shortName: "M", kanjiName: "明治"},
		{g: taisho, name: "taisho", shortName: "T", kanjiName: "大正"},
		{g: showa, name: "showa", shortName: "S", kanjiName: "昭和"},
		{g: heisei, name: "heisei", shortName: "H", kanjiName: "平成"},
		{g: reiwa, name: "reiwa", shortName: "R", kanjiName: "令和"},
	}

	for _, p := range params {
		actual := p.g

		if actual.Name() != p.name {
			t.Errorf("Name() = %s; want %s", actual.Name(), p.name)
		}

		if actual.ShortName() != p.shortName {
			t.Errorf("ShortName() = %s; want %s", actual.ShortName(), p.shortName)
		}

		if actual.KanjiName() != p.kanjiName {
			t.Errorf("KanjiName() = %s; want %s", actual.KanjiName(), p.kanjiName)
		}
	}
}

func TestBetween(t *testing.T) {
	now := time.Now()

	if meiji.Between(now) == true ||
		taisho.Between(now) == true ||
		showa.Between(now) == true ||
		heisei.Between(now) == true {
		t.Errorf("showa.Between(%v) = true; want false", now)
	}

	if reiwa.Between(now) == false {
		t.Errorf("heisei.Between(%v) = false; want true", now)
	}
}

func TestValues(t *testing.T) {
	actual := Values()
	expect := []*Gengo{meiji, taisho, showa, heisei, reiwa}

	if len(actual) != len(expect) {
		t.Errorf("len(Values()) = %d; want %d", len(actual), len(expect))
	}

	for i, actualGengo := range actual {
		expectGengo := expect[i]
		if actualGengo != expectGengo {
			t.Errorf("got %+v; want %+v", actualGengo, expectGengo)
		}
	}
}

func TestDate(t *testing.T) {
	params := []struct {
		t      time.Time
		expect *Gengo
		err    error
	}{
		{t: time.Date(1868, 1, 24, 0, 0, 0, 0, jst), expect: nil,
			err: errors.New("range error. must specify date: greater than 1868/01/24")},
		{t: time.Date(1868, 1, 25, 23, 59, 59, 0, jst), expect: meiji, err: nil},
		{t: time.Date(1912, 7, 29, 23, 59, 59, 0, jst), expect: meiji, err: nil},
		{t: time.Date(1912, 7, 30, 0, 0, 0, 0, jst), expect: taisho, err: nil},
		{t: time.Date(1926, 12, 24, 23, 59, 59, 0, jst), expect: taisho, err: nil},
		{t: time.Date(1926, 12, 25, 0, 0, 0, 0, jst), expect: showa, err: nil},
		{t: time.Date(1989, 1, 7, 23, 59, 59, 0, jst), expect: showa, err: nil},
		{t: time.Date(1989, 1, 8, 0, 0, 0, 0, jst), expect: heisei, err: nil},
		{t: time.Date(2019, 4, 30, 23, 59, 59, 0, jst), expect: heisei, err: nil},
		{t: time.Date(2019, 5, 1, 0, 0, 0, 0, jst), expect: reiwa, err: nil},
		{t: time.Date(9999, 12, 31, 0, 0, 0, 0, jst), expect: reiwa, err: nil},
	}

	for _, p := range params {
		actual, err := Date(p.t)
		if err != nil && err.Error() != p.err.Error() {
			t.Errorf("Date(%v) failed (%+v); want %+v", p.t, err, p.err)
		}

		expect := p.expect
		if actual != expect {
			t.Errorf("Date(%v) = %+v; want %+v", p.t, actual, expect)
		}
	}
}

func TestNow(t *testing.T) {
	actual, _ := Now()
	expect := reiwa

	if actual != expect {
		t.Errorf("Now() = %+v; want %+v", actual, expect)
	}
}

func TestConvert(t *testing.T) {
	params := []struct {
		t      time.Time
		g      *Gengo
		expect int
	}{
		{t: time.Date(1868, 1, 25, 23, 59, 59, 0, jst), g: meiji, expect: 1},
		{t: time.Date(1912, 7, 29, 23, 59, 59, 0, jst), g: meiji, expect: 45},
		{t: time.Date(1912, 7, 30, 0, 0, 0, 0, jst), g: taisho, expect: 1},
		{t: time.Date(1926, 12, 24, 23, 59, 59, 0, jst), g: taisho, expect: 15},
		{t: time.Date(1926, 12, 25, 0, 0, 0, 0, jst), g: showa, expect: 1},
		{t: time.Date(1989, 1, 7, 23, 59, 59, 0, jst), g: showa, expect: 64},
		{t: time.Date(1989, 1, 8, 0, 0, 0, 0, jst), g: heisei, expect: 1},
		{t: time.Date(2019, 4, 30, 23, 59, 59, 0, jst), g: heisei, expect: 31},
		{t: time.Date(2019, 5, 1, 0, 0, 0, 0, jst), g: reiwa, expect: 1},
		{t: time.Date(9999, 12, 31, 0, 0, 0, 0, jst), g: reiwa, expect: 7981},
	}

	for _, p := range params {
		g := p.g
		actual := g.Convert(p.t)
		expect := p.expect

		if actual != expect {

			t.Errorf("Convert(%v) = %+v; want %+v", p.t, actual, expect)
		}
	}
}

func TestToAC(t *testing.T) {
	params := []struct {
		g      *Gengo
		year   int
		expect int
	}{
		{g: meiji, year: 1, expect: 1868},
		{g: meiji, year: 45, expect: 1912},
		{g: taisho, year: 1, expect: 1912},
		{g: taisho, year: 15, expect: 1926},
		{g: showa, year: 1, expect: 1926},
		{g: showa, year: 64, expect: 1989},
		{g: heisei, year: 1, expect: 1989},
		{g: heisei, year: 30, expect: 2018},
		{g: reiwa, year: 1, expect: 2019},
		{g: reiwa, year: 2, expect: 2020},
	}

	for _, p := range params {
		g := p.g
		actual := g.ToAC(p.year)
		expect := p.expect

		if actual != expect {
			t.Errorf("ToAC(%d) = %d; want %d", p.year, actual, expect)
		}
	}
}
