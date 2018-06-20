package controllers_front

import (
	"src/core"
	"src/helper"
	"src/models"

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
// (ex) http://localhost:18080/?themeKey=theme01&styleKey=style01&itemKey=item01
func Index(c web.C, w http.ResponseWriter, r *http.Request) {

	theme_key := r.FormValue("themeKey")
	style_key := r.FormValue("styleKey")
	item_key := r.FormValue("itemKey")

	page := models.GetPage(theme_key, style_key, item_key)
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
		helper.Error_F(ee)
	}

}

// Sitemap
func Sitemap(c web.C, w http.ResponseWriter, r *http.Request) {

	theme_page_info := models.GetPages()
	view_param := map[string]interface{}{
		"page_info": theme_page_info,
	}

	f, e := helper.GetMacroTemplate("sitemap.xml").ParseFiles(
		core.GetRelativePath() + "src/views/front/sitemap.xml",
	)

	// Set XML Header
	w.Header().Set("Content-Type", "application/xml")

	t := template.Must(f, e)
	if ee := t.Execute(w, view_param); ee != nil {
		helper.Error_F(ee)
	}

}
