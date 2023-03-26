package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/postgresql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Group) {
	orderRepository := repositories.RepositoryOrder(postgresql.DB)
	userRepository := repositories.RepositoryUser(postgresql.DB)
	h := handlers.HandlerOrder(orderRepository, userRepository)

	e.GET("/orders", middleware.Auth(h.FindOrders))
	e.GET("/order/:id", middleware.Auth(h.GetOrder))
	e.POST("/order/:vendor_id", middleware.Auth(h.CreateOrder))
	e.PATCH("/order/:id", middleware.Auth(h.UpdateOrderStatus))
	e.POST("/notification", h.Notification)
}
