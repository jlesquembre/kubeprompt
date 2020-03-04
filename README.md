# kubeprompt

Show the current Kubernetes context and namespace in your prompt

![prompt](imgs/kubeprompt.png)

## Installation

### Manual install

Pre-built binaries are available for linux and MacOS on the
[releases](https://github.com/jlesquembre/kubeprompt/releases) page.

### Homebrew (macOS and Linux)

```bash
brew tap jlesquembre/kubeprompt https://github.com/jlesquembre/kubeprompt/
brew install kubeprompt
```

### NixOS

`kubeprompt` is available in the
[Nix Packages collection](https://github.com/NixOS/nixpkgs/blob/master/pkgs/development/tools/kubeprompt/default.nix).

To install it globally, add it to your `systemPackages`. If you just want to try
it, you can do it in a Nix shell:

```bash
nix-shell -p kubeprompt
```

## Usage

`kubeprompt` is shell-agnostic and thus works on any shell, but every shell has
a different way to customize its prompt. Some examples:

fish shell:

```sh
function fish_prompt
  echo (kubeprompt -f default) '>'
end
```

Zsh:

```sh
PROMPT='$(kubeprompt -f default)'$PROMPT
```

Bash:

```sh
PS1="[\u@\h \W \$(kubeprompt -f default)]\$ "
```

`kubeprompt` will print to stdout information about the current cluster, but
first it needs to be enabled. It is considered enabled if the environment
variable `KUBECONFIG` is set to a file in the `$TMP/kubeprompt` directory. To
enable it, just execute `kubeprompt`.

`kubeprompt` command will print the current K8S context and namespace, if
kubeprompt is enabled. If not, it will start a sub shell with kubeprompt
enabled. The sub shell to launch depends on the value of the environment
variable `SHELL`, starting a `bash` shell if is not defined.

Valid flags:

- `-f`, `--format` custom format string
- `-c`, `--check` print information about kubeprompt status
- `-h`, `--help` help for kubeprompt
- `-v`, `--version` print the version

## Customize prompt

It is possible to customize the prompt using go templates. The templates have
access to 3 variables, `Ctx`, `Ns` and `Enabled`, and to the
[color functions provided by Aurora](https://github.com/logrusorgru/aurora#standard-and-bright-colors),
plus the `Bold`, `Faint`, `Italic` and `Underline` functions.

The default format is

```go
{{if .Enabled}}(K8S {{.Ctx | Yellow | Bold}}|{{.Ns | Magenta | Bold}}){{end}}
```

You have 2 options to provide your format string, with `-f` / `--format` option,
or setting the environment variable `KUBEPROMPT_FORMAT`. If both are provided,
`-f` is used.

Some examples:

- No colors:

```go
'(⎈ {{.Ctx}}|{{.Ns}})'
```

- Print only if kubeprompt is enabled:

```go
'{{if .Enabled}}{{"⎈" | Black | BgWhite }} {{.Ctx}}|{{.Ns}}{{end}}'
```

- Print `k8s` string with a different color if kubeprompt is enabled:

```go
'{{if .Enabled}}{{"k8s"|Green|Bold}}{{else}}{{"k8s"|Red}}{{end}} {{.Ctx}}|{{.Ns}}'
```

## Workflows

Usually you'll call `kubeprompt -f default` in your dot files. In the terminal,
you usually want to call `kubeprompt` to enable it, because after it, you can be
confident about the information in your terminal, since every terminal will have
its own kubeconfig.

You can decide if you want to show the kubernetes information always or only
when _kubeprompt_ is enabled. If you want to show the information always, I
recommend to use different colors to know if _kubeprompt_ is enabled or not. I
consider the disabled color as a warning, because in this case, you cannot be
sure if the information is accurate. Since your kubeconfig is global, other
applications (or yourself in other terminal) can change the global kubernetes
context.

## F.A.Q.

### Why to copy kubeconfig and start a sub shell?

If you are working with multiple contexts/namespaces in multiple terminals, all
of them will modify the same kubeconfig file when you update the
context/namespace. For more context, see
[kubernetes PR #60044](https://github.com/kubernetes/kubernetes/pull/60044#issuecomment-405420482)

Let's suppose that you are working in 2 terminals with `kubectl` in context
_foo_. Since you are using `kubeprompt`, your prompt shows that the context is
_foo_. Now, you change the context in one terminal to _bar_, and after it you
change the focus to the other terminal. Since the prompt information is static,
in that terminal the prompt still says that the context is _foo_, but that
information is old, and you could execute a command in the wrong context.

To avoid it, `kubeprompt` creates a copy of your current kubeconfig per terminal
and sets the `KUBECONFIG` environment variable to that copy. If now you change
the context or the namespace, only that terminal will be affected. If you want
to disable `kubeprompt` on one terminal, you just need to press `CTRL+d`

## Related tools

- [kubectx](https://github.com/ahmetb/kubectx)
- [kube-ps1](https://github.com/jonmosco/kube-ps1)
- [fish-kubectl-prompt](https://github.com/Ladicle/fish-kubectl-prompt)
