package output

import "fmt"

const LevelInfo = 0
const LevelDebug = 1

var level = LevelInfo

func SetLevel(l int) {
	level = l
}

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

func isLevel(l interface{}) bool {
	return level == l
}

func Info(out string) {
	output(LevelInfo, out)
}

func Infof(out string, vars ...interface{}) {
	outputf(LevelInfo, out, vars...)
}

func Debug(out string) {
	output(LevelDebug, out)
}

func Debugf(out string, vars ...interface{}) {
	outputf(LevelDebug, out, vars...)
}
