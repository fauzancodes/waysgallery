package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/postgresql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func ProjectRoutes(e *echo.Group) {
	projectRepository := repositories.RepositoryProject(postgresql.DB)
	h := handlers.HandlerProject(projectRepository)

	e.GET("/projects", h.FindProjects)
	e.GET("/project/:id", h.GetProject)
	e.POST("/project/:order_id", middleware.Auth(middleware.UploadFiles(h.CreateProject)))
}
