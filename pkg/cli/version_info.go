package cli

import (
	"errors"
	"fmt"
	"github.com/timo-reymann/deterministic-zip/pkg/buildinfo"
	"os"
	"runtime"
	"text/tabwriter"
)

var ErrAbort = errors.New("abort")

func addLine(w *tabwriter.Writer, heading string, val string) {
	_, _ = fmt.Fprintf(w, heading+"\t%s\n", val)
}

// PrintVersionInfo prints a tabular list with build info
func PrintVersionInfo() {
	PrintCompactInfo()
	println()
	println("Build information")
	w := tabwriter.NewWriter(os.Stderr, 10, 1, 10, byte(' '), tabwriter.TabIndent)
	addLine(w, "GitSha", buildinfo.GitSha)
	addLine(w, "Version", buildinfo.Version)
	addLine(w, "BuildTime", buildinfo.BuildTime)
	addLine(w, "Go-Version", runtime.Version())
	addLine(w, "OS/Arch", runtime.GOOS+"/"+runtime.GOARCH)
	_ = w.Flush()
}

// PrintCompactInfo outputs a online header with basic infos
func PrintCompactInfo() {
	fmt.Printf("deterministic-zip %s (%s) by Timo Reymann\n", buildinfo.Version, buildinfo.BuildTime)
}
