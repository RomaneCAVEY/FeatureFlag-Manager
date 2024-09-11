package infrastructure

import (
	"testing"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
)

func Test_ValidateRequestFromCompagnyUser_ShouldNotValidateUser_WhenWrongEmailWithoutSymbolAt(t *testing.T) {
	user := entities.User{Email: "romaneatCompagny.io"}
	error := ValidateRequestFromCompagnyUser(user)
	if error == nil {
		t.Fatal("The user shoudn't be allowed to connect")
	}

}

func Test_ValidateRequestFromCompagnyUser_ShouldNotValidateUser_WhenWrongEmailAtPamplemousseIo(t *testing.T) {
	user := entities.User{Email: "romanea@pamplemousse.io"}
	error := ValidateRequestFromCompagnyUser(user)
	if error == nil {
		t.Fatal("The user shoudn't be allowed to connect")
	}

}
func Test_ValidateRequestFromCompagnyUser_ShouldNotValidateUser_WhenWrongEmailAtCompagnyCom(t *testing.T) {
	user := entities.User{Email: "romanea@Compagny.com"}
	error := ValidateRequestFromCompagnyUser(user)
	if error == nil {
		t.Fatal("The user shoudn't be allowed to connect")
	}

}

func Test_ValidateRequestFromCompagnyUser_ShouldValidateUser_WhenRightEmailAtCompagnyIo(t *testing.T) {
	user := entities.User{Email: "romane@Compagny.io"}
	error := ValidateRequestFromCompagnyUser(user)
	if error != nil {
		t.Fatal("The user shoudn't be allowed to connect")
	}

}
