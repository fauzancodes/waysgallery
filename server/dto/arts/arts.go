package artsdto

type ArtRequest struct {
	Image string `json:"image" form:"image"`
}

type ArtResponse struct {
	ID            int    `json:"id"`
	Image         string `json:"image"`
	ImagePublicID string `json:"image_public_id"`
}
