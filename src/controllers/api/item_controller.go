package controllers_api

import (
	"encoding/json"
	"net/http"

	"src/models"
	"src/utils"

	"github.com/zenazn/goji/web"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------

//---------------------------------------------------------------------
// Method / Create
//---------------------------------------------------------------------
// Create Item
// curl -X POST -d "id=ASDF&title=ASNM" -u user-name:jedbLCxb5XcRi "http://localhost:8000/item/create" --verbose
func CreateItemStage(c web.C, w http.ResponseWriter, r *http.Request) {

	request := getItemRequest(w, r)
	if request == nil {
		return
	}

	// Check ID is existed.
	if existItemStageRecord(request) {
		return
	}

	// Check result
	if error := models.CreateItemStage(request); error != nil {
		w.Write(utils.CreateErrorResponse(error))
	}

	// Process success
	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Update
//---------------------------------------------------------------------
// Update item
// curl -X POST -d "id=ASDF&title=ASNM" -u user-name:jedbLCxb5XcRi "http://localhost:8000/style/update" --verbose
func UpdateItemStage(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getItemRequest(w, r)
	if request == nil {
		return
	}

	// Check ID is existed.
	if !existItemStageRecord(request) {
		w.Write(utils.CreateErrorResponse([]string{
			"指定されたIDは存在しません。"}))
		return
	}

	// Update item_stage
	if error := models.UpdateItemStage(request); error != nil {
		w.Write(utils.CreateErrorResponse(error))
	}

	// Success
	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Delete
//---------------------------------------------------------------------
// Delete item
func DeleteItem(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getItemRequest(w, r)
	if request == nil {
		return
	}

	// delete item
	models.DeleteItem(request.Id)
	models.DeleteItemStage(request.Id)

	w.Write(utils.CreateSuccessResponse())
}

// Update item deleted
func UpdateItemDeleted(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getItemRequest(w, r)
	if request == nil {
		return
	}

	// delete item
	models.UpdateItemDeleted(request.Id, request.Deleted)
	models.UpdateItemStageDeleted(request.Id, request.Deleted)

	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Sync
//---------------------------------------------------------------------
// Sync item
func SyncItem(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getItemRequest(w, r)
	if request == nil {
		return
	}

	item_stage := models.GetItemStage(request.Id)
	if item_stage.Id == "" {
		w.Write(utils.CreateErrorResponse([]string{
			"指定されたIDは存在しません。"}))
		return
	}

	item := copyItem(item_stage)
	if existItemStageRecord(&item_stage) {
		models.UpdateItem(&item)
	} else {
		models.CreateItem(&item)
	}

	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Sub module
//---------------------------------------------------------------------
// Check item record
func existItemStageRecord(request *models.Item_Stage) bool {

	if t := models.GetItemStage(request.Id); t.Id == "" {
		return false
	}

	return true
}

// Get request json from body
func getItemRequest(w http.ResponseWriter, r *http.Request) *models.Item_Stage {

	// convert request to json
	request := models.Item_Stage{}
	error := json.NewDecoder(r.Body).Decode(&request)
	if error != nil {
		errorStr, _ := json.Marshal(error)
		w.Write(utils.CreateErrorResponse([]string{
			"リクエストの解析に失敗しました。 Reason : " + string(errorStr)}))
		return nil
	}

	return &request
}

func copyItem(item_stage models.Item_Stage) models.Item {

	item := models.Item{}

	item.Id = item_stage.Id
	item.Name = item_stage.Name

	item.Message1 = item_stage.Message1
	item.Message2 = item_stage.Message2

	item.Version = item_stage.Version
	item.Deleted = item_stage.Deleted
	item.UpdatedBy = item_stage.UpdatedBy
	item.UpdateAt = item_stage.UpdateAt
	item.CreatedBy = item_stage.CreatedBy
	item.CreateAt = item_stage.CreateAt

	return item
}
