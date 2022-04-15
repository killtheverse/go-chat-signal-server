package main

import (
	"fmt"
	"log"
	"os"
)

var (
    logger *log.Logger 
)

func init() {
    logger = log.New(os.Stdout, "go-chat-server", log.LstdFlags)
}

func Write(format string, v ...interface{}) {
    fmt.Printf(format, v...)
}

func Fatal(format string, v ...interface{}) {
    logger.Fatalf(format, v...)
}
