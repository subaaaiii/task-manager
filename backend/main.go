package main

import (
	"net/http"

	"backend/config"
	"backend/database"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Membuat instance default dari Gin
	config.LoadEnv()
	router := gin.Default()
	database.InitDB()

	// 2. Route sederhana untuk tes halaman utama
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to task Backend!")
	})

	// 3. Panggil fungsi setup routes yang telah kita pisah
	r := routes.SetupRoutes()

	// 4. Jalankan server di port 8080
	r.Run(":8080")

}
