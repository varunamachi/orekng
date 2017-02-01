package olog

import "log"

//ConsoleWriter - Log writer that writes to console
type ConsoleWriter struct {
	enabled bool
}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{
		enabled: true,
	}
}

//UniqueID - identifier for console writer
func (cw *ConsoleWriter) UniqueID() string {
	return "console"
}

//Write - writes message to console
func (cw *ConsoleWriter) Write(message string) {
	if cw.enabled {
		log.Printf(message)
	}
}

//Enable - enables or disables console logger based on the passed value
func (cw *ConsoleWriter) Enable(value bool) {
	cw.enabled = value
}

//IsEnabled - tells if the writer is enabled
func (cw *ConsoleWriter) IsEnabled() (value bool) {
	return cw.enabled
}
