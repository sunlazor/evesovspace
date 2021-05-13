package main

import (
	"archive/zip"
	"evesovspace/internal/structs"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	const width, height = 1000, 1000

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{R: 0, G: 0, B: 0, A: 255})
		}
	}

	// for y := 0; y < height; y++ {
	// 	for x := 0; x < width; x++ {
	// 		img.Set(x, y, color.NRGBA{
	// 			R: uint8((x + y) & 255),
	// 			G: uint8((x + y) << 1 & 255),
	// 			B: uint8((x + y) << 2 & 255),
	// 			A: 255,
	// 		})
	// 	}
	// }

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// archive()
}

func archive() {
	zipR, err := zip.OpenReader("./data/sde.zip")
	switch err := err.(type) {
	case *fs.PathError:
		fmt.Println("SDE file wasn't found in data folder. Exit...")
		os.Exit(0)
	case nil:
	default:
		panic(err)
	}

	var region structs.Region
	var constellation structs.Constellation
	var solarsystem structs.SolarSystem

	for _, file := range zipR.File {
		if strings.Contains(file.Name, ".staticdata") && strings.Contains(file.Name, "sde/fsd/universe/eve/") {
			// fmt.Println(file.Name)

			if strings.Contains(file.Name, "region.staticdata") {
				rc, _ := file.Open()
				err = yaml.NewDecoder(rc).Decode(&region)
				if err != nil {
					fmt.Printf("Unmarshall error #%v", err)
					continue
				}
				fmt.Printf("file.Name: %v %+v\n", file.Name, region)
			}

			if strings.Contains(file.Name, "constellation.staticdata") {
				rc, _ := file.Open()
				err = yaml.NewDecoder(rc).Decode(&constellation)
				if err != nil {
					fmt.Printf("Unmarshall error #%v", err)
					continue
				}
				fmt.Printf("file.Name: %v %+v\n", file.Name, constellation)
			}

			if strings.Contains(file.Name, "solarsystem.staticdata") {
				rc, _ := file.Open()
				err = yaml.NewDecoder(rc).Decode(&solarsystem)
				if err != nil {
					fmt.Printf("Unmarshall error #%v", err)
					continue
				}
				fmt.Printf("file.Name: %v %+v\n", file.Name, solarsystem)
			}
		}
	}
}
