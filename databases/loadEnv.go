package databases

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Tidak dapat memuat .env")
	}

}
