# AWS Lambda Layer with kubectl (and helm)

This module exports a single class called `KubectlLayer` which is a `lambda.Layer` that bundles the [`kubectl`](https://kubernetes.io/docs/reference/kubectl/kubectl/) and the [`helm`](https://helm.sh/) command line.

> * Helm Version: 3.5.4
> * Kubectl Version: 1.20.0

Usage:

```go
// KubectlLayer bundles the 'kubectl' and 'helm' command lines
import "github.com/aws/aws-cdk-go/awscdk"

var fn function

fn.addLayers(awscdk.NewKubectlLayer(this, jsii.String("KubectlLayer")))
```

`kubectl` will be installed under `/opt/kubectl/kubectl`, and `helm` will be installed under `/opt/helm/helm`.
