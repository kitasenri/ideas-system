package models

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------
type Page struct {
	Theme_id string
	Style_id string
	Item_id  string
	Theme    Theme
	Style    Style
	Item     Item
}

type PageStage struct {
	Theme_id string
	Style_id string
	Item_id  string
	Theme    Theme_Stage
	Style    Style_Stage
	Item     Item_Stage
}

type ThemePage struct {
	Theme_id string
	Styles   []Style
	Items    []Item
}

type ThemePageStage struct {
	Theme_id string
	Styles   []Style_Stage
	Items    []Item_Stage
}

type ThemePageInfo struct {
	ThemePage
	Theme Theme
}

//---------------------------------------------------------------------
// Method / Validate
//---------------------------------------------------------------------
// Get page
func GetPage(theme_id string, style_id string, item_id string) Page {

	style_whitelist := GetWhitelist(theme_id, style_id, "")
	item_whitelist := GetWhitelist(theme_id, "", item_id)
	if style_whitelist.ThemeId == "" || item_whitelist.ItemId == "" {
		return Page{}
	}

	theme := GetTheme(theme_id)
	style := GetStyle(style_id)
	item := GetItem(item_id)

	page := Page{
		Theme_id: theme_id,
		Style_id: style_id,
		Item_id:  item_id,
		Theme:    theme,
		Style:    style,
		Item:     item,
	}

	return page
}

func GetPageStage(theme_id string, style_id string, item_id string) PageStage {

	style_whitelist := GetWhitelist(theme_id, style_id, "")
	item_whitelist := GetWhitelist(theme_id, "", item_id)
	if style_whitelist.ThemeId == "" || item_whitelist.ItemId == "" {
		return PageStage{}
	}

	theme := GetThemeStage(theme_id)
	style := GetStyleStage(style_id)
	item := GetItemStage(item_id)

	page_stage := PageStage{
		Theme_id: theme_id,
		Style_id: style_id,
		Item_id:  item_id,
		Theme:    theme,
		Style:    style,
		Item:     item,
	}

	return page_stage
}

// Get theme page
func GetThemePage(theme_id string) ThemePage {

	style_whitelists := GetStyleWhitelists(theme_id)
	item_whitelists := GetItemWhitelists(theme_id)

	style_stages := GetStylesFromWhitelist(style_whitelists)
	item_stages := GetItemsFromWhitelist(item_whitelists)

	theme_page_stage := ThemePage{
		Theme_id: theme_id,
		Styles:   style_stages,
		Items:    item_stages,
	}

	return theme_page_stage
}

// Get theme page
func GetThemePageStage(theme_id string) ThemePageStage {

	style_whitelists := GetStyleWhitelists(theme_id)
	item_whitelists := GetItemWhitelists(theme_id)

	styles := GetStyleStagesFromWhitelist(style_whitelists)
	items := GetItemStagesFromWhitelist(item_whitelists)

	theme_page_stage := ThemePageStage{
		Theme_id: theme_id,
		Styles:   styles,
		Items:    items,
	}

	return theme_page_stage
}

// Get Pages
func GetPages() []ThemePageInfo {

	theme_page_info := []ThemePageInfo{}

	theme_list := GetThemeList()
	style_list := GetStyleList()
	item_list := GetItemList()

	whitelists := GetAllWhitelist()
	for _, theme := range theme_list {

		styles := getStylesFromWhitelist(theme.Id, whitelists, style_list)
		items := getItemsFromWhitelist(theme.Id, whitelists, item_list)

		tpi := ThemePageInfo{}
		tpi.Theme_id = theme.Id
		tpi.Styles = styles
		tpi.Items = items
		tpi.Theme = theme

		theme_page_info = append(theme_page_info, tpi)
	}

	return theme_page_info
}

func getStylesFromWhitelist(theme_id string, whitelists []Whitelist, styles []Style) []Style {

	// create whitelist temp map
	temp_whitelists := map[string]Whitelist{}
	for _, wl := range whitelists {

		if wl.StyleId == "" {
			continue
		} else if wl.ThemeId == theme_id {
			temp_whitelists[wl.StyleId] = wl
		}

	}

	temp_styles := []Style{}
	for _, style := range styles {

		if _, ok := temp_whitelists[style.Id]; ok {
			temp_styles = append(temp_styles, style)
		}

	}

	return temp_styles
}

func getItemsFromWhitelist(theme_id string, whitelists []Whitelist, items []Item) []Item {

	// create whitelist temp map
	temp_whitelists := map[string]Whitelist{}
	for _, wl := range whitelists {

		if wl.ItemId == "" {
			continue
		} else if wl.ThemeId == theme_id {
			temp_whitelists[wl.ItemId] = wl
		}

	}

	temp_items := []Item{}
	for _, item := range items {

		if _, ok := temp_whitelists[item.Id]; ok {
			temp_items = append(temp_items, item)
		}

	}

	return temp_items
}
