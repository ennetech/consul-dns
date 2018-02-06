package logger

import (
	"fmt"
)

func log(level string, string string) {
	fmt.Println("["+level+"]", string)
}

func Debug(string string, what string) {
	log("DEBUG", what+"\n"+string)
}

func Info(string string) {
	log("INFO", string)
}

func Warn(string string) {
	log("WARN", string)
}

func Error(string string) {
	log("ERROR", string)
}

func Fatal(string string) {
	log("FATAL", string)
	panic(string)
}

func Custom(level, string string) {
	log(level, string)
	panic(string)
}
