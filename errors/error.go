package errors

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/yogameleniawan/golang-error-lines/logs"
)

func Error(message string) {
	// get log file path
	logFilePath := logs.WriteLogFile()

	// Open file to write log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed open file log: %v", err)
	}
	defer logFile.Close()

	// Configure logging to set output to the log file
	log.SetOutput(logFile)

	logger := log.New(logFile, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

func Info(message string) {
	// get log file path
	logFilePath := logs.WriteLogFile()

	// Open file to write log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed open file log: %v", err)
	}
	defer logFile.Close()

	// Configure logging to set output to the log file
	log.SetOutput(logFile)

	logger := log.New(logFile, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

func Warning(message string) {
	// get log file path
	logFilePath := logs.WriteLogFile()

	// Open file to write log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed open file log: %v", err)
	}
	defer logFile.Close()

	// Configure logging to set output to the log file
	log.SetOutput(logFile)

	logger := log.New(logFile, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

func ErrorLines(err error) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		fmt.Printf("Error: %s\nFile: %s\nLine: %d\n", err, file, line)
	} else {
		fmt.Printf("Error: %s\n", err)
	}
}
