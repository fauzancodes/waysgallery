package database

import (
	"fmt"
	"waysgallery/models"
	"waysgallery/pkg/postgresql"
)

func RunMigration() {
	err := postgresql.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Art{},
		&models.Following{},
		&models.Post{},
		&models.Order{},
		&models.Project{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
