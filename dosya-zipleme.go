package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {

	var yazi string
	fmt.Println("Ziplenecek Dosyaları Seçin:")
	fmt.Scan(&yazi)

	fmt.Println("creating zip archive...")
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	fmt.Println("opening first file...")
	f1, err := os.Open(yazi)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fmt.Println("writing first file to archive...")
	w1, err := zipWriter.Create(yazi)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}

	fmt.Println("closing zip archive...")
	zipWriter.Close()
}
