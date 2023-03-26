package profilesdto

type ProfileRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Greeting string `json:"greeting" form:"greeting" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
}

type ProfileResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Greeting      string `json:"greeting"`
	Image         string `json:"image"`
	ImagePublicID string `json:"image_public_id"`
}
