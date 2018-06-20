package models

import (
	"time"

	"src/core"

	"github.com/wcl48/valval"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------

type Theme struct {
	CommonPartsHeader

	Title  string `sql:"type:text;" gorm:"not null;"`
	Css    string `sql:"type:text;" gorm:"not null;"`
	Script string `sql:"type:text;" gorm:"not null;"`
	AddTags string `sql:"type:text;" gorm:"not null;"`

	SearchWindowTitle    string `sql:"type:text;" gorm:"not null;"`
	SearchWindowMessage  string `sql:"type:text;" gorm:"not null;"`
	SearchWindowImageURL string `sql:"type:text;" gorm:"not null;"`

	CommonPartsMessage1
	CommonPartsMessage2

	CommonPartsFooter
}

type Theme_Stage struct {
	Theme
}

type Theme_Info struct {
	Id string

	Theme_Name     string
	Theme_Deleted  bool
	Theme_UpdateAt time.Time
	Theme_Version  int

	ThemeStage_Name     string
	ThemeStage_Deleted  bool
	ThemeStage_UpdateAt time.Time
	ThemeStage_Version  int
}

//---------------------------------------------------------------------
// Method / Validate
//---------------------------------------------------------------------
// Validate Theme
func ValidateThemeStage(theme_stage *Theme_Stage) error {

	Validator := valval.Object(valval.M{
		"Id":   VALIDATE_KEY,
		"Name": VALIDATE_NAME,
	})

	return Validator.Validate(theme_stage)
}

//---------------------------------------------------------------------
// Method / Get
//---------------------------------------------------------------------
// Get theme
func GetTheme(id string) Theme {

	theme := Theme{}

	db := core.GetSlaveConnection()
	db.Where("id = ?", id).First(&theme)

	return theme
}

// Get theme stage
func GetThemeStage(id string) Theme_Stage {

	theme_stage := Theme_Stage{}

	db := core.GetSlaveConnection()
	db.Where("id = ?", id).First(&theme_stage)

	return theme_stage
}

// Get theme stage list
func GetThemeStageList() []Theme_Stage {

	db := core.GetSlaveConnection()

	themes := []Theme_Stage{}
	db.Find(&themes)

	return themes
}

// Get theme list
func GetThemeList() []Theme {

	db := core.GetSlaveConnection()

	themes := []Theme{}
	db.Find(&themes)

	return themes
}

// Get theme list
func GetThemeInfoList() []Theme_Info {

	db := core.GetSlaveConnection()

	themes := []Theme{}
	db.Find(&themes)

	theme_stages := []Theme_Stage{}
	db.Find(&theme_stages)

	// merge theme and theme_stage
	count_stage := len(theme_stages)
	theme_joins := make([]Theme_Info, count_stage, count_stage)
	for ii := 0; ii < count_stage; ii++ {

		theme_stage := theme_stages[ii]
		theme_joins[ii].Id = theme_stage.Id
		theme_joins[ii].ThemeStage_Name = theme_stage.Name
		theme_joins[ii].ThemeStage_Deleted = theme_stage.Deleted
		theme_joins[ii].ThemeStage_UpdateAt = theme_stage.UpdateAt
		theme_joins[ii].ThemeStage_Version = theme_stage.Version

		count_prod := len(themes)
		for jj := 0; jj < count_prod; jj++ {

			theme := themes[jj]
			if theme.Id == theme_stage.Id {

				// theme is found.
				theme_joins[ii].Theme_Name = theme.Name
				theme_joins[ii].Theme_Deleted = theme.Deleted
				theme_joins[ii].Theme_UpdateAt = theme.UpdateAt
				theme_joins[ii].Theme_Version = theme.Version

				break
			}

		}

	}

	return theme_joins
}

//---------------------------------------------------------------------
// Method / Create
//---------------------------------------------------------------------
// Insert theme record
func CreateTheme(theme *Theme) {
	db := core.GetMasterConnection()
	db.Create(&theme)
}

// Insert theme record
func CreateThemeStage(theme_stage *Theme_Stage) []string {

	if error := ValidateThemeStage(theme_stage); error != nil {
		return []string{error.Error()}
	}

	theme_stage.Version = 1

	// Success
	db := core.GetMasterConnection()
	db.Create(&theme_stage)

	return nil
}

//---------------------------------------------------------------------
// Method / Update
//---------------------------------------------------------------------
// Update theme record
func UpdateTheme(theme *Theme) []string {
	db := core.GetMasterConnection()
	db.Save(&theme)
	return nil
}

// Update theme record
func UpdateThemeStage(theme_stage *Theme_Stage) []string {

	if error := ValidateThemeStage(theme_stage); error != nil {
		return []string{error.Error()}
	}

	// Update
	theme_stage.Version = theme_stage.Version + 1

	db := core.GetMasterConnection()
	db.Save(&theme_stage)

	return nil
}

// Update theme record
func UpdateThemeDeleted(id string, deleted bool) {
	theme := Theme{}
	db := core.GetMasterConnection()
	db.Model(theme).Where("id = ?", id).Update("deleted", deleted)
}

// Update theme record
func UpdateThemeStageDeleted(id string, deleted bool) {
	theme_stage := Theme_Stage{}
	db := core.GetMasterConnection()
	db.Model(theme_stage).Where("id = ?", id).Update("deleted", deleted)
}

//---------------------------------------------------------------------
// Method / Delete
//---------------------------------------------------------------------
// Delete theme
func DeleteTheme(id string) {
	theme := Theme{}
	db := core.GetMasterConnection()
	db.Where("id = ?", id).Delete(&theme)
}

func DeleteThemeStage(id string) {
	theme_stage := Theme_Stage{}
	db := core.GetMasterConnection()
	db.Where("id = ?", id).Delete(&theme_stage)
}
