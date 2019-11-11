package prompt

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"

	. "github.com/logrusorgru/aurora"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
)

var tempDir = filepath.Join(os.TempDir(), "kubeprompt")

func isPromptActive(path string) bool {
	return strings.HasPrefix(path, tempDir+"/")
}

func exit(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

func subShell(extraEnv map[string]string) {

	for k, v := range extraEnv {
		os.Setenv(k, v)
	}
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/usr/bin/env bash"
	}

	shellCmd := strings.Fields(shell)
	cmd := exec.Command(shellCmd[0], shellCmd[1:]...)

	// catch and forwards all signals
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	go func() {
		for {
			s := <-c
			cmd.Process.Signal(s)
		}
	}()

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func enableKubeprompt(config clientcmd.ClientConfig) {
	err := os.MkdirAll(tempDir, os.ModePerm)
	exit(err)
	tmpfile, err := ioutil.TempFile(tempDir, "kubeconfig.*.yaml")
	configFile := tmpfile.Name()
	exit(err)
	defer os.Remove(configFile)

	rawConfig, err := config.RawConfig()
	exit(err)
	clientcmd.WriteToFile(rawConfig, configFile)

	subShell(map[string]string{
		"KUBECONFIG": configFile})
}

func printPrompt(config clientcmd.ClientConfig, monochrome bool) {
	rawConfig, err := config.RawConfig()
	ctx := rawConfig.CurrentContext
	ns, _, err := config.Namespace()
	exit(err)
	if ctx == "" {
		ctx = "N/A"
		ns = "N/A"
	}

	if monochrome {
		fmt.Printf("(K8S %s:%s)\n", ctx, ns)
		return
	}

	fmt.Printf("(K8S %s:%s)\n", Bold(Yellow(ctx)), Bold(Magenta(ns)))
}

func Run(force bool, printOnly bool, check bool, monochrome bool) {
	config := genericclioptions.NewConfigFlags(true).ToRawKubeConfigLoader()
	kubeconfigPath := config.ConfigAccess().GetDefaultFilename()
	isActive := isPromptActive(kubeconfigPath)

	if check {
		if isActive {
			fmt.Println("kubeprompt is", Bold("active"))
		} else {
			fmt.Println("kubeprompt is", Bold("NOT"), "active")
		}
		return
	}

	if isActive || force {
		printPrompt(config, monochrome)
	} else if !printOnly {
		enableKubeprompt(config)
	}

}
