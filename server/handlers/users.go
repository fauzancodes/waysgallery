package handlers

import (
	"net/http"
	"strconv"
	"time"
	dto "waysgallery/dto/result"
	usersdto "waysgallery/dto/users"
	"waysgallery/models"
	"waysgallery/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handler struct {
	UserRepository    repositories.UserRepository
	ProfileRepository repositories.ProfileRepository
	ArtRepository     repositories.ArtRepository
}

func HandlerUser(UserRepository repositories.UserRepository, ProfileRepository repositories.ProfileRepository, ArtRepository repositories.ArtRepository) *handler {
	return &handler{
		UserRepository:    UserRepository,
		ProfileRepository: ProfileRepository,
		ArtRepository:     ArtRepository,
	}
}

func (h *handler) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Error from Repository: " + err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Data for all users was successfully obtained", Data: users})
}

func (h *handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "User data successfully obtained", Data: user})
}

func (h *handler) UpdateUser(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := new(usersdto.UserUpdateRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	user.UpdatedAt = time.Now()

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "User data updated successfully", Data: convertResponse(data)})
}

func (h *handler) DeleteUser(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	profiles, err := h.ProfileRepository.FindProfiles()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	for _, profile := range profiles {
		if profile.UserID == int(userId) {
			arts, err := h.ArtRepository.FindArts()
			if err != nil {
				return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			}
			for _, art := range arts {
				if art.ProfileID == profile.ID {
					artToDelete, err := h.ArtRepository.GetArt(art.ID)
					if err != nil {
						return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
					}

					_, err = h.ArtRepository.DeleteArt(artToDelete)
					if err != nil {
						return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
					}
				}
			}

			profileToDelete, err := h.ProfileRepository.GetProfile(profile.ID)
			if err != nil {
				return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			}

			_, err = h.ProfileRepository.DeleteProfile(profileToDelete)
			if err != nil {
				return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			}
		}
	}

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "User data deleted successfully", Data: convertResponse(data)})
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
