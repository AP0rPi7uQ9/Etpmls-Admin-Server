package library

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Library_Logrus *logrus.Logger

func InitLogrus() {
	Library_Logrus = logrus.New()
	// Library_Logrus as JSON instead of the default ASCII formatter.
	Library_Logrus.Formatter = new(logrus.JSONFormatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Library_Logrus.Out = &lumberjack.Logger{
		Filename:   "storage/log/app.log",
		MaxSize:    500, // megabytes
		MaxAge:     30, //days
		Compress:   true, // disabled by default
	}

	// Only log the warning severity or above.
	level, err := logrus.ParseLevel(Config.Log.Level)
	if err != nil {
		level = logrus.WarnLevel
		Library_Logrus.Warning("Set Library_Logrus Level Failed!")
	} else {
		Library_Logrus.Info("Logrus initialized successfully.")
	}
	Library_Logrus.Level = level
}

type Logrus struct {

}

func (this *Logrus) Panic(args ...interface{}) {
	Library_Logrus.Panic(args)
	return
}


func (this *Logrus) Fatal(args ...interface{}) {
	Library_Logrus.Fatal(args)
	return
}


func (this *Logrus) Error(args ...interface{}) {
	Library_Logrus.Error(args)
	return
}


func (this *Logrus) Warning(args ...interface{}) {
	Library_Logrus.Warning(args)
	return
}


func (this *Logrus) Info(args ...interface{}) {
	Library_Logrus.Info(args)
	return
}


func (this *Logrus) Debug(args ...interface{}) {
	Library_Logrus.Debug(args)
	return
}


func (this *Logrus) Trace(args ...interface{}) {
	Library_Logrus.Trace(args)
	return
}


/*func GeneralLog(content interface{})  {
	if strings.ToUpper(os.Getenv("DEBUG")) == "TRUE" {
		fmt.Println(content)
	}
}*/