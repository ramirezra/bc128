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
func Encode128(data string) {
	code, err := code128.Encode(data)
	if err != nil {
		fmt.Printf("string %s cannot be encoded\n", data)
		os.Exit(1)
	}
	barcodeImage, err := barcode.Scale(code, 710, 150)
	if err != nil {
		log.Fatal(err)
	}
	writeImage("output/"+data+".png", barcodeImage)
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
