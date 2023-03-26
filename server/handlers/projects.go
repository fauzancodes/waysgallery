package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	projectsdto "waysgallery/dto/projects"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

type handlerProject struct {
	ProjectRepository repositories.ProjectRepository
}

func HandlerProject(ProjectRepository repositories.ProjectRepository) *handlerProject {
	return &handlerProject{ProjectRepository}
}

func (h *handlerProject) FindProjects(c echo.Context) error {
	projects, err := h.ProjectRepository.FindProjects()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Data for all projects was successfully obtained", Data: projects})
}

func (h *handlerProject) GetProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var project models.Project
	project, err := h.ProjectRepository.GetProject(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Project data successfully obtained", Data: project})
}

func (h *handlerProject) CreateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("order_id"))
	filepath := c.Get("dataFiles").([]string)

	request := projectsdto.ProjectRequest{
		Description: c.FormValue("description"),
		Image1:      filepath[0],
		Image2:      filepath[1],
		Image3:      filepath[2],
		Image4:      filepath[3],
		Image5:      filepath[4],
	}

	var project models.Project

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	if request.Description != "" {
		project.Description = request.Description
	}
	if request.Image1 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[0], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		project.Image1 = resp.SecureURL
		project.ImagePublicID1 = resp.PublicID
	}
	if request.Image2 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[1], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		project.Image2 = resp.SecureURL
		project.ImagePublicID2 = resp.PublicID
	}
	if request.Image3 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[2], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		project.Image3 = resp.SecureURL
		project.ImagePublicID3 = resp.PublicID
	}
	if request.Image4 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[3], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		project.Image4 = resp.SecureURL
		project.ImagePublicID4 = resp.PublicID
	}
	if request.Image5 != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, filepath[4], uploader.UploadParams{Folder: "waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		project.Image5 = resp.SecureURL
		project.ImagePublicID5 = resp.PublicID
	}
	project.OrderID = id

	project, err := h.ProjectRepository.CreateProject(project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Project data created successfully", Data: convertResponseProject(project)})
}

func convertResponseProject(u models.Project) projectsdto.ProjectResponse {
	return projectsdto.ProjectResponse{
		ID:             u.ID,
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
	}
}
