package go_error_lines

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

	// Membuka file log untuk penulisan
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Gagal membuka file log: %v", err)
	}
	defer logFile.Close()

	// Konfigurasi logger untuk menulis ke file
	log.SetOutput(logFile)

	logger := log.New(logFile, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

func Info(message string) {
	// get log file path
	logFilePath := logs.WriteLogFile()

	// Membuka file log untuk penulisan
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Gagal membuka file log: %v", err)
	}
	defer logFile.Close()

	// Konfigurasi logger untuk menulis ke file
	log.SetOutput(logFile)

	logger := log.New(logFile, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

func Warning(message string) {
	// get log file path
	logFilePath := logs.WriteLogFile()

	// Membuka file log untuk penulisan
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Gagal membuka file log: %v", err)
	}
	defer logFile.Close()

	// Konfigurasi logger untuk menulis ke file
	log.SetOutput(logFile)

	logger := log.New(logFile, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

func ErrorLines(err error) {
	_, file, line, ok := runtime.Caller(1)

	var message string
	if ok {
		message = fmt.Sprintf("Error: %s\nFile: %s\nLine: %d\n", err, file, line)

		fmt.Printf("Error: %s\nFile: %s\nLine: %d\n", err, file, line)
	} else {
		message = fmt.Sprintf("Error: %s\n", err)
		fmt.Printf("Error: %s\n", err)
	}

	Error(message)
}
