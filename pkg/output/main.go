package output

import "fmt"

// LevelSilence hides all output
const LevelSilence = -1

// LevelInfo shows informal messages
const LevelInfo = 0

// LevelDebug shows verbose output
const LevelDebug = 1

var level = LevelInfo

// SetLevel to the given log level
func SetLevel(l int) {
	level = l
}

// Level returns the current
func Level() int {
	return level
}

func output(level int, out string) {
	if isLevel(level) {
		println(out)
	}
}

func outputf(level int, out string, vars ...interface{}) {
	if isLevel(level) {
		fmt.Printf(out+"\n", vars...)
	}
}

func isLevel(l int) bool {
	return level >= l
}

// Info outputs informal message
func Info(out string) {
	output(LevelInfo, out)
}

// Infof same as Info but with format string
func Infof(out string, vars ...interface{}) {
	outputf(LevelInfo, out, vars...)
}

// Debug outputs debug message
func Debug(out string) {
	output(LevelDebug, out)
}

// Debugf same as Debug but with format string
func Debugf(out string, vars ...interface{}) {
	outputf(LevelDebug, out, vars...)
}
