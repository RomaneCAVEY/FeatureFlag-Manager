package service

import (
	"errors"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/dto"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/entities"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/tree/main/domain/port"

)

type FeatureFlagManagerService struct {
	repository            port.FeatureFlagRepositoryPort
	applicationRepository port.ApplicationRepositoryPort
}

func (s *FeatureFlagManagerService) CreateAFeatureFlag(d dto.CreateAFeatureFlagDTO, user entities.User) (*entities.FeatureFlag, error) {
	flag := entities.MakeFeatureFlag(d.Projects, d.Label, d.IsEnabled, d.Application, d.Owners, d.Description, user)
	_, error := s.applicationRepository.FindByLabel(d.Application)
	if error != nil {
		return nil, error
	}

	createdFlag, error := s.repository.Save(flag)
	if error != nil {
		return nil, error
	}
	return &createdFlag, nil
}

func (s *FeatureFlagManagerService) GetAllFeatureFlags(start int, end int) (*[]entities.FeatureFlag, int, error) {
	if end < start {
		return nil, 0, errors.New("Error in pagination: start greater than end")
	}
	listAllFeatureFlags, count, error := s.repository.FindAll(start, end)
	if error != nil {
		return nil, 0, error
	}

	return &listAllFeatureFlags, count, nil
}

func (s *FeatureFlagManagerService) GetFeatureFlagsByApplication(application string, start int, end int) (*[]entities.FeatureFlag, int, error) {
	if end < start {
		return nil, 0, errors.New("Error in pagination: start greater than end")
	}
	_, error := s.applicationRepository.FindByLabel(application)
	if error != nil {
		return nil, 0, error
	}
	listFeatureFlagsByApplication, count, error := s.repository.FindByApplication(application, start, end)

	return &listFeatureFlagsByApplication, count, nil
}

func (s *FeatureFlagManagerService) GetFeatureFlagsById(id uint32) (*entities.FeatureFlag, error) {
	getFeatureFlag, error := s.repository.FindById(id)
	if error != nil {
		return nil, error
	}

	return &getFeatureFlag, nil
}

func (s *FeatureFlagManagerService) ModifyFeatureFlag(id uint32, d dto.ModifyFeatureFlagDTO, user entities.User) (*entities.FeatureFlag, error) {
	modifiedFeatureFlag, error := s.repository.SaveChangesFeatureFlag(id, d.Label, d.IsEnabled, user.GivenName+" "+user.FamilyName)
	if error != nil {
		return nil, error
	}
	return &modifiedFeatureFlag, nil
}

func MakeFeatureFlagManagerService(repoFeatureFlags port.FeatureFlagRepositoryPort, repoApplication port.ApplicationRepositoryPort) FeatureFlagManagerService {
	var s FeatureFlagManagerService
	s.repository = repoFeatureFlags
	s.applicationRepository = repoApplication
	return s
}

func (s *FeatureFlagManagerService) DeleteFeatureFlag(id uint32) error {
	error := s.repository.RemoveFeatureFlag(id)
	return error
}
