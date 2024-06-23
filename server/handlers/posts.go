package handlers

import (
	"net/http"
	"strconv"
	"time"
	postsdto "waysgallery/dto/posts"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

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

	var post models.WaysGalleryPost
	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Post data successfully obtained", Data: post})
}

func (h *handlerPost) CreatePost(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	filepath := c.Get("cloudinarySecureURL").([]string)
	publicID := c.Get("cloudinaryPublicID").([]string)

	request := postsdto.PostRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Image1:      filepath[0],
		Image2:      filepath[1],
		Image3:      filepath[2],
		Image4:      filepath[3],
		Image5:      filepath[4],
	}

	var post models.WaysGalleryPost

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	if request.Title != "" {
		post.Title = request.Title
	}
	if request.Description != "" {
		post.Description = request.Description
	}

	if request.Image1 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[0], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image1 = request.Image1
		post.ImagePublicID1 = publicID[0]
	}
	if request.Image2 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[1], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image2 = request.Image2
		post.ImagePublicID2 = publicID[1]
	}
	if request.Image3 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[2], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image3 = request.Image3
		post.ImagePublicID3 = publicID[2]
	}
	if request.Image4 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[3], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image4 = request.Image4
		post.ImagePublicID4 = publicID[3]
	}
	if request.Image5 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[4], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image5 = request.Image5
		post.ImagePublicID5 = publicID[4]
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
	filepath := c.Get("cloudinarySecureURL").([]string)
	publicID := c.Get("cloudinaryPublicID").([]string)

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

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	if request.Title != "" {
		post.Title = request.Title
	}
	if request.Description != "" {
		post.Description = request.Description
	}

	if request.Image1 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[0], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image1 = request.Image1
		post.ImagePublicID1 = publicID[0]
	}
	if request.Image2 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[1], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image2 = request.Image2
		post.ImagePublicID2 = publicID[1]
	}
	if request.Image3 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[2], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image3 = request.Image3
		post.ImagePublicID3 = publicID[2]
	}
	if request.Image4 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[3], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image4 = request.Image4
		post.ImagePublicID4 = publicID[3]
	}
	if request.Image5 != "" {
		// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		// resp, err := cld.Upload.Upload(ctx, filepath[4], uploader.UploadParams{Folder: "waysgallery"})
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		post.Image5 = request.Image5
		post.ImagePublicID5 = publicID[4]
	}

	post.UpdatedAt = time.Now()

	data, err := h.PostRepository.UpdatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Post data updated successfully", Data: convertResponsePost(data)})
}

func convertResponsePost(u models.WaysGalleryPost) postsdto.PostResponse {
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
