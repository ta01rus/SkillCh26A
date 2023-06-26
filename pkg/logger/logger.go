package logger

import (
	"fmt"
	"io"
	"os"
)

type ConsoleLogger struct {
}

func NewConsoleLoger() ConsoleLogger {
	return ConsoleLogger{}

}

func (c *ConsoleLogger) Info(f string, arg ...any) (n int, err error) {
	s := fmt.Sprintf(f, arg...)
	return io.WriteString(os.Stderr, s)
}
func (c *ConsoleLogger) Error(f string, arg ...any) (n int, err error) {
	s := fmt.Sprintf(f, arg...)
	return io.WriteString(os.Stdout, s)
}
