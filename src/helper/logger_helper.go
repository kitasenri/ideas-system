package helper

import (
	"os"
	"src/consts"
	"src/core"

	"github.com/sirupsen/logrus"
)

var g_front_info *logrus.Logger
var g_front_error *logrus.Logger

var g_back_info *logrus.Logger
var g_back_error *logrus.Logger

/**
 * Initialize
 */
func init() {

	config := core.GetConfig()

	// create back logger
	f1, _ := os.OpenFile(
		consts.RELATIVE_PATH+config.Log.Front.Info,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	g_front_info = logrus.New()
	g_front_info.Formatter = new(logrus.JSONFormatter)
	g_front_info.Out = f1

	// create front logger
	f2, _ := os.OpenFile(
		consts.RELATIVE_PATH+config.Log.Front.Error,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	g_front_error = logrus.New()
	g_front_error.Formatter = new(logrus.JSONFormatter)
	g_front_error.Out = f2

	// create back logger
	f3, _ := os.OpenFile(
		consts.RELATIVE_PATH+config.Log.Back.Info,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	g_back_info = logrus.New()
	g_back_info.Formatter = new(logrus.JSONFormatter)
	g_back_info.Out = f3

	// create front logger
	f4, _ := os.OpenFile(
		consts.RELATIVE_PATH+config.Log.Back.Error,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	g_back_error = logrus.New()
	g_back_error.Formatter = new(logrus.JSONFormatter)
	g_back_error.Out = f4

}

//------------------------------------------------------------
// Get Logger
//------------------------------------------------------------
/**
 * Get logger
 */
func GetFrontLoggerInfo() *logrus.Logger {
	return g_front_info
}

/**
 * Get logger
 */
func GetBackLoggerInfo() *logrus.Logger {
	return g_back_info
}

//------------------------------------------------------------
// Get Logger
//------------------------------------------------------------
func Info_F(value interface{}) {
	g_front_info.Info(value)
}

func Info_B(value interface{}) {
	g_back_info.Info(value)
}

func Error_F(value interface{}) {
	g_front_error.Info(value)
}

func Error_B(value interface{}) {
	g_back_error.Info(value)
}
