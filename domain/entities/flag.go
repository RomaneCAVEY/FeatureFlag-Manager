package entities

import (
	"time"

	"github.com/gosimple/slug"
)

type FeatureFlag struct {
	Id          uint32    `json:"id"`
	Slug        string    `json:"slug"`
	Label       string    `json:"label"`
	IsEnabled   bool      `json:"isEnabled"`
	Application string    `json:"application"`
	Projects    []string  `json:"projects"`
	Owners      []string  `json:"owners"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedBy   string    `json:"updatedBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Description string    `json:"description"`
}

func MakeFeatureFlag(projects []string, label string, isEnabled *bool, application string, owners []string, description string, user User) FeatureFlag {
	var featureFlag = FeatureFlag{
		Slug:        slug.Make(label),
		Label:       label,
		IsEnabled:   *isEnabled,
		Application: application,
		Projects:    projects,
		Owners:      owners,
		CreatedBy:   user.GivenName + " " + user.FamilyName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UpdatedBy:   user.GivenName + " " + user.FamilyName,
		Description: description,
	}

	return featureFlag
}
