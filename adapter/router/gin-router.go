package router

import (
	"github.com/gin-gonic/gin"
	"github.com/robertsan96/go-amiflow/adapter/presenter"
	"github.com/robertsan96/go-amiflow/usecase"
)

// Gin Specific Routes
// -------------------
//
// Loads every app route and sets up the dependency chain.
//

func NewGinEngine() *gin.Engine {
	engine := gin.Default()

	setupPsgRoutes(engine)

	return engine
}

func setupPsgRoutes(engine *gin.Engine) {
	psgService := usecase.NewPersistentStorageGroupService()
	psgJsonPresenter := presenter.NewPersistentStorageGroupJsonPresenter(psgService)
	ginPsgRouter := NewGinPersistentStorageGroupRouter(psgJsonPresenter)

	engine.GET("/psg", ginPsgRouter.ReadPersistentStorageGroup)
}
