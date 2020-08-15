package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func Info(msg string, ctx ...interface{}) {
	if len(ctx) > 0 {
		logger.Printf("INFO: %s %v\n", msg, ctx)
	} else {
		logger.Printf("INFO: %s\n", msg)
	}
}

func Error(msg string, ctx ...interface{}) {
	if len(ctx) > 0 {
		logger.Printf("😡: %s %v\n", msg, ctx)
	} else {
		logger.Printf("😡: %s\n", msg)
	}
}

func init() {
	logger = log.New(os.Stdout, "[DNS]:", log.LstdFlags)
}
