package models

import (
	"time"

	"src/core"

	"github.com/wcl48/valval"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------

type Whitelist struct {
	ThemeId   string    `sql:"type:varchar(255);" gorm:"primary_key;not null;"`
	StyleId   string    `sql:"type:varchar(255);" gorm:"primary_key;not null;"`
	ItemId    string    `sql:"type:varchar(255);" gorm:"primary_key;not null;"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`
	CreatedBy string    `sql:"type:varchar(255);"`
}

//---------------------------------------------------------------------
// Method
//---------------------------------------------------------------------
// Validate whitelist
func ValidateWhitelist(whitelist *Whitelist) error {

	Validator := valval.Object(valval.M{
		"Theme": VALIDATE_KEY,
		"Style": VALIDATE_KEY,
		"Item":  VALIDATE_KEY,
	})

	return Validator.Validate(whitelist)
}

//---------------------------------------------------------------------
// Method / Get
//---------------------------------------------------------------------
// Get whitelist
func GetWhitelist(theme_id string, style_id string, item_id string) Whitelist {

	whitelist := Whitelist{}

	db := core.GetSlaveConnection()
	db.Where("theme_id = ? AND style_id = ? AND item_id = ?", theme_id, style_id, item_id).Find(&whitelist)

	return whitelist
}

// Get whitelists
func GetStyleWhitelists(theme_id string) []Whitelist {

	whitelists := []Whitelist{}

	db := core.GetSlaveConnection()
	db.Where("theme_id = ? AND item_id = ?", theme_id, "").Find(&whitelists)

	return whitelists
}

// Get whitelists
func GetItemWhitelists(theme_id string) []Whitelist {

	whitelists := []Whitelist{}

	db := core.GetSlaveConnection()
	db.Where("theme_id = ? AND style_id = ?", theme_id, "").Find(&whitelists)

	return whitelists
}

// Get whitelists
func GetWhitelists(theme_id string) []Whitelist {

	whitelists := []Whitelist{}

	db := core.GetSlaveConnection()
	db.Where("theme_id = ?", theme_id).Find(&whitelists)

	return whitelists
}

func GetAllWhitelist() []Whitelist {

	whitelists := []Whitelist{}

	db := core.GetSlaveConnection()
	db.Find(&whitelists)

	return whitelists
}

//---------------------------------------------------------------------
// Method / Create
//---------------------------------------------------------------------
// Create whitelist
func CreateWhitelist(whitelist *Whitelist) {
	db := core.GetMasterConnection()
	db.Create(&whitelist)
}

//---------------------------------------------------------------------
// Method / Delete
//---------------------------------------------------------------------
// Delete whitelist
func DeleteWhitelist(whitelist *Whitelist) {
	db := core.GetMasterConnection()
	db.Delete(&whitelist)
}
