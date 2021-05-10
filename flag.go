package main

import (
	"flag"
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"
)

func parseFlags() (string, string, int, qrcode.RecoveryLevel) {
	var (
		content    string
		outputName string
		level      string
		size       int
	)
	const (
		defaultOutputName = "qr.png"
		usageOutputName   = "Specify the output file name"
		defaultSize       = 256
		usageSize         = "Specify the frame size of output file"
		defaultContent    = ""
		usageContent      = "Specify the content for QRCode"
		defaultLevel      = "medium"
		usageLevel        = "Specify the recovery level"
	)
	flag.StringVar(&outputName, "output", defaultOutputName, usageOutputName)
	flag.StringVar(&outputName, "o", defaultOutputName, usageOutputName+" (shorthand)")
	flag.IntVar(&size, "size", defaultSize, usageSize)
	flag.IntVar(&size, "s", defaultSize, usageSize+" (shorthand)")
	flag.StringVar(&content, "content", defaultContent, usageContent)
	flag.StringVar(&content, "c", defaultContent, usageContent+" (shorthand)")
	flag.StringVar(&level, "level", defaultLevel, usageLevel)
	flag.StringVar(&level, "l", defaultLevel, usageLevel+" (shorthand)")
	flag.Parse()

	if content == defaultContent {
		fmt.Println("Please set -content <QRCode content>")
		os.Exit(1)
	}

	return content, outputName, size, validateLevel(level)
}

func validateLevel(level string) qrcode.RecoveryLevel {
	switch level {
	case "low":
		return qrcode.Low
	case "middle":
		return qrcode.Medium
	case "high":
		return qrcode.High
	case "highest":
		return qrcode.Highest
	default:
		return qrcode.Medium
	}
}
