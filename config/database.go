package config

import (
	"github.com/emejotaw/voting-service/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {

	dsn := "database.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Survey{}, &entity.Option{}, &entity.Vote{})

	if err != nil {
		panic(err)
	}

	return db
}
