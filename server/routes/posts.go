package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/postgresql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Group) {
	postRepository := repositories.RepositoryPost(postgresql.DB)
	h := handlers.HandlerPost(postRepository)

	e.GET("/posts", h.FindPosts)
	e.GET("/post/:id", h.GetPost)
	e.POST("/post", middleware.Auth(middleware.UploadFiles(h.CreatePost)))
	e.DELETE("/post/:id", middleware.Auth(h.DeletePost))
	e.PATCH("/post/:id", middleware.Auth(middleware.UploadFiles(h.UpdatePost)))
}