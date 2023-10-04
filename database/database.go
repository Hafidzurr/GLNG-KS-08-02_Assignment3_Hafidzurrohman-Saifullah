package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB adalah variabel yang akan digunakan untuk koneksi ke database MySQL.
var DB *sql.DB

// validateDBConnection memeriksa apakah koneksi database telah diinisialisasi dan valid.
func validateDBConnection() error {
	if DB == nil {
		return errors.New("koneksi database belum diinisialisasi")
	}
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("koneksi database gagal: %v", err)
	}
	return nil
}

// InitDB digunakan untuk menginisialisasi koneksi ke database.
func InitDB() error {
	// Buat koneksi ke database MySQL
	var err error
	DB, err = sql.Open("mysql", "root:Hafidzurr1@tcp(localhost:3306)/weather")
	if err != nil {
		return err
	}

	// Pastikan koneksi berhasil
	if err := validateDBConnection(); err != nil {
		return err
	}

	log.Println("Koneksi database sukses")

	// Buat tabel jika belum ada

	// Eksekusi perintah SQL untuk membuat tabel 'weather_data' jika belum ada.
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS weather_data (
            id INT AUTO_INCREMENT PRIMARY KEY,
            water INT,
            wind INT,
            water_status VARCHAR(255),
            wind_status VARCHAR(255),
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	return nil
}

// CloseDB digunakan untuk menutup koneksi database jika masih terbuka.
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
