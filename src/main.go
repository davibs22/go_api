package main

import (
	"fmt"
	"log"
	"os"

	"api/routes"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	port := goDotEnvVariable("PORT")
	fmt.Printf("Iniciando o servidor Rest com Go na porta %s\n", port)
	routes.HandleResquest()
}
