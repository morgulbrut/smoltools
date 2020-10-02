package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	fn := "hello"
	text := "hellogo"

	if len(os.Args) == 2 {
		text = os.Args[1]
	}
	if len(os.Args) == 3 {
		text = os.Args[1]
		fn = os.Args[2]
	}
	generateQR(text, fn)
}

func generateQR(s string, n string) {
	// Create the barcode
	qrCode, _ := qr.Encode(s, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create(n + ".png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
