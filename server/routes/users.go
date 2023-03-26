package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/postgresql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(postgresql.DB)
	profileRepository := repositories.RepositoryProfile(postgresql.DB)
	artRepository := repositories.RepositoryArt(postgresql.DB)
	h := handlers.HandlerUser(userRepository, profileRepository, artRepository)

	e.GET("/users", middleware.Auth(h.FindUsers))
	e.GET("/user/:id", middleware.Auth(h.GetUser))
	e.PATCH("/user", middleware.Auth(h.UpdateUser))
	e.DELETE("/user", middleware.Auth(h.DeleteUser))
}
