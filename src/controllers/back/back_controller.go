package controllers_back

import (
	"src/core"
	"src/helper"
	"src/models"
	"src/utils"

	template "html/template"
	"net/http"

	"github.com/zenazn/goji/web"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------

//---------------------------------------------------------------------
// Method / Index
//---------------------------------------------------------------------
// Render index page.
func Index(c web.C, w http.ResponseWriter, r *http.Request) {

	themes := models.GetThemeInfoList()
	styles := models.GetStyleInfoList()
	items := models.GetItemInfoList()

	view_param := map[string]interface{}{
		"themeList": themes,
		"styleList": styles,
		"itemList":  items,
	}

	f, e := helper.GetMacroTemplate("index.html").ParseFiles(
		core.GetRelativePath()+"src/views/back/index.html",
		core.GetRelativePath()+"src/views/back/macro/head.html",
		core.GetRelativePath()+"src/views/back/macro/navigation.html",
		core.GetRelativePath()+"src/views/back/macro/list_theme.html",
		core.GetRelativePath()+"src/views/back/macro/list_style.html",
		core.GetRelativePath()+"src/views/back/macro/list_item.html",
	)

	t := template.Must(f, e)
	if ee := t.Execute(w, view_param); ee != nil {
		helper.Error_B(ee)
	}

}

// Render Edit Theme Page
func EditTheme(c web.C, w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("themeKey")
	theme := models.GetThemeStage(id)
	if id != "" && theme.Id == "" {
		panic("Key Not Found")
	}

	view_param := map[string]interface{}{
		"theme":  theme,
		"isEdit": utils.IsEdit(id),
	}

	f, e := helper.GetMacroTemplate("edit_theme.html").ParseFiles(
		core.GetRelativePath()+"src/views/back/edit_theme.html",
		core.GetRelativePath()+"src/views/back/macro/head.html",
		core.GetRelativePath()+"src/views/back/macro/navigation.html",
		core.GetRelativePath()+"src/views/back/macro/edit_footer.html",
	)

	t := template.Must(f, e)
	if ee := t.Execute(w, view_param); ee != nil {
		helper.Error_B(ee)
	}

}

func EditStyle(c web.C, w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("styleKey")
	style := models.GetStyleStage(id)
	if id != "" && style.Id == "" {
		panic("Key Not Found")
	}

	view_param := map[string]interface{}{
		"style":  style,
		"isEdit": utils.IsEdit(id)}

	f, e := helper.GetMacroTemplate("edit_style.html").ParseFiles(
		core.GetRelativePath()+"src/views/back/edit_style.html",
		core.GetRelativePath()+"src/views/back/macro/head.html",
		core.GetRelativePath()+"src/views/back/macro/navigation.html",
		core.GetRelativePath()+"src/views/back/macro/edit_footer.html",
	)

	t := template.Must(f, e)
	if ee := t.Execute(w, view_param); ee != nil {
		helper.Error_B(ee)
	}

}

func EditItem(c web.C, w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("itemKey")
	item := models.GetItemStage(id)
	if id != "" && item.Id == "" {
		panic("Key Not Found")
	}

	view_param := map[string]interface{}{
		"item":   item,
		"isEdit": utils.IsEdit(id),
	}

	f, e := helper.GetMacroTemplate("edit_item.html").ParseFiles(
		core.GetRelativePath()+"src/views/back/edit_item.html",
		core.GetRelativePath()+"src/views/back/macro/head.html",
		core.GetRelativePath()+"src/views/back/macro/navigation.html",
		core.GetRelativePath()+"src/views/back/macro/edit_footer.html",
	)

	t := template.Must(f, e)
	if ee := t.Execute(w, view_param); ee != nil {
		helper.Error_B(ee)
	}

}

// Edit whitelist
func EditWhiteList(c web.C, w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("themeKey")

	// check theme
	theme := models.GetThemeStage(id)
	if id != "" && theme.Id == "" {
		panic("Key Not Found")
	}

	style_stages := models.GetStyleStageList()
	item_stages := models.GetItemStageList()
	whitelists := models.GetWhitelists(id)

	view_param := map[string]interface{}{
		"theme_id":     id,
		"whitelists":   whitelists,
		"style_stages": style_stages,
		"item_stages":  item_stages,
	}

	f, e := helper.GetMacroTemplate("edit_whitelist.html").ParseFiles(
		core.GetRelativePath()+"src/views/back/edit_whitelist.html",
		core.GetRelativePath()+"src/views/back/macro/head.html",
		core.GetRelativePath()+"src/views/back/macro/navigation.html",
		core.GetRelativePath()+"src/views/back/macro/whitelist_style.html",
		core.GetRelativePath()+"src/views/back/macro/whitelist_item.html",
	)

	t := template.Must(f, e)
	if ee := t.Execute(w, view_param); ee != nil {
		helper.Error_B(ee)
	}

}

// Theme Pages
func ThemePages(c web.C, w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("themeKey")

	// check theme
	theme := models.GetThemeStage(id)
	if id != "" && theme.Id == "" {
		panic("Key Not Found")
	}

	style_whitelists := models.GetStyleWhitelists(id)
	item_whitelists := models.GetItemWhitelists(id)

	view_param := map[string]interface{}{
		"theme_id":         id,
		"style_whitelists": style_whitelists,
		"item_whitelists":  item_whitelists,
	}

	f, e := helper.GetMacroTemplate("theme_pages.html").ParseFiles(
		core.GetRelativePath()+"src/views/back/theme_pages.html",
		core.GetRelativePath()+"src/views/back/macro/head.html",
		core.GetRelativePath()+"src/views/back/macro/navigation.html",
		core.GetRelativePath()+"src/views/back/macro/whitelist_style.html",
	)

	t := template.Must(f, e)
	if ee := t.Execute(w, view_param); ee != nil {
		helper.Error_B(ee)
	}

}

// Show page stage
func StagePage(c web.C, w http.ResponseWriter, r *http.Request) {

	theme_key := r.FormValue("themeKey")
	style_key := r.FormValue("styleKey")
	item_key := r.FormValue("itemKey")

	page := models.GetPageStage(theme_key, style_key, item_key)
	if page.Theme_id == "" || page.Style_id == "" || page.Item_id == "" {
		return
	}

	view_param := map[string]interface{}{
		"page": page,
	}

	f, e := helper.GetMacroTemplate("index.html").ParseFiles(
		core.GetRelativePath()+"src/views/common/index.html",
		core.GetRelativePath()+"src/views/common/macro/head.html",
		core.GetRelativePath()+"src/views/common/macro/contents.html",
		core.GetRelativePath()+"src/views/common/macro/footer.html",
		core.GetRelativePath()+"src/views/common/macro/navigation.html",
	)

	t := template.Must(f, e)
	if ee := t.Execute(w, view_param); ee != nil {
		helper.Error_B(ee)
	}

}

//---------------------------------------------------------------------
// Method / Option
//---------------------------------------------------------------------
func EditUniqueContent(c web.C, w http.ResponseWriter, r *http.Request) {

}

func IndexUniqueContent(c web.C, w http.ResponseWriter, r *http.Request) {

}

func IndexParts(c web.C, w http.ResponseWriter, r *http.Request) {

}

func EditParts(c web.C, w http.ResponseWriter, r *http.Request) {

}
