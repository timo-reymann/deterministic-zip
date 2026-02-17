package main

import "github.com/timo-reymann/deterministic-zip/cmd"
import _ "embed"

//go:embed NOTICE
var noticeContent string

func main() {
	cmd.Execute(noticeContent)
}
