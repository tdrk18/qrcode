package main

import (
	"github.com/skip2/go-qrcode"
	"log"
)

func main() {
	err := qrcode.WriteFile("https://google.com", qrcode.Medium, 256, "qr.png")
	if err != nil {
		log.Fatal(err)
	}
}
