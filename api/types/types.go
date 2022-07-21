package types

type FormRequest struct {
	Name string `json:"name" binding:"required"`
}
