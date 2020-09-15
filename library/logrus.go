package library

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

var Log *logrus.Logger

func InitLogrus() {
	Log = logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	Log.Formatter = new(logrus.JSONFormatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Log.Out = &lumberjack.Logger{
		Filename:   "storage/log/app.log",
		MaxSize:    500, // megabytes
		MaxAge:     30, //days
		Compress:   true, // disabled by default
	}

	// Only log the warning severity or above.
	level, err := logrus.ParseLevel(Config.Log.Level)
	if err != nil {
		level = logrus.WarnLevel
		Log.Warning("Set Log Level Failed!")
	} else {
		Log.Info("Logrus initialized successfully.")
	}
	Log.Level = level


}

func GeneralLog(content interface{})  {
	if strings.ToUpper(os.Getenv("DEBUG")) == "TRUE" {
		fmt.Println(content)
	}
}