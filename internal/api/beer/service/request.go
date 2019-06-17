package beer

type CreateBeerRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Strength    int    `json:"strength"`
}

type UpdateBeerRequest struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Strength    *int    `json:"strength"`
}
