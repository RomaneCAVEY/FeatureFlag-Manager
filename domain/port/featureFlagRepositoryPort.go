package port

import "github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/entities"

type FeatureFlagRepositoryPort interface {
	Save(entities.FeatureFlag) (entities.FeatureFlag, error)
	FindAll(int, int) ([]entities.FeatureFlag, int, error)
	FindByApplication(string, int, int) ([]entities.FeatureFlag, int, error)
	FindById(uint32) (entities.FeatureFlag, error)
	RemoveFeatureFlag(uint32) error
	SaveChangesFeatureFlag(uint32, string, bool, string) (entities.FeatureFlag, error)

}
