package utils

import (
	"src/consts"

	"encoding/json"
)

//---------------------------------------------------------------------
// Method
//---------------------------------------------------------------------
func ToStringByteFromObjectByte(jobj []byte) []byte {
	return []byte(string(jobj))
}

// Check indexOf
func IndexOf(element string, items []string) int {

	for ii, vv := range items {
		if element == vv {
			return ii
		}
	}
	return -1
}

// Check deleted param
func GetDeleted(deleted string) bool {

	ret := false
	if deleted == consts.PARAM_DELETED_YES {
		return true
	}

	return ret
}

// Check new or not
func IsEdit(Id string) bool {

	isEdit := true
	if Id == "" {
		isEdit = false
	}

	return isEdit
}

func CreateSuccessResponse() []byte {

	jsonObj := map[string]interface{}{
		"status": "OK",
	}

	jsonString, _ := json.Marshal(jsonObj)
	return jsonString
}

func CreateErrorResponse(errorMessage []string) []byte {

	jsonObj := map[string]interface{}{
		"status":    "NG",
		"errorInfo": errorMessage,
	}

	jsonString, _ := json.Marshal(jsonObj)
	return jsonString
}
