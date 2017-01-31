package olog

import (
	"fmt"
)

//DirectLogger - logger that writes directly to all registered writers
type DirectLogger struct {
	writers map[string]Writer
}

//Log - logs a message with given level and module
func (dl *DirectLogger) Log(level Level,
	module string,
	fmtstr string,
	args ...interface{}) {
	// args = append(args, level, module)
	msg := fmt.Sprintf(string(level)+" "+module+" "+fmtstr, args...)
	fmt.Println(msg)
	for _, writer := range dl.writers {
		if writer.IsEnabled() {
			writer.Write(msg)
		}
	}
}

//RegisterWriter - registers a writer
func (dl *DirectLogger) RegisterWriter(writer Writer) {
	if writer != nil {
		dl.writers[writer.UniqueID()] = writer
	}
}

//RemoveWriter - removes a writer with given ID
func (dl *DirectLogger) RemoveWriter(uniqueID string) {
	delete(dl.writers, uniqueID)
}

//GetWriter - gives the writer with given ID
func (dl *DirectLogger) GetWriter(uniqueID string) (writer Writer) {
	return dl.writers[uniqueID]
}
