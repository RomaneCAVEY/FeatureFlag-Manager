package entities

import (
	"time"
)

type Application struct {
	Id          uint32    `json:"id"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func MakeApplication(label string, description string) Application {
	var application = Application{
		Label:       label,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now()}

	return application
}
