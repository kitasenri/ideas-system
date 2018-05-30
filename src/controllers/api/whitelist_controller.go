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
// Create whitelist
func CreateWhitelist(c web.C, w http.ResponseWriter, r *http.Request) {

	request := getWhitelistRequest(w, r)
	if request == nil {
		return
	}

	whitelist := models.GetWhitelist(request.ThemeId, request.StyleId, request.ItemId)
	if whitelist.ThemeId == "" {
		// there is not whitelist
		models.CreateWhitelist(request)
	}

	// Success
	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Method / Delete
//---------------------------------------------------------------------
// Delete whitelist
func DeleteWhitelist(c web.C, w http.ResponseWriter, r *http.Request) {

	request := getWhitelistRequest(w, r)
	if request == nil {
		return
	}

	whitelist := models.GetWhitelist(request.ThemeId, request.StyleId, request.ItemId)
	if whitelist.ThemeId != "" {
		// there is not whitelist
		models.DeleteWhitelist(request)
	}

	// Success
	w.Write(utils.CreateSuccessResponse())
}

//---------------------------------------------------------------------
// Sub module
//---------------------------------------------------------------------
// Get request json from body
func getWhitelistRequest(w http.ResponseWriter, r *http.Request) *models.Whitelist {

	// convert request to json
	request := models.Whitelist{}
	error := json.NewDecoder(r.Body).Decode(&request)
	if error != nil {
		errorStr, _ := json.Marshal(error)
		w.Write(utils.CreateErrorResponse([]string{
			"リクエストの解析に失敗しました。 Reason : " + string(errorStr)}))
		return nil
	}

	return &request
}
