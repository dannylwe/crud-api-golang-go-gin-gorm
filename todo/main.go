package main

import (
	"fmt"

	"github.com/danny/todo/todo/common"
	"github.com/danny/todo/todo/migrations"
	"github.com/danny/todo/todo/routers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	var err error
	common.DB, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Printf("could not create sqlitedb %v", err)
	}
	defer common.DB.Close()
	migrations.Migrate()

	r := routers.SetupRouter()

	r.Run(":" + common.PORT)
}
