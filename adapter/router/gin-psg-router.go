package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/robertsan96/go-amiflow/router"
)

type ginPersistentStorageGroupRouter struct{}

func NewGinPersistentStorageGroupRouter() router.PersistentStorageGroupRouter[*gin.Context] {
	return &ginPersistentStorageGroupRouter{}
}

func (router *ginPersistentStorageGroupRouter) ReadPersistentStorageGroup(ctx *gin.Context) {
	fmt.Println("Contexttt", ctx.Params)
}
