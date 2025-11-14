package utils

import (
	"log/slog"
	"os"
)

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		slog.Warn("Failed to close file:", "error:", err)
	}
}
