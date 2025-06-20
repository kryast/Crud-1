package main

import (
	"fmt"
	"log"

	"github.com/kryast/Crud-1.git/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Gagal mendapatkan koneksi DB: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Database tidak merespon: %v", err)
	}

	fmt.Println("Koneksi ke database berhasil")
}
