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
	Print    bool
	printVer bool
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
			prompt.Run(Print)
		},
	}
	rootCmd.PersistentFlags().BoolVarP(&Print, "print", "p", false, "print without checking if enabled")
	rootCmd.Flags().BoolVarP(&printVer, "version", "v", false, "print the version")
	rootCmd.Execute()
}
