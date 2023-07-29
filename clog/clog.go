package clog

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	simplelogger int = iota
	stdlogger
)

type ILogger interface {
	RunLogger() *log.Logger
}
type LoggerFactory struct{}
type SimpleLogger struct{}
type StdLogger struct{}

func (*SimpleLogger) RunLogger() *log.Logger {
	fmt.Println("Simple Log Mode")
	simplelog := log.New(io.MultiWriter(os.Stdout), "", log.Lmicroseconds|log.LstdFlags)
	return simplelog
}

func (*StdLogger) RunLogger() *log.Logger {
	fmt.Println("Std Log Mode")
	osFile, err := os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("Create log file failed: %v", err)
	}
	stdlog := log.New(io.MultiWriter(os.Stdout, osFile), "", log.Lmicroseconds|log.LstdFlags)
	return stdlog
}

func (t *LoggerFactory) new(typ int) ILogger {
	switch typ {
	case simplelogger:
		return new(SimpleLogger)
	case stdlogger:
		return new(StdLogger)
	default:
		return nil
	}
}

func NewSimpleLogger() *log.Logger {
	var factory LoggerFactory
	newsimplelog := factory.new(simplelogger)
	newlog := newsimplelog.RunLogger()
	return newlog
}
func NewStdLogger() *log.Logger {
	var factory LoggerFactory
	newstdlog := factory.new(stdlogger)
	newlog := newstdlog.RunLogger()
	return newlog
}
