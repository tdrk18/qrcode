package main

import (
	"github.com/skip2/go-qrcode"
	"log"
)

func createQRCode(content string, outputName string, size int, level qrcode.RecoveryLevel) {
	err := qrcode.WriteFile(content, level, size, outputName)
	if err != nil {
		log.Fatal(err)
	}
}
