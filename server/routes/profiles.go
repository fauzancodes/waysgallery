package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/postgresql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(e *echo.Group) {
	profileRepository := repositories.RepositoryProfile(postgresql.DB)
	h := handlers.HandlerProfile(profileRepository)

	e.GET("/profiles", h.FindProfiles)
	e.GET("/profile", middleware.Auth(h.GetProfile))
	e.PATCH("/profile", middleware.Auth(middleware.UploadFile(h.UpdateProfile)))
}
