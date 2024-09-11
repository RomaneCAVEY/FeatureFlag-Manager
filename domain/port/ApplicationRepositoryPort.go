package port

import (
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
)

type ApplicationRepositoryPort interface {
	Save(entities.Application) (entities.Application, error)
	FindAll(int, int) ([]entities.Application, int, error)
	FindById(uint32) (entities.Application, error)
	FindByLabel(string) (entities.Application, error)
	RemoveApplication(uint32) error
	UpdateApplication(uint32, string, string) (entities.Application, error)

}
