package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/entities"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/entities"
	"github.com/stretchr/testify/assert"
)

var TestApplicationRepository = ApplicationRepository{Collection: ConnectDBTestApplication()}

var configTestApplication = Config{
	host:     "localhost",
	port:     "5432",
	user:     "postgres",
	password: "docker",
	dbname:   "postgres",
}

func ConnectDBTestApplication() *sql.DB {
	port, _ := strconv.Atoi(configTestApplication.port)
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configTestApplication.host, port, configTestApplication.user, configTestApplication.password, configTestApplication.dbname)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to db!")
	return db
}
func initDBTestsApplication() {
	TestApplicationRepository.Collection.Exec("create table if not exists applications (Id  serial,Label VARCHAR(50),Description VARCHAR(50),CreatedAt TIMESTAMP WITH TIME ZONE,UpdatedAt TIMESTAMP WITH TIME ZONE,PRIMARY KEY (label));")

}

func CleanDbApplication() {
	TestApplicationRepository.Collection.Exec("DELETE FROM applications;")
}

func BuildApplication() entities.Application {
	var application = entities.Application{
		Label:       "label",
		Description: "description",
	}
	return application

}
func Test_SaveApplication_ShouldReturnTheSavedApplication_WhenLabelIsUnique(t *testing.T) {
	initDBTestsApplication()
	application := BuildApplication()
	savedApplication, errorSave := TestApplicationRepository.Save(application)
	getApplication, errGet := TestApplicationRepository.FindByLabel(application.Label)
	CleanDbApplication()

	assert.Equal(t, savedApplication, getApplication)
	assert.NoError(t, errorSave)
	assert.NoError(t, errGet)

}

func Test_SaveApplication_ShouldReturnAnError_WhenLabelIsNotUnique(t *testing.T) {
	application := BuildApplication()
	_, err := TestApplicationRepository.Save(application)
	_, errorSave := TestApplicationRepository.Save(application)
	CleanDbApplication()

	assert.NoError(t, err)
	assert.Equal(t, errorSave, errDuplicationKey)

}

func Test_FindAllApplications_ShouldReturnListOf1Applications_When1ApplicationWasSaved(t *testing.T) {
	application := BuildApplication()

	savedApplication, errorSave := TestApplicationRepository.Save(application)
	getApplications, count, errorGet := TestApplicationRepository.FindAll(0, 10)
	CleanDbApplication()

	assert.Equal(t, savedApplication, getApplications[0])
	assert.Equal(t, len(getApplications), 1)
	assert.Equal(t, count, 1)
	assert.NoError(t, errorSave)
	assert.NoError(t, errorGet)
}

func Test_FindById_ShouldBuildApplicationWithRequiredId_When1ApplicationWasSavedWithRequiredId(t *testing.T) {
	application := BuildApplication()

	savedApplication, errorSave := TestApplicationRepository.Save(application)
	getApplication, _ := TestApplicationRepository.FindByLabel(savedApplication.Label)
	id := getApplication.Id
	FoundApplicationById, errorGet := TestApplicationRepository.FindById(id)
	CleanDbApplication()

	assert.Equal(t, savedApplication, FoundApplicationById)
	assert.NoError(t, errorSave)
	assert.NoError(t, errorGet)

}

func Test_FindById_ShouldReturnAnError_WhenNoApplicationWasSavedWithRequiredId(t *testing.T) {
	_, errorGet := TestApplicationRepository.FindById(100)
	CleanDbApplication()

	assert.ErrorContains(t, errorGet, "no Application with this id")

}

func Test_FindByLabel_ShouldBuildApplicationWithRequiredLabel_When1ApplicationWasSavedWithRequiredLabel(t *testing.T) {
	application := BuildApplication()
	savedApplication, errorSave := TestApplicationRepository.Save(application)
	FoundApplicationById, errorGet := TestApplicationRepository.FindByLabel(application.Label)
	CleanDbApplication()

	assert.Equal(t, savedApplication, FoundApplicationById)
	assert.NoError(t, errorGet)
	assert.NoError(t, errorSave)

}

func Test_FindByLabel_ShouldReturnAnError_WhenNoApplicationWasSavedWithRequiredLabel(t *testing.T) {
	_, errorGet := TestApplicationRepository.FindByLabel("doesn't exist")
	CleanDbApplication()

	assert.ErrorContains(t, errorGet, "no Application with this label")

}

func Test_UpdateApplication_ShouldReturnModifiedApplication_WhenApplicationWasSavedWithRequiredId(t *testing.T) {
	application := BuildApplication()
	savedApplication, errorSave := TestApplicationRepository.Save(application)
	savedApplication.Label = "test"
	savedApplication.Description = "description test"

	modifiedApplication, errorModification := TestApplicationRepository.UpdateApplication(savedApplication.Id, "test", "description test")
	CleanDbApplication()

	assert.Equal(t, savedApplication.Label, modifiedApplication.Label)
	assert.Equal(t, savedApplication.Description, modifiedApplication.Description)
	assert.Equal(t, savedApplication.CreatedAt, modifiedApplication.CreatedAt)

	assert.NoError(t, errorModification)
	assert.NoError(t, errorSave)
}

func Test_UpdateApplication_ShouldReturAnError_WhenNoApplicationWereSaved(t *testing.T) {
	_, errorModification := TestApplicationRepository.UpdateApplication(10000, "test", "test")
	CleanDbApplication()

	assert.ErrorContains(t, errorModification, "no Application with this id")
}

func Test_RemoveApplication_ShouldRemoveApplication_WhenApplicationWasSavedWithRequiredId(t *testing.T) {
	application := BuildApplication()
	applicationSaved, errorSave := TestApplicationRepository.Save(application)
	errorRemove := TestApplicationRepository.RemoveApplication(applicationSaved.Id)
	NoMoreFlag, errorFind := TestApplicationRepository.FindByLabel(applicationSaved.Label)
	CleanDbApplication()

	assert.Equal(t, NoMoreFlag, entities.Application{})
	assert.NoError(t, errorSave)
	assert.ErrorContains(t, errorFind, "no Application with this label")
	assert.NoError(t, errorRemove)
}

func Test_RemoveApplication_ShouldReturnError_WhenNoApplicationWasSavedWithRequiredId(t *testing.T) {
	errorRemove := TestApplicationRepository.RemoveApplication(10000)
	CleanDbApplication()

	assert.ErrorContains(t, errorRemove, "no Application with this id")
}
