package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/postgresql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func FollowingRoutes(e *echo.Group) {
	followingRepository := repositories.RepositoryFollowing(postgresql.DB)
	h := handlers.HandlerFollowing(followingRepository)

	e.GET("/followings", middleware.Auth(h.FindFollowings))
	e.GET("/following/:id", middleware.Auth(h.GetFollowing))
	e.POST("/following/:target_id", middleware.Auth(h.CreateFollowing))
	e.DELETE("/following", middleware.Auth(h.DeleteFollowing))
}
