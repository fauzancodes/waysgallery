package handlers

import (
	"log"
	"net/http"
	"time"
	dto "waysgallery/dto/result"
	usersdto "waysgallery/dto/users"

	"waysgallery/models"
	"waysgallery/pkg/bcrypt"
	jwtToken "waysgallery/pkg/jwt"
	"waysgallery/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository    repositories.AuthRepository
	ProfileRepository repositories.ProfileRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository, ProfileRepository repositories.ProfileRepository) *handlerAuth {
	return &handlerAuth{
		AuthRepository:    AuthRepository,
		ProfileRepository: ProfileRepository,
	}
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(usersdto.RegisterRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.WaysGalleryUser{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	_, err = h.AuthRepository.Login(user.Email)
	if err != nil {
		data, err := h.AuthRepository.Register(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		}

		registerResponse := usersdto.UserResponse{
			ID:    data.ID,
			Name:  data.Name,
			Email: data.Email,
		}

		profile := models.WaysGalleryProfile{
			ID:       registerResponse.ID,
			Name:     registerResponse.Name,
			Greeting: "",
			Image:    "",
			UserID:   registerResponse.ID,
		}
		_, err = h.ProfileRepository.CreateProfile(profile)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Your registration is successful", Data: registerResponse})
	} else {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "This email is already registered"})
	}
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(usersdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.WaysGalleryUser{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "wrong password"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := usersdto.LoginResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: isValid,
		Token:    token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "You have successfully logged in", Data: loginResponse})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "The authentication was successfully examined", Data: user})
}
