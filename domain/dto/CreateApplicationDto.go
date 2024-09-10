package dto

type CreateAnApplicationDTO struct {
	Label       string `json:"label" binding:"required"`
	Description string `json:"description" binding:"required"`
}
