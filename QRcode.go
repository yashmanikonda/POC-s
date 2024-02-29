package main

import (
	"fmt"
	"log"
	"github.com/skip2/go-qrcode"
)

func generateQRCode(deviceID, filename string) {
	err := qrcode.WriteFile(deviceID, qrcode.Low, 256, filename)
	if err != nil {
		log.Fatal("Error generating QR code:", err)
	}
	fmt.Printf("QR code for Device ID '%s' has been generated and saved as '%s'.\n", deviceID, filename)
}

func main() {
	// Replace 'your_device_id_here' with the actual device ID
	deviceID := "your_device_id_here"

	// Replace 'output_qr_code.png' with the desired output filename
	filename := "output_qr_code.png"

	// Generate the QR code
	generateQRCode(deviceID, filename)
}
