package utils

import (
	"log"
	"os"
)

// Logger defines a struct for logging throughout the application.
type Logger struct {
	logFile *os.File
	logger  *log.Logger
}

// NewLogger creates a new instance of Logger.
func NewLogger(logFilePath string) (*Logger, error) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := log.New(file, "", log.Ldate|log.Ltime)

	return &Logger{
		logFile: file,
		logger:  logger,
	}, nil
}

// LogInfo logs an informational message.
func (l *Logger) LogInfo(message string) {
	l.logger.Printf("[INFO] %s", message)
}

// LogError logs an error message.
func (l *Logger) LogError(message string) {
	l.logger.Printf("[ERROR] %s", message)
}

// Close closes the logger and the log file.
func (l *Logger) Close() error {
	err := l.logFile.Close()
	if err != nil {
		return err
	}

	return nil
}
