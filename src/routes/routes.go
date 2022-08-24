package routes

import (
	"log"
	"net/http"
	"os"

	"api/controllers"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func HandleResquest() {
	port := goDotEnvVariable("PORT")
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
