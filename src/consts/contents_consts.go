package consts

// For contents sequence
type Content struct {
	Value string
	Label string
}

var SEARCH_WINDOW = Content{Value: "SW", Label: "検索窓"}

var THEME_MESSAGE1 = Content{Value: "TM1", Label: "テーマ・メッセージ１"}
var THEME_MESSAGE2 = Content{Value: "TM2", Label: "テーマ・メッセージ２"}
var THEME_MESSAGE3 = Content{Value: "TM3", Label: "テーマ・メッセージ３"}
var THEME_MESSAGE4 = Content{Value: "TM4", Label: "テーマ・メッセージ４"}
var THEME_MESSAGE5 = Content{Value: "TM5", Label: "テーマ・メッセージ５"}

var STYLE_MESSAGE1 = Content{Value: "SM1", Label: "スタイル・メッセージ１"}
var STYLE_MESSAGE2 = Content{Value: "SM2", Label: "スタイル・メッセージ２"}

var ITEM_MESSAGE1 = Content{Value: "IM1", Label: "アイテム・メッセージ１"}
var ITEM_MESSAGE2 = Content{Value: "IM2", Label: "アイテム・メッセージ２"}

var ITEM_PARTS1 = Content{Value: "IP1", Label: "アイテム・部品１"}
var ITEM_PARTS2 = Content{Value: "IP2", Label: "アイテム・部品２"}
var ITEM_PARTS3 = Content{Value: "IP3", Label: "アイテム・部品３"}

var SITE_MAP = Content{Value: "SM", Label: "サイトマップ"}

var CONTENT_ITEMS = []Content{
	SEARCH_WINDOW,
	THEME_MESSAGE1,
	THEME_MESSAGE2,
	THEME_MESSAGE3,
	THEME_MESSAGE4,
	THEME_MESSAGE5,
	STYLE_MESSAGE1,
	STYLE_MESSAGE2,
	ITEM_MESSAGE1,
	ITEM_MESSAGE2,
	ITEM_PARTS1,
	ITEM_PARTS2,
	ITEM_PARTS3,
	SITE_MAP}

var CONTENT_ITEMS_DEFAULT = []string{
	SEARCH_WINDOW.Value,
	THEME_MESSAGE1.Value,
	THEME_MESSAGE2.Value,
	ITEM_PARTS1.Value,
	SITE_MAP.Value}
