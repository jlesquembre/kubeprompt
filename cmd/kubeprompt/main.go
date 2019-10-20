package main

import (
	"github.com/jlesquembre/kubeprompt/pkg"
	"github.com/spf13/cobra"
)

func main() {
	var Print bool
	var rootCmd = &cobra.Command{
		Use:   "kubeprompt",
		Short: "K8S info prompt",
		Run: func(cmd *cobra.Command, args []string) {
			prompt.Run(Print)
		},
	}
	rootCmd.PersistentFlags().BoolVarP(&Print, "print", "p", false, "print without checking if enabled")
	rootCmd.Execute()
}
