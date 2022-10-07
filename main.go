package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertsan96/go-amiflow/adapter/presenter"
	"github.com/robertsan96/go-amiflow/adapter/router"
	"github.com/robertsan96/go-amiflow/usecase"
)

func startGinServer() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		psgService := usecase.NewPersistentStorageGroupService()
		psgJsonPresenter := presenter.NewPersistentStorageGroupJsonPresenter(psgService)
		ctx.JSON(http.StatusOK, gin.H{"data": psgJsonPresenter.ReadPersistentStorageGroup()})
		return
	})

	ginPsgRouter := router.NewGinPersistentStorageGroupRouter()

	r.GET("/test", ginPsgRouter.ReadPersistentStorageGroup)
	r.Run(":8080")
}

func main() {
	startGinServer()
}
