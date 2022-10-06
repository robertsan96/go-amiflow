package usecase

import "github.com/robertsan96/go-amiflow/entity"

type PersistentStorageGroupService interface {
	ReadPersistentStorageGroup() *entity.PeristentStorageGroupEntity
	DeletePersistentStorageGroup()
}

type persistentStorageGroupService struct {
	// Repos
}

func NewPersistentStorageGroupService(
// repos
) PersistentStorageGroupService {
	return &persistentStorageGroupService{}
}

func (service *persistentStorageGroupService) ReadPersistentStorageGroup() *entity.PeristentStorageGroupEntity {
	return &entity.PeristentStorageGroupEntity{
		Id:        "1",
		Name:      "TableX",
		CreatedAt: "now",
		UpdatedAt: "now",
	}
}

func (service *persistentStorageGroupService) DeletePersistentStorageGroup() {}
