package service

import (
	"errors"
	"testing"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/dto"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
	mock_port 	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/port/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var flag = entities.FeatureFlag{
	Label:       "label",
	Application: "Application",
	IsEnabled:   true,
	Description: "description",
	Projects:    []string{"project"},
	Owners:      []string{"test_owner"},
}

func Test_CreateAFeatureFlag_ShouldReturnTheSavedFlag_WhenApplicationExists(t *testing.T) {
	owners := []string{"test_owner"}
	projects := []string{"project"}
	value := true
	var json = dto.CreateAFeatureFlagDTO{
		Label:       "label",
		Application: "application",
		IsEnabled:   &value,
		Description: "description",
		Projects:    projects,
		Owners:      owners,
	}
	var Application = entities.Application{
		Label:       "Application",
		Description: "description"}
	var user = entities.User{GivenName: "Romane", FamilyName: "Cavey"}

	flagTest := entities.MakeFeatureFlag(json.Projects, json.Label, json.IsEnabled, json.Application, json.Owners, json.Description, user)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().FindByLabel(json.Application).Return(Application, nil)
	MockFeatureFlagRepository.EXPECT().Save(gomock.Any()).Return(flagTest, nil)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	featureFlagCreated, error := managerApplicationFeatureFlag.CreateAFeatureFlag(json, user)

	assert.NoError(t, error)
	assert.Equal(t, *featureFlagCreated, flagTest)

}

func Test_CreateAFeatureFlag_ShouldReturnAnErrorAndEmptyFlag_WhenApplicationDoesntExists(t *testing.T) {
	owners := []string{"test_owner"}
	projects := []string{"project"}
	value := true
	var json = dto.CreateAFeatureFlagDTO{
		Label:       "label",
		Application: "Application",
		IsEnabled:   &value,
		Description: "description",
		Projects:    projects,
		Owners:      owners,
	}
	var user = entities.User{GivenName: "Romane", FamilyName: "Cavey"}
	errExpected := errors.New("no application with this label")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	//TODO: ignore some parameters (as time)
	MockApplicationRepository.EXPECT().FindByLabel(json.Application).Return(entities.Application{}, errExpected)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	_, error := managerApplicationFeatureFlag.CreateAFeatureFlag(json, user)

	assert.ErrorContains(t, error, "no application with this label")
}

func Test_CreateAFeatureFlag_ShouldReturnAnError_WhenErrorHappenedWithDb(t *testing.T) {
	owners := []string{"test_owner"}
	projects := []string{"project"}
	value := true
	var json = dto.CreateAFeatureFlagDTO{
		Label:       "label",
		Application: "Application",
		IsEnabled:   &value,
		Description: "description",
		Projects:    projects,
		Owners:      owners,
	}
	var Application = entities.Application{
		Label:       "Application",
		Description: "description"}

	var user = entities.User{GivenName: "Romane", FamilyName: "Cavey"}

	flagTest := entities.MakeFeatureFlag(json.Projects, json.Label, json.IsEnabled, json.Application, json.Owners, json.Description, user)
	errorBD := errors.New("error with DB")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().FindByLabel(json.Application).Return(Application, nil)
	MockFeatureFlagRepository.EXPECT().Save(gomock.Any()).Return(flagTest, errorBD)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	_, error := managerApplicationFeatureFlag.CreateAFeatureFlag(json, user)

	assert.Equal(t, errorBD, error)
	assert.ErrorContains(t, error, "error with DB")

}
func Test_GetAllFeatureFlag_ShouldReturnEmptyArray_WhenNoFlagsWereSaved(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	emptyFlagsArray := []entities.FeatureFlag{}
	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	MockFeatureFlagRepository.EXPECT().FindAll(0, 1).Return(emptyFlagsArray, 0, nil)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	GetAllFeatureFlags, count, error := managerApplicationFeatureFlag.GetAllFeatureFlags(0, 1)

	assert.NoError(t, error)
	assert.Equal(t, count, 0)
	assert.Equal(t, *GetAllFeatureFlags, emptyFlagsArray)

}

func Test_GetAllFeatureFlag_ShouldReturnAllFeatureFlags_WhenFlagsWereSaved(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	AllFlags := []entities.FeatureFlag{flag, flag}
	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	MockFeatureFlagRepository.EXPECT().FindAll(0, 10).Return(AllFlags, 2, nil)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	GetAllFeatureFlags, count, error := managerApplicationFeatureFlag.GetAllFeatureFlags(0, 10)

	assert.NoError(t, error)
	assert.Equal(t, *GetAllFeatureFlags, AllFlags)
	assert.Equal(t, len(*GetAllFeatureFlags), 2)
	assert.Equal(t, count, 2)

}
func Test_GetAllFeatureFlag_ShouldReturnAnError_WhenAnErrorHappenedWithInsert(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	errorExpected := errors.New("something go wrong with the request insert in the db")

	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	MockFeatureFlagRepository.EXPECT().FindAll(0, 10).Return([]entities.FeatureFlag{}, 0, errorExpected)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	_, count, error := managerApplicationFeatureFlag.GetAllFeatureFlags(0, 10)

	assert.ErrorContains(t, error, "something go wrong with the request insert in the db")
	assert.Equal(t, count, 0)

}

func Test_GetAllFeatureFlag_ShouldReturnAnError_WhenStartIsGreaterThanEnd(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	_, count, error := managerApplicationFeatureFlag.GetAllFeatureFlags(10, 0)

	assert.ErrorContains(t, error, "Error in pagination: start greater than end")
	assert.Equal(t, count, 0)

}

func Test_GetFeatureFlagByApplication_ShouldReturnFeatureFlagsWithRequiredApplication_WhenRequiredApplicationExists(t *testing.T) {
	owners := []string{"test_owner"}
	projects := []string{"project"}
	value := true
	var json = dto.CreateAFeatureFlagDTO{
		Label:       "label",
		Application: "Application",
		IsEnabled:   &value,
		Description: "description",
		Projects:    projects,
		Owners:      owners,
	}
	var user = entities.User{GivenName: "Romane", FamilyName: "Cavey"}

	flagTest := entities.MakeFeatureFlag(json.Projects, json.Label, json.IsEnabled, json.Application, json.Owners, json.Description, user)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	listFeatureFlagsByApplication := []entities.FeatureFlag{flagTest}
	application := entities.Application{}
	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	MockFeatureFlagRepository.EXPECT().FindByApplication(flagTest.Application, 0, 1).Return(listFeatureFlagsByApplication, 1, nil)
	MockApplicationRepository.EXPECT().FindByLabel(flagTest.Application).Return(application, nil)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	GetFlags, count, error := managerApplicationFeatureFlag.GetFeatureFlagsByApplication(flagTest.Application, 0, 1)

	assert.NoError(t, error)
	assert.Equal(t, *GetFlags, listFeatureFlagsByApplication)
	assert.Equal(t, count, 1)

}

func Test_GetFeatureFlagByApplication_ShouldReturnAnError_WhenRequiredApplicationDoesntExists(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	errorExpected := errors.New("no Application with this label")

	MockApplicationRepository.EXPECT().FindByLabel("Application which doesn't exist").Return(entities.Application{}, errorExpected)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	_, count, error := managerApplicationFeatureFlag.GetFeatureFlagsByApplication("Application which doesn't exist", 0, 1)

	assert.Equal(t, error, errorExpected)
	assert.Equal(t, count, 0)
	assert.ErrorContains(t, error, "no Application with this label")

}

func Test_ModifyFeatureFlag_ShouldReturnTheModifiedFeatureFlagWithRequiredId_WhenTheIdExists(t *testing.T) {

	owners := []string{"test_owner"}
	projects := []string{"project"}
	var expectedFlagModified = entities.FeatureFlag{
		Label:       "label",
		Slug:        "label",
		Application: "Application",
		IsEnabled:   true,
		Description: "description",
		Projects:    projects,
		Owners:      owners,
	}
	var modifyFeatureFlagDTO = dto.ModifyFeatureFlagDTO{
		Label:     "label",
		IsEnabled: true,
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var id uint32 = 0
	var user = entities.User{GivenName: "Romane", FamilyName: "Cavey"}

	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	MockFeatureFlagRepository.EXPECT().SaveChangesFeatureFlag(id, "label", true, "Romane Cavey").Return(expectedFlagModified, nil)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	flagModified, error := managerApplicationFeatureFlag.ModifyFeatureFlag(0, modifyFeatureFlagDTO, user)

	assert.NoError(t, error)
	assert.Equal(t, *flagModified, expectedFlagModified)

}

func Test_ModifyFeatureFlag_ShouldReturnAnErrord_WhenTheIdDoesntExists(t *testing.T) {

	var modifyFeatureFlagDTO = dto.ModifyFeatureFlagDTO{
		Label:     "label",
		IsEnabled: true,
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var id uint32 = 111
	var user = entities.User{GivenName: "Romane", FamilyName: "Cavey"}

	errorExpected := errors.New("no feature-flag with this id")
	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	MockFeatureFlagRepository.EXPECT().SaveChangesFeatureFlag(id, "label", true, "Romane Cavey").Return(entities.FeatureFlag{}, errorExpected)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	_, error := managerApplicationFeatureFlag.ModifyFeatureFlag(id, modifyFeatureFlagDTO, user)

	assert.ErrorContains(t, error, "no feature-flag with this id")

}

func Test_DeleteFeatureFlag_ShouldReturnNoError_WhenRequiredIdExists(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var id uint32 = 0
	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	MockFeatureFlagRepository.EXPECT().RemoveFeatureFlag(id).Return(nil)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	error := managerApplicationFeatureFlag.DeleteFeatureFlag(id)

	assert.NoError(t, error)
}

func Test_DeleteFeatureFlag_ShouldReturnAnError_WhenRequiredIdDoesntExists(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var id uint32 = 0
	expectedError := errors.New("no feature-flag with this id")
	MockFeatureFlagRepository := mock_port.NewMockFeatureFlagRepositoryPort(mockCtrl)
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)

	MockFeatureFlagRepository.EXPECT().RemoveFeatureFlag(id).Return(expectedError)
	managerApplicationFeatureFlag := MakeFeatureFlagManagerService(MockFeatureFlagRepository, MockApplicationRepository)
	error := managerApplicationFeatureFlag.DeleteFeatureFlag(id)

	assert.ErrorContains(t, error, "no feature-flag with this id")

}
