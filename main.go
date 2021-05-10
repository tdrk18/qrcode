package main

func main() {
	content, outputName, size := parseFlags()
	createQRCode(content, outputName, size)
}
