package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const GridSize int = 256
const ExitError int = 1
const ExitSuccess int = 0

func usage() {
	fmt.Println("./binviz <FILENAME> [<PNG_IMAGE_FILENAME>]")
	fmt.Println("If <IMAGE_FILENAME> is ommitted, <FILENAME>.png will be used")
}

func buildGrid(data []byte) [GridSize][GridSize]uint64 {
	grid := [GridSize][GridSize]uint64{}
	for i := 0; i < len(data)-1; i++ {
		grid[data[i+1]][data[i]]++
	}
	return grid
}

func buildImage(grid [GridSize][GridSize]uint64) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, GridSize, GridSize))
	var pixColor color.Color
	for y, r := range grid {
		for x, v := range r {
			if v > 0 {
				pixColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255 - uint8(v)}
			} else {
				pixColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
			}
			img.Set(x, y, pixColor)
		}
	}
	return img
}

func writeImage(img image.Image, fn string) {
	f, err := os.Create(fn)
	if err != nil {
		fmt.Printf("Cannot create image file [%s]\n", fn)
		os.Exit(ExitError)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		fmt.Printf("Cannot encode to image file [%s]\n", fn)
		fmt.Println(err)
		os.Exit(ExitError)
	}
	if err := f.Close(); err != nil {
		fmt.Printf("Could not close file [%s]\n", fn)
		fmt.Println(err)
		os.Exit(ExitError)
	}
}

func parseCommandLine() (string, string) {
	var fnIn, fnOut string
	if len(os.Args) < 2 {
		usage()
		os.Exit(ExitError)
	}
	fnIn = os.Args[1]
	if len(os.Args) > 2 {
		fnOut = os.Args[2]
	} else {
		fnOut = fnIn + ".png"
	}
	return fnIn, fnOut
}

func main() {
	fnIn, fnOut := parseCommandLine()
	data, err := os.ReadFile(fnIn)
	if err != nil {
		fmt.Printf("Could not read file [%s]\n", fnIn)
		os.Exit(ExitError)
	}
	writeImage(buildImage(buildGrid(data)), fnOut)
	os.Exit(ExitSuccess)
}
