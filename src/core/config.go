package core

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net"
	"os"

	"src/consts"

	_ "github.com/go-sql-driver/mysql"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------
// Config instance
var g_config *Config
var g_is_production bool = false
var g_listener net.Listener
var g_relative_path string

type Config struct {
	Master DBInfo     `json:"master"`
	Slave  DBInfo     `json:"slave"`
	Log    LogInfo    `json:"log"`
	Launch LaunchInfo `json:"launch"`
}

type DBInfo struct {
	Host   string `json:"host"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Schema string `json:"schema"`
}

type LogInfo struct {
	Front LogPath `json:"front"`
	Back  LogPath `json:"back"`
}

type LogPath struct {
	Info  string `json:"info"`
	Error string `json:"error"`
}

type LaunchInfo struct {
	Front ServerInfo `json:"front"`
	Back  ServerInfo `json:"back"`
}

type ServerInfo struct {
	Port    string `json:"port"`
	Scheme  string `json:"scheme"`
	Domain  string `json:"domain"`
	Context string `json:"context"`
}

//---------------------------------------------------------------------
// Method
//---------------------------------------------------------------------
func init() {

	fd := flag.Uint("fd", 0, "File descriptor to listen and serve.")
	flag.Parse()

	// check resource path
	resource_path := consts.PATH_TEST
	if *fd != 0 {

		listener, err := net.FileListener(os.NewFile(uintptr(*fd), ""))
		if err != nil {
			panic(err)
		}

		g_listener = listener
		g_is_production = true
		resource_path = consts.PATH_PRODUCTION
	}

	// parse resource
	file, error := ioutil.ReadFile(resource_path)
	if error != nil {
		panic(error)
	}

	json.Unmarshal(file, &g_config)
}

// Get config from resources
func GetConfig() *Config {
	return g_config
}

// Is production
func IsProduction() bool {
	return g_is_production
}

// Get listener
func GetListener() net.Listener {
	return g_listener
}

// Get relative path
func GetRelativePath() string {

	if g_is_production {
		return "./../"
	} else {
		return "./../../"
	}

}