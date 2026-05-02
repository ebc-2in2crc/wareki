package main

import (
	"context"
	"os"
)

const appName = "wareki"
const version = "1.2.1"

func main() {
	cli := &CLO{inputStream: os.Stdin, outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(context.Background(), os.Args))
}
