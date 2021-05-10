package main

func main() {
	content, outputName, size, level := parseFlags()
	createQRCode(content, outputName, size, level)
}
