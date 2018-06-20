package main

import (
	"src/core"
	"src/models"
)

//---------------------------------------------------------------------
// Consts
//---------------------------------------------------------------------

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------

//---------------------------------------------------------------------
// Method
//---------------------------------------------------------------------
/**
 * Update tables by based on model structs.
 * (*) $ go run ./migrate.go
 */
func main() {

	db := core.GetMasterConnection()

	// close connection after function is finished.
	defer db.Close()

	// Theme
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&models.Theme{})
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&models.Theme_Stage{})

	// Style
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&models.Style{})
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&models.Style_Stage{})

	// Item
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&models.Item{})
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&models.Item_Stage{})

	// Whitelist
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&models.Whitelist{})

	// UniqueContents
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&models.UniqueContent{})
	db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&models.UniqueContent_Stage{})

}
