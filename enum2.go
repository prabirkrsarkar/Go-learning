package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

// Note: the type of LevelTrace is LogLevel while its underlying type is int

var levelNames = []string{
	"Trace",
	"Debug",
	"Info",
	"Warning",
	"Error",
}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError { // Check if the log level is out of range
		return "Unknown"
	}
	return levelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log Level: %d %s\n", level, level.String())
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(10)
}
