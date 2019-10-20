package prompt

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	. "github.com/logrusorgru/aurora"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
)

var tempDir string = filepath.Join(os.TempDir(), "kubeprompt")

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
		shell = "/bin/bash"
	}

	cmd := exec.Command(shell)
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

func printPrompt(config clientcmd.ClientConfig) {
	rawConfig, err := config.RawConfig()
	ctx := rawConfig.CurrentContext
	ns, _, err := config.Namespace()
	exit(err)
	fmt.Printf("(K8S %s:%s)\n", Bold(Yellow(ctx)), Bold(Magenta(ns)))
}

func Run(printOnly bool) {
	config := genericclioptions.NewConfigFlags(true).ToRawKubeConfigLoader()
	kubeconfigPath := config.ConfigAccess().GetDefaultFilename()
	if isPromptActive(kubeconfigPath) || printOnly {
		printPrompt(config)
	} else {
		enableKubeprompt(config)
	}
}
