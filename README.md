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
./go masterCreate --name=master-01
# or
./go masterCreate --memory=2G --disk=25G --name=master-01
```

### Create a k8s worker node

```sh
./go masterCreate --name=worker-01
# or
./go masterCreate --memory=2G --disk=25G --name=worker-01
```
