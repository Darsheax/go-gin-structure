package envRead

import (
	"log"

	"github.com/joho/godotenv"
)

func Read(file string) map[string]string {

	var envs map[string]string
	envs, err := godotenv.Read(file)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return envs
}
