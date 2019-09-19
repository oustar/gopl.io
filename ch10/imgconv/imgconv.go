//
// imgxch [-f="jpeg|png|gif"] [infile]
//	-o: the name of output image file(default: stdout)
//	-f: format of output image file (default: jpeg)
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

var format = flag.String("f", "jpeg", "format of output image")

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		if err := conv(os.Stdin, *format, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "imgconv:%v \n", err)
			os.Exit(1)
		}
		return
	}

	for _, arg := range flag.Args() {

		log.Printf("input file name: %s\n", arg)
		fin, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "imagconv: %v \n", err)
			continue
		}
		fout, err := os.Create(arg + "." + *format)
		if err != nil {
			fin.Close()
			fmt.Fprintf(os.Stderr, "imagconv: %v \n", err)
			continue
		}

		err = conv(bufio.NewReader(fin), *format, bufio.NewWriter(fout))
		if err != nil {
			fin.Close()
			fout.Close()
			fmt.Fprintf(os.Stderr, "imgconv: %v\n", err)
			continue
		}
		fin.Close()
		fout.Close()
	}

	log.Printf("well done\n")
}

func conv(in io.Reader, f string, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	log.Printf("Input format = %s\n", kind)
	log.Printf("Output format = %s\n", f)

	switch f {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)
	}
	return nil
}
