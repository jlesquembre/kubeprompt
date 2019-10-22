# kubeprompt

Show information about your current kubernetes context

![prompt](imgs/kubeprompt.png)

## Installing

Download it from the
[releases](https://github.com/jlesquembre/kubeprompt/releases) page or build it
locally running `make build`

## Usage

`kubeprompt` is shell-agnostic and thus works on any shell, but every shell has
a different way to customize its prompt. Some examples:

fish shell:

```sh
function fish_prompt
  echo (kubeprompt -p) '>'
end
```

Zsh:

```sh
PROMPT='$(kubeprompt -p)'$PROMPT
```

Bash:

```sh
PS1='[\u@\h \W $(kubeprompt -p)]\$ '
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

- `-p`, `--print-only` print if kubeprompt is enabled, don't do anything if not
- `-f`, `--force` print without checking if kubeprompt is enabled
- `-c`, `--check` print information about kubeprompt status
- `-h`, `--help` help for kubeprompt
- `-v`, `--version` print the version

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

To avoid it, `kubepromt` creates a copy of your current kubeconfig per terminal
and sets the `KUBECONFIG` environment variable to that copy. If now you change
the context or the namespace, only that terminal will be affected. If you want
to disable `kubeprompt` on one terminal, you just need to press `CTRL+d`

## Related tools

- [kubectx](https://github.com/ahmetb/kubectx)
- [kube-ps1](https://github.com/jonmosco/kube-ps1)
- [fish-kubectl-prompt](https://github.com/Ladicle/fish-kubectl-prompt)
