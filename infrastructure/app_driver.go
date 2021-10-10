package infrastructure

import (
	"github.com/joho/godotenv"
	"os"
)

type DriverAppService struct {
	Port   string
	Secret string
}

func NewDriverApp() (DriverAppService, error) {
	if err := godotenv.Load(".env"); err != nil {
		return DriverAppService{}, err
	}

	return DriverAppService{
		Port:   os.Getenv("PORT"),
		Secret: os.Getenv("APP_SECRET"),
	}, nil
}
