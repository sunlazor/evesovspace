package main

import (
	"archive/zip"
	"evesovspace/internal/structs"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	archive()
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

	// var region structs.Region
	// var constellation structs.Constellation
	var solarsystem structs.SolarSystem

	for _, file := range zipR.File {
		if strings.Contains(file.Name, ".staticdata") && strings.Contains(file.Name, "sde/fsd/universe/eve/") {
			// fmt.Println(file.Name)

			// if strings.Contains(file.Name, "region.staticdata") {
			// 	rc, _ := file.Open()
			// 	err = yaml.NewDecoder(rc).Decode(&region)
			// 	if err != nil {
			// 		fmt.Printf("Unmarshall error #%v", err)
			// 		continue
			// 	}
			// 	fmt.Printf("file.Name: %v %+v\n", file.Name, region)
			// }

			// if strings.Contains(file.Name, "constellation.staticdata") {
			// 	rc, _ := file.Open()
			// 	err = yaml.NewDecoder(rc).Decode(&constellation)
			// 	if err != nil {
			// 		fmt.Printf("Unmarshall error #%v", err)
			// 		continue
			// 	}
			// 	fmt.Printf("file.Name: %v %+v\n", file.Name, constellation)
			// }

			if strings.Contains(file.Name, "solarsystem.staticdata") {
				rc, _ := file.Open()
				err = yaml.NewDecoder(rc).Decode(&solarsystem)
				if err != nil {
					fmt.Printf("Unmarshall error #%v", err)
					break
				}
				fmt.Printf("file.Name: %v %+v\n", file.Name, solarsystem)

				break
			}
		}
	}
}
