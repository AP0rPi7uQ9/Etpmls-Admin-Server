package core

import (
	"Etpmls-Admin-Server/library"
	"errors"
	"fmt"
)

var (
	LogPanic = OutputLog{Level:PanicLevel}
	LogFatal = OutputLog{Level:FatalLevel}
	LogError = OutputLog{Level:ErrorLevel}
	LogWarn = OutputLog{Level:WarnLevel}
	LogInfo = OutputLog{Level:InfoLevel}
	LogDebug = OutputLog{Level:DebugLevel}
	LogTrace = OutputLog{Level:TraceLevel}
)


const (
	LOG_MODE_ONLY = 1
	CONSOLE_MODE_ONLY = 2
	LOG_CONSOLE_MODE = 3
)

type Level uint32
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)



// Debug errors and custom errors are used as parameters at the same time, and appropriate errors are output according to environment variables.
// 把Debug错误和自定义错误同时作为参数，根据环境变量输出适合的错误。
func GetErrorByIfDebug(err error, msg string) error {
	if IsDebug() {
		return errors.New(msg)
	}
	return err
}

// Generate errors with both custom messages and error messages
// 生成同时带有自定义信息和错误信息的错误
func GenerateErrorWithMessage(msg string, err error) error {
	return errors.New(msg + err.Error())
}


type OutputLog struct {
	Level Level
}

// No matter whether it is in Debug mode, it will output an message
// 无论是否为Debug模式，都输出信息
func (o OutputLog) Output (info interface{}) {
	switch o.Level {
	case PanicLevel:
		switch library.Config.App.Log.Panic {
		case LOG_MODE_ONLY:
			library.Log.Panic(info)
		case CONSOLE_MODE_ONLY:
			fmt.Println(info)
		case LOG_CONSOLE_MODE:
			fmt.Println(info)
			library.Log.Panic(info)
		default:
			fmt.Println(info)
			library.Log.Panic(info)
		}

	case FatalLevel:
		switch library.Config.App.Log.Fatal {
		case LOG_MODE_ONLY:
			library.Log.Fatal(info)
		case CONSOLE_MODE_ONLY:
			fmt.Println(info)
		case LOG_CONSOLE_MODE:
			fmt.Println(info)
			library.Log.Fatal(info)
		default:
			fmt.Println(info)
			library.Log.Fatal(info)
		}

	case ErrorLevel:
		switch library.Config.App.Log.Error {
		case LOG_MODE_ONLY:
			library.Log.Error(info)
		case CONSOLE_MODE_ONLY:
			fmt.Println(info)
		case LOG_CONSOLE_MODE:
			fmt.Println(info)
			library.Log.Error(info)
		default:
			fmt.Println(info)
			library.Log.Error(info)
		}

	case WarnLevel:
		switch library.Config.App.Log.Warning {
		case LOG_MODE_ONLY:
			library.Log.Warning(info)
		case CONSOLE_MODE_ONLY:
			fmt.Println(info)
		case LOG_CONSOLE_MODE:
			fmt.Println(info)
			library.Log.Warning(info)
		default:
			fmt.Println(info)
			library.Log.Warning(info)
		}

	case InfoLevel:
		switch library.Config.App.Log.Info {
		case LOG_MODE_ONLY:
			library.Log.Info(info)
		case CONSOLE_MODE_ONLY:
			fmt.Println(info)
		case LOG_CONSOLE_MODE:
			fmt.Println(info)
			library.Log.Info(info)
		default:
			fmt.Println(info)
			library.Log.Info(info)
		}

	case DebugLevel:
		switch library.Config.App.Log.Debug {
		case LOG_MODE_ONLY:
			library.Log.Debug(info)
		case CONSOLE_MODE_ONLY:
			fmt.Println(info)
		case LOG_CONSOLE_MODE:
			fmt.Println(info)
			library.Log.Debug(info)
		default:
			fmt.Println(info)
			library.Log.Debug(info)
		}

	case TraceLevel:
		switch library.Config.App.Log.Trace {
		case LOG_MODE_ONLY:
			library.Log.Trace(info)
		case CONSOLE_MODE_ONLY:
			fmt.Println(info)
		case LOG_CONSOLE_MODE:
			fmt.Println(info)
			library.Log.Trace(info)
		default:
			fmt.Println(info)
			library.Log.Trace(info)
		}

	}
}

// If it is currently in Debug mode, it will output an return message, if it is in production mode, it will output a custom message
// 若当前为Debug模式，则输出返回信息，若为生产模式，则输出自定义信息
func (o OutputLog) OutputDebug (err error, msg interface{}) {
	var m interface{}
	if IsDebug() {
		m = err
	} else {
		m = msg
	}

	switch o.Level {
	case PanicLevel:
		switch library.Config.App.Log.Panic {
		case LOG_MODE_ONLY:
			library.Log.Panic(m)
		case CONSOLE_MODE_ONLY:
			fmt.Println(m)
		case LOG_CONSOLE_MODE:
			fmt.Println(m)
			library.Log.Panic(m)
		default:
			fmt.Println(m)
			library.Log.Panic(m)
		}

	case FatalLevel:
		switch library.Config.App.Log.Fatal {
		case LOG_MODE_ONLY:
			library.Log.Fatal(m)
		case CONSOLE_MODE_ONLY:
			fmt.Println(m)
		case LOG_CONSOLE_MODE:
			fmt.Println(m)
			library.Log.Fatal(m)
		default:
			fmt.Println(m)
			library.Log.Fatal(m)
		}

	case ErrorLevel:
		switch library.Config.App.Log.Error {
		case LOG_MODE_ONLY:
			library.Log.Error(m)
		case CONSOLE_MODE_ONLY:
			fmt.Println(m)
		case LOG_CONSOLE_MODE:
			fmt.Println(m)
			library.Log.Error(m)
		default:
			fmt.Println(m)
			library.Log.Error(m)
		}

	case WarnLevel:
		switch library.Config.App.Log.Warning {
		case LOG_MODE_ONLY:
			library.Log.Warning(m)
		case CONSOLE_MODE_ONLY:
			fmt.Println(m)
		case LOG_CONSOLE_MODE:
			fmt.Println(m)
			library.Log.Warning(m)
		default:
			fmt.Println(m)
			library.Log.Warning(m)
		}

	case InfoLevel:
		switch library.Config.App.Log.Info {
		case LOG_MODE_ONLY:
			library.Log.Info(m)
		case CONSOLE_MODE_ONLY:
			fmt.Println(m)
		case LOG_CONSOLE_MODE:
			fmt.Println(m)
			library.Log.Info(m)
		default:
			fmt.Println(m)
			library.Log.Info(m)
		}

	case DebugLevel:
		switch library.Config.App.Log.Debug {
		case LOG_MODE_ONLY:
			library.Log.Debug(m)
		case CONSOLE_MODE_ONLY:
			fmt.Println(m)
		case LOG_CONSOLE_MODE:
			fmt.Println(m)
			library.Log.Debug(m)
		default:
			fmt.Println(m)
			library.Log.Debug(m)
		}

	case TraceLevel:
		switch library.Config.App.Log.Trace {
		case LOG_MODE_ONLY:
			library.Log.Trace(m)
		case CONSOLE_MODE_ONLY:
			fmt.Println(m)
		case LOG_CONSOLE_MODE:
			fmt.Println(m)
			library.Log.Trace(m)
		default:
			fmt.Println(m)
			library.Log.Trace(m)
		}

	}
}

// Automatically output Debug, if it is a debug environment, it will output custom information + Error, if it is not a Debug environment, it will output custom information
// 自动输出Debug，如果是debug环境，则输出自定义信息+Error，如果不是Debug环境，输出自定义信息
func (o OutputLog) AutoOutputDebug (msg interface{}, err error) {
	v, ok := msg.(string);
	if !ok {
		o.OutputDebug(err, msg)
		return
	}

	o.OutputDebug(GenerateErrorWithMessage(v + "Error: ", err), msg)
	return
}
