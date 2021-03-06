package main

import (
	"packages.hetic.net/gocqrs/bus/queries"
	"packages.hetic.net/gocqrs/models"
	"packages.hetic.net/gocqrs/router"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	env, _ := godotenv.Read(".env")

	if len(env["DB_PASSWORD"]) == 0 {
		panic("Ajouter les variables d'environnement au niveaux du .env.example")
	}

	queries.InitRoutine()

	models.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], env["DB_PORT"])

	router.StartRouter(env["API_PORT"])
}
