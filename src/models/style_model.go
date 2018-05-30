package models

import (
	"time"

	"src/core"

	"github.com/wcl48/valval"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------

type Style struct {
	CommonPartsHeader

	CommonPartsMessage1

	ContentsSequence string `sql:"type:varchar(512);" gorm:"not null;"`

	CommonPartsFooter
}

type Style_Stage struct {
	Style
}

type Style_Info struct {
	Id string

	Style_Name     string
	Style_Deleted  bool
	Style_UpdateAt time.Time
	Style_Version  int

	StyleStage_Name     string
	StyleStage_Deleted  bool
	StyleStage_UpdateAt time.Time
	StyleStage_Version  int
}

//---------------------------------------------------------------------
// Method / Validate
//---------------------------------------------------------------------
// Validate Style
func ValidateStyle(style_stage *Style_Stage) error {

	Validator := valval.Object(valval.M{
		"Id":   VALIDATE_KEY,
		"Name": VALIDATE_NAME,
	})

	return Validator.Validate(style_stage)
}

//---------------------------------------------------------------------
// Method / Get
//---------------------------------------------------------------------
// Get style
func GetStyle(id string) Style {

	style := Style{}

	db := core.GetSlaveConnection()
	db.Where("id = ?", id).First(&style)

	return style
}

// Get style stage
func GetStyleStage(id string) Style_Stage {

	style_stage := Style_Stage{}

	db := core.GetSlaveConnection()
	db.Where("id = ?", id).First(&style_stage)

	return style_stage
}

// Get style stage list
func GetStyleStageList() []Style_Stage {

	db := core.GetSlaveConnection()

	styles := []Style_Stage{}
	db.Find(&styles)

	return styles
}

// Get style stage list
func GetStyleList() []Style {

	db := core.GetSlaveConnection()

	styles := []Style{}
	db.Find(&styles)

	return styles
}

// Get style list
func GetStyleInfoList() []Style_Info {

	db := core.GetSlaveConnection()

	styles := []Style{}
	db.Find(&styles)

	style_stages := []Style_Stage{}
	db.Find(&style_stages)

	// merge style and style_stage
	count_stage := len(style_stages)
	style_joins := make([]Style_Info, count_stage, count_stage)
	for ii := 0; ii < count_stage; ii++ {

		style_stage := style_stages[ii]
		style_joins[ii].Id = style_stage.Id
		style_joins[ii].StyleStage_Name = style_stage.Name
		style_joins[ii].StyleStage_Deleted = style_stage.Deleted
		style_joins[ii].StyleStage_UpdateAt = style_stage.UpdateAt
		style_joins[ii].StyleStage_Version = style_stage.Version

		count_prod := len(styles)
		for jj := 0; jj < count_prod; jj++ {

			style := styles[jj]
			if style.Id == style_stage.Id {

				// style is found.
				style_joins[ii].Style_Name = style.Name
				style_joins[ii].Style_Deleted = style.Deleted
				style_joins[ii].Style_UpdateAt = style.UpdateAt
				style_joins[ii].Style_Version = style.Version

				break
			}

		}

	}

	return style_joins
}

// Get styles from whitelist
func GetStylesFromWhitelist(whitelists []Whitelist) []Style {

	cc := len(whitelists)
	style_ids := make([]string, cc, cc)

	for ii, vv := range whitelists {
		style_ids[ii] = vv.StyleId
	}

	styles := []Style{}

	db := core.GetMasterConnection()
	db.Raw("SELECT * FROM styles WHERE id IN (?)", style_ids).Scan(&styles)

	return styles
}

// Get style stages from whitelist
func GetStyleStagesFromWhitelist(whitelists []Whitelist) []Style_Stage {

	cc := len(whitelists)
	style_stage_ids := make([]string, cc, cc)

	for ii, vv := range whitelists {
		style_stage_ids[ii] = vv.StyleId
	}

	style_stages := []Style_Stage{}

	db := core.GetMasterConnection()
	db.Raw("SELECT * FROM style_stages WHERE id IN (?)", style_stage_ids).Scan(&style_stages)

	return style_stages
}

//---------------------------------------------------------------------
// Method / Create
//---------------------------------------------------------------------
// Insert style record
func CreateStyle(style *Style) {
	db := core.GetMasterConnection()
	db.Create(&style)
}

// Insert style record
func CreateStyleStage(style_stage *Style_Stage) []string {

	if error := ValidateStyle(style_stage); error != nil {
		return []string{error.Error()}
	}

	style_stage.Version = 1

	// Success
	db := core.GetMasterConnection()
	db.Create(&style_stage)

	return nil
}

//---------------------------------------------------------------------
// Method / Update
//---------------------------------------------------------------------
// Update style record
func UpdateStyle(style *Style) []string {
	db := core.GetMasterConnection()
	db.Save(&style)
	return nil
}

// Update style record
func UpdateStyleStage(style_stage *Style_Stage) []string {

	if error := ValidateStyle(style_stage); error != nil {
		return []string{error.Error()}
	}

	// Update
	style_stage.Version = style_stage.Version + 1

	db := core.GetMasterConnection()
	db.Save(&style_stage)

	return nil
}

// Update style record
func UpdateStyleDeleted(id string, deleted bool) {
	style := Style{}
	db := core.GetMasterConnection()
	db.Model(style).Where("id = ?", id).Update("deleted", deleted)
}

// Update style record
func UpdateStyleStageDeleted(id string, deleted bool) {
	style_stage := Style_Stage{}
	db := core.GetMasterConnection()
	db.Model(style_stage).Where("id = ?", id).Update("deleted", deleted)
}

//---------------------------------------------------------------------
// Method / Delete
//---------------------------------------------------------------------
// Delete style
func DeleteStyle(id string) {
	style := Style{}
	db := core.GetMasterConnection()
	db.Where("id = ?", id).Delete(&style)
}

func DeleteStyleStage(id string) {
	style_stage := Style_Stage{}
	db := core.GetMasterConnection()
	db.Where("id = ?", id).Delete(&style_stage)
}
