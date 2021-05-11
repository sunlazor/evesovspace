package main

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"os"
	"strings"
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

	for _, file := range zipR.File {
		if strings.Contains(file.Name, ".staticdata") && strings.Contains(file.Name, "sde/fsd/universe/eve/") {
			fmt.Println(file.Name)
		}

		// fmt.Println("Файл " + file.Name + " содержит следующее:")
		// r, err := file.Open()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// _, err = io.Copy(os.Stdout, r)
		// if err != nil {
		// 	panic(err)
		// }
		// err = r.Close()
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println()

		// break
	}

}
