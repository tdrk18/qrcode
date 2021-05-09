package main

import (
	"github.com/skip2/go-qrcode"
	"log"
)

func main() {
	qr, err := qrcode.New("https://google.com", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}
	println(qr.ToString(false))
}
