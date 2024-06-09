package postsdto

type PostRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Image1      string `json:"image_1" form:"image1"`
	Image2      string `json:"image_2" form:"image2"`
	Image3      string `json:"image_3" form:"image3"`
	Image4      string `json:"image_4" form:"image4"`
	Image5      string `json:"image_5" form:"image5"`
}

type PostResponse struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
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
	UserID         int    `json:"user_id"`
}
