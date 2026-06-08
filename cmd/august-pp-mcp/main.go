// Copyright 2026 vera. Licensed under Apache-2.0. See LICENSE.

// Command august-pp-mcp exposes the August API as MCP tools over stdio.
package main

import (
	"fmt"
	"os"

	"github.com/AIF-Of-Counsel/august-cli/internal/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	serverName    = "august-pp-mcp"
	serverVersion = "1.0.0"
)

func main() {
	s := server.NewMCPServer(serverName, serverVersion)
	mcp.RegisterTools(s)
	mcp.RegisterNovelFeatureTools(s)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "august-pp-mcp: %v\n", err)
		os.Exit(1)
	}
}
