package main

import (
	"os"
)

const AppName = "wareki"
const Version = "0.9.0"

func main() {
	cli := &CLO{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
