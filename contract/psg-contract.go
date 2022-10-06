package contract

type PersistentStorageGroupContract interface {
	CreatePersistentStorageGroup(name string)
	DeletePersistentStorageGroup(name string)
}
