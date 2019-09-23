package migrations

import (
	"github.com/danny/todo/todo/common"
	"github.com/danny/todo/todo/models"
)

func Migrate() {
	common.DB.AutoMigrate(&models.Person{})
}
