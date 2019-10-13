package cmd

import (
	"fmt"
	"os"
	"path"
	// "os/exec"
	// "runtime"
	// "strings"
	// "errors"
	"io/ioutil"
	// "github.com/kjk/u"
)

// From:
// https://presstige.io/p/Using-Go-instead-of-bash-for-scripts-6b51885c1f6940aeb40476000d0eb0fc#90fdae84-2318-487f-931e-b3b3a03aeb1e

// func must(err error) {
// 	u.Must(err)
// }

func exit(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

func getHomeDir() string {
	s, err := os.UserHomeDir()
	exit(err)
	return s
}

func cpFile(dstPath, srcPath string) {
	d, err := ioutil.ReadFile(srcPath)
	exit(err)
	err = ioutil.WriteFile(dstPath, d, 0666)
	exit(err)
}

// func setEnv(name string, value string) {
// 	cmd.Env = os.Environ()
// 	cmd.Env = append(cmd.Env, "GOOS=linux")
// }

// func subShell() {
// 	subShell(map[string]string{})
// }
func subShell(extraEnv map[string]string) {

	for k, v := range extraEnv {
		os.Setenv(k, v)
	}
	shell := os.Getenv("SHELL")
	// must(err)
	if shell == "" {
		shell = "/bin/bash"
		// exit(errors.New("'SHELL' environment variable is not set"))
	}
	fmt.Println("Current shell:", shell)
	cpConfig()

	// cmd := exec.Command(shell)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	// cmd.Stdin = os.Stdin
	// cmd.Run()
}

func cpConfig() {
	tempDir := path.Join(os.TempDir(), "kubeprompt")
	err := os.MkdirAll(tempDir, os.ModePerm)
	exit(err)
	file, err := ioutil.TempFile(tempDir, "kubeconfig.*.yaml")
	exit(err)
	// defer os.Remove(file.Name())

	fmt.Println(file.Name())
}
