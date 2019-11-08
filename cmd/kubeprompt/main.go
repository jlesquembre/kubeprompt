package main

import (
	"fmt"

	"github.com/jlesquembre/kubeprompt/pkg/prompt"
	"github.com/jlesquembre/kubeprompt/pkg/version"
	"github.com/spf13/cobra"
)

func printVersion(name string) {
	fmt.Printf("%s version %s\n", name, version.Version)
}

var (
	force      bool
	printOnly  bool
	printVer   bool
	check      bool
	monochrome bool
	format     string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "kubeprompt",
		Short: "K8S info prompt",
		Run: func(cmd *cobra.Command, args []string) {
			if printVer {
				printVersion(cmd.Name())
				return
			}
			prompt.Run(force, printOnly, check, monochrome, format)
		},
	}

	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "print without checking if enabled")
	rootCmd.Flags().BoolVarP(&printOnly, "print-only", "p", false, "print only if enabled")
	rootCmd.Flags().BoolVarP(&check, "check", "c", false, "print information about kubeprompt status")
	rootCmd.Flags().BoolVarP(&monochrome, "monochrome", "m", false, "disables colors in output")
	rootCmd.Flags().StringVarP(&format, "format", "", "⎈ %s:%s", "custom format string")
	rootCmd.Flags().BoolVarP(&printVer, "version", "v", false, "print the version")
	rootCmd.Execute()
}
