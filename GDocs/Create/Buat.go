package create

import (
	"fmt"
	"os"
)

// BuatFile adalah fungsi untuk membuat file baru
func BuatFile(namaFile string) error {
	file, err := os.Create(namaFile)
	if err != nil {
		return fmt.Errorf("gagal membuat file: %v", err)
	}
	defer file.Close()

	// Lakukan operasi lain terkait pembuatan file

	return nil
}

// BuatDirektori adalah fungsi untuk membuat direktori baru
func BuatDirektori(namaDirektori string) error {
	err := os.Mkdir(namaDirektori, 0755)
	if err != nil {
		return fmt.Errorf("gagal membuat direktori: %v", err)
	}

	// Lakukan operasi lain terkait pembuatan direktori

	return nil
}