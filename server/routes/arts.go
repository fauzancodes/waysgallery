package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/postgresql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func ArtRoutes(e *echo.Group) {
	artRepository := repositories.RepositoryArt(postgresql.DB)
	h := handlers.HandlerArt(artRepository)

	e.GET("/arts", middleware.Auth(h.FindArts))
	e.GET("/art/:id", middleware.Auth(h.GetArt))
	e.POST("/art", middleware.Auth(middleware.UploadFile(h.CreateArt)))
	e.DELETE("/art/:id", middleware.Auth(h.DeleteArt))
}
