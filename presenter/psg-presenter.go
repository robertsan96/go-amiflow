package presenter

import "github.com/robertsan96/go-amiflow/dto"

type PersistentStorageGroupPresenter interface {
	ReadPersistentStorageGroup() *dto.ReadPersistentStorageGroupDto
}
