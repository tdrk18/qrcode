package main

import (
	"github.com/skip2/go-qrcode"
	"log"
)

func createQRCode(content string, outputName string, size int) {
	err := qrcode.WriteFile(content, qrcode.Medium, size, outputName)
	if err != nil {
		log.Fatal(err)
	}
}
