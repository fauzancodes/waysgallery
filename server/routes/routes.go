package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	UserRoutes(e)
	ProfileRoutes(e)
	ArtRoutes(e)
	FollowingRoutes(e)
	OrderRoutes(e)
	PostRoutes(e)
	ProjectRoutes(e)
}
