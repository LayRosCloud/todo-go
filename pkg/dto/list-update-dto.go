package dto

type ListUpdateDto struct {
	Id          int64  `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
