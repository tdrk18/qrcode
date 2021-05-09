package main

import (
	"flag"
	"github.com/skip2/go-qrcode"
	"log"
)

func main() {
	var (
		outputName string
		size       int
	)
	const (
		defaultOutputName = "qr.png"
		usageOutputName   = "Specify the output file name"
		defaultSize       = 256
		usageSize         = "Specify the frame size of output file"
	)
	flag.StringVar(&outputName, "output", defaultOutputName, usageOutputName)
	flag.StringVar(&outputName, "o", defaultOutputName, usageOutputName+" (shorthand)")
	flag.IntVar(&size, "size", defaultSize, usageSize)
	flag.IntVar(&size, "s", defaultSize, usageSize+" (shorthand)")
	flag.Parse()
	err := qrcode.WriteFile("https://google.com", qrcode.Medium, size, outputName)
	if err != nil {
		log.Fatal(err)
	}
}
