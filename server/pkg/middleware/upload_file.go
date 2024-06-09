package middleware

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	dto "waysgallery/dto/result"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var method = c.Request().Method
		file, err := c.FormFile("image")

		if err != nil {
			if method == "PATCH" && err.Error() == "http: no such file" {
				c.Set("cloudinarySecureURL", "")
				c.Set("cloudinaryPublicID", "")
				return next(c)
			}
		}
		if err != nil {
			if method == "POST" && err.Error() == "http: no such file" {
				c.Set("cloudinarySecureURL", "")
				c.Set("cloudinaryPublicID", "")
				return next(c)
			}
		}

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		ext := filepath.Ext(file.Filename)
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()

			// tempFile, err := ioutil.TempFile("uploads", "image-*.png")
			// if err != nil {
			// 	return c.JSON(http.StatusBadRequest, err)
			// }
			// defer tempFile.Close()

			// if _, err = io.Copy(tempFile, src); err != nil {
			// 	return c.JSON(http.StatusBadRequest, err)
			// }

			// data := tempFile.Name()

			// c.Set("dataFile", data)

			var ctx = context.Background()
			var CLOUD_NAME = os.Getenv("CLOUD_NAME")
			var API_KEY = os.Getenv("API_KEY")
			var API_SECRET = os.Getenv("API_SECRET")
			cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
			resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "waysgallery"})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
			}

			c.Set("cloudinarySecureURL", resp.SecureURL)
			c.Set("cloudinaryPublicID", resp.PublicID)

			return next(c)
		} else {
			return c.JSON(http.StatusBadRequest, "The file extension is wrong. Allowed file extensions are images (.png, .jpg, .jpeg, .webp)")
		}
	}
}
