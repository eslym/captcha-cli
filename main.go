package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/dchest/captcha"
	"image/color"
	"os"
)

func main() {
	cli := flag.NewFlagSet("captcha", flag.ContinueOnError)

	var (
		dataUrl bool
		width   int
		height  int
	)

	cli.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s [options] digits\n", os.Args[0])
		_, _ = fmt.Fprintf(os.Stderr, "Options:\n")
		cli.PrintDefaults()
	}

	cli.IntVar(&width, "width", 240, "width of the image")
	cli.IntVar(&height, "height", 80, "height of the image")
	cli.BoolVar(&dataUrl, "data-url", false, "output data url")

	err := cli.Parse(os.Args[1:])

	if err != nil {
		cli.Usage()
		os.Exit(1)
	}

	if cli.NArg() != 1 {
		cli.Usage()
		os.Exit(1)
	}

	digits := []byte(cli.Arg(0))

	for i, d := range digits {
		if d < '0' || d > '9' {
			_, _ = fmt.Fprintf(os.Stderr, "error: invalid digit %q\n", d)
			os.Exit(1)
		}
		digits[i] = d - '0'
	}

	img := captcha.NewImage("", digits, width, height)

	img.Paletted.Palette[0] = color.White

	if dataUrl {
		w := &bytesWriter{}
		_, err = img.WriteTo(w)
		if err != nil {
			panic(err)
		}
		fmt.Printf("data:image/png;base64,")
		fmt.Printf("%s", base64.StdEncoding.EncodeToString(w.bytes))
		return
	}

	_, err = img.WriteTo(os.Stdout)
	if err != nil {
		panic(err)
	}
}

type bytesWriter struct {
	bytes []byte
}

func (w *bytesWriter) Write(p []byte) (n int, err error) {
	w.bytes = append(w.bytes, p...)
	return len(p), nil
}
