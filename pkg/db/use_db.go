package db

import "gorm.io/gorm"

type AppDb interface {
	UseMysql() *gorm.DB
}

type appDb struct {
}

func (appDb *appDb) UseMysql() *gorm.DB {
	return MySql()
}

func NewAppDb() AppDb {
	return &appDb{}
}
