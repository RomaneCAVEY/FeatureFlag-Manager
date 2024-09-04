package dto

type ModifyApplicationDTO struct {
	Label       string `json:"label" binding:"required"`
	Description string `json:"description" binding:""`
}
