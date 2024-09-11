package infrastructure

import (
	"errors"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
)

type FeatureFlagRepositiryStupid struct {
	flags []entities.FeatureFlag
	count uint32
}

func (r *FeatureFlagRepositiryStupid) Save(f entities.FeatureFlag) (entities.FeatureFlag, error) {
	f.Id = r.count
	r.flags = append(r.flags, f)
	r.count += 1
	return r.flags[r.count-1], nil

}

func (r *FeatureFlagRepositiryStupid) FindAll() ([]entities.FeatureFlag, error) {
	return r.flags, nil
}

func (r *FeatureFlagRepositiryStupid) FindByApplication(s string) ([]entities.FeatureFlag, error) {
	var listFeatureFlagsByapplication = []entities.FeatureFlag{}
	for i := 0; i < (len(r.flags)); i++ {
		if r.flags[i].Application == s {
			listFeatureFlagsByapplication = append(listFeatureFlagsByapplication, r.flags[i])
		}
	}
	return listFeatureFlagsByapplication, nil
}

func (r *FeatureFlagRepositiryStupid) SaveChangesFeatureFlag(id uint32, label string, slug string, isEnabled bool) (entities.FeatureFlag, error) {
	for i := 0; i < (len(r.flags)); i++ {
		if r.flags[i].Id == id {
			r.flags[i].Label = label
			r.flags[i].Slug = slug
			r.flags[i].IsEnabled = isEnabled
			return r.flags[i], nil
		}
	}
	return entities.FeatureFlag{}, errors.New("no feature flag with this id in our data base")
}

func (r *FeatureFlagRepositiryStupid) RemoveFeatureFlag(id uint32) (entities.FeatureFlag, error) {
	for i := 0; i < (len(r.flags)); i++ {
		if r.flags[i].Id == id {
			var deletedFeatureFlag entities.FeatureFlag = r.flags[i]
			r.flags = append(r.flags[:i], r.flags[(i+1):]...)
			return deletedFeatureFlag, nil
		}
	}
	return entities.FeatureFlag{}, errors.New("no feature flag with this id in our data base")
}
