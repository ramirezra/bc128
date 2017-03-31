package bc

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

// Combine exported to main
func Combine(i int, data []Data, lblPart image.Image, bcPart image.Image, bcSerial image.Image, lblSerial image.Image) {
	// Combine template and barcode
	template, err := os.Open("sample.png")
	if err != nil {
		fmt.Println(err)
	}
	// lblPart, err := os.Open("output/" + data[i].PartNo + "-Label.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// bcPart, err := os.Open("output/" + data[i].PartNo + ".png")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// lblSerial, err := os.Open("output/" + data[i].SerialNo + "-Label.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// bcSerial, err := os.Open("output/" + data[i].SerialNo + ".png")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	label, _, err := image.Decode(template)
	if err != nil {
		fmt.Println(err)
	}
	// partLabel, _, err := image.Decode(lblPart)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// partBarCode, _, err := image.Decode(bcPart)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// serialLabel, _, err := image.Decode(lblSerial)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// serialBarCode, _, err := image.Decode(bcSerial)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// create a new image
	combined := image.NewRGBA(label.Bounds())
	draw.Draw(combined, label.Bounds(), label, image.Point{0, 0}, draw.Src)

	draw.Draw(combined, label.Bounds(), lblSerial, image.Point{-375, -1400}, draw.Src)
	draw.Draw(combined, label.Bounds(), bcSerial, image.Point{-375, -1475}, draw.Src)
	draw.Draw(combined, label.Bounds(), lblPart, image.Point{-720, -1725}, draw.Src)
	draw.Draw(combined, label.Bounds(), bcPart, image.Point{-715, -1800}, draw.Src)

	// Export image
	output, err := os.Create("labels/" + data[i].PartNo + "-" + data[i].SerialNo + ".png")
	if err != nil {
		fmt.Println(err)
	}
	png.Encode(output, combined)
}
