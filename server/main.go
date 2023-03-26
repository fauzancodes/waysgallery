package main

import (
	"fmt"
	"os"
	"waysgallery/database"
	"waysgallery/pkg/postgresql"
	"waysgallery/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	postgresql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	var port = os.Getenv("PORT")

	fmt.Println("server running localhost:" + port)
	e.Logger.Fatal(e.Start("localhost:" + port))
}
