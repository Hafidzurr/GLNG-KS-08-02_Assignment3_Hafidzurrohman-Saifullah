package handlers

import (
	"net/http"

	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment3_Hafidzurrohman-Saifullah.git/database"
	"github.com/Hafidzurr/GLNG-KS-08-02_Assignment3_Hafidzurrohman-Saifullah.git/models"
	"github.com/gin-gonic/gin"
)

// SetupRouter adalah fungsi yang digunakan untuk mengatur rute-rute aplikasi web.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Menambahkan rute untuk /weatherupdate
	router.GET("/weatherupdate", GetWeatherUpdate)

	return router
}

// GetWeatherUpdate adalah fungsi yang akan dipanggil ketika permintaan GET diterima pada rute /weatherupdate.
func GetWeatherUpdate(c *gin.Context) {
	// Inisialisasi database jika belum diinisialisasi
	if database.DB == nil {
		if err := database.InitDB(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer database.CloseDB()
	}

	// Dapatkan data terbaru dari database menggunakan model GetLatestData
	data, err := models.GetLatestData(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kirimkan respons JSON dengan data cuaca terbaru
	c.JSON(http.StatusOK, data)
}
