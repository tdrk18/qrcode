package main

import (
	"flag"
	"github.com/skip2/go-qrcode"
	"log"
)

func main() {
	var outputName string
	const (
		defaultOutputName = "qr.png"
		usageOutputName   = "Specify the output file name"
	)
	flag.StringVar(&outputName, "output", defaultOutputName, usageOutputName)
	flag.StringVar(&outputName, "o", defaultOutputName, usageOutputName+" (shorthand)")
	flag.Parse()
	err := qrcode.WriteFile("https://google.com", qrcode.Medium, 256, outputName)
	if err != nil {
		log.Fatal(err)
	}
}
