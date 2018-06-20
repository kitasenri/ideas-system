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
 * Create tables by based on model structs.
 * (*) $ go run ./create.go
 */
func main() {

	db := core.GetMasterConnection()

	// close connection after function is finished.
	defer db.Close()

	// Theme
	db.DropTableIfExists(&models.Theme{})
	db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&models.Theme{})

	db.DropTableIfExists(&models.Theme_Stage{})
	db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&models.Theme_Stage{})

	// Style
	db.DropTableIfExists(&models.Style{})
	db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&models.Style{})

	db.DropTableIfExists(&models.Style_Stage{})
	db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&models.Style_Stage{})

	// Item
	db.DropTableIfExists(&models.Item{})
	db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&models.Item{})

	db.DropTableIfExists(&models.Item_Stage{})
	db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&models.Item_Stage{})

	// Whitelist
	db.DropTableIfExists(&models.Whitelist{})
	db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&models.Whitelist{})

	// UniqueContents
	db.DropTableIfExists(&models.UniqueContent{})
	db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&models.UniqueContent{})

	db.DropTableIfExists(&models.UniqueContent_Stage{})
	db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&models.UniqueContent_Stage{})

}
