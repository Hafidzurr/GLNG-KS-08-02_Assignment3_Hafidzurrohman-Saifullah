package main

import (
	"log"
	"math/rand" // Import math/rand untuk menghasilkan angka acak
	"net/http"
	"time"

	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment3_Hafidzurrohman-Saifullah.git/database" // Import package database untuk mengelola database
	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment3_Hafidzurrohman-Saifullah.git/handlers" // Import package handlers untuk mengelola rute HTTP
)

func main() {
	// Inisialisasi database
	if err := database.InitDB(); err != nil { // Menginisialisasi koneksi ke database
		log.Fatal(err)
	}
	defer database.CloseDB() // Menutup koneksi database setelah program selesai

	// Inisialisasi router untuk mengelola rute HTTP
	router := handlers.SetupRouter()

	// Mulai server HTTP
	srv := &http.Server{
		Addr:         ":8080",          // Server akan berjalan pada port 8080
		Handler:      router,           // Router akan mengelola permintaan HTTP
		ReadTimeout:  10 * time.Second, // Batasan waktu untuk membaca permintaan
		WriteTimeout: 10 * time.Second, // Batasan waktu untuk menulis respons
	}

	log.Println("Server berjalan di :8080") // Pesan log untuk menunjukkan server telah dimulai
	go updateDataEvery15Seconds()           // Memulai goroutine untuk memperbarui data setiap 15 detik
	log.Fatal(srv.ListenAndServe())         // Memulai server HTTP dan menunggu permintaan
}

func updateDataEvery15Seconds() {
	for {
		// Panggil fungsi yang menghasilkan dan memperbarui data cuaca
		err := generateAndSaveData()
		if err != nil {
			log.Println("Gagal mengupdate data:", err)
		}

		// Tunggu selama 15 detik sebelum mengupdate lagi
		time.Sleep(15 * time.Second)
	}
}

func generateAndSaveData() error {
	// Generate angka acak untuk water (air) dan wind (angin)
	rand.Seed(time.Now().UnixNano()) // Inisialisasi generator angka acak dengan waktu sekarang
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	// Tentukan status air (water)
	var waterStatus string
	if water < 5 {
		waterStatus = "aman"
	} else if water >= 5 && water <= 8 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	// Tentukan status angin (wind)
	var windStatus string
	if wind < 6 {
		windStatus = "aman"
	} else if wind >= 6 && wind <= 15 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	// Simpan data ke database
	_, err := database.DB.Exec("INSERT INTO weather_data (water, wind, water_status, wind_status) VALUES (?, ?, ?, ?)", water, wind, waterStatus, windStatus)
	if err != nil {
		return err
	}

	// Tampilkan status dalam log
	log.Printf("water: %d, wind: %d\n", water, wind)
	log.Printf("status water: %s, status wind: %s\n", waterStatus, windStatus)

	return nil
}
