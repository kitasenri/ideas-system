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
// Create style
// curl -X POST -d "id=ASDF&title=ASNM" -u user-name:jedbLCxb5XcRi "http://localhost:8000/style/create" --verbose
func CreateStyleStage(c web.C, w http.ResponseWriter, r *http.Request) {

	request := getStyleRequest(w, r)
	if request == nil {
		return
	}

	// Check ID is existed.
	if existStyleStageRecord(request) {
		return
	}

	// Check result
	if error := models.CreateStyleStage(request); error != nil {
		w.Write(utils.CreateErrorResponse(error))
	}

	// Process success
	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Update
//---------------------------------------------------------------------
// Update style
// curl -X POST -d "id=ASDF&title=ASNM" -u user-name:jedbLCxb5XcRi "http://localhost:8000/style/update" --verbose
func UpdateStyleStage(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getStyleRequest(w, r)
	if request == nil {
		return
	}

	// Check ID is existed.
	if !existStyleStageRecord(request) {
		w.Write(utils.CreateErrorResponse([]string{
			"指定されたIDは存在しません。"}))
		return
	}

	// Update style_stage
	if error := models.UpdateStyleStage(request); error != nil {
		w.Write(utils.CreateErrorResponse(error))
	}

	// Success
	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Delete
//---------------------------------------------------------------------
// Delete style
func DeleteStyle(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getStyleRequest(w, r)
	if request == nil {
		return
	}

	// delete style
	models.DeleteStyle(request.Id)
	models.DeleteStyleStage(request.Id)

	w.Write(utils.CreateSuccessResponse())
}

// Update style deleted
func UpdateStyleDeleted(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getStyleRequest(w, r)
	if request == nil {
		return
	}

	// delete style
	models.UpdateStyleDeleted(request.Id, request.Deleted)
	models.UpdateStyleStageDeleted(request.Id, request.Deleted)

	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Sync
//---------------------------------------------------------------------
// Sync style
func SyncStyle(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getStyleRequest(w, r)
	if request == nil {
		return
	}

	style_stage := models.GetStyleStage(request.Id)
	if style_stage.Id == "" {
		w.Write(utils.CreateErrorResponse([]string{
			"指定されたIDは存在しません。"}))
		return
	}

	style := copyStyle(style_stage)
	if existStyleStageRecord(&style_stage) {
		models.UpdateStyle(&style)
	} else {
		models.CreateStyle(&style)
	}

	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Sub module
//---------------------------------------------------------------------
// Check style record
func existStyleStageRecord(request *models.Style_Stage) bool {

	if t := models.GetStyleStage(request.Id); t.Id == "" {
		return false
	}

	return true
}

// Get request json from body
func getStyleRequest(w http.ResponseWriter, r *http.Request) *models.Style_Stage {

	// convert request to json
	request := models.Style_Stage{}
	error := json.NewDecoder(r.Body).Decode(&request)
	if error != nil {
		errorStr, _ := json.Marshal(error)
		w.Write(utils.CreateErrorResponse([]string{
			"リクエストの解析に失敗しました。 Reason : " + string(errorStr)}))
		return nil
	}

	return &request
}

func copyStyle(style_stage models.Style_Stage) models.Style {

	style := models.Style{}

	style.Id = style_stage.Id
	style.Name = style_stage.Name

	style.ContentsSequence = style_stage.ContentsSequence

	style.Message1 = style_stage.Message1
	style.Message2 = style_stage.Message2

	style.Version = style_stage.Version
	style.Deleted = style_stage.Deleted
	style.UpdatedBy = style_stage.UpdatedBy
	style.UpdateAt = style_stage.UpdateAt
	style.CreatedBy = style_stage.CreatedBy
	style.CreateAt = style_stage.CreateAt

	return style
}
