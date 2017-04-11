// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	// Decode the JPEG data. If reading from file, create a reader with
	//
	// reader, err := os.Open("testdata/video-001.q50.420.jpeg")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer reader.Close()

	fileToBeUploaded := "images/menu1.jpg"
	file, err := os.Open(fileToBeUploaded)

	if err != nil {
		fmt.Println("ohno", err)
		os.Exit(1)
	}

	defer file.Close()

	// fileInfo, _ := file.Stat()
	// var size int64 = fileInfo.Size()
	// imageBytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	// _, err = buffer.Read(imageBytes)

	fmt.Println(*file, *buffer)

	reader := base64.NewDecoder(base64.StdEncoding, buffer)
	m, st, err := image.Decode(reader)
	if err != nil {
		fmt.Println("waaah", st, m, err)
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
	//
	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}

	// Print the results.
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}
}
