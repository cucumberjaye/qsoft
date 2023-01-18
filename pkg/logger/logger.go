package logger

import (
	"log"
	"os"
)

var (
	ErrorLogger *log.Logger
)

func InitLogger() {
	ErrorLogger = log.New(os.Stderr, "ERROR:  ", log.Ldate|log.Ltime|log.Lmsgprefix)
}
