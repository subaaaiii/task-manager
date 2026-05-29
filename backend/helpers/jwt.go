package helpers

import (
	"backend/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Nilai secret diambil dari environment variable JWT_SECRET
var jwtKey = config.GetJWTKey()

func GenerateToken() string {

	// Mengatur waktu kedaluwarsa token, di sini kita set 60 menit dari waktu sekarang
	expirationTime := time.Now().Add(60 * time.Minute)

	// Membuat klaim (claims) JWT
	// Subject berisi , dan ExpiresAt menentukan waktu expired token
	claims := &JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Membuat token baru dengan klaim yang telah dibuat
	// Menggunakan algoritma HS256 untuk menandatangani token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	// Mengembalikan token dalam bentuk string
	return token
}

func GenerateResetToken() string {

	// Mengatur waktu kedaluwarsa token, di sini kita set 10 menit dari waktu sekarang
	expirationTime := time.Now().Add(10 * time.Minute)

	// Membuat klaim (claims) JWT
	// Subject berisi , dan ExpiresAt menentukan waktu expired token
	claims := &JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Membuat token baru dengan klaim yang telah dibuat
	// Menggunakan algoritma HS256 untuk menandatangani token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	// Mengembalikan token dalam bentuk string
	return token
}
