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
	printVer   bool
	check      bool
	format     string
	tempConfig bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "kubeprompt",
		Short: "K8S info prompt. Call without flags to enable it",
		Run: func(cmd *cobra.Command, args []string) {
			if printVer {
				printVersion(cmd.Name())
				return
			}
			prompt.Run(tempConfig, check, format, cmd.Flags().Changed("format"))
		},
	}

	rootCmd.Flags().BoolVarP(&check, "check", "c", false, "print information about kubeprompt status")
	rootCmd.Flags().BoolVarP(&tempConfig, "temp-config", "t", false, "copy current KUBECONFIG to a temporary file")
	rootCmd.Flags().StringVarP(&format, "format", "f", prompt.DefaultFormat, "print using custom format string")
	rootCmd.Flags().BoolVarP(&printVer, "version", "v", false, "print the version")

	rootCmd.Execute()
}
