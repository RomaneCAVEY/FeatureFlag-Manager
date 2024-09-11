package port

import (
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/dto"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
)

type FeatureFlagManagerPort interface {
	CreateAFeatureFlag(dto.CreateAFeatureFlagDTO, entities.User) (*entities.FeatureFlag, error)
	GetAllFeatureFlags(int, int) (*[]entities.FeatureFlag, int, error)
	GetFeatureFlagsByApplication(string, int, int) (*[]entities.FeatureFlag, int, error)
	GetFeatureFlagsById(uint32) (*entities.FeatureFlag, error)
	DeleteFeatureFlag(uint32) (*entities.FeatureFlag, error)
	ModifyFeatureFlag(uint32, dto.ModifyFeatureFlagDTO, entities.User) (*entities.FeatureFlag, error)

}
