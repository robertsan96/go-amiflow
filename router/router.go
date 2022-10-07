package router

import "github.com/gin-gonic/gin"

type ginHandler struct {
	GinContext gin.Context
}
type PersistentStorageGroupRouter[T any] interface {
	ReadPersistentStorageGroup(p T)
}
