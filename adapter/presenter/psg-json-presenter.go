package presenter

import (
	"github.com/robertsan96/go-amiflow/dto"
	"github.com/robertsan96/go-amiflow/presenter"
	"github.com/robertsan96/go-amiflow/usecase"
)

type persistentStorageGroupJsonPresenter struct {
	PersistentStorageGroupService usecase.PersistentStorageGroupService
}

func NewPersistentStorageGroupJsonPresenter(
	persistentStorageGroupService usecase.PersistentStorageGroupService,
) presenter.PersistentStorageGroupPresenter {
	return &persistentStorageGroupJsonPresenter{
		PersistentStorageGroupService: persistentStorageGroupService,
	}
}

func (presenter *persistentStorageGroupJsonPresenter) ReadPersistentStorageGroup() *dto.ReadPersistentStorageGroupDto {
	psgEntity := presenter.PersistentStorageGroupService.ReadPersistentStorageGroup()
	return &dto.ReadPersistentStorageGroupDto{
		Id:        psgEntity.Id,
		Name:      psgEntity.Name,
		CreatedAt: psgEntity.CreatedAt,
		UpdatedAt: psgEntity.UpdatedAt,
	}
}
