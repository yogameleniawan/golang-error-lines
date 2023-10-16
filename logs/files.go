package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func WriteLogFile() string {
	// Mendapatkan direktori kerja saat ini (working directory)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Gagal mendapatkan working directory: %v", err)
	}

	// Target "modules" ini merupakan folder yang pasti digunakan
	// karena semua proses logging dilakukan di dalam folder modules

	target := os.Getenv("APP_PROJECT_ROOT")

	index := strings.Index(wd, target)

	var working_directory string
	if index != -1 {
		working_directory = wd[:index]
	}

	working_directory = working_directory + "/" + target

	// Path relatif ke folder "logs" dari working directory saat ini
	logFolder := filepath.Join(working_directory, "logs")

	// Membuat folder log (jika belum ada)
	if _, err := os.Stat(logFolder); os.IsNotExist(err) {
		err := os.Mkdir(logFolder, 0755) // 0755 permission mode
		if err != nil {
			log.Fatalf("Gagal membuat folder log: %v", err)
		}
	}

	// Get tanggal sekarang
	tanggalSekarang := time.Now().Format("2006-01-02")

	// Nama file log
	logFileName := fmt.Sprintf("%s.log", tanggalSekarang)

	// Penggabungan path folder log dengan nama file log
	logFilePath := filepath.Join(logFolder, logFileName)

	return logFilePath
}
