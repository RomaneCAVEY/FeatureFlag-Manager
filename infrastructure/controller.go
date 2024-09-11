package infrastructure

import (
	"errors"
	"log"
	"strings"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
)

/*
Validate if the user is from the Compagny , else return an error
*/

func ValidateRequestFromCompagnyUser(user entities.User) error {
	email := user.Email

	if !strings.HasSuffix(email, "@Compagny.com") {
		log.Println("An invalid user has tried to create a feature flag. The email of the user is: ", email)
		return errors.New("invalid user")
	} else {
		return nil
	}
}
