package router

type PersistentStorageGroupRouter[T any] interface {
	ReadPersistentStorageGroup(p T)
}
