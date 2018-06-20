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
// Create theme
// curl -X POST -d "id=ASDF&title=ASNM" -u user-name:jedbLCxb5XcRi "http://localhost:8000/theme/create" --verbose
func CreateThemeStage(c web.C, w http.ResponseWriter, r *http.Request) {

	request := getThemeRequest(w, r)
	if request == nil {
		return
	}

	// Check ID is existed.
	if existThemeStageRecord(request) {
		return
	}

	// Check result
	if error := models.CreateThemeStage(request); error != nil {
		w.Write(utils.CreateErrorResponse(error))
	}

	// Process success
	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Update
//---------------------------------------------------------------------
// Update theme
// curl -X POST -d "id=ASDF&title=ASNM" -u user-name:jedbLCxb5XcRi "http://localhost:8000/theme/update" --verbose
func UpdateThemeStage(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getThemeRequest(w, r)
	if request == nil {
		return
	}

	// Check ID is existed.
	if !existThemeStageRecord(request) {
		w.Write(utils.CreateErrorResponse([]string{
			"指定されたIDは存在しません。"}))
		return
	}

	// Update theme_stage
	if error := models.UpdateThemeStage(request); error != nil {
		w.Write(utils.CreateErrorResponse(error))
	}

	// Success
	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Delete
//---------------------------------------------------------------------
// Delete theme
func DeleteTheme(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getThemeRequest(w, r)
	if request == nil {
		return
	}

	// delete theme
	models.DeleteTheme(request.Id)
	models.DeleteThemeStage(request.Id)

	w.Write(utils.CreateSuccessResponse())
}

// Update theme deleted
func UpdateThemeDeleted(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getThemeRequest(w, r)
	if request == nil {
		return
	}

	// delete theme
	models.UpdateThemeDeleted(request.Id, request.Deleted)
	models.UpdateThemeStageDeleted(request.Id, request.Deleted)

	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Sync
//---------------------------------------------------------------------
// Sync theme
func SyncTheme(c web.C, w http.ResponseWriter, r *http.Request) {

	// convert request to json
	request := getThemeRequest(w, r)
	if request == nil {
		return
	}

	theme_stage := models.GetThemeStage(request.Id)
	if theme_stage.Id == "" {
		w.Write(utils.CreateErrorResponse([]string{
			"指定されたIDは存在しません。"}))
		return
	}

	theme := copyTheme(theme_stage)
	if existThemeStageRecord(&theme_stage) {
		models.UpdateTheme(&theme)
	} else {
		models.CreateTheme(&theme)
	}

	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Sub module
//---------------------------------------------------------------------
// Check theme record
func existThemeStageRecord(request *models.Theme_Stage) bool {

	if t := models.GetThemeStage(request.Id); t.Id == "" {
		return false
	}

	return true
}

// Get request json from body
func getThemeRequest(w http.ResponseWriter, r *http.Request) *models.Theme_Stage {

	// convert request to json
	request := models.Theme_Stage{}
	error := json.NewDecoder(r.Body).Decode(&request)
	if error != nil {
		errorStr, _ := json.Marshal(error)
		w.Write(utils.CreateErrorResponse([]string{
			"リクエストの解析に失敗しました。 Reason : " + string(errorStr)}))
		return nil
	}

	return &request
}

func copyTheme(theme_stage models.Theme_Stage) models.Theme {

	theme := models.Theme{}

	theme.Id = theme_stage.Id
	theme.Name = theme_stage.Name

	theme.Title = theme_stage.Title
	theme.Css = theme_stage.Css
	theme.Script = theme_stage.Script
	theme.AddTags = theme_stage.AddTags

	theme.SearchWindowTitle = theme_stage.SearchWindowTitle
	theme.SearchWindowMessage = theme_stage.SearchWindowMessage
	theme.SearchWindowImageURL = theme_stage.SearchWindowImageURL

	theme.Message1 = theme_stage.Message1
	theme.Message2 = theme_stage.Message2
	theme.Message3 = theme_stage.Message3
	theme.Message4 = theme_stage.Message4
	theme.Message5 = theme_stage.Message5

	theme.Version = theme_stage.Version
	theme.Deleted = theme_stage.Deleted
	theme.UpdatedBy = theme_stage.UpdatedBy
	theme.UpdateAt = theme_stage.UpdateAt
	theme.CreatedBy = theme_stage.CreatedBy
	theme.CreateAt = theme_stage.CreateAt

	return theme
}
