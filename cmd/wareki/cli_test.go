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
		{argstr: appName + " --version"},
		{argstr: appName + " -v"},
	}

	for _, p := range params {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		clo := &CLO{outStream: outStream, errStream: errStream}

		args := strings.Split(p.argstr, " ")
		status := clo.Run(args)
		if status != exitCodeOK {
			t.Errorf("Run(%s): ExitStatus = %d; want %d", p.argstr, status, exitCodeOK)
		}

		actual := outStream.String()
		expect := fmt.Sprintf(appName+" version %s", version)
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
		{argstr: appName + " --meiji 1", expect: "1868"},
		{argstr: appName + " -M 1", expect: "1868"},
		{argstr: appName + " --taisho 1", expect: "1912"},
		{argstr: appName + " -T 1", expect: "1912"},
		{argstr: appName + " --showa 1", expect: "1926"},
		{argstr: appName + " -S 1", expect: "1926"},
		{argstr: appName + " --heisei 1", expect: "1989"},
		{argstr: appName + " -H 1", expect: "1989"},
		{argstr: appName + " --reiwa 1", expect: "2019"},
		{argstr: appName + " -R 1", expect: "2019"},
	}

	for _, p := range params {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		clo := &CLO{outStream: outStream, errStream: errStream}

		args := strings.Split(p.argstr, " ")
		status := clo.Run(args)
		if status != exitCodeOK {
			t.Errorf("Run(%s): ExitStatus = %d; want %d", p.argstr, status, exitCodeOK)
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
		{argstr: appName + " 1868/01/25", expect: "M1"},
		{argstr: appName + " 1912/07/30", expect: "T1"},
		{argstr: appName + " 1926/12/25", expect: "S1"},
		{argstr: appName + " 1989/01/08", expect: "H1"},
		{argstr: appName + " 2019/05/01", expect: "R1"},
		{argstr: appName + " 1989/01", expect: "S64"},
		{argstr: appName + " 1989", expect: "S64"},
	}

	for _, p := range params {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		clo := &CLO{outStream: outStream, errStream: errStream}

		args := strings.Split(p.argstr, " ")
		status := clo.Run(args)
		if status != exitCodeOK {
			t.Errorf("Run(%s): ExitStatus = %d; want %d", p.argstr, status, exitCodeOK)
		}

		actual := outStream.String()
		if strings.Contains(actual, p.expect) == false {
			t.Errorf("Run(%s): Output = %q; want %q", p.argstr, actual, p.expect)
		}
	}
}

func TestRun_err(t *testing.T) {
	params := []struct {
		argstr string
		expect string
	}{
		{argstr: appName + " 1989-01-01",
			expect: "invalid date format. must specify date: e.g.) 2018 or 2018/01 or 2018/01/01"},
		{argstr: appName + " 1868/01/24",
			expect: "range error. must specify date: greater than 1868/01/24"},
	}

	for _, p := range params {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		clo := &CLO{outStream: outStream, errStream: errStream}

		args := strings.Split(p.argstr, " ")
		status := clo.Run(args)
		if status != exitCodeError {
			t.Errorf("Run(%s): ExitStatus = %d; want %d", p.argstr, status, exitCodeError)
		}

		actual := errStream.String()
		if strings.Contains(actual, p.expect) == false {
			t.Errorf("Run(%s): Output = %q; want %q", p.argstr, actual, p.expect)
		}
	}
}
