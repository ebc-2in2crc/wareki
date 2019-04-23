package main

import (
	"os"
)

const appName = "wareki"
const version = "0.10.0"

func main() {
	cli := &CLO{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
