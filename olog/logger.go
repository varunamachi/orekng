package olog

import "os"

//Level - gives log level
type Level int

const (
	//TraceLevel - low level debug message
	TraceLevel Level = 1

	//DebugLevel - a debug message
	DebugLevel Level = 2

	//InfoLevel - information message
	InfoLevel Level = 3

	//WarnLevel - warning message
	WarnLevel Level = 4

	//ErrorLevel - error message
	ErrorLevel Level = 5

	//FatalLevel - fatal messages
	FatalLevel Level = 6

	//PrintLevel - prints a output message
	PrintLevel Level = 7
)

//Writer - interface that takes a message and writes it based on
//the implementation
type Writer interface {
	UniqueID() string
	Write(message string)
	Enable(value bool)
	IsEnabled() (value bool)
}

//Logger - interface that defines a logger implementation
type Logger interface {
	//Log - logs a message with given level and module
	Log(level Level,
		module string,
		fmtstr string,
		args ...interface{})

	//RegisterWriter - registers a writer
	RegisterWriter(writer Writer)

	//RemoveWriter - removes a writer with given ID
	RemoveWriter(uniqueID string)

	//GetWriter - gives the writer with given ID
	GetWriter(uniqueID string) (writer Writer)
}

func (level *Level) String() string {
	switch *level {
	case TraceLevel:
		return "[TRACE]"
	case DebugLevel:
		return "[DEBUG]"
	case InfoLevel:
		return "[ INFO]"
	case WarnLevel:
		return "[ WARN]"
	case ErrorLevel:
		return "[ERROR]"
	case FatalLevel:
		return "[FATAL]"
	}
	return "[     ]"
}

var logger = &DirectLogger{}
var logConsole = false
var filterLevel = InfoLevel

//SetLevel - sets the filter level
func SetLevel(level Level) {
	filterLevel = level
}

//GetLevel - gets the filter level
func GetLevel() (level Level) {
	return filterLevel
}

//Trace - trace logs
func Trace(module, fmtStr string, args ...interface{}) {
	logger.Log(TraceLevel, module, fmtStr, args...)
}

//Debug - debug logs
func Debug(module, fmtStr string, args ...interface{}) {
	logger.Log(DebugLevel, module, fmtStr, args...)
}

//Info - information logs
func Info(module, fmtStr string, args ...interface{}) {
	logger.Log(InfoLevel, module, fmtStr, args...)
}

//Warn - warning logs
func Warn(module, fmtStr string, args ...interface{}) {
	logger.Log(WarnLevel, module, fmtStr, args...)
}

//Error - error logs
func Error(module, fmtStr string, args ...interface{}) {
	logger.Log(ErrorLevel, module, fmtStr, args...)
}

//Fatal - error logs
func Fatal(module, fmtStr string, args ...interface{}) {
	logger.Log(FatalLevel, module, fmtStr, args...)
	Print(module, fmtStr, args...)
	os.Exit(-1)
}

//PrintError - error logs
func PrintError(module string, err error) {
	logger.Log(ErrorLevel, module, "%v", err)
}

//PrintFatal - error logs
func PrintFatal(module string, err error) {
	logger.Log(FatalLevel, module, "%v", err)
	os.Exit(-1)
}

//Print - prints the message on console
func Print(module, fmtStr string, args ...interface{}) {
	logger.Log(PrintLevel, module, fmtStr, args)
	// fmt.Printf(fmtStr+"\n", args...)
}
