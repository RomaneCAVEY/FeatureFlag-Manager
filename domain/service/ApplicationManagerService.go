package service

import (
	"errors"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/dto"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/port"
)

type ApplicationManagerService struct {
	repository port.ApplicationRepositoryPort
}

func (s *ApplicationManagerService) CreateAnApplication(d dto.CreateAnApplicationDTO) (*entities.Application, error) {
	Application := entities.MakeApplication(d.Label, d.Description)
	createdApplication, error := s.repository.Save(Application)
	if error != nil {
		return nil, error
	}
	return &createdApplication, nil
}

func MakeApplicationManagerService(r port.ApplicationRepositoryPort) ApplicationManagerService {
	var s ApplicationManagerService
	s.repository = r
	return s
}

func (s *ApplicationManagerService) GetAllApplications(start int, end int) (*[]entities.Application, int, error) {
	if end < start {
		return nil, 0, errors.New("Error in pagination: start greater than end")
	}
	listAllApplications, count, error := s.repository.FindAll(start, end)
	if error != nil {
		return nil, 0, error
	}

	return &listAllApplications, count, nil
}

func (s *ApplicationManagerService) GetApplicationById(id uint32) (entities.Application, error) {
	ApplicationById, error := s.repository.FindById(id)
	if error != nil {
		return entities.Application{}, error
	}

	return ApplicationById, nil
}

func (s *ApplicationManagerService) ModifyApplication(id uint32, d dto.ModifyApplicationDTO) (*entities.Application, error) {
	modifiedApplication, error := s.repository.UpdateApplication(id, d.Label, d.Description)
	if error != nil {
		return nil, error
	}
	return &modifiedApplication, nil
}

func (s *ApplicationManagerService) DeleteApplication(id uint32) error {
	error := s.repository.RemoveApplication(id)
	return error
}
