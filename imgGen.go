package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

var width = 32
var height = 32

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Incorrect number of arguments")
	} else {
		data, _ := ioutil.ReadAll(os.Stdin)
		name := []rune(strings.Trim(strings.ToLower(string(data)), "\n"))

		image := image.NewRGBA(image.Rect(0, 0, width, height))

		color1 := color.RGBA{
			toPro(float32(name[0])),
			toPro(float32(name[1])),
			toPro(float32(name[2])),
			255,
		}

		color2 := color.RGBA{
			255 - color1.R,
			255 - color1.G,
			255 - color1.B,
			255,
		}

		var imageSlice [32][32]int

		var count = 0
		var numCycle = 0

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				if numCycle == 0 {
					imageSlice[i][j] = int(name[count])
					count++
				} else {
					imageSlice[i][j] = int(name[count])
					count += 2
				}
				if count >= len(name) {
					if numCycle == 0 {
						count = 0
						numCycle++
					} else if numCycle == 1 {
						count = 1
						numCycle++
					} else {
						count = 0
						numCycle = 0
					}

				}
			}
		}
		if len(name)%2 != 0 {
			for i := 0; i < width/2; i++ {
				for j := 0; j < height; j++ {
					imageSlice[height-1-i][j] = imageSlice[i][j]
				}
			}
		}

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				if imageSlice[i][j]%2 == 0 {
					image.SetRGBA(i, j, color1)
				} else {
					image.SetRGBA(i, j, color2)
				}
			}
		}

		output, err := os.Create("test.png")
		if err != nil {
			fmt.Println("An error occured")
		}

		png.Encode(output, image)
		output.Close()
	}
}

func toPro(num float32) uint8 {
	num -= 97
	return uint8((num / 26) * 255)
}
