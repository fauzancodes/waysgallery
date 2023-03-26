package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	postsdto "waysgallery/dto/posts"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerPost struct {
	PostRepository repositories.PostRepository
}

func HandlerPost(PostRepository repositories.PostRepository) *handlerPost {
	return &handlerPost{PostRepository}
}

func (h *handlerPost) FindPosts(c echo.Context) error {
	posts, err := h.PostRepository.FindPosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Data for all posts was successfully obtained", Data: posts})
}

func (h *handlerPost) GetPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var post models.Post
	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Post data successfully obtained", Data: post})
}

func (h *handlerPost) CreatePost(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	filepath := c.Get("dataFiles").([]string)

	request := postsdto.PostRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Image1:      filepath[0],
		Image2:      filepath[1],
		Image3:      filepath[2],
		Image4:      filepath[3],
		Image5:      filepath[4],
	}

	var post models.Post

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	if request.Title != "" {
		post.Title = request.Title
	}
	if request.Description != "" {
		post.Description = request.Description
	}
	if request.Image1 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[0], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image1 = resp.SecureURL
		post.ImagePublicID1 = resp.PublicID
	}
	if request.Image2 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[1], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image2 = resp.SecureURL
		post.ImagePublicID2 = resp.PublicID
	}
	if request.Image3 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[2], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image3 = resp.SecureURL
		post.ImagePublicID3 = resp.PublicID
	}
	if request.Image4 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[3], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image4 = resp.SecureURL
		post.ImagePublicID4 = resp.PublicID
	}
	if request.Image5 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[4], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image5 = resp.SecureURL
		post.ImagePublicID5 = resp.PublicID
	}
	post.UserID = int(userId)

	post, err := h.PostRepository.CreatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Post data created successfully", Data: convertResponsePost(post)})
}

func (h *handlerPost) DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.PostRepository.DeletePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Post data deleted successfully", Data: convertResponsePost(data)})
}

func (h *handlerPost) UpdatePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	filepath := c.Get("dataFiles").([]string)

	request := postsdto.PostRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Image1:      filepath[0],
		Image2:      filepath[1],
		Image3:      filepath[2],
		Image4:      filepath[3],
		Image5:      filepath[4],
	}

	post, err := h.PostRepository.GetPost(int(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	if request.Title != "" {
		post.Title = request.Title
	}
	if request.Description != "" {
		post.Description = request.Description
	}
	if request.Image1 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[0], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image1 = resp.SecureURL
		post.ImagePublicID1 = resp.PublicID
	}
	if request.Image2 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[1], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image2 = resp.SecureURL
		post.ImagePublicID2 = resp.PublicID
	}
	if request.Image3 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[2], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image3 = resp.SecureURL
		post.ImagePublicID3 = resp.PublicID
	}
	if request.Image4 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[3], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image4 = resp.SecureURL
		post.ImagePublicID4 = resp.PublicID
	}
	if request.Image5 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[4], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		post.Image5 = resp.SecureURL
		post.ImagePublicID5 = resp.PublicID
	}

	post.UpdatedAt = time.Now()

	data, err := h.PostRepository.UpdatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Post data updated successfully", Data: convertResponsePost(data)})
}

func convertResponsePost(u models.Post) postsdto.PostResponse {
	return postsdto.PostResponse{
		ID:             u.ID,
		Title:          u.Title,
		Description:    u.Description,
		Image1:         u.Image1,
		ImagePublicID1: u.ImagePublicID1,
		Image2:         u.Image2,
		ImagePublicID2: u.ImagePublicID2,
		Image3:         u.Image3,
		ImagePublicID3: u.ImagePublicID3,
		Image4:         u.Image4,
		ImagePublicID4: u.ImagePublicID4,
		Image5:         u.Image5,
		ImagePublicID5: u.ImagePublicID5,
		UserID:         u.UserID,
	}
}
