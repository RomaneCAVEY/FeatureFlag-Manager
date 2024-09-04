package dto

type CreateFeatureFlagDTO struct {
	Label       string   `json:"label" binding:"required"`
	Application string   `json:"application" binding:"required"`
	IsEnabled   *bool    `json:"isEnabled" binding:"required"`
	Description string   `json:"description" binding:""`
	Projects    []string `json:"projects" binding:""`
	Owners      []string `json:"owners" binding:"required"`
}
