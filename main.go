package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var fNoColor = flag.Bool("nocolor", false, "turn off color output")

func main() {
	flag.Parse()

	var source *os.File
	if fn := flag.Arg(0); fn != "" {
		var err error
		source, err = os.Open(fn)
		if err != nil {
			fmt.Println("Couldn't open input: ", err)
			return
		}
	} else {
		source = os.Stdin
	}

	in := bufio.NewScanner(source)
	cw := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: *fNoColor}

	for in.Scan() {
		line := in.Bytes()
		cw.Write(line)

		if err := in.Err(); err != nil {
			fmt.Println("Err: ", err)
			return
		}
	}
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [file]\n", os.Args[0])
		flag.PrintDefaults()
	}
}
