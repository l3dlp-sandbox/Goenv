// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// main package
package main

import (
	"os"
	"strings"

	"github.com/clivern/goenv/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	cmd.Version = version
	cmd.Commit = commit
	cmd.Date = date
	cmd.BuiltBy = builtBy
	cmd.HOME = strings.TrimSpace(os.Getenv("HOME"))

	cmd.Execute()
}
