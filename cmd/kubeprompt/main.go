package main

import (
	"os"

	"github.com/jlesquembre/kubeprompt/pkg/cmd"
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
	cmd.JlCmd2()
}
