package port

import (
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/dto"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
)

type ApplicationManagerPort interface {
	CreateAnApplication(dto.CreateAnApplicationDTO) (*entities.Application, error)
	GetAllApplications(int, int) (*[]entities.Application, int, error)
	GetApplicationById(uint32) (*entities.Application, error)
	DeleteApplication(uint32) (*entities.Application, error)
	ModifyApplication(uint32, dto.ModifyApplicationDTO) (*entities.Application, error)

}
