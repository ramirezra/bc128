package bc

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

// Combine exported to main
func Combine(i int, data []Data) {
	// Combine template and barcode
	fmt.Println(data[i].PartNo)
	template, err := os.Open("sample.png")
	if err != nil {
		fmt.Println(err)
	}
	bcPart, err := os.Open("output/" + data[i].PartNo + ".png")
	if err != nil {
		fmt.Println(err)
	}
	bcSerial, err := os.Open("output/" + data[i].SerialNo + ".png")
	if err != nil {
		fmt.Println(err)
	}
	label, _, err := image.Decode(template)
	if err != nil {
		fmt.Println(err)
	}
	partBarCode, _, err := image.Decode(bcPart)
	if err != nil {
		fmt.Println(err)
	}
	serialBarCode, _, err := image.Decode(bcSerial)
	if err != nil {
		fmt.Println(err)
	}

	// create a new image
	combined := image.NewRGBA(label.Bounds())
	draw.Draw(combined, label.Bounds(), label, image.Point{0, 0}, draw.Src)
	// draw.Draw(rgba, r2, bayercode, image.Point{0, 0}, draw.Src)
	draw.Draw(combined, label.Bounds(), serialBarCode, image.Point{-400, -1475}, draw.Src)
	draw.Draw(combined, label.Bounds(), partBarCode, image.Point{-750, -1800}, draw.Src)

	// Export image
	output, err := os.Create("labels/" + data[i].SerialNo + ".png")
	if err != nil {
		fmt.Println(err)
	}
	png.Encode(output, combined)
}
