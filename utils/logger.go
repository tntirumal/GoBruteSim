package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

var (
	instance *log.Logger
	once     sync.Once
	level    = Info
)

var logfile string = fmt.Sprintf("/mnt/HDD_data/sdb4_virtual_backup/projects/go_projects/src/miniproject/logs/%s", "test1.log")

// LogLevel defines logging verbosity used by the package helpers.
//
// The constants Debug, Info, Warn and Error represent increasing
// severity levels; the package helper functions will respect the global
// `level` when printing messages.

// Init initializes the package logger once and sets the desired LogLevel.
// It creates a multi-writer that writes to stdout and a logfile. Calling
// Init multiple times is safe; initialization happens only once.
func Init(l LogLevel) {
	once.Do(func() {
		level = l

		// 1. File-a open pannu (Append mode | Create if not exists | Write Only)
		file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Fatal("Could not open log file")
		}

		// 2. MultiWriter: Console (Stdout) and File renduukkum link pannu
		multi := io.MultiWriter(os.Stdout, file)

		// 3. Logger-a MultiWriter vachi initialize pannu
		instance = log.New(
			multi,
			"Brute-Force Simulator: ",
			log.LstdFlags|log.Lshortfile, // Lshortfile helps to see file name & line number
		)
	})
}

// get returns the singleton logger
func get() *log.Logger {
	if instance == nil {
		Init(Info)
	}
	return instance
}

// ---- Logging helpers ----

func Debugf(format string, v ...any) {
	if level <= Debug {
		message := fmt.Sprintf("[DEBUG] "+format, v...)
		get().Output(2, message)
	}
}

// Infof logs an informational message if the current level is Info or
// lower (more verbose). It preserves file/line information via the
// logger's Output call.
func Infof(format string, v ...any) {
	if level <= Info {
		message := fmt.Sprintf("[INFO] "+format, v...)

		// Output(calldepth, message)
		// 2 means: skip this function (Info) and the log internal call
		get().Output(2, message)
	}
}

func Warnf(format string, v ...any) {
	if level <= Warn {
		message := fmt.Sprintf("[WARN] "+format, v...)
		get().Output(2, message)
	}
}

// Errorlog logs an error message. Note: the helper is named Errorlog to
// avoid colliding with the standard library `log` package symbol.
func Errorlog(format string, v ...any) {
	if level <= Error {
		message := fmt.Sprintf("[ERROR] "+format, v...)
		get().Output(2, message)
	}
}
