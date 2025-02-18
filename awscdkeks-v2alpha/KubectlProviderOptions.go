package awscdkeks-v2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"
//
//
//   cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_32(),
//   	KubectlProviderOptions: &KubectlProviderOptions{
//   		KubectlLayer: kubectlv32.NewKubectlV32Layer(this, jsii.String("kubectl")),
//   		Environment: map[string]*string{
//   			"http_proxy": jsii.String("http://proxy.myproxy.com"),
//   		},
//   	},
//   })
//
// Experimental.
type KubectlProviderOptions struct {
	// An AWS Lambda layer that includes `kubectl` and `helm`.
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `field:"required" json:"kubectlLayer" yaml:"kubectlLayer"`
	// An AWS Lambda layer that contains the `aws` CLI.
	//
	// If not defined, a default layer will be used containing the AWS CLI 2.x.
	// Experimental.
	AwscliLayer awslambda.ILayerVersion `field:"optional" json:"awscliLayer" yaml:"awscliLayer"`
	// Custom environment variables when running `kubectl` against this cluster.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The amount of memory allocated to the kubectl provider's lambda function.
	// Experimental.
	Memory awscdk.Size `field:"optional" json:"memory" yaml:"memory"`
	// Subnets to host the `kubectl` compute resources.
	//
	// If not specified, the k8s
	// endpoint is expected to be accessible publicly.
	// Experimental.
	PrivateSubnets *[]awsec2.ISubnet `field:"optional" json:"privateSubnets" yaml:"privateSubnets"`
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	//
	// This role is directly passed to the lambda handler that sends Kube Ctl commands to the cluster.
	// Default: - if not specified, the default role created by a lambda function will
	// be used.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// A security group to use for `kubectl` execution.
	// Default: - If not specified, the k8s endpoint is expected to be accessible
	// publicly.
	//
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

