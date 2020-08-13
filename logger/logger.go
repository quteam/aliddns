package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func Info(msg string, ctx ...interface{}) {
	logger.Printf("INFO %s %v\n", msg, ctx)
}

func Error(msg string, ctx ...interface{}) {
	logger.Printf("ERROR %s %v\n", msg, ctx)
}

func init() {
	logger = log.New(os.Stdout, "[ALIDDNS]", log.LstdFlags)
}
