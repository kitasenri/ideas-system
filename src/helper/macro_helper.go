package helper

import (
	template "html/template"
	"strings"
	"time"

	"src/consts"
	"src/core"
	"src/models"
	"src/utils"
)

// Macro Map
// "and":      and,
// "call":     call,
// "html":     HTMLEscaper,
// "index":    index,
// "js":       JSEscaper,
// "len":      length,
// "not":      not,
// "or":       or,
// "print":    fmt.Sprint,
// "printf":   fmt.Sprintf,
// "println":  fmt.Sprintln,
// "urlquery": URLQueryEscaper,
//
// Comparisons
// "eq": eq, // ==
// "ge": ge, // >=
// "gt": gt, // >
// "le": le, // <=
// "lt": lt, // <
// "ne": ne, // !=
var MacroMap = template.FuncMap{
	"upper":     strings.ToUpper,
	"lower":     strings.ToLower,
	"increment": func(value int) int { return value + 1 },
	"safe":      func(text string) template.HTML { return template.HTML(text) },
	"is_whitelist_style": func(style_id string, whitelists []models.Whitelist) bool {
		isWhitelist := false
		for _, vv := range whitelists {
			if vv.StyleId == style_id {
				isWhitelist = true
				break
			}
		}
		return isWhitelist
	},
	"is_whitelist_item": func(item_id string, whitelists []models.Whitelist) bool {
		isWhitelist := false
		for _, vv := range whitelists {
			if vv.ItemId == item_id {
				isWhitelist = true
				break
			}
		}
		return isWhitelist
	},
	"get_contents_sequence": func(contents_sequence string) [][]consts.Content {

		items := strings.Split(contents_sequence, "_")

		showLength := 0
		if len(items) == 1 && items[0] == "" {
			items = consts.CONTENT_ITEMS_DEFAULT
		}

		showLength = len(items)
		hideLength := len(consts.CONTENT_ITEMS) - showLength

		showContents := make([]consts.Content, showLength)
		hideContents := make([]consts.Content, hideLength)

		ii, jj := 0, 0
		for _, vv := range consts.CONTENT_ITEMS {

			ii = utils.IndexOf(vv.Value, items)
			if ii != -1 {
				// show
				showContents[ii] = vv
			} else {
				// hide
				hideContents[jj] = vv
				jj = jj + 1
			}

		}

		retContents := make([][]consts.Content, 2)
		retContents[0] = showContents
		retContents[1] = hideContents

		return retContents
	},
	"get_contents": func(value string) map[string]consts.Content {

		contents := map[string]consts.Content{
			"SEARCH_WINDOW":  consts.SEARCH_WINDOW,
			"THEME_MESSAGE1": consts.THEME_MESSAGE1,
			"THEME_MESSAGE2": consts.THEME_MESSAGE2,
			"THEME_MESSAGE3": consts.THEME_MESSAGE3,
			"THEME_MESSAGE4": consts.THEME_MESSAGE4,
			"THEME_MESSAGE5": consts.THEME_MESSAGE5,
			"STYLE_MESSAGE1": consts.STYLE_MESSAGE1,
			"STYLE_MESSAGE2": consts.STYLE_MESSAGE2,
			"ITEM_MESSAGE1":  consts.ITEM_MESSAGE1,
			"ITEM_MESSAGE2":  consts.ITEM_MESSAGE2,
			"ITEM_PARTS1":    consts.ITEM_PARTS1,
			"ITEM_PARTS2":    consts.ITEM_PARTS2,
			"ITEM_PARTS3":    consts.ITEM_PARTS3,
			"SITE_MAP":       consts.SITE_MAP,
		}

		return contents
	},
	"get_front_server_info": func(value string) core.ServerInfo {
		return core.GetConfig().Launch.Front
	},
	"get_last_modified": func(value string) string {
		return time.Now().Format("2006-01-02")
	},
}

// Get Template Object with macro
func GetMacroTemplate(baseFile string) *template.Template {
	return template.New(baseFile).Funcs(MacroMap)
}
