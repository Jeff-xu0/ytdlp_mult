package yt_dlp

import (
	"fmt"
	"os/exec"
	"strconv"
)

// Format command line
func (y *Ytdlp) toCmdString(index int) *exec.Cmd {
	args := y.buildArgs(index)
	return exec.Command("yt-dlp", args...)
}

// handle the arguments
func (y *Ytdlp) buildArgs(index int) []string {
	var args []string

	// URL is processed separately
	if y.Url != "" {
		args = append(args, fmt.Sprintf("%v?p=%v", y.Url, strconv.Itoa(index)))
	}

	// Helper function to append arguments
	addArg := func(field, argName string) {
		if field != "" {
			args = append(args, argName, field)
		}
	}

	// Handling each field
	addArg(y.CookieFromBrowser, "--cookies-from-browser")

	//TODO ... You can add the rest yourself

	return args
}
