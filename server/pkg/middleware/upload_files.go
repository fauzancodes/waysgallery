package middleware

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	dto "waysgallery/dto/result"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func UploadFiles(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var method = c.Request().Method
		form, err := c.MultipartForm()

		if err != nil {
			if method == "PATCH" && err.Error() == "http: no such file" {
				c.Set("cloudinarySecureURL", "")
				c.Set("cloudinaryPublicID", "")
				return next(c)
			}
			if method == "POST" && err.Error() == "http: no such file" {
				c.Set("cloudinarySecureURL", "")
				c.Set("cloudinaryPublicID", "")
				return next(c)
			}
			return c.JSON(http.StatusBadRequest, err)
		}

		// var dataFiles []string
		var cloudinarySecureURL []string
		var cloudinaryPublicID []string

		for i := 0; i < 5; i++ {
			file := form.File["image"+strconv.Itoa(i+1)]
			if len(file) == 0 {
				cloudinarySecureURL = append(cloudinarySecureURL, "")
				cloudinaryPublicID = append(cloudinaryPublicID, "")
				continue
			}

			ext := filepath.Ext(file[0].Filename)
			if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
				src, err := file[0].Open()
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

				// dataFiles = append(dataFiles, tempFile.Name())

				var ctx = context.Background()
				var CLOUD_NAME = os.Getenv("CLOUD_NAME")
				var API_KEY = os.Getenv("API_KEY")
				var API_SECRET = os.Getenv("API_SECRET")
				cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
				resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "waysgallery"})
				if err != nil {
					return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
				}

				cloudinarySecureURL = append(cloudinarySecureURL, resp.SecureURL)
				cloudinaryPublicID = append(cloudinaryPublicID, resp.PublicID)

			} else {
				return c.JSON(http.StatusBadRequest, "The file extension is wrong. Allowed file extensions are images (.png, .jpg, .jpeg, .webp)")
			}
		}

		// c.Set("dataFiles", dataFiles)
		c.Set("cloudinarySecureURL", cloudinarySecureURL)
		c.Set("cloudinaryPublicID", cloudinaryPublicID)

		return next(c)
	}
}

// package middleware

// import (
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"path/filepath"

// 	"github.com/labstack/echo/v4"
// )

// func UploadFiles(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var method = c.Request().Method
// 		form, err := c.MultipartForm()

// 		if err != nil {
// 			if method == "PATCH" && err.Error() == "http: no such file" {
// 				c.Set("dataFile", "")
// 				return next(c)
// 			}
// 			if method == "POST" && err.Error() == "http: no such file" {
// 				c.Set("dataFile", "")
// 				return next(c)
// 			}
// 			return c.JSON(http.StatusBadRequest, err)
// 		}

// 		files := form.File["image"]

// 		var dataFiles []string

// 		for _, file := range files {
// 			ext := filepath.Ext(file.Filename)
// 			if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
// 				src, err := file.Open()
// 				if err != nil {
// 					return c.JSON(http.StatusBadRequest, err)
// 				}
// 				defer src.Close()

// 				tempFile, err := ioutil.TempFile("uploads", "image-*.png")
// 				if err != nil {
// 					return c.JSON(http.StatusBadRequest, err)
// 				}
// 				defer tempFile.Close()

// 				if _, err = io.Copy(tempFile, src); err != nil {
// 					return c.JSON(http.StatusBadRequest, err)
// 				}

// 				dataFiles = append(dataFiles, tempFile.Name())
// 			} else {
// 				return c.JSON(http.StatusBadRequest, "The file extension is wrong. Allowed file extensions are images (.png, .jpg, .jpeg, .webp)")
// 			}
// 		}

// 		c.Set("dataFiles", dataFiles)

// 		return next(c)
// 	}
// }
