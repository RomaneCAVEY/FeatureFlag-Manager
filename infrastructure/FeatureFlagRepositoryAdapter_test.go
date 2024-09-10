package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/entities"
	"github.com/stretchr/testify/assert"
)

var featureFlagRepository = FeatureFlagRepository{Collection: ConnectDBTest()}

var configTest = Config{
	host:     "localhost",
	port:     "5432",
	user:     "postgres",
	password: "docker",
	dbname:   "postgres",
}

func ConnectDBTest() *sql.DB {
	port, _ := strconv.Atoi(configTest.port)
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configTest.host, port, configTest.user, configTest.password, configTest.dbname)

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

func initDBTests() {
	featureFlagRepository.Collection.Exec("create table if not exists feature_flags (Id  serial,slug VARCHAR(50),Label VARCHAR(50),isEnabled BOOL,Application VARCHAR(50),Projects VARCHAR(50),Owners VARCHAR(50),description VARCHAR(50),CreatedAt TIMESTAMP WITH TIME ZONE,UpdatedAt TIMESTAMP WITH TIME ZONE,CreatedBy VARCHAR(200),UpdatedBy VARCHAR(50),PRIMARY KEY (slug, application));")
}

func CleanDb() {
	featureFlagRepository.Collection.Exec("DELETE FROM feature_flags;")
}
func CloseDb() {
	featureFlagRepository.Collection.Close()
}

func BuildFeatureFlag() entities.FeatureFlag {
	var flag = entities.FeatureFlag{
		Label:       "label",
		Slug:        "label",
		Application: "Application",
		IsEnabled:   true,
		Description: "description",
		Projects:    []string{"iot"},
		Owners:      []string{"iot"},
		UpdatedBy:   "Romane",
		CreatedBy:   "Romane",
	}
	return flag

}

func Test_SaveFeatureFlag_ReturnTheSavedFlag_WhenTheCoupleApplicationAndLabelIsUnique(t *testing.T) {
	initDBTests()
	flag := BuildFeatureFlag()
	savedFlag, errorSave := featureFlagRepository.Save(flag)
	getFlag, count, errGet := featureFlagRepository.FindByApplication(flag.Application, 0, 10)
	CleanDb()

	assert.Equal(t, savedFlag, getFlag[0])
	assert.Equal(t, count, 1)
	assert.NoError(t, errorSave)
	assert.NoError(t, errGet)

}

func Test_SaveFeatureFlag_ShouldReturnAnError_WhenTheCoupleApplicationAndLabelIsNotUnique(t *testing.T) {
	initDBTests()
	flag := BuildFeatureFlag()
	_, err := featureFlagRepository.Save(flag)
	_, errorSave := featureFlagRepository.Save(flag)
	CleanDb()

	assert.NoError(t, err)
	assert.ErrorContains(t, errorSave, "Label already exists for this application. Find another label")

}

func Test_FindAllFeatureFlags_ShouldReturnListOf1Flag_When1FlagWasSaved(t *testing.T) {
	flag := BuildFeatureFlag()

	savedFlag, errorSave := featureFlagRepository.Save(flag)
	getFlags, count, errorGet := featureFlagRepository.FindAll(0, 10)
	CleanDb()

	assert.Equal(t, savedFlag, getFlags[0])
	assert.Equal(t, len(getFlags), 1)
	assert.Equal(t, count, 1)
	assert.NoError(t, errorSave)
	assert.NoError(t, errorGet)

}

func Test_FindByApplication_ShouldReturnListOf2Flags_When2FlagsWereSavedWithRequiredApplication(t *testing.T) {
	flag1 := BuildFeatureFlag()
	flag2 := BuildFeatureFlag()
	flag2.Label = "test"
	flag2.Slug = "test"
	savedFlag1, errorSave1 := featureFlagRepository.Save(flag1)
	savedFlag2, errorSave2 := featureFlagRepository.Save(flag2)
	getFlags, count, errorGet := featureFlagRepository.FindByApplication(flag1.Application, 0, 10)
	CleanDb()

	assert.Equal(t, savedFlag1, getFlags[0])
	assert.Equal(t, savedFlag2, getFlags[1])
	assert.Equal(t, len(getFlags), 2)
	assert.Equal(t, count, 2)
	assert.NoError(t, errorSave1)
	assert.NoError(t, errorSave2)
	assert.NoError(t, errorGet)

}

func Test_FindByApplication_ShouldReturnEmptyArray_WhenNoFlagWereSavedWithRequiredApplication(t *testing.T) {
	flag := BuildFeatureFlag()
	flag.Application = "iot_test"

	_, errorSave := featureFlagRepository.Save(flag)
	getFlags, count, errorGet := featureFlagRepository.FindByApplication("Application", 0, 10)
	CleanDb()

	assert.Equal(t, count, 0)
	assert.Equal(t, len(getFlags), 0)
	assert.NoError(t, errorSave)
	assert.NoError(t, errorGet)

}

func Test_SaveChangesFeatureFlag_ShouldReturnModifiedFeatureFlag_WhenFlagWasSavedWithRequiredId(t *testing.T) {
	flag := BuildFeatureFlag()
	_, errorSave := featureFlagRepository.Save(flag)
	getFlags, count, _ := featureFlagRepository.FindByApplication(flag.Application, 0, 10)
	getFlag := getFlags[0]
	flag.Label = "test"
	flag.IsEnabled = false

	modifiedFlag, errorModification := featureFlagRepository.SaveChangesFeatureFlag(getFlag.Id, "test", false, "Romane")
	flag.Id = getFlag.Id
	CleanDb()

	assert.Equal(t, count, 1)
	assert.Equal(t, flag.Label, modifiedFlag.Label)
	assert.Equal(t, flag.Description, modifiedFlag.Description)
	assert.Equal(t, flag.Application, modifiedFlag.Application)
	assert.Equal(t, flag.Owners, modifiedFlag.Owners)
	assert.Equal(t, flag.Projects, modifiedFlag.Projects)

	assert.NoError(t, errorSave)
	assert.NoError(t, errorModification)
}

func TestSaveChangesFeatureFlag_ShouldReturAnError_WhenNoFlagWereSavedWithRequiredId(t *testing.T) {
	flag := BuildFeatureFlag()
	_, errorSave := featureFlagRepository.Save(flag)
	_, errorModification := featureFlagRepository.SaveChangesFeatureFlag(10000, "test", false, "Romane")
	CleanDb()

	assert.ErrorContains(t, errorModification, "no feature-flag with this id")
	assert.NoError(t, errorSave)
}

func Test_RemoveFeatureFlag_ShouldRemoveFlag_WhenFlagWasSavedWithRequiredId(t *testing.T) {
	flag := BuildFeatureFlag()
	savedFlag, errorSave := featureFlagRepository.Save(flag)
	beforeDeletion, beforeCount, errBeforeDeletion := featureFlagRepository.FindAll(0, 10)
	errorRemove := featureFlagRepository.RemoveFeatureFlag(savedFlag.Id)
	afterDeletion, afterCount, errAfterDeletion := featureFlagRepository.FindAll(0, 10)
	CleanDb()

	assert.NoError(t, errorSave)
	assert.NoError(t, errorRemove)
	assert.NoError(t, errAfterDeletion)
	assert.NoError(t, errBeforeDeletion)
	assert.Equal(t, len(beforeDeletion), 1)
	assert.Equal(t, len(afterDeletion), 0)
	assert.Equal(t, beforeCount, 1)
	assert.Equal(t, afterCount, 0)
}

func Test_RemoveFeatureFlag_ShouldReturnAnError_WhenNoFlagWasSavedWithRequiredId(t *testing.T) {
	errorRemove := featureFlagRepository.RemoveFeatureFlag(1)
	assert.ErrorContains(t, errorRemove, "no feature-flag with this id")

	CleanDb()
	CloseDb()

}
