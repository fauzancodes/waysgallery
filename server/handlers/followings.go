package handlers

import (
	"strconv"
	followingsdto "waysgallery/dto/followings"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerFollowing struct {
	FollowingRepository repositories.FollowingRepository
}

func HandlerFollowing(FollowingRepository repositories.FollowingRepository) *handlerFollowing {
	return &handlerFollowing{FollowingRepository}
}

func (h *handlerFollowing) FindFollowings(c echo.Context) error {
	followings, err := h.FollowingRepository.FindFollowings()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Data for all followings was successfully obtained", Data: followings})
}

func (h *handlerFollowing) GetFollowing(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	following, err := h.FollowingRepository.GetFollowing(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Following data successfully obtained", Data: following})
}

func (h *handlerFollowing) CreateFollowing(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("target_id"))
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	followings, err := h.FollowingRepository.FindFollowings()
	for _, followingData := range followings {
		if followingData.UserID == int(userId) && followingData.FollowingID == id {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Sorry, you are already following this user"})
		}
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	following := models.Following{
		FollowingID: id,
		UserID:      int(userId),
	}
	data, err := h.FollowingRepository.CreateFollowing(following)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Following data created successfully", Data: convertResponseFollowing(data)})
}

func (h *handlerFollowing) DeleteFollowing(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	following, err := h.FollowingRepository.GetFollowing(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FollowingRepository.DeleteFollowing(following)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Following data deleted successfully", Data: convertResponseFollowing(data)})
}

func convertResponseFollowing(u models.Following) followingsdto.FollowingResponse {
	return followingsdto.FollowingResponse{
		ID:          u.ID,
		FollowingID: u.FollowingID,
	}
}
