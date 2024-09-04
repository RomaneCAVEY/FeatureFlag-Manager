package dto

type ModifyFeatureFlagDTO struct {
	Label     string `json:"label" binding:"required"`
	IsEnabled bool   `json:"isEnabled" binding:""`
}
