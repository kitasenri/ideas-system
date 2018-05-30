package models

import (
	"src/core"

	"time"

	"github.com/wcl48/valval"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------

type Item struct {
	CommonPartsHeader

	CommonPartsMessage1

	CommonPartsFooter
}

type Item_Stage struct {
	Item
}

type Item_Info struct {
	Id string

	Item_Name     string
	Item_Deleted  bool
	Item_UpdateAt time.Time
	Item_Version  int

	ItemStage_Name     string
	ItemStage_Deleted  bool
	ItemStage_UpdateAt time.Time
	ItemStage_Version  int
}

//---------------------------------------------------------------------
// Method / Validate
//---------------------------------------------------------------------
// Validate Item
func ValidateItem(item_stage *Item_Stage) error {

	Validator := valval.Object(valval.M{
		"Id":   VALIDATE_KEY,
		"Name": VALIDATE_NAME,
	})

	return Validator.Validate(item_stage)
}

//---------------------------------------------------------------------
// Method / Get
//---------------------------------------------------------------------
// Get item
func GetItem(id string) Item {

	item := Item{}

	db := core.GetSlaveConnection()
	db.Where("id = ?", id).First(&item)

	return item
}

// Get item
func GetItemStage(id string) Item_Stage {

	item_stage := Item_Stage{}

	db := core.GetSlaveConnection()
	db.Where("id = ?", id).First(&item_stage)

	return item_stage
}

// Get item stage list
func GetItemStageList() []Item_Stage {

	db := core.GetSlaveConnection()

	items := []Item_Stage{}
	db.Find(&items)

	return items
}

// Get item list
func GetItemList() []Item {

	db := core.GetSlaveConnection()

	items := []Item{}
	db.Find(&items)

	return items
}

// Get item list
func GetItemInfoList() []Item_Info {

	db := core.GetSlaveConnection()

	items := []Item{}
	db.Find(&items)

	item_stages := []Item_Stage{}
	db.Find(&item_stages)

	// merge item and item_stage
	count_stage := len(item_stages)
	item_joins := make([]Item_Info, count_stage, count_stage)
	for ii := 0; ii < count_stage; ii++ {

		item_stage := item_stages[ii]
		item_joins[ii].Id = item_stage.Id
		item_joins[ii].ItemStage_Name = item_stage.Name
		item_joins[ii].ItemStage_Deleted = item_stage.Deleted
		item_joins[ii].ItemStage_UpdateAt = item_stage.UpdateAt
		item_joins[ii].ItemStage_Version = item_stage.Version

		count_prod := len(items)
		for jj := 0; jj < count_prod; jj++ {

			item := items[jj]
			if item.Id == item_stage.Id {

				// item is found.
				item_joins[ii].Item_Name = item.Name
				item_joins[ii].Item_Deleted = item.Deleted
				item_joins[ii].Item_UpdateAt = item.UpdateAt
				item_joins[ii].Item_Version = item.Version

				break
			}

		}

	}

	return item_joins
}

// Get items from whitelist
func GetItemsFromWhitelist(whitelists []Whitelist) []Item {

	cc := len(whitelists)
	item_ids := make([]string, cc, cc)

	for ii, vv := range whitelists {
		item_ids[ii] = vv.ItemId
	}

	items := []Item{}

	db := core.GetMasterConnection()
	db.Raw("SELECT * FROM items WHERE id IN (?)", item_ids).Scan(&items)

	return items
}

// Get items from whitelist
func GetItemStagesFromWhitelist(whitelists []Whitelist) []Item_Stage {

	cc := len(whitelists)
	item_stage_ids := make([]string, cc, cc)

	for ii, vv := range whitelists {
		item_stage_ids[ii] = vv.ItemId
	}

	item_stages := []Item_Stage{}

	db := core.GetMasterConnection()
	db.Raw("SELECT * FROM item_stages WHERE id IN (?)", item_stage_ids).Scan(&item_stages)

	return item_stages
}

//---------------------------------------------------------------------
// Method / Create
//---------------------------------------------------------------------
// Insert item record
func CreateItem(item *Item) {
	db := core.GetMasterConnection()
	db.Create(&item)
}

// Insert item record
func CreateItemStage(item_stage *Item_Stage) []string {

	if error := ValidateItem(item_stage); error != nil {
		return []string{error.Error()}
	}

	item_stage.Version = 1

	// Success
	db := core.GetMasterConnection()
	db.Create(&item_stage)

	return nil
}

//---------------------------------------------------------------------
// Method / Update
//---------------------------------------------------------------------
// Update item record
func UpdateItem(item *Item) []string {
	db := core.GetMasterConnection()
	db.Save(&item)
	return nil
}

// Update item record
func UpdateItemStage(item_stage *Item_Stage) []string {

	if error := ValidateItem(item_stage); error != nil {
		return []string{error.Error()}
	}

	// Update
	item_stage.Version = item_stage.Version + 1

	db := core.GetMasterConnection()
	db.Save(&item_stage)

	return nil
}

// Update item record
func UpdateItemDeleted(id string, deleted bool) {
	item := Item{}
	db := core.GetMasterConnection()
	db.Model(item).Where("id = ?", id).Update("deleted", deleted)
}

// Update item record
func UpdateItemStageDeleted(id string, deleted bool) {
	item_stage := Item_Stage{}
	db := core.GetMasterConnection()
	db.Model(item_stage).Where("id = ?", id).Update("deleted", deleted)
}

//---------------------------------------------------------------------
// Method / Delete
//---------------------------------------------------------------------
// Delete item
func DeleteItem(id string) {
	item := Item{}
	db := core.GetMasterConnection()
	db.Where("id = ?", id).Delete(&item)
}

func DeleteItemStage(id string) {
	item_stage := Item_Stage{}
	db := core.GetMasterConnection()
	db.Where("id = ?", id).Delete(&item_stage)
}
