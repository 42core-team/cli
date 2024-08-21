package logging

import (
	"log"
	"os"
)

var logFile *os.File

func SetupLogToFile() {
	var err error
	logFile, err = os.OpenFile("application.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Default().Fatalf("error opening file: %v", err)
	}

	log.Default().SetOutput(logFile)
}

func CloseLogToFile() {
	log.Default().SetOutput(os.Stdout)
	logFile.Close()
}
