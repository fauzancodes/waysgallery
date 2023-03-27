package handlers

import (
	"fmt"
	"log"
	"os"
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
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
	UserRepository  repositories.UserRepository
}

func HandlerOrder(OrderRepository repositories.OrderRepository, UserRepository repositories.UserRepository) *handlerOrder {
	return &handlerOrder{
		OrderRepository: OrderRepository,
		UserRepository:  UserRepository,
	}
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

	var orderIsMatch = false
	var orderId int
	for !orderIsMatch {
		orderId = int(time.Now().Unix())
		orderData, _ := h.OrderRepository.GetOrder(orderId)
		if orderData.ID == 0 {
			orderIsMatch = true
		}
	}

	order := models.Order{
		ID:          orderId,
		Title:       request.Title,
		Description: request.Description,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Price:       request.Price,
		VendorID:    id,
		ClientID:    int(userId),
		UserID:      int(userId),
		Status:      "cancel",
	}
	data, err := h.OrderRepository.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	// return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Order data created successfully", Data: convertResponseOrder(data)})

	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(data.ID),
			GrossAmt: int64(data.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: snapResp})
}

func (h *handlerOrder) UpdateOrderStatus(c echo.Context) error {
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

	data, err := h.OrderRepository.UpdateOrderStatus(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Order data updated successfully", Data: convertResponseOrder(data)})
}

func (h *handlerOrder) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	order_id, _ := strconv.Atoi(orderId)
	order, _ := h.OrderRepository.GetOrder(order_id)
	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			h.OrderRepository.UpdateOrder("cancel", order_id)
		} else if fraudStatus == "accept" {
			SendMail("waiting", order)
			h.OrderRepository.UpdateOrder("waiting", order_id)
		}
	} else if transactionStatus == "settlement" {
		SendMail("waiting", order)
		h.OrderRepository.UpdateOrder("waiting", order_id)
	} else if transactionStatus == "deny" {
		h.OrderRepository.UpdateOrder("cancel", order_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		h.OrderRepository.UpdateOrder("cancel", order_id)
	} else if transactionStatus == "pending" {
		h.OrderRepository.UpdateOrder("cancel", order_id)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Payment finished", Data: notificationPayload})
}

func SendMail(status string, order models.Order) {

	if status != order.Status && (status == "waiting") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "WaysGallery <waysgallery.admin@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var price = strconv.Itoa(order.Price)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", "mfauzan.murtadho@gmail.com")
		mailer.SetHeader("Subject", "WaysGallery Order Payment")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
    <html lang="en">
      <head>
      <meta charset="UTF-8" />
      <meta http-equiv="X-UA-Compatible" content="IE=edge" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Document</title>
      </head>
      <body>
      <h2>Product Payment</h2>
      <ul style="list-style-type:none;">
				<li>Title : %s</li>
				<li>Description : %s</li>
				<li>Start Date : %s</li>
				<li>End Date : %s</li>
				<li>Price : %s</li>
				<li>Status : %s approvement from vendor</li>
      </ul>
			<h4>&copy; 2023. <a href="https://waysgallery.vercel.app">WaysGallery</a>.</h4>
      </body>
    </html>`, order.Title, order.Description, order.StartDate, order.EndDate, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + CONFIG_AUTH_EMAIL)
	}
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
