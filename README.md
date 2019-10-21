Current context / NS is saved in the kubeconfig file, see:

https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/
https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/

The only source of true is the kubeconfig file. kubectl uses the default
kubeconfig file, \$HOME/.kube/config, unless the KUBECONFIG environment variable
does exist. KUBECONFIG is the only environment variable readed by kubectl.

If we are working with multiple contexts/NS in multiple terminals, all of them
will modify the same kubeconfig file when we update the context/NS.

One possible trick to have different configurations per terminal, is to write
the output of `kubectl config view --flatten` to a temporal file (unique per
terminal), and set KUBECONFIG value to that temporal file. We can also create an
env variable (e.g.: LOCAL_KUBECONFIG) to know if KUBECONFIG is pointing to a
temporal file

See https://github.com/kubernetes/kubernetes/pull/60044#issuecomment-405420482

TODO if the namespace is not set, we get an error Error: no namespace is set for
your current context: "gke_spr-sandbox-1_us-central1-a_micro-cluster-1"

```
go run cmd/kubeprompt.go
```

based on: https://github.com/kubernetes/sample-cli-plugin

# Usage

We offer to commands, `kubeprompt` and `kubeon`.

`kubeprompt` will print to stdout information about the current cluster, but
first it needs to be enabled. It is considered enabled if the environment
variable `KUBECONFIG` is set to a file in the `$TMP/kubeprompt` directory

`kubeon` will enable the prompt. To do that, it will copy your current
`KUBECONFIG` to a temporal file in the `$TMP/kubeprompt` directory and launch a
new shell.

`kubeprompt`: it will print the current K8S context and namespace, if kubeprompt
is enable. If not, it will start a sub shell with kubeprompt enabled.

Flags:

- `-p`, `--print-only` print if kubeprompt is enabled, don't do anything if not
- `-f`, `--force` print without checking if kubeprompt is enabled
- `-c`, `--check` print information about kubeprompt status
- `-h`, `--help` help for kubeprompt
- `-v`, `--version` print the version
