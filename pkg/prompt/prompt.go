package prompt

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/logrusorgru/aurora"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
)

var tempDir = filepath.Join(os.TempDir(), "kubeprompt")
var defaultFormat = `{{if .Enabled}}(K8S {{.Ctx | Yellow | Bold}}|{{.Ns | Magenta | Bold}}){{end}}`

// KubeData holds k8s data to render the template
type KubeData struct {
	Ctx, Ns string
	Enabled bool
}

var funcMap = template.FuncMap{
	"Bold":      aurora.Bold,
	"Faint":     aurora.Faint,
	"Italic":    aurora.Italic,
	"Underline": aurora.Underline,

	"Black":   aurora.Black,
	"Red":     aurora.Red,
	"Green":   aurora.Green,
	"Yellow":  aurora.Yellow,
	"Blue":    aurora.Blue,
	"Magenta": aurora.Magenta,
	"Cyan":    aurora.Cyan,
	"White":   aurora.White,

	"BrightBlack":   aurora.BrightBlack,
	"BrightRed":     aurora.BrightRed,
	"BrightGreen":   aurora.BrightGreen,
	"BrightYellow":  aurora.BrightYellow,
	"BrightBlue":    aurora.BrightBlue,
	"BrightMagenta": aurora.BrightMagenta,
	"BrightCyan":    aurora.BrightCyan,
	"BrightWhite":   aurora.BrightWhite,

	"BgBlack":   aurora.BgBlack,
	"BgRed":     aurora.BgRed,
	"BgGreen":   aurora.BgGreen,
	"BgYellow":  aurora.BgYellow,
	"BgBlue":    aurora.BgBlue,
	"BgMagenta": aurora.BgMagenta,
	"BgCyan":    aurora.BgCyan,
	"BgWhite":   aurora.BgWhite,

	"BgBrightBlack":   aurora.BgBrightBlack,
	"BgBrightRed":     aurora.BgBrightRed,
	"BgBrightGreen":   aurora.BgBrightGreen,
	"BgBrightYellow":  aurora.BgBrightYellow,
	"BgBrightBlue":    aurora.BgBrightBlue,
	"BgBrightMagenta": aurora.BgBrightMagenta,
	"BgBrightCyan":    aurora.BgBrightCyan,
	"BgBrightWhite":   aurora.BgBrightWhite,
}

// TEMPLATE Default template used to render k8s data
var TEMPLATE = template.New("prompt").Funcs(funcMap)

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

func getFormatStr(format string) string {
	if format != "" {
		return format
	}

	envFormat := os.Getenv("KUBEPROMPT_FORMAT")

	if envFormat != "" {
		return envFormat
	}

	return defaultFormat
}

func printPrompt(config clientcmd.ClientConfig, isActive bool, format string) {
	rawConfig, err := config.RawConfig()
	ctx := rawConfig.CurrentContext
	ns, _, err := config.Namespace()
	if ctx == "" || err != nil {
		ctx = "N/A"
		ns = "N/A"
	}

	t, err := TEMPLATE.Parse(format)
	exit(err)
	r := KubeData{ctx, ns, isActive}
	exit(t.Execute(os.Stdout, r))
}

// Run CLI entry point
func Run(printInfo bool, check bool, format string) {
	config := genericclioptions.NewConfigFlags(true).ToRawKubeConfigLoader()
	kubeconfigPath := config.ConfigAccess().GetDefaultFilename()
	isActive := isPromptActive(kubeconfigPath)

	if check {
		if isActive {
			fmt.Println("kubeprompt is", aurora.Bold("active"))
		} else {
			fmt.Println("kubeprompt is", aurora.Bold("NOT"), "active")
		}
		return
	}

	if isActive || printInfo {
		printPrompt(config, isActive, getFormatStr(format))
	} else {
		enableKubeprompt(config)
	}
}
