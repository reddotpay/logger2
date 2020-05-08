package logger2

import (
	"fmt"
	"runtime"
	"strings"
)

// WhereAmI retrieves call line where function is invoked
func WhereAmI(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}

	function, _, line, _ := runtime.Caller(depth)
	// file.function:line
	return fmt.Sprintf("%s:%d", runtime.FuncForPC(function).Name(), line)
	// return fmt.Sprintf("File: %s  Function: %s Line: %d", chopPath(file), runtime.FuncForPC(function).Name(), line)
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	}
	return original[i+1:]

}
