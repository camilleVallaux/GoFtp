package main

import (
	"log"
	"os"
)

var logger *log.Logger

func logInit(out string) {
	logFile, error := os.Create(out)
	if error != nil {
		panic(error)
	}

	logger = log.New(logFile, "Ftp server:", log.Ldate|log.Ltime)
}

func logMsg(msg string) {
	logger.Println(msg)
}

func logError(msg string) {
	logger.Println("[Error] ", msg)
}

func logFatal(msg string) {
	logger.Fatalln(msg)
}

func logPanic(msg string) {
	logger.Panicln(msg)
}
