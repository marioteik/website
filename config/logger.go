package config

import (
	"log"
	"os"
)

type Logger struct{}

var InfoLog *log.Logger
var ErrorLog *log.Logger

func InitLoggers() *Logger {
	InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{}
}

func (l *Logger) Info(message string) {
	InfoLog.Println(message)
}

func (l *Logger) Error(error error) {
	ErrorLog.Println(error)
}

func (l *Logger) Fatal(error error) {
	ErrorLog.Fatal(error)
}
