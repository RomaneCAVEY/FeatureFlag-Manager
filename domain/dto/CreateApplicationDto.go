package dto

type CreateApplicationDTO struct {
	Label       string `json:"label" binding:"required"`
	Description string `json:"description" binding:"required"`
}
