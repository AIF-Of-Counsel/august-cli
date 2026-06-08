// Copyright 2026 vera. Licensed under Apache-2.0. See LICENSE.

// Command august-pp-cli is the entrypoint for the August Printing Press CLI.
package main

import (
	"os"

	"github.com/AIF-Of-Counsel/august-cli/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		// cli.Execute (and Cobra) already write the error to stderr; here we
		// only translate it into the conventional typed exit code.
		os.Exit(cli.ExitCode(err))
	}
}
