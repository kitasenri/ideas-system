package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"

	"src/controllers/api"
	"src/controllers/back"
	"src/controllers/common"
	"src/core"
	"src/helper"

	"github.com/goji/glogrus"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

// Framework : http://goji.io/
// Golang HTTP : https://golang.org/pkg/net/http/
// OR Mapper : http://doc.gorm.io/database.html
func main() {

	// (0-1) reserve close db connection
	master := core.GetMasterConnection()
	slave := core.GetSlaveConnection()

	// close db connection
	defer master.Close()
	defer slave.Close()

	// Create SIGINT channel
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println("Signal is comming : ", sig)
			master.Close()
			slave.Close()
		}
	}()

	// (1-1) Common Routing
	createBackCommonRouting()

	// (2-1) Theme Routing
	createThemeRouting()

	// (2-2) Style Routing
	createStyleRouting()

	// (2-3) Item Routing
	createItemRouting()

	// (2-4) Whitelist Routing
	createWhitelistRouting()

	// (3-1) Page Routing
	createBackPageRouting()

	// Start server
	if core.IsProduction() {
		// Production environment
		goji.ServeListener(core.GetListener())
	} else {
		// Test environment
		config := core.GetConfig()
		flag.Set("bind", ":"+config.Launch.Back.Port)
		goji.Serve()
	}
}

// Create common routing
func createBackCommonRouting() {

	// (1) set document root
	docRoot := core.GetRelativePath() + "webroot"
	docPath := http.FileServer(http.Dir(docRoot))
	staticPattern := regexp.MustCompile("^/(css|js|images)")

	goji.Handle(staticPattern, docPath)

	// (2) set logger
	logs := helper.GetBackLoggerInfo()

	goji.Abandon(middleware.Logger)
	goji.Use(glogrus.NewGlogrus(logs, "app-back"))
}

// Create theme routing
func createThemeRouting() {

	theme := web.New()
	goji.Handle("/theme/*", theme)

	theme.Use(middleware.SubRouter)
	theme.Use(controllers_common.CheckBasicAuth)

	theme.Post("/create", controllers_api.CreateThemeStage)
	theme.Post("/update", controllers_api.UpdateThemeStage)
	theme.Post("/delete", controllers_api.DeleteTheme)
	theme.Post("/sync", controllers_api.SyncTheme)
	theme.Post("/updateDeleted", controllers_api.UpdateThemeDeleted)
}

// Create style routing
func createStyleRouting() {

	style := web.New()
	goji.Handle("/style/*", style)

	style.Use(middleware.SubRouter)
	style.Use(controllers_common.CheckBasicAuth)

	style.Post("/create", controllers_api.CreateStyleStage)
	style.Post("/update", controllers_api.UpdateStyleStage)
	style.Post("/delete", controllers_api.DeleteStyle)
	style.Post("/sync", controllers_api.SyncStyle)
	style.Post("/updateDeleted", controllers_api.UpdateStyleDeleted)
}

//
func createItemRouting() {

	item := web.New()
	goji.Handle("/item/*", item)

	item.Use(middleware.SubRouter)
	item.Use(controllers_common.CheckBasicAuth)

	item.Post("/create", controllers_api.CreateItemStage)
	item.Post("/update", controllers_api.UpdateItemStage)
	item.Post("/delete", controllers_api.DeleteItem)
	item.Post("/sync", controllers_api.SyncItem)
	item.Post("/updateDeleted", controllers_api.UpdateItemDeleted)
}

func createWhitelistRouting() {

	item := web.New()
	goji.Handle("/whitelist/*", item)

	item.Use(middleware.SubRouter)
	item.Use(controllers_common.CheckBasicAuth)

	item.Post("/create", controllers_api.CreateWhitelist)
	item.Post("/delete", controllers_api.DeleteWhitelist)
}

// Back pages
func createBackPageRouting() {

	goji.Get("/", controllers_back.Index)
	goji.Get("/edit_theme", controllers_back.EditTheme)
	goji.Get("/edit_style", controllers_back.EditStyle)
	goji.Get("/edit_item", controllers_back.EditItem)
	goji.Get("/edit_whitelist", controllers_back.EditWhiteList)

	goji.Get("/theme_pages", controllers_back.ThemePages)
	goji.Get("/stage_page", controllers_back.StagePage)

}
