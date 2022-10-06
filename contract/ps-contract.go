package contract

type PersistentStorageContract interface {
	CreatePersistentStorageGroup(name string)
	DeletePersistentStorageGroup(name string)
}
