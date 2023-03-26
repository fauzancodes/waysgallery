package followingsdto

type FollowingRequest struct {
	FollowingID int `json:"following_id" form:"following_id" validate:"required"`
}

type FollowingResponse struct {
	ID          int `json:"id"`
	FollowingID int `json:"following_id"`
}
