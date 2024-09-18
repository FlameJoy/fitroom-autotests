package initializers

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// LandFunURL string
	// LandRuURL  string
	// AppFunURL  string
	// AppRuURL   string
	// AdmFunURL  string
	// AdmRuURL   string
	AdmTel   string
	UserTel  string
	AdmCode  string
	UserCode string
)

func LoadENV() {
	if err := godotenv.Load("./.env"); err != nil {
		panic(err)
	}
	// // FUN
	// LandFunURL = os.Getenv("LAND_FUN")
	// AppFunURL = os.Getenv("APP_FUN")
	// AdmFunURL = os.Getenv("ADM_FUN")
	// // RU
	// LandRuURL = os.Getenv("LAND_RU")
	// AppRuURL = os.Getenv("APP_RU")
	// AdmRuURL = os.Getenv("ADM_RU")
	// Users
	AdmTel = os.Getenv("ADMIN_PHONE")
	UserTel = os.Getenv("USER_PHONE")
	AdmCode = os.Getenv("ADMIN_CODE")
	UserCode = os.Getenv("USER_CODE")
}
