package service

import (
	"errors"
	"testing"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/dto"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/entities"
	mock_port "github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/port/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateAnApplication_ShoulReturnTheCreatedApplication_WhenNoProblemHappenedWithDB(t *testing.T) {
	var json = dto.CreateAnApplicationDTO{
		Label:       "label",
		Description: "description"}
	var expectedApplication = entities.Application{
		Label:       "label",
		Description: "description"}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().Save(gomock.Any()).Return(expectedApplication, nil)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	applicationCreated, error := applicationManagerService.CreateAnApplication(json)

	assert.Equal(t, *applicationCreated, expectedApplication)
	assert.NoError(t, error)
}

func Test_CreateAnApplication_ShouldReturnAnError_WhenLabelIsAlreadyUsed(t *testing.T) {
	var json = dto.CreateAnApplicationDTO{
		Label:       "label",
		Description: "description"}
	var expectedApplication = entities.Application{}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	errorExpected := errors.New("there is already a an application with this label")

	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().Save(gomock.Any()).Return(expectedApplication, errorExpected)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	_, error := applicationManagerService.CreateAnApplication(json)

	assert.ErrorContains(t, error, "there is already a an application with this label")

}

func Test_GetAllApplication_ShouldReturnAnEmptyArray_WhenNoApplicationWereSaved(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	AllApplications := []entities.Application{}
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().FindAll(0, 10).Return(AllApplications, 0, nil)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	getApplications, count, error := applicationManagerService.GetAllApplications(0, 10)

	assert.NoError(t, error)
	assert.Equal(t, *getApplications, AllApplications)
	assert.Equal(t, count, 0)

}

func Test_GetAllApplication_ShouldReturnAllApplication_whenThereAreSavedApplications(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var application1 = entities.Application{
		Label:       "label_test_1",
		Description: "description"}
	var application2 = entities.Application{
		Label:       "label_test_2",
		Description: "description"}

	AllApplications := []entities.Application{application1, application2}
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().FindAll(0, 10).Return(AllApplications, 2, nil)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	getApplications, count, error := applicationManagerService.GetAllApplications(0, 10)

	assert.NoError(t, error)
	assert.Equal(t, *getApplications, AllApplications)
	assert.Equal(t, count, 2)

}

func Test_GetAllApplication_ShouldReturnAnError_WhenAProblemHappenedWithDB(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	errorExpected := errors.New("A problem happened with the DB to get applications ")

	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().FindAll(0, 10).Return([]entities.Application{}, 0, errorExpected)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	_, count, error := applicationManagerService.GetAllApplications(0, 10)

	assert.ErrorContains(t, error, "A problem happened with the DB to get applications ")
	assert.Equal(t, count, 0)

}

func Test_GetAllApplication_ShouldReturnAnError_WhenStartGreaterThanEnd(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	_, count, error := applicationManagerService.GetAllApplications(10, 0)

	assert.ErrorContains(t, error, "Error in pagination: start greater than end")
	assert.Equal(t, count, 0)

}

func Test_GetApplicationgById_ShouldReturnTheApplicationWithRequiredId_WhenhRequiredIdExists(t *testing.T) {
	var expectedApplication = entities.Application{
		Id:          0,
		Label:       "label",
		Description: "description"}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var id uint32 = 0
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().FindById(id).Return(expectedApplication, nil)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	getApplication, error := applicationManagerService.GetApplicationById(id)

	assert.NoError(t, error)
	assert.Equal(t, getApplication, expectedApplication)

}

func Test_GetApplicationgById_ShouldReturnAnError_WhenhRequiredIdDoesntExists(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var id uint32 = 9999
	errorExpected := errors.New("no Application with this id")

	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().FindById(id).Return(entities.Application{}, errorExpected)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	_, error := applicationManagerService.GetApplicationById(id)

	assert.ErrorContains(t, error, "no Application with this id")

}

func Test_ModifyApplication_ShouldReturnTheModifiedFlag_WhenRequiredIdExists(t *testing.T) {

	var json = dto.ModifyApplicationDTO{
		Label:       "label",
		Description: "description"}
	var expectedApplication = entities.Application{
		Id:          0,
		Label:       "label",
		Description: "description"}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var id uint32 = 0
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().UpdateApplication(id, expectedApplication.Label, expectedApplication.Description).Return(expectedApplication, nil)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	modifiedApplication, error := applicationManagerService.ModifyApplication(id, json)

	assert.NoError(t, error)
	assert.Equal(t, *modifiedApplication, expectedApplication)

}
func Test_ModifyApplication_ShouldReturnAnError_WhenRequiredIdDoesntExists(t *testing.T) {

	var json = dto.ModifyApplicationDTO{
		Label:       "label",
		Description: "description"}

	var expectedApplication = entities.Application{
		Id:          0,
		Label:       "label",
		Description: "description"}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var id uint32 = 0
	errorExpected := errors.New("no application with this Id")
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().UpdateApplication(id, expectedApplication.Label, expectedApplication.Description).Return(expectedApplication, errorExpected)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	_, error := applicationManagerService.ModifyApplication(id, json)

	assert.ErrorContains(t, error, "no application with this Id")

}

func Test_DeleteApplication_ShouldReturnNoError_WhenRequiredIdExists(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var id uint32 = 0
	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().RemoveApplication(id).Return(nil)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	error := applicationManagerService.DeleteApplication(id)

	assert.NoError(t, error)
}

func Test_DeleteApplication_ShouldReturnAnError_WhenRequiredIdDoesntExists(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var id uint32 = 0
	errorExpected := errors.New("no application with this Id")

	MockApplicationRepository := mock_port.NewMockApplicationRepositoryPort(mockCtrl)
	MockApplicationRepository.EXPECT().RemoveApplication(id).Return(errorExpected)
	applicationManagerService := MakeApplicationManagerService(MockApplicationRepository)
	error := applicationManagerService.DeleteApplication(id)

	assert.ErrorContains(t, error, "no application with this Id")
}
