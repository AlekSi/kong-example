package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

func main() {
	var cli struct {
		Flag string `default:"foo" env:"FLAG"`
	}

	kong.Parse(&cli)
	fmt.Printf("flag = %q\n", cli.Flag)
}
