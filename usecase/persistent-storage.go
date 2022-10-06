package usecase

type PersistentStorageService interface {
	CreatePersistentStorage()
	DeletePersistentStorage()
}

type persistentStorageService struct {
	// Repos
}

func NewPersistentStorageService(
// repos
) PersistentStorageService {
	return &persistentStorageService{}
}

func (service *persistentStorageService) CreatePersistentStorage() {}
func (service *persistentStorageService) DeletePersistentStorage() {}
