package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//---------------------------------------------------------------------
// Consts
//---------------------------------------------------------------------

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------
var g_master_connection *gorm.DB
var g_slave_connection *gorm.DB

//---------------------------------------------------------------------
// Method
//---------------------------------------------------------------------
// Initialize
func init() {
	g_master_connection = createMasterConnect()
	g_slave_connection = createSlaveConnect()
}

// Get master connection
func GetMasterConnection() *gorm.DB {
	return g_master_connection
}

// Get slave connection
func GetSlaveConnection() *gorm.DB {
	return g_slave_connection
}

// Create Master connection
func createMasterConnect() *gorm.DB {
	config := GetConfig()
	return createConnect(&config.Master)
}

// Create Slave Connection
func createSlaveConnect() *gorm.DB {
	config := GetConfig()
	return createConnect(&config.Slave)
}

// Create Database Connection
func createConnect(dbInfo *DBInfo) *gorm.DB {

	DBMS := "mysql"
	USER := dbInfo.User
	PASS := dbInfo.Pass
	HOST := dbInfo.Host
	SCHEMA := dbInfo.Schema

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + SCHEMA + "?charset=utf8&parseTime=True&loc=Local"
	db, error := gorm.Open(DBMS, CONNECT)

	if error != nil {
		panic(error.Error())
	}

	return db
}
