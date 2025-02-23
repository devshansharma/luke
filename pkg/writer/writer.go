package writer

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	instance *OutputWriter
	once     sync.Once
)

type OutputWriter struct {
	responseWriter io.Writer
	logWriter      io.Writer
}

func (w *OutputWriter) Close() {
	if closer, ok := w.responseWriter.(io.Closer); ok {
		closer.Close()
	}

	if closer, ok := w.logWriter.(io.Closer); ok {
		closer.Close()
	}
}

func GetInstance() *OutputWriter {
	if instance == nil {
		fmt.Fprintln(os.Stderr, "OutputWriter not initialized. Call InitOutputWriter first.")
		os.Exit(1)
	}

	return instance
}

func InitOutputWriter(outputFile, logFile string, silent, verbose bool) *OutputWriter {
	once.Do(func() {
		writer := &OutputWriter{
			responseWriter: os.Stdout,
			logWriter:      os.Stderr,
		}

		if verbose {
			writer.logWriter = os.Stdout
		}

		if !silent {
			if outputFile != "" {
				file, err := os.Create(outputFile)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error creating response file:", err)
					os.Exit(1)
				}
				writer.responseWriter = file
			}

			if logFile != "" {
				file, err := os.Create(logFile)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error creating log file:", err)
					os.Exit(1)
				}
				writer.logWriter = file
			}
		} else {
			writer.logWriter = io.Discard
		}

		instance = writer
	})

	return instance
}

// Log writes logs to the logWriter
func (o *OutputWriter) Log(format string, a ...interface{}) {
	fmt.Fprintf(o.logWriter, format+"\n", a...)
}

// Error writes errors to the logWriter
func (o *OutputWriter) Error(format string, a ...interface{}) {
	fmt.Fprintf(o.logWriter, "ERROR: "+format+"\n", a...)
}

// Response writes responses to the responseWriter
func (o *OutputWriter) Response(format string, a ...interface{}) {
	fmt.Fprintf(o.responseWriter, format+"\n", a...)
}
