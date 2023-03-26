package handlers

import (
	"strconv"
	"time"
	ordersdto "waysgallery/dto/orders"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
}

func HandlerOrder(OrderRepository repositories.OrderRepository) *handlerOrder {
	return &handlerOrder{OrderRepository}
}

func (h *handlerOrder) FindOrders(c echo.Context) error {
	orders, err := h.OrderRepository.FindOrders()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Data for all orders was successfully obtained", Data: orders})
}

func (h *handlerOrder) GetOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Order data successfully obtained", Data: order})
}

func (h *handlerOrder) CreateOrder(c echo.Context) error {
	StartDateInput, _ := time.Parse("2006-01-02", c.FormValue("start_date"))
	EndDateInput, _ := time.Parse("2006-01-02", c.FormValue("end_date"))
	price, _ := strconv.Atoi(c.FormValue("price"))

	request := ordersdto.OrderRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		StartDate:   StartDateInput,
		EndDate:     EndDateInput,
		Price:       price,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("vendor_id"))
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	order := models.Order{
		Title:       request.Title,
		Description: request.Description,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Price:       request.Price,
		VendorID:    id,
		ClientID:    int(userId),
		UserID:      int(userId),
		Status:      "pending",
	}
	data, err := h.OrderRepository.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Order data created successfully", Data: convertResponseOrder(data)})
}

func (h *handlerOrder) UpdateOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(ordersdto.OrderStatusRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Status != "" {
		order.Status = request.Status
	}

	order.UpdatedAt = time.Now()

	data, err := h.OrderRepository.UpdateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Order data updated successfully", Data: convertResponseOrder(data)})
}

func convertResponseOrder(u models.Order) ordersdto.OrderResponse {
	return ordersdto.OrderResponse{
		ID:          u.ID,
		Title:       u.Title,
		Description: u.Description,
		StartDate:   u.StartDate,
		EndDate:     u.EndDate,
		Price:       u.Price,
		VendorID:    u.VendorID,
		ClientID:    u.ClientID,
		Status:      u.Status,
	}
}
