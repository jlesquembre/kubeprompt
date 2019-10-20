package main

import (
	"fmt"
	"os"

	"github.com/jlesquembre/kubeprompt/pkg/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func mainOO() {

	// kubeconfig := path.Join(
	// 	os.Getenv("HOME"), ".kube", "config",
	// )
	// config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(config)
	// fmt.Println(config.rawConfig.CurrentContext)

	//----
	//----
	//----
	//----

	flags := pflag.NewFlagSet("kubeprompt", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdNamespace(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	var Print bool
	var rootCmd = &cobra.Command{
		Use:   "kubeprompt",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		Run: func(ccmd *cobra.Command, args []string) {
			fmt.Println("flag:", Print)
			cmd.JlCmd2()
			// Do Stuff Here
		},
	}
	rootCmd.PersistentFlags().BoolVarP(&Print, "print", "p", false, "print without checking if enabled")
	rootCmd.Execute()
	// cmd.JlCmd2()
}
