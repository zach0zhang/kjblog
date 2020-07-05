package kjlog

import (
	"errors"
	"log"
	"os"
)

// Define log Level to control log output.
const (
	LevelDebug   = iota // 0
	LevelInfo           // 1
	LevelWarning        // 2
	LevelError          // 3
)

// set default log lever to LevelInfo
var outputLever = LevelInfo

var logger *log.Logger

// InitLog init the kjlog
func InitLog(file string, level int) error {
	err := SetLogLevel(level)
	if err != nil {
		return err
	}
	if file == "" {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		file, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		logger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return nil
}

// SetLogLevel set output level
func SetLogLevel(level int) error {
	if level < LevelDebug || level > LevelError {
		return errors.New("input error level")
	}
	outputLever = level
	return nil
}

// Debug logs
func Debug(args ...interface{}) {
	if outputLever <= LevelDebug {
		logger.SetPrefix("Debug ")
		logger.Println(args...)
	}
}

// Info logs
func Info(args ...interface{}) {
	if outputLever <= LevelInfo {
		logger.SetPrefix("Info ")
		logger.Println(args...)
	}
}

// Warning logs
func Warning(args ...interface{}) {
	if outputLever <= LevelWarning {
		logger.SetPrefix("Warning ")
		logger.Println(args...)
	}
}

// Error logs
func Error(args ...interface{}) {
	if outputLever <= LevelError {
		logger.SetPrefix("Error ")
		logger.Println(args...)
	}
}
