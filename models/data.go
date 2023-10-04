package models

import (
	"database/sql"
)

// WeatherData adalah struktur data yang digunakan untuk menyimpan informasi cuaca.
type WeatherData struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

// GetLatestData adalah fungsi yang digunakan untuk mengambil data cuaca terbaru dari database.
func GetLatestData(db *sql.DB) (*WeatherData, error) {
	// Query untuk mengambil data terbaru dari tabel 'weather_data' berdasarkan kolom 'updated_at'.
	query := "SELECT water, wind, water_status, wind_status FROM weather_data ORDER BY updated_at DESC LIMIT 1"

	// Eksekusi query dan menyimpan hasilnya pada 'row'.
	row := db.QueryRow(query)

	// Inisialisasi variabel untuk menyimpan data dari hasil query.
	var water, wind int
	var waterStatus, windStatus string

	// Ambil nilai dari baris yang dihasilkan oleh query dan masukkan ke dalam variabel yang sesuai.
	err := row.Scan(&water, &wind, &waterStatus, &windStatus)
	if err != nil {
		return nil, err
	}

	// Membuat objek WeatherData dari hasil query.
	data := &WeatherData{
		Water:       water,
		Wind:        wind,
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}

	// Mengembalikan data cuaca terbaru dan tidak ada error.
	return data, nil
}
