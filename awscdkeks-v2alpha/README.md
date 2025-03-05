# Amazon EKS V2 Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

The eks-v2-alpha module is a rewrite of the existing aws-eks module (https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_eks-readme.html). This new iteration leverages native L1 CFN resources, replacing the previous custom resource approach for creating EKS clusters and Fargate Profiles.

Compared to the original EKS module, it has the following major changes:

* Use native L1 AWS::EKS::Cluster resource to replace custom resource Custom::AWSCDK-EKS-Cluster
* Use native L1 AWS::EKS::FargateProfile resource to replace custom resource Custom::AWSCDK-EKS-FargateProfile
* Kubectl Handler will not be created by default. It will only be created if users specify it.
* Remove AwsAuth construct. Permissions to the cluster will be managed by Access Entry.
* Remove the limit of 1 cluster per stack
* Remove nested stacks
* API changes to make them more ergonomic.

## Quick start

Here is the minimal example of defining an AWS EKS cluster

```go
cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
})
```

## Architecture

```text
 +-----------------------------------------------+
 | EKS Cluster      | kubectl |  |
 | -----------------|<--------+| Kubectl Handler |
 | AWS::EKS::Cluster             (Optional)      |
 | +--------------------+    +-----------------+ |
 | |                    |    |                 | |
 | | Managed Node Group |    | Fargate Profile | |
 | |                    |    |                 | |
 | +--------------------+    +-----------------+ |
 +-----------------------------------------------+
    ^
    | connect self managed capacity
    +
 +--------------------+
 | Auto Scaling Group |
 +--------------------+
```

In a nutshell:

* EKS Cluster - The cluster endpoint created by EKS.
* Managed Node Group - EC2 worker nodes managed by EKS.
* Fargate Profile - Fargate worker nodes managed by EKS.
* Auto Scaling Group - EC2 worker nodes managed by the user.
* Kubectl Handler (Optional) - Custom resource (i.e Lambda Function) for invoking kubectl commands on the
  cluster - created by CDK

## Provisioning cluster

Creating a new cluster is done using the `Cluster` constructs. The only required property is the kubernetes version.

```go
eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
})
```

You can also use `FargateCluster` to provision a cluster that uses only fargate workers.

```go
eks.NewFargateCluster(this, jsii.String("HelloEKS"), &FargateClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
})
```

**Note: Unlike the previous EKS cluster, `Kubectl Handler` will not
be created by default. It will only be deployed when `kubectlProviderOptions`
property is used.**

```go
import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"


eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	KubectlProviderOptions: &KubectlProviderOptions{
		KubectlLayer: kubectlv32.NewKubectlV32Layer(this, jsii.String("kubectl")),
	},
})
```

## EKS Auto Mode

[Amazon EKS Auto Mode](https://aws.amazon.com/eks/auto-mode/) extends AWS management of Kubernetes clusters beyond the cluster itself, allowing AWS to set up and manage the infrastructure that enables the smooth operation of your workloads.

### Using Auto Mode

While `aws-eks` uses `DefaultCapacityType.NODEGROUP` by default, `aws-eks-v2` uses `DefaultCapacityType.AUTOMODE` as the default capacity type.

Auto Mode is enabled by default when creating a new cluster without specifying any capacity-related properties:

```go
// Create EKS cluster with Auto Mode implicitly enabled
cluster := eks.NewCluster(this, jsii.String("EksAutoCluster"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
})
```

You can also explicitly enable Auto Mode using `defaultCapacityType`:

```go
// Create EKS cluster with Auto Mode explicitly enabled
cluster := eks.NewCluster(this, jsii.String("EksAutoCluster"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	DefaultCapacityType: eks.DefaultCapacityType_AUTOMODE,
})
```

### Node Pools

When Auto Mode is enabled, the cluster comes with two default node pools:

* `system`: For running system components and add-ons
* `general-purpose`: For running your application workloads

These node pools are managed automatically by EKS. You can configure which node pools to enable through the `compute` property:

```go
cluster := eks.NewCluster(this, jsii.String("EksAutoCluster"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	DefaultCapacityType: eks.DefaultCapacityType_AUTOMODE,
	Compute: &ComputeConfig{
		NodePools: []*string{
			jsii.String("system"),
			jsii.String("general-purpose"),
		},
	},
})
```

For more information, see [Create a Node Pool for EKS Auto Mode](https://docs.aws.amazon.com/eks/latest/userguide/create-node-pool.html).

### Node Groups as the default capacity type

If you prefer to manage your own node groups instead of using Auto Mode, you can use the traditional node group approach by specifying `defaultCapacityType` as `NODEGROUP`:

```go
// Create EKS cluster with traditional managed node group
cluster := eks.NewCluster(this, jsii.String("EksCluster"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	DefaultCapacityType: eks.DefaultCapacityType_NODEGROUP,
	DefaultCapacity: jsii.Number(3),
	 // Number of instances
	DefaultCapacityInstance: ec2.InstanceType_Of(ec2.InstanceClass_T3, ec2.InstanceSize_LARGE),
})
```

You can also create a cluster with no initial capacity and add node groups later:

```go
cluster := eks.NewCluster(this, jsii.String("EksCluster"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	DefaultCapacityType: eks.DefaultCapacityType_NODEGROUP,
	DefaultCapacity: jsii.Number(0),
})

// Add node groups as needed
cluster.AddNodegroupCapacity(jsii.String("custom-node-group"), &NodegroupOptions{
	MinSize: jsii.Number(1),
	MaxSize: jsii.Number(3),
	InstanceTypes: []instanceType{
		ec2.*instanceType_Of(ec2.InstanceClass_T3, ec2.InstanceSize_LARGE),
	},
})
```

Read [Managed node groups](#managed-node-groups) for more information on how to add node groups to the cluster.

### Mixed with Auto Mode and Node Groups

You can combine Auto Mode with traditional node groups for specific workload requirements:

```go
cluster := eks.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	DefaultCapacityType: eks.DefaultCapacityType_AUTOMODE,
	Compute: &ComputeConfig{
		NodePools: []*string{
			jsii.String("system"),
			jsii.String("general-purpose"),
		},
	},
})

// Add specialized node group for specific workloads
cluster.AddNodegroupCapacity(jsii.String("specialized-workload"), &NodegroupOptions{
	MinSize: jsii.Number(1),
	MaxSize: jsii.Number(3),
	InstanceTypes: []instanceType{
		ec2.*instanceType_Of(ec2.InstanceClass_C5, ec2.InstanceSize_XLARGE),
	},
	Labels: map[string]*string{
		"workload": jsii.String("specialized"),
	},
})
```

### Important Notes

1. Auto Mode and traditional capacity management are mutually exclusive at the default capacity level. You cannot opt in to Auto Mode and specify `defaultCapacity` or `defaultCapacityInstance`.
2. When Auto Mode is enabled:

   * The cluster will automatically manage compute resources
   * Node pools cannot be modified, only enabled or disabled
   * EKS will handle scaling and management of the node pools
3. Auto Mode requires specific IAM permissions. The construct will automatically attach the required managed policies.

### Managed node groups

Amazon EKS managed node groups automate the provisioning and lifecycle management of nodes (Amazon EC2 instances) for Amazon EKS Kubernetes clusters.
With Amazon EKS managed node groups, you don't need to separately provision or register the Amazon EC2 instances that provide compute capacity to run your Kubernetes applications. You can create, update, or terminate nodes for your cluster with a single operation. Nodes run using the latest Amazon EKS optimized AMIs in your AWS account while node updates and terminations gracefully drain nodes to ensure that your applications stay available.

> For more details visit [Amazon EKS Managed Node Groups](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).

By default, when using `DefaultCapacityType.NODEGROUP`, this library will allocate a managed node group with 2 *m5.large* instances (this instance type suits most common use-cases, and is good value for money).

```go
eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	DefaultCapacityType: eks.DefaultCapacityType_NODEGROUP,
})
```

At cluster instantiation time, you can customize the number of instances and their type:

```go
eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	DefaultCapacityType: eks.DefaultCapacityType_NODEGROUP,
	DefaultCapacity: jsii.Number(5),
	DefaultCapacityInstance: ec2.InstanceType_Of(ec2.InstanceClass_M5, ec2.InstanceSize_SMALL),
})
```

To access the node group that was created on your behalf, you can use `cluster.defaultNodegroup`.

Additional customizations are available post instantiation. To apply them, set the default capacity to 0, and use the `cluster.addNodegroupCapacity` method:

```go
cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	DefaultCapacityType: eks.DefaultCapacityType_NODEGROUP,
	DefaultCapacity: jsii.Number(0),
})

cluster.AddNodegroupCapacity(jsii.String("custom-node-group"), &NodegroupOptions{
	InstanceTypes: []instanceType{
		ec2.NewInstanceType(jsii.String("m5.large")),
	},
	MinSize: jsii.Number(4),
	DiskSize: jsii.Number(100),
})
```

### Fargate profiles

AWS Fargate is a technology that provides on-demand, right-sized compute
capacity for containers. With AWS Fargate, you no longer have to provision,
configure, or scale groups of virtual machines to run containers. This removes
the need to choose server types, decide when to scale your node groups, or
optimize cluster packing.

You can control which pods start on Fargate and how they run with Fargate
Profiles, which are defined as part of your Amazon EKS cluster.

See [Fargate Considerations](https://docs.aws.amazon.com/eks/latest/userguide/fargate.html#fargate-considerations) in the AWS EKS User Guide.

You can add Fargate Profiles to any EKS cluster defined in your CDK app
through the `addFargateProfile()` method. The following example adds a profile
that will match all pods from the "default" namespace:

```go
var cluster cluster

cluster.AddFargateProfile(jsii.String("MyProfile"), &FargateProfileOptions{
	Selectors: []selector{
		&selector{
			Namespace: jsii.String("default"),
		},
	},
})
```

You can also directly use the `FargateProfile` construct to create profiles under different scopes:

```go
var cluster cluster

eks.NewFargateProfile(this, jsii.String("MyProfile"), &FargateProfileProps{
	Cluster: Cluster,
	Selectors: []selector{
		&selector{
			Namespace: jsii.String("default"),
		},
	},
})
```

To create an EKS cluster that **only** uses Fargate capacity, you can use `FargateCluster`.
The following code defines an Amazon EKS cluster with a default Fargate Profile that matches all pods from the "kube-system" and "default" namespaces. It is also configured to [run CoreDNS on Fargate](https://docs.aws.amazon.com/eks/latest/userguide/fargate-getting-started.html#fargate-gs-coredns).

```go
cluster := eks.NewFargateCluster(this, jsii.String("MyCluster"), &FargateClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
})
```

`FargateCluster` will create a default `FargateProfile` which can be accessed via the cluster's `defaultProfile` property. The created profile can also be customized by passing options as with `addFargateProfile`.

**NOTE**: Classic Load Balancers and Network Load Balancers are not supported on
pods running on Fargate. For ingress, we recommend that you use the [ALB Ingress
Controller](https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html)
on Amazon EKS (minimum version v1.1.4).

### Endpoint Access

When you create a new cluster, Amazon EKS creates an endpoint for the managed Kubernetes API server that you use to communicate with your cluster (using Kubernetes management tools such as `kubectl`)

You can configure the [cluster endpoint access](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) by using the `endpointAccess` property:

```go
cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	EndpointAccess: eks.EndpointAccess_PRIVATE(),
})
```

The default value is `eks.EndpointAccess.PUBLIC_AND_PRIVATE`. Which means the cluster endpoint is accessible from outside of your VPC, but worker node traffic and `kubectl` commands issued by this library stay within your VPC.

### Alb Controller

Some Kubernetes resources are commonly implemented on AWS with the help of the [ALB Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/).

From the docs:

> AWS Load Balancer Controller is a controller to help manage Elastic Load Balancers for a Kubernetes cluster.
>
> * It satisfies Kubernetes Ingress resources by provisioning Application Load Balancers.
> * It satisfies Kubernetes Service resources by provisioning Network Load Balancers.

To deploy the controller on your EKS cluster, configure the `albController` property:

```go
eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	AlbController: &AlbControllerOptions{
		Version: eks.AlbControllerVersion_V2_8_2(),
	},
})
```

The `albController` requires `defaultCapacity` or at least one nodegroup. If there's no `defaultCapacity` or available
nodegroup for the cluster, the `albController` deployment would fail.

Querying the controller pods should look something like this:

```console
❯ kubectl get pods -n kube-system
NAME                                            READY   STATUS    RESTARTS   AGE
aws-load-balancer-controller-76bd6c7586-d929p   1/1     Running   0          109m
aws-load-balancer-controller-76bd6c7586-fqxph   1/1     Running   0          109m
...
...
```

Every Kubernetes manifest that utilizes the ALB Controller is effectively dependant on the controller.
If the controller is deleted before the manifest, it might result in dangling ELB/ALB resources.
Currently, the EKS construct library does not detect such dependencies, and they should be done explicitly.

For example:

```go
var cluster cluster

manifest := cluster.addManifest(jsii.String("manifest"), map[string]interface{}{
})
if cluster.AlbController {
	manifest.Node.AddDependency(cluster.AlbController)
}
```

You can specify the VPC of the cluster using the `vpc` and `vpcSubnets` properties:

```go
var vpc vpc


eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	Vpc: Vpc,
	VpcSubnets: []subnetSelection{
		&subnetSelection{
			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
		},
	},
})
```

If you do not specify a VPC, one will be created on your behalf, which you can then access via `cluster.vpc`. The cluster VPC will be associated to any EKS managed capacity (i.e Managed Node Groups and Fargate Profiles).

Please note that the `vpcSubnets` property defines the subnets where EKS will place the *control plane* ENIs. To choose
the subnets where EKS will place the worker nodes, please refer to the **Provisioning clusters** section above.

If you allocate self managed capacity, you can specify which subnets should the auto-scaling group use:

```go
var vpc vpc
var cluster cluster

cluster.AddAutoScalingGroupCapacity(jsii.String("nodes"), &AutoScalingGroupCapacityOptions{
	VpcSubnets: &SubnetSelection{
		Subnets: vpc.PrivateSubnets,
	},
	InstanceType: ec2.NewInstanceType(jsii.String("t2.medium")),
})
```

There is an additional components you might want to provision within the VPC.

The `KubectlHandler` is a Lambda function responsible to issuing `kubectl` and `helm` commands against the cluster when you add resource manifests to the cluster.

The handler association to the VPC is derived from the `endpointAccess` configuration. The rule of thumb is: *If the cluster VPC can be associated, it will be*.

Breaking this down, it means that if the endpoint exposes private access (via `EndpointAccess.PRIVATE` or `EndpointAccess.PUBLIC_AND_PRIVATE`), and the VPC contains **private** subnets, the Lambda function will be provisioned inside the VPC and use the private subnets to interact with the cluster. This is the common use-case.

If the endpoint does not expose private access (via `EndpointAccess.PUBLIC`) **or** the VPC does not contain private subnets, the function will not be provisioned within the VPC.

If your use-case requires control over the IAM role that the KubeCtl Handler assumes, a custom role can be passed through the ClusterProps (as `kubectlLambdaRole`) of the EKS Cluster construct.

### Kubectl Support

You can choose to have CDK create a `Kubectl Handler` - a Python Lambda Function to
apply k8s manifests using `kubectl apply`. This handler will not be created by default.

To create a `Kubectl Handler`, use `kubectlProviderOptions` when creating the cluster.
`kubectlLayer` is the only required property in `kubectlProviderOptions`.

```go
import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"


eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	KubectlProviderOptions: &KubectlProviderOptions{
		KubectlLayer: kubectlv32.NewKubectlV32Layer(this, jsii.String("kubectl")),
	},
})
```

`Kubectl Handler` created along with the cluster will be granted admin permissions to the cluster.

If you want to use an existing kubectl provider function, for example with tight trusted entities on your IAM Roles - you can import the existing provider and then use the imported provider when importing the cluster:

```go
import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"


handlerRole := iam.Role_FromRoleArn(this, jsii.String("HandlerRole"), jsii.String("arn:aws:iam::123456789012:role/lambda-role"))
// get the serivceToken from the custom resource provider
functionArn := lambda.Function_FromFunctionName(this, jsii.String("ProviderOnEventFunc"), jsii.String("ProviderframeworkonEvent-XXX")).FunctionArn
kubectlProvider := eks.KubectlProvider_FromKubectlProviderAttributes(this, jsii.String("KubectlProvider"), &KubectlProviderAttributes{
	ServiceToken: functionArn,
	Role: handlerRole,
})

cluster := eks.Cluster_FromClusterAttributes(this, jsii.String("Cluster"), &ClusterAttributes{
	ClusterName: jsii.String("cluster"),
	KubectlProvider: KubectlProvider,
})
```

#### Environment

You can configure the environment of this function by specifying it at cluster instantiation. For example, this can be useful in order to configure an http proxy:

```go
import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"


cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	KubectlProviderOptions: &KubectlProviderOptions{
		KubectlLayer: kubectlv32.NewKubectlV32Layer(this, jsii.String("kubectl")),
		Environment: map[string]*string{
			"http_proxy": jsii.String("http://proxy.myproxy.com"),
		},
	},
})
```

#### Runtime

The kubectl handler uses `kubectl`, `helm` and the `aws` CLI in order to
interact with the cluster. These are bundled into AWS Lambda layers included in
the `@aws-cdk/lambda-layer-awscli` and `@aws-cdk/lambda-layer-kubectl` modules.

The version of kubectl used must be compatible with the Kubernetes version of the
cluster. kubectl is supported within one minor version (older or newer) of Kubernetes
(see [Kubernetes version skew policy](https://kubernetes.io/releases/version-skew-policy/#kubectl)).
Depending on which version of kubernetes you're targeting, you will need to use one of
the `@aws-cdk/lambda-layer-kubectl-vXY` packages.

```go
import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"


cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	KubectlProviderOptions: &KubectlProviderOptions{
		KubectlLayer: kubectlv32.NewKubectlV32Layer(this, jsii.String("kubectl")),
	},
})
```

#### Memory

By default, the kubectl provider is configured with 1024MiB of memory. You can use the `memory` option to specify the memory size for the AWS Lambda function:

```go
import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"


eks.NewCluster(this, jsii.String("MyCluster"), &ClusterProps{
	KubectlProviderOptions: &KubectlProviderOptions{
		KubectlLayer: kubectlv32.NewKubectlV32Layer(this, jsii.String("kubectl")),
		Memory: awscdk.Size_Gibibytes(jsii.Number(4)),
	},
	Version: eks.KubernetesVersion_V1_32(),
})
```

### ARM64 Support

Instance types with `ARM64` architecture are supported in both managed nodegroup and self-managed capacity. Simply specify an ARM64 `instanceType` (such as `m6g.medium`), and the latest
Amazon Linux 2 AMI for ARM64 will be automatically selected.

```go
var cluster cluster

// add a managed ARM64 nodegroup
cluster.AddNodegroupCapacity(jsii.String("extra-ng-arm"), &NodegroupOptions{
	InstanceTypes: []instanceType{
		ec2.NewInstanceType(jsii.String("m6g.medium")),
	},
	MinSize: jsii.Number(2),
})

// add a self-managed ARM64 nodegroup
cluster.AddAutoScalingGroupCapacity(jsii.String("self-ng-arm"), &AutoScalingGroupCapacityOptions{
	InstanceType: ec2.NewInstanceType(jsii.String("m6g.medium")),
	MinCapacity: jsii.Number(2),
})
```

### Masters Role

When you create a cluster, you can specify a `mastersRole`. The `Cluster` construct will associate this role with `AmazonEKSClusterAdminPolicy` through [Access Entry](https://docs.aws.amazon.com/eks/latest/userguide/access-policy-permissions.html).

```go
var role role

eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	MastersRole: role,
})
```

If you do not specify it, you won't have access to the cluster from outside of the CDK application.

### Encryption

When you create an Amazon EKS cluster, envelope encryption of Kubernetes secrets using the AWS Key Management Service (AWS KMS) can be enabled.
The documentation on [creating a cluster](https://docs.aws.amazon.com/eks/latest/userguide/create-cluster.html)
can provide more details about the customer master key (CMK) that can be used for the encryption.

You can use the `secretsEncryptionKey` to configure which key the cluster will use to encrypt Kubernetes secrets. By default, an AWS Managed key will be used.

> This setting can only be specified when the cluster is created and cannot be updated.

```go
secretsKey := kms.NewKey(this, jsii.String("SecretsKey"))
cluster := eks.NewCluster(this, jsii.String("MyCluster"), &ClusterProps{
	SecretsEncryptionKey: secretsKey,
	Version: eks.KubernetesVersion_V1_32(),
})
```

You can also use a similar configuration for running a cluster built using the FargateCluster construct.

```go
secretsKey := kms.NewKey(this, jsii.String("SecretsKey"))
cluster := eks.NewFargateCluster(this, jsii.String("MyFargateCluster"), &FargateClusterProps{
	SecretsEncryptionKey: secretsKey,
	Version: eks.KubernetesVersion_V1_32(),
})
```

The Amazon Resource Name (ARN) for that CMK can be retrieved.

```go
var cluster cluster

clusterEncryptionConfigKeyArn := cluster.ClusterEncryptionConfigKeyArn
```

## Permissions and Security

In the new EKS module, `ConfigMap` is deprecated. Clusters created by the new module will use `API` as authentication mode. Access Entry will be the only way for granting permissions to specific IAM users and roles.

### Access Entry

An access entry is a cluster identity—directly linked to an AWS IAM principal user or role that is used to authenticate to
an Amazon EKS cluster. An Amazon EKS access policy authorizes an access entry to perform specific cluster actions.

Access policies are Amazon EKS-specific policies that assign Kubernetes permissions to access entries. Amazon EKS supports
only predefined and AWS managed policies. Access policies are not AWS IAM entities and are defined and managed by Amazon EKS.
Amazon EKS access policies include permission sets that support common use cases of administration, editing, or read-only access
to Kubernetes resources. See [Access Policy Permissions](https://docs.aws.amazon.com/eks/latest/userguide/access-policies.html#access-policy-permissions) for more details.

Use `AccessPolicy` to include predefined AWS managed policies:

```go
// AmazonEKSClusterAdminPolicy with `cluster` scope
eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSClusterAdminPolicy"), &AccessPolicyNameOptions{
	AccessScopeType: eks.AccessScopeType_CLUSTER,
})
// AmazonEKSAdminPolicy with `namespace` scope
eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSAdminPolicy"), &AccessPolicyNameOptions{
	AccessScopeType: eks.AccessScopeType_NAMESPACE,
	Namespaces: []*string{
		jsii.String("foo"),
		jsii.String("bar"),
	},
})
```

Use `grantAccess()` to grant the AccessPolicy to an IAM principal:

```go
import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"
var vpc vpc


clusterAdminRole := iam.NewRole(this, jsii.String("ClusterAdminRole"), &RoleProps{
	AssumedBy: iam.NewArnPrincipal(jsii.String("arn_for_trusted_principal")),
})

eksAdminRole := iam.NewRole(this, jsii.String("EKSAdminRole"), &RoleProps{
	AssumedBy: iam.NewArnPrincipal(jsii.String("arn_for_trusted_principal")),
})

cluster := eks.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	Vpc: Vpc,
	MastersRole: clusterAdminRole,
	Version: eks.KubernetesVersion_V1_32(),
	KubectlProviderOptions: &KubectlProviderOptions{
		KubectlLayer: kubectlv32.NewKubectlV32Layer(this, jsii.String("kubectl")),
		Memory: awscdk.Size_Gibibytes(jsii.Number(4)),
	},
})

// Cluster Admin role for this cluster
cluster.GrantAccess(jsii.String("clusterAdminAccess"), clusterAdminRole.RoleArn, []iAccessPolicy{
	eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSClusterAdminPolicy"), &AccessPolicyNameOptions{
		AccessScopeType: eks.AccessScopeType_CLUSTER,
	}),
})

// EKS Admin role for specified namespaces of this cluster
cluster.GrantAccess(jsii.String("eksAdminRoleAccess"), eksAdminRole.RoleArn, []iAccessPolicy{
	eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSAdminPolicy"), &AccessPolicyNameOptions{
		AccessScopeType: eks.AccessScopeType_NAMESPACE,
		Namespaces: []*string{
			jsii.String("foo"),
			jsii.String("bar"),
		},
	}),
})
```

By default, the cluster creator role will be granted the cluster admin permissions. You can disable it by setting
`bootstrapClusterCreatorAdminPermissions` to false.

> **Note** - Switching `bootstrapClusterCreatorAdminPermissions` on an existing cluster would cause cluster replacement and should be avoided in production.

### Cluster Security Group

When you create an Amazon EKS cluster, a [cluster security group](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html)
is automatically created as well. This security group is designed to allow all traffic from the control plane and managed node groups to flow freely
between each other.

The ID for that security group can be retrieved after creating the cluster.

```go
var cluster cluster

clusterSecurityGroupId := cluster.ClusterSecurityGroupId
```

## Applying Kubernetes Resources

To apply kubernetes resource, kubectl provider needs to be created for the cluster. You can use `kubectlProviderOptions` to create the kubectl Provider.

The library supports several popular resource deployment mechanisms, among which are:

### Kubernetes Manifests

The `KubernetesManifest` construct or `cluster.addManifest` method can be used
to apply Kubernetes resource manifests to this cluster.

> When using `cluster.addManifest`, the manifest construct is defined within the cluster's stack scope. If the manifest contains
> attributes from a different stack which depend on the cluster stack, a circular dependency will be created and you will get a synth time error.
> To avoid this, directly use `new KubernetesManifest` to create the manifest in the scope of the other stack.

The following examples will deploy the [paulbouwer/hello-kubernetes](https://github.com/paulbouwer/hello-kubernetes)
service on the cluster:

```go
var cluster cluster

appLabel := map[string]*string{
	"app": jsii.String("hello-kubernetes"),
}

deployment := map[string]interface{}{
	"apiVersion": jsii.String("apps/v1"),
	"kind": jsii.String("Deployment"),
	"metadata": map[string]*string{
		"name": jsii.String("hello-kubernetes"),
	},
	"spec": map[string]interface{}{
		"replicas": jsii.Number(3),
		"selector": map[string]map[string]*string{
			"matchLabels": appLabel,
		},
		"template": map[string]map[string]map[string]*string{
			"metadata": map[string]map[string]*string{
				"labels": appLabel,
			},
			"spec": map[string][]map[string]interface{}{
				"containers": []map[string]interface{}{
					map[string]interface{}{
						"name": jsii.String("hello-kubernetes"),
						"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
						"ports": []map[string]*f64{
							map[string]*f64{
								"containerPort": jsii.Number(8080),
							},
						},
					},
				},
			},
		},
	},
}

service := map[string]interface{}{
	"apiVersion": jsii.String("v1"),
	"kind": jsii.String("Service"),
	"metadata": map[string]*string{
		"name": jsii.String("hello-kubernetes"),
	},
	"spec": map[string]interface{}{
		"type": jsii.String("LoadBalancer"),
		"ports": []map[string]*f64{
			map[string]*f64{
				"port": jsii.Number(80),
				"targetPort": jsii.Number(8080),
			},
		},
		"selector": appLabel,
	},
}

// option 1: use a construct
// option 1: use a construct
eks.NewKubernetesManifest(this, jsii.String("hello-kub"), &KubernetesManifestProps{
	Cluster: Cluster,
	Manifest: []map[string]interface{}{
		deployment,
		service,
	},
})

// or, option2: use `addManifest`
cluster.addManifest(jsii.String("hello-kub"), service, deployment)
```

#### ALB Controller Integration

The `KubernetesManifest` construct can detect ingress resources inside your manifest and automatically add the necessary annotations
so they are picked up by the ALB Controller.

> See [Alb Controller](#alb-controller)

To that end, it offers the following properties:

* `ingressAlb` - Signal that the ingress detection should be done.
* `ingressAlbScheme` - Which ALB scheme should be applied. Defaults to `internal`.

#### Adding resources from a URL

The following example will deploy the resource manifest hosting on remote server:

```text
// This example is only available in TypeScript

import * as yaml from 'js-yaml';
import * as request from 'sync-request';

declare const cluster: eks.Cluster;
const manifestUrl = 'https://url/of/manifest.yaml';
const manifest = yaml.safeLoadAll(request('GET', manifestUrl).getBody());
cluster.addManifest('my-resource', manifest);
```

#### Dependencies

There are cases where Kubernetes resources must be deployed in a specific order.
For example, you cannot define a resource in a Kubernetes namespace before the
namespace was created.

You can represent dependencies between `KubernetesManifest`s using
`resource.node.addDependency()`:

```go
var cluster cluster

namespace := cluster.addManifest(jsii.String("my-namespace"), map[string]interface{}{
	"apiVersion": jsii.String("v1"),
	"kind": jsii.String("Namespace"),
	"metadata": map[string]*string{
		"name": jsii.String("my-app"),
	},
})

service := cluster.addManifest(jsii.String("my-service"), map[string]interface{}{
	"metadata": map[string]*string{
		"name": jsii.String("myservice"),
		"namespace": jsii.String("my-app"),
	},
	"spec": map[string]interface{}{
	},
})

service.Node.AddDependency(namespace)
```

**NOTE:** when a `KubernetesManifest` includes multiple resources (either directly
or through `cluster.addManifest()`) (e.g. `cluster.addManifest('foo', r1, r2, r3,...)`), these resources will be applied as a single manifest via `kubectl`
and will be applied sequentially (the standard behavior in `kubectl`).

---


Since Kubernetes manifests are implemented as CloudFormation resources in the
CDK. This means that if the manifest is deleted from your code (or the stack is
deleted), the next `cdk deploy` will issue a `kubectl delete` command and the
Kubernetes resources in that manifest will be deleted.

#### Resource Pruning

When a resource is deleted from a Kubernetes manifest, the EKS module will
automatically delete these resources by injecting a *prune label* to all
manifest resources. This label is then passed to [`kubectl apply --prune`](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/declarative-config/#alternative-kubectl-apply-f-directory-prune-l-your-label).

Pruning is enabled by default but can be disabled through the `prune` option
when a cluster is defined:

```go
eks.NewCluster(this, jsii.String("MyCluster"), &ClusterProps{
	Version: eks.KubernetesVersion_V1_32(),
	Prune: jsii.Boolean(false),
})
```

#### Manifests Validation

The `kubectl` CLI supports applying a manifest by skipping the validation.
This can be accomplished by setting the `skipValidation` flag to `true` in the `KubernetesManifest` props.

```go
var cluster cluster

eks.NewKubernetesManifest(this, jsii.String("HelloAppWithoutValidation"), &KubernetesManifestProps{
	Cluster: Cluster,
	Manifest: []map[string]interface{}{
		map[string]interface{}{
			"foo": jsii.String("bar"),
		},
	},
	SkipValidation: jsii.Boolean(true),
})
```

### Helm Charts

The `HelmChart` construct or `cluster.addHelmChart` method can be used
to add Kubernetes resources to this cluster using Helm.

> When using `cluster.addHelmChart`, the manifest construct is defined within the cluster's stack scope. If the manifest contains
> attributes from a different stack which depend on the cluster stack, a circular dependency will be created and you will get a synth time error.
> To avoid this, directly use `new HelmChart` to create the chart in the scope of the other stack.

The following example will install the [NGINX Ingress Controller](https://kubernetes.github.io/ingress-nginx/) to your cluster using Helm.

```go
var cluster cluster

// option 1: use a construct
// option 1: use a construct
eks.NewHelmChart(this, jsii.String("NginxIngress"), &HelmChartProps{
	Cluster: Cluster,
	Chart: jsii.String("nginx-ingress"),
	Repository: jsii.String("https://helm.nginx.com/stable"),
	Namespace: jsii.String("kube-system"),
})

// or, option2: use `addHelmChart`
cluster.addHelmChart(jsii.String("NginxIngress"), &HelmChartOptions{
	Chart: jsii.String("nginx-ingress"),
	Repository: jsii.String("https://helm.nginx.com/stable"),
	Namespace: jsii.String("kube-system"),
})
```

Helm charts will be installed and updated using `helm upgrade --install`, where a few parameters
are being passed down (such as `repo`, `values`, `version`, `namespace`, `wait`, `timeout`, etc).
This means that if the chart is added to CDK with the same release name, it will try to update
the chart in the cluster.

Additionally, the `chartAsset` property can be an `aws-s3-assets.Asset`. This allows the use of local, private helm charts.

```go
import s3Assets "github.com/aws/aws-cdk-go/awscdk"

var cluster cluster

chartAsset := s3Assets.NewAsset(this, jsii.String("ChartAsset"), &AssetProps{
	Path: jsii.String("/path/to/asset"),
})

cluster.addHelmChart(jsii.String("test-chart"), &HelmChartOptions{
	ChartAsset: chartAsset,
})
```

Nested values passed to the `values` parameter should be provided as a nested dictionary:

```go
var cluster cluster


cluster.addHelmChart(jsii.String("ExternalSecretsOperator"), &HelmChartOptions{
	Chart: jsii.String("external-secrets"),
	Release: jsii.String("external-secrets"),
	Repository: jsii.String("https://charts.external-secrets.io"),
	Namespace: jsii.String("external-secrets"),
	Values: map[string]interface{}{
		"installCRDs": jsii.Boolean(true),
		"webhook": map[string]*f64{
			"port": jsii.Number(9443),
		},
	},
})
```

Helm chart can come with Custom Resource Definitions (CRDs) defined that by default will be installed by helm as well. However in special cases it might be needed to skip the installation of CRDs, for that the property `skipCrds` can be used.

```go
var cluster cluster

// option 1: use a construct
// option 1: use a construct
eks.NewHelmChart(this, jsii.String("NginxIngress"), &HelmChartProps{
	Cluster: Cluster,
	Chart: jsii.String("nginx-ingress"),
	Repository: jsii.String("https://helm.nginx.com/stable"),
	Namespace: jsii.String("kube-system"),
	SkipCrds: jsii.Boolean(true),
})
```

### OCI Charts

OCI charts are also supported.
Also replace the `${VARS}` with appropriate values.

```go
var cluster cluster

// option 1: use a construct
// option 1: use a construct
eks.NewHelmChart(this, jsii.String("MyOCIChart"), &HelmChartProps{
	Cluster: Cluster,
	Chart: jsii.String("some-chart"),
	Repository: jsii.String("oci://${ACCOUNT_ID}.dkr.ecr.${ACCOUNT_REGION}.amazonaws.com/${REPO_NAME}"),
	Namespace: jsii.String("oci"),
	Version: jsii.String("0.0.1"),
})
```

Helm charts are implemented as CloudFormation resources in CDK.
This means that if the chart is deleted from your code (or the stack is
deleted), the next `cdk deploy` will issue a `helm uninstall` command and the
Helm chart will be deleted.

When there is no `release` defined, a unique ID will be allocated for the release based
on the construct path.

By default, all Helm charts will be installed concurrently. In some cases, this
could cause race conditions where two Helm charts attempt to deploy the same
resource or if Helm charts depend on each other. You can use
`chart.node.addDependency()` in order to declare a dependency order between
charts:

```go
var cluster cluster

chart1 := cluster.addHelmChart(jsii.String("MyChart"), &HelmChartOptions{
	Chart: jsii.String("foo"),
})
chart2 := cluster.addHelmChart(jsii.String("MyChart"), &HelmChartOptions{
	Chart: jsii.String("bar"),
})

chart2.Node.AddDependency(chart1)
```

#### Custom CDK8s Constructs

You can also compose a few stock `cdk8s+` constructs into your own custom construct. However, since mixing scopes between `aws-cdk` and `cdk8s` is currently not supported, the `Construct` class
you'll need to use is the one from the [`constructs`](https://github.com/aws/constructs) module, and not from `aws-cdk-lib` like you normally would.
This is why we used `new cdk8s.App()` as the scope of the chart above.

```go
import constructs "github.com/aws/constructs-go/constructs"
import "github.com/cdk8s-team/cdk8s-core-go/cdk8s"
import "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25"


type loadBalancedWebService struct {
	port *f64
	image *string
	replicas *f64
}

app := cdk8s.NewApp()
chart := cdk8s.NewChart(app, jsii.String("my-chart"))

type loadBalancedWebService struct {
	construct
}

func newLoadBalancedWebService(scope construct, id *string, props loadBalancedWebService) *loadBalancedWebService {
	this := &loadBalancedWebService{}
	constructs.NewConstruct_Override(this, scope, id)

	deployment := kplus.NewDeployment(chart, jsii.String("Deployment"), &DeploymentProps{
		Replicas: props.replicas,
		Containers: []containerProps{
			kplus.NewContainer(&containerProps{
				Image: props.image,
			}),
		},
	})

	deployment.ExposeViaService(&DeploymentExposeViaServiceOptions{
		Ports: []servicePort{
			&servicePort{
				Port: props.port,
			},
		},
		ServiceType: kplus.ServiceType_LOAD_BALANCER,
	})
	return this
}
```

#### Manually importing k8s specs and CRD's

If you find yourself unable to use `cdk8s+`, or just like to directly use the `k8s` native objects or CRD's, you can do so by manually importing them using the `cdk8s-cli`.

See [Importing kubernetes objects](https://cdk8s.io/docs/latest/cli/import/) for detailed instructions.

## Patching Kubernetes Resources

The `KubernetesPatch` construct can be used to update existing kubernetes
resources. The following example can be used to patch the `hello-kubernetes`
deployment from the example above with 5 replicas.

```go
var cluster cluster

eks.NewKubernetesPatch(this, jsii.String("hello-kub-deployment-label"), &KubernetesPatchProps{
	Cluster: Cluster,
	ResourceName: jsii.String("deployment/hello-kubernetes"),
	ApplyPatch: map[string]interface{}{
		"spec": map[string]*f64{
			"replicas": jsii.Number(5),
		},
	},
	RestorePatch: map[string]interface{}{
		"spec": map[string]*f64{
			"replicas": jsii.Number(3),
		},
	},
})
```

## Querying Kubernetes Resources

The `KubernetesObjectValue` construct can be used to query for information about kubernetes objects,
and use that as part of your CDK application.

For example, you can fetch the address of a [`LoadBalancer`](https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer) type service:

```go
var cluster cluster

// query the load balancer address
myServiceAddress := eks.NewKubernetesObjectValue(this, jsii.String("LoadBalancerAttribute"), &KubernetesObjectValueProps{
	Cluster: cluster,
	ObjectType: jsii.String("service"),
	ObjectName: jsii.String("my-service"),
	JsonPath: jsii.String(".status.loadBalancer.ingress[0].hostname"),
})

// pass the address to a lambda function
proxyFunction := lambda.NewFunction(this, jsii.String("ProxyFunction"), &FunctionProps{
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("my-code")),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Environment: map[string]*string{
		"myServiceAddress": myServiceAddress.value,
	},
})
```

Specifically, since the above use-case is quite common, there is an easier way to access that information:

```go
var cluster cluster

loadBalancerAddress := cluster.GetServiceLoadBalancerAddress(jsii.String("my-service"))
```

## Add-ons

[Add-ons](https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html) is a software that provides supporting operational capabilities to Kubernetes applications. The EKS module supports adding add-ons to your cluster using the `eks.Addon` class.

```go
var cluster cluster


eks.NewAddon(this, jsii.String("Addon"), &AddonProps{
	Cluster: Cluster,
	AddonName: jsii.String("aws-guardduty-agent"),
	AddonVersion: jsii.String("v1.6.1"),
	// whether to preserve the add-on software on your cluster but Amazon EKS stops managing any settings for the add-on.
	PreserveOnDelete: jsii.Boolean(false),
})
```

## Using existing clusters

The EKS library allows defining Kubernetes resources such as [Kubernetes
manifests](#kubernetes-resources) and [Helm charts](#helm-charts) on clusters
that are not defined as part of your CDK app.

First you will need to import the kubectl provider and cluster created in another stack

```go
handlerRole := iam.Role_FromRoleArn(this, jsii.String("HandlerRole"), jsii.String("arn:aws:iam::123456789012:role/lambda-role"))

kubectlProvider := eks.KubectlProvider_FromKubectlProviderAttributes(this, jsii.String("KubectlProvider"), &KubectlProviderAttributes{
	ServiceToken: jsii.String("arn:aws:lambda:us-east-2:123456789012:function:my-function:1"),
	Role: handlerRole,
})

cluster := eks.Cluster_FromClusterAttributes(this, jsii.String("Cluster"), &ClusterAttributes{
	ClusterName: jsii.String("cluster"),
	KubectlProvider: KubectlProvider,
})
```

Then, you can use `addManifest` or `addHelmChart` to define resources inside
your Kubernetes cluster.

```go
var cluster cluster

cluster.addManifest(jsii.String("Test"), map[string]interface{}{
	"apiVersion": jsii.String("v1"),
	"kind": jsii.String("ConfigMap"),
	"metadata": map[string]*string{
		"name": jsii.String("myconfigmap"),
	},
	"data": map[string]*string{
		"Key": jsii.String("value"),
		"Another": jsii.String("123454"),
	},
})
```

## Logging

EKS supports cluster logging for 5 different types of events:

* API requests to the cluster.
* Cluster access via the Kubernetes API.
* Authentication requests into the cluster.
* State of cluster controllers.
* Scheduling decisions.

You can enable logging for each one separately using the `clusterLogging`
property. For example:

```go
cluster := eks.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	// ...
	Version: eks.KubernetesVersion_V1_32(),
	ClusterLogging: []clusterLoggingTypes{
		eks.*clusterLoggingTypes_API,
		eks.*clusterLoggingTypes_AUTHENTICATOR,
		eks.*clusterLoggingTypes_SCHEDULER,
	},
})
```
