package models

import (
	"github.com/wcl48/valval"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------

type UniqueContent struct {
	ThemeId string `sql:"type:varchar(255);" gorm:"primary_key;"`
	StyleId string `sql:"type:varchar(255);" gorm:"primary_key;"`
	ItemId  string `sql:"type:varchar(255);" gorm:"primary_key;"`

	CommonPartsFooter
}

type UniqueContent_Stage struct {
	UniqueContent
}

//---------------------------------------------------------------------
// Method
//---------------------------------------------------------------------
// Validate UniqueContents
func ValidateUniqueContents(uniqueContent UniqueContent) error {

	Validator := valval.Object(valval.M{
		"Theme": VALIDATE_KEY,
		"Style": VALIDATE_KEY,
		"Item":  VALIDATE_KEY,
	})

	return Validator.Validate(uniqueContent)
}
