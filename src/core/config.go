package core

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"strconv"

	"src/consts"

	_ "github.com/go-sql-driver/mysql"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------
// Config instance
var g_config *Config

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

	flag.Parse()
	arg, error := strconv.Atoi(flag.Arg(0))

	// check resource path
	resource_path := consts.PATH_TEST
	if (error != nil) && (arg == consts.ENV_PRODUCTION) {
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
