// Apps.go

package main

import (
	"fmt"
	"log"

	"github.com/EtoniaThea/Docs/read"
	"github.com/EtoniaThea/Docs/create"
)

func main() {
	// Memanggil fungsi dari package read
	err := read.BacaFile("nama_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Memanggil fungsi dari package create
	err = create.BuatFile("file_baru.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = create.BuatDirektori("direktori_baru")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Operasi selesai.")
}
