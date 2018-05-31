package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"

	"src/consts"
	"src/controllers/front"
	"src/core"
	"src/helper"

	"github.com/goji/glogrus"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
)

// Framework : http://goji.io/
// Golang HTTP : https://golang.org/pkg/net/http/
// OR Mapper : http://doc.gorm.io/database.html
func main() {

	// (0-1) reserve close db connection
	master := core.GetMasterConnection()
	slave := core.GetSlaveConnection()

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
	createFrontCommonRouting()

	// (2-1) Page Routing
	createFrontPageRouting()

	// Start server
	if core.IsProduction() {
		// Production environment
		goji.ServeListener(core.GetListener())
	} else {
		// Test environment
		config := core.GetConfig()
		flag.Set("bind", ":"+config.Launch.Front.Port)
		goji.Serve()
	}
}

// Create common routing
func createFrontCommonRouting() {

	// (1) set document root
	docRoot := consts.RELATIVE_PATH + "webroot"
	docPath := http.FileServer(http.Dir(docRoot))
	staticPattern := regexp.MustCompile("^/(css|js|images)")

	goji.Handle(staticPattern, docPath)

	// (2) set logger
	logs := helper.GetFrontLoggerInfo()

	goji.Abandon(middleware.Logger)
	goji.Use(glogrus.NewGlogrus(logs, "app-front"))
}

//
func createFrontPageRouting() {

	goji.Get("/", controllers_front.Index)
	goji.Get("/sitemap.xml", controllers_front.Sitemap)
}
