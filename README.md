# Setup a Kubernetes cluster using multipass on MacOS

A simple set of scripts and cloud-config to configure a k8s master node (control plane) and a worker node (data plane).

## Dependencies

- MacOS
- brew
- multipass

### Install multipass

```sh
brew install multipass
```

### Create a k8s master node

```sh
./go master --name=master-01
# or
./go master --memory=2G --disk=25G --name=master-01
```

### Create a k8s worker node

```sh
./go worker --name=worker-01
# or
./go worker --memory=2G --disk=25G --name=worker-01
```
