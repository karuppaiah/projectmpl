package main

import (
	"fmt"

	"github.com/Depado/projectmpl/cmd"
	"github.com/Depado/projectmpl/renderer"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// Main command that will be run when no other command is provided on the
// command-line
var rootCmd = &cobra.Command{
	Use:   "projectmpl [renderer] <options>",
	Short: "projectmpl is a boilerplate engine",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		renderer.Render(args[0])
	},
}

// Version command that will display the build number and version (if any)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run:   func(cmd *cobra.Command, args []string) { fmt.Printf("Build: %s\nVersion: %s\n", Build, Version) },
}

func main() {
	// Initialize Cobra and Viper
	cobra.OnInitialize(cmd.Initialize)
	cmd.AddRendererFlags(rootCmd)
	rootCmd.AddCommand(versionCmd)

	// Run the command
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Couldn't start")
	}
}
