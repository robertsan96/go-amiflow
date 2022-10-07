package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertsan96/go-amiflow/presenter"
	"github.com/robertsan96/go-amiflow/router"
)

type ginPersistentStorageGroupRouter struct {
	Presenter presenter.PersistentStorageGroupPresenter
}

func NewGinPersistentStorageGroupRouter(
	presenter presenter.PersistentStorageGroupPresenter,
) router.PersistentStorageGroupRouter[*gin.Context] {
	return &ginPersistentStorageGroupRouter{
		Presenter: presenter,
	}
}

func (router *ginPersistentStorageGroupRouter) ReadPersistentStorageGroup(ctx *gin.Context) {
	dto := router.Presenter.ReadPersistentStorageGroup()
	ctx.JSON(http.StatusOK, gin.H{"data": dto})
}
