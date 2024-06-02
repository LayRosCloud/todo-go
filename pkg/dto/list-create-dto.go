package dto

type ListCreateDto struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
