package projectssdto

type ProjectRequest struct {
	Description string `json:"description"`
	Image1      string `json:"image_1" form:"image1" validate:"required"`
	Image2      string `json:"image_2" form:"image2" validate:"required"`
	Image3      string `json:"image_3" form:"image3" validate:"required"`
	Image4      string `json:"image_4" form:"image4" validate:"required"`
	Image5      string `json:"image_5" form:"image5" validate:"required"`
}

type ProjectResponse struct {
	ID             int    `json:"id"`
	Description    string `json:"description"`
	Image1         string `json:"image_1"`
	ImagePublicID1 string `json:"image_public_id_1"`
	Image2         string `json:"image_2"`
	ImagePublicID2 string `json:"image_public_id_2"`
	Image3         string `json:"image_3"`
	ImagePublicID3 string `json:"image_public_id_3"`
	Image4         string `json:"image_4"`
	ImagePublicID4 string `json:"image_public_id_4"`
	Image5         string `json:"image_5"`
	ImagePublicID5 string `json:"image_public_id_5"`
}
