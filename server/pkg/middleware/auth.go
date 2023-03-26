package middleware

import (
	"net/http"
	"strings"
	dto "waysgallery/dto/result"
	jwtToken "waysgallery/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusBadRequest, Message: "unauthorized"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "unauthorized"})
		}

		c.Set("userLogin", claims)
		return next(c)
	}
}
