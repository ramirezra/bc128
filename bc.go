package bc

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/fogleman/gg"
)

// Data exported to main function
type Data struct {
	PartNo   string
	SerialNo string
}

// GetData exported to main function
func GetData() []Data {
	file, err := os.Open("sampledata.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	var data []Data
	for _, item := range records {
		record := Data{PartNo: item[0], SerialNo: item[1]}
		data = append(data, record)
	}

	return data
}

// Encode128 exported to main
func Encode128(data string) *os.File {
	code, err := code128.Encode(data)
	if err != nil {
		fmt.Printf("string %s cannot be encoded\n", data)
		os.Exit(1)
	}
	barcodeImage, err := barcode.Scale(code, 710, 150)
	if err != nil {
		log.Fatal(err)
	}
	// writeImage("output/"+data+".png", barcodeImage)
	file, err := os.Create("output/" + data + ".png")
	if err != nil {
		log.Fatalln(err)
	}
	png.Encode(file, barcodeImage)
	return file
}

// EncodeLabel exported to main
func EncodeLabel(data string) *os.File {
	const W = 700
	const H = 75
	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	if err := dc.LoadFontFace("/Windows/Fonts/Arial.ttf", 72); err != nil {
		log.Fatalln(err)
		panic(err)
	}
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(data, W/2, H/2, 0.5, 0.5)

	// dc.SavePNG("output/" + data + "-Label.png")
	// Return image file
	file, err := os.Create("output/" + data + ".png")
	if err != nil {
		log.Fatalln(err)
	}
	err = dc.EncodePNG(file)
	if err != nil {
		log.Fatalln(err)
	}
	return file
}

func writeImage(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatalln(err)
	}
	file.Close()
}
