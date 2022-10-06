package router

import "github.com/robertsan96/go-amiflow/router"

type ginPersistentStorageGroupRouter struct {
}

func NewGinPersistentStorageGroupRouter() router.PersistentStorageGroupRouter {
	return &ginPersistentStorageGroupRouter{}
}

func (router *ginPersistentStorageGroupRouter) ReadPersistentStorageGroup(id string) {

}
