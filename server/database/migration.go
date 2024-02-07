package database

import (
	"fmt"
	"waysgallery/models"
	"waysgallery/pkg/postgresql"
)

func RunMigration() {
	err := postgresql.DB.AutoMigrate(
		&models.WaysGalleryUser{},
		&models.WaysGalleryProfile{},
		&models.WaysGalleryArt{},
		&models.WaysGalleryFollowing{},
		&models.WaysGalleryPost{},
		&models.WaysGalleryOrder{},
		&models.WaysGalleryProject{},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Migration Success")
}
