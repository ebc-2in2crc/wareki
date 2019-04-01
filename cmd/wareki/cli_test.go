package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun_versionFlag(t *testing.T) {
	params := []struct {
		argstr string
	}{
		{argstr: AppName + " --version"},
		{argstr: AppName + " -v"},
	}

	for _, p := range params {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		clo := &CLO{outStream: outStream, errStream: errStream}

		args := strings.Split(p.argstr, " ")
		status := clo.Run(args)
		if status != ExitCodeOK {
			t.Errorf("Run(%s): ExitStatus = %d; want %d", p.argstr, status, ExitCodeOK)
		}

		actual := outStream.String()
		expect := fmt.Sprintf(AppName+" version %s", Version)
		if strings.Contains(actual, expect) == false {
			t.Errorf("Run(%s): Output = %q; want %q", p.argstr, actual, expect)
		}
	}
}

func TestRun_warekiToACFlag(t *testing.T) {
	params := []struct {
		argstr string
		expect string
	}{
		{argstr: AppName + " --meiji 1", expect: "1868"},
		{argstr: AppName + " -M 1", expect: "1868"},
		{argstr: AppName + " --taisho 1", expect: "1912"},
		{argstr: AppName + " -T 1", expect: "1912"},
		{argstr: AppName + " --showa 1", expect: "1926"},
		{argstr: AppName + " -S 1", expect: "1926"},
		{argstr: AppName + " --heisei 1", expect: "1989"},
		{argstr: AppName + " -H 1", expect: "1989"},
		{argstr: AppName + " --reiwa 1", expect: "2019"},
		{argstr: AppName + " -R 1", expect: "2019"},
	}

	for _, p := range params {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		clo := &CLO{outStream: outStream, errStream: errStream}

		args := strings.Split(p.argstr, " ")
		status := clo.Run(args)
		if status != ExitCodeOK {
			t.Errorf("Run(%s): ExitStatus = %d; want %d", p.argstr, status, ExitCodeOK)
		}

		actual := outStream.String()
		expect := p.expect
		if strings.Contains(actual, expect) == false {
			t.Errorf("Run(%s): Output = %q; want %q", p.argstr, actual, expect)
		}
	}
}

func TestRun_acToWareki(t *testing.T) {
	params := []struct {
		argstr string
		expect string
	}{
		{argstr: AppName + " 1868/01/25", expect: "M1"},
		{argstr: AppName + " 1912/07/30", expect: "T1"},
		{argstr: AppName + " 1926/12/25", expect: "S1"},
		{argstr: AppName + " 1989/01/08", expect: "H1"},
		{argstr: AppName + " 2019/05/01", expect: "R1"},
		{argstr: AppName + " 1989/01", expect: "S64"},
		{argstr: AppName + " 1989", expect: "S64"},
	}

	for _, p := range params {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		clo := &CLO{outStream: outStream, errStream: errStream}

		args := strings.Split(p.argstr, " ")
		status := clo.Run(args)
		if status != ExitCodeOK {
			t.Errorf("Run(%s): ExitStatus = %d; want %d", p.argstr, status, ExitCodeOK)
		}

		actual := outStream.String()
		expect := fmt.Sprintf("%s", p.expect)
		if strings.Contains(actual, expect) == false {
			t.Errorf("Run(%s): Output = %q; want %q", p.argstr, actual, expect)
		}
	}
}

func TestRun_err(t *testing.T) {
	params := []struct {
		argstr string
		expect string
	}{
		{argstr: AppName + " 1989-01-01",
			expect: "invalid date format. must specify date: e.g.) 2018 or 2018/01 or 2018/01/01"},
		{argstr: AppName + " 1868/01/24",
			expect: "range error. must specify date: greater than 1868/01/24"},
	}

	for _, p := range params {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		clo := &CLO{outStream: outStream, errStream: errStream}

		args := strings.Split(p.argstr, " ")
		status := clo.Run(args)
		if status != ExitCodeError {
			t.Errorf("Run(%s): ExitStatus = %d; want %d", p.argstr, status, ExitCodeError)
		}

		actual := errStream.String()
		expect := fmt.Sprintf("%s", p.expect)
		if strings.Contains(actual, expect) == false {
			t.Errorf("Run(%s): Output = %q; want %q", p.argstr, actual, expect)
		}
	}
}
