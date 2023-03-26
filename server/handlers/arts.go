package handlers

import (
	"context"
	"os"
	"strconv"
	artsdto "waysgallery/dto/arts"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerArt struct {
	ArtRepository repositories.ArtRepository
}

func HandlerArt(ArtRepository repositories.ArtRepository) *handlerArt {
	return &handlerArt{ArtRepository}
}

func (h *handlerArt) FindArts(c echo.Context) error {
	arts, err := h.ArtRepository.FindArts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Data for all arts was successfully obtained", Data: arts})
}

func (h *handlerArt) GetArt(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	art, err := h.ArtRepository.GetArt(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Art data successfully obtained", Data: art})
}

func (h *handlerArt) CreateArt(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	filepath := c.Get("dataFile").(string)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysgallery"})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	art := models.Art{
		Image:         resp.SecureURL,
		ImagePublicID: resp.PublicID,
		ProfileID:     int(userId),
	}
	data, err := h.ArtRepository.CreateArt(art)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Art data created successfully", Data: convertResponseArt(data)})
}

func (h *handlerArt) DeleteArt(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	art, err := h.ArtRepository.GetArt(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ArtRepository.DeleteArt(art)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Art data deleted successfully", Data: convertResponseArt(data)})
}

func convertResponseArt(u models.Art) artsdto.ArtResponse {
	return artsdto.ArtResponse{
		ID:            u.ID,
		Image:         u.Image,
		ImagePublicID: u.ImagePublicID,
	}
}
