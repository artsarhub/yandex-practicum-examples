package main

import (
	"sandbox/interfaces/logger"
	"sandbox/interfaces/zoo"
)

func main() {
	log := logger.NewConsoleLogger()
	//log := logger.NewFileLogger("logs/service.log")

	//log1 := logger.NewConsoleLogger()
	//log2 := logger.NewFileLogger("logs/service.log")
	//
	//log1.Info("test info message")
	//log1.Warn("test warn message")
	//log1.Error("test error message")
	//
	//log2.Info("test info message")
	//log2.Warn("test warn message")
	//log2.Error("test error message")

	zoo.EmulateZoo(log)
}
