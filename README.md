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
